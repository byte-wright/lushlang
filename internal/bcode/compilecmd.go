package bcode

import (
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

var (
	stringType = &common.BasicType{Type: common.String}
	intType    = &common.BasicType{Type: common.Int}
	boolType   = &common.BasicType{Type: common.Bool}
)

// cmdRun is a cmd whose errors are checked but not handled yet
type cmdRun struct {
	results []Value
	// failed and message are nil when it can't fail
	failed  *VarValue
	message *VarValue
}

// evalCmd runs a cmd whose error goes on, like a call with try.
func (c *Compiler) evalCmd(block *Block, x *ast.Cmd, loop *ast.ForIn) []Value {
	run := c.runCmd(block, x, loop)

	if run.failed != nil {
		fail := block.subBlock()
		fail.add(&Fail{Message: run.message})
		block.add(&If{Condition: run.failed, Block: fail})
	}

	return run.results
}

// catchCmd runs a cmd, assign gets its results. When it fails, the block of catch runs
// instead.
func (c *Compiler) catchCmd(block *Block, x *ast.Cmd, catch *ast.Catch, assign func(then *Block, results []Value)) {
	run := c.runCmd(block, x, nil)

	st := &If{Condition: run.failed, Block: block.subBlock(), Else: block.subBlock()}

	if catch.Error.Name != "_" {
		st.Block.set(&VarValue{Name: varName(catch.Error.Name), T: stringType}, run.message)
	}

	c.evalBlock(st.Block, catch.Block)

	if assign != nil {
		assign(st.Else, run.results)
	}

	block.add(st)
}

// runCmd adds the Cmd and the checks of its errors. The body is the block after the call,
// or loop, the for that reads the pipe.
func (c *Compiler) runCmd(block *Block, x *ast.Cmd, loop *ast.ForIn) *cmdRun {
	st := &Cmd{}
	main := c.proc(block, x, st)

	if x.Body != nil || loop != nil {
		st.Pipe = block.nextTmp(&common.PipeType{})
		st.Body = block.subBlock()
		main.Stopped = block.nextTmp(boolType)

		if loop != nil {
			c.readLines(st.Body, st.Pipe, loop)
		} else {
			st.Body.vars[varName(x.Pipe())] = st.Pipe
			c.evalBlock(st.Body, x.Body)
		}
	}

	block.add(st)

	run := &cmdRun{}
	for _, out := range []*ProcOutput{main.Stdout, main.Stderr} {
		if out != nil && out.Kind == OutCollect {
			run.results = append(run.results, out.Var)
		}
	}

	switch x.Exit {
	case ast.ExitInt:
		run.results = append(run.results, main.Code)
	case ast.ExitBool:
		run.results = append(run.results, block.setTmp(&Equal{Left: main.Code, Right: &NumberValue{Value: 0}}))
	}

	if x.CanFail {
		run.failed, run.message = c.cmdErrors(block, st, x)
	}

	return run
}

// proc adds the proc of x and the ones of its sources to st, the sources come first.
func (c *Compiler) proc(block *Block, x *ast.Cmd, st *Cmd) *Proc {
	p := &Proc{Code: block.nextTmp(intType)}

	for i, arg := range x.Args {
		inner, isCmd := arg.Value.(*ast.Cmd)
		if isCmd && inner.Use == ast.UseSource {
			p.Args = append(p.Args, &ProcArg{Source: c.source(block, inner, st)})
			continue
		}

		val := c.evalExpression(block, arg.Value)[0]
		if arg.Spread {
			val = block.asVar(val)
		}

		if i == 0 {
			p.Name = val
			continue
		}

		p.Args = append(p.Args, &ProcArg{Value: val, Spread: arg.Spread})
	}

	// the message of a failure has the collected stderr
	needsText := x.Exit == ast.ExitFails || x.Exit == ast.ExitBool
	p.Stdout = c.output(block, x.Stdout, false)
	p.Stderr = c.output(block, x.Stderr, needsText)

	if x.Stdin != nil {
		switch x.Stdin.Kind {
		case ast.InputString:
			p.Stdin = &ProcInput{Kind: InputString, Value: c.evalExpression(block, x.Stdin.Value)[0]}
		case ast.InputFile:
			p.Stdin = &ProcInput{Kind: InputFile, Value: c.evalExpression(block, x.Stdin.Value)[0]}
		case ast.InputCmd:
			p.Stdin = &ProcInput{Kind: InputSource, Source: c.source(block, x.Stdin.Value.(*ast.Cmd), st)}
		}
	}

	if x.Dir != nil {
		p.Dir = c.evalExpression(block, x.Dir)[0]
	}

	for _, env := range x.Env {
		p.Env = append(p.Env, &ProcEnv{Name: env.Name, Value: c.evalExpression(block, env.Value)[0]})
	}

	st.Procs = append(st.Procs, p)

	return p
}

func (c *Compiler) source(block *Block, x *ast.Cmd, st *Cmd) *Proc {
	p := c.proc(block, x, st)
	p.Stopped = block.nextTmp(boolType)

	return p
}

func (c *Compiler) output(block *Block, t *ast.Target, needsText bool) *ProcOutput {
	if t == nil {
		return nil
	}

	switch t.Kind {
	case ast.CollectString:
		v := block.nextTmp(stringType)
		out := &ProcOutput{Kind: OutCollect, Var: v}
		if needsText {
			out.Text = v
		}
		return out

	case ast.CollectLines:
		out := &ProcOutput{Kind: OutCollect, Var: block.nextTmp(&common.ArrayType{ElementType: stringType})}
		if needsText {
			out.Text = block.nextTmp(stringType)
		}
		return out

	case ast.PipeTarget:
		return &ProcOutput{Kind: OutPipe}

	case ast.FileTarget:
		return &ProcOutput{Kind: OutFile, Path: c.evalExpression(block, t.Path)[0]}

	case ast.AppendTarget:
		return &ProcOutput{Kind: OutAppend, Path: c.evalExpression(block, t.Path)[0]}

	case ast.DiscardTarget:
		return &ProcOutput{Kind: OutDiscard}
	}

	return &ProcOutput{Kind: OutStdout}
}

// cmdErrors returns whether one of the procs failed, and the message of the first one that
// did. A proc that was stopped didn't fail.
func (c *Compiler) cmdErrors(block *Block, st *Cmd, x *ast.Cmd) (*VarValue, *VarValue) {
	failed := block.setTmp(&BoolValue{Value: false})
	message := block.setTmp(&StringValue{Value: ""})

	exits := []ast.ExitMode{}
	sourceExits(x, &exits)

	for i, p := range st.Procs {
		exit := exits[i]
		if exit != ast.ExitFails && exit != ast.ExitBool {
			continue
		}

		conds := []Atom{&NotEqual{Left: p.Code, Right: &NumberValue{Value: 0}}}
		if exit == ast.ExitBool {
			conds = append(conds, &NotEqual{Left: p.Code, Right: &NumberValue{Value: 1}})
		}
		if p.Stopped != nil {
			conds = append(conds, &Not{Expression: p.Stopped})
		}
		conds = append(conds, &Not{Expression: failed})

		inner := block
		for _, cond := range conds {
			condVar := inner.setTmp(cond)
			then := inner.subBlock()
			inner.add(&If{Condition: condVar, Block: then})
			inner = then
		}

		text := inner.setTmp(&Add{Left: p.Name, Right: &StringValue{Value: " failed with exit code "}})
		inner.set(message, &Add{Left: text, Right: p.Code})

		if p.Stderr != nil && p.Stderr.Text != nil {
			hasText := inner.setTmp(&NotEqual{Left: p.Stderr.Text, Right: &StringValue{Value: ""}})
			withText := inner.subBlock()
			sep := withText.setTmp(&Add{Left: message, Right: &StringValue{Value: ": "}})
			withText.set(message, &Add{Left: sep, Right: p.Stderr.Text})
			inner.add(&If{Condition: hasText, Block: withText})
		}

		inner.set(failed, &BoolValue{Value: true})
	}

	return failed, message
}

// sourceExits lists the exit modes of x and its sources in the order of the procs
func sourceExits(x *ast.Cmd, exits *[]ast.ExitMode) {
	for _, arg := range x.Args {
		inner, isCmd := arg.Value.(*ast.Cmd)
		if isCmd && inner.Use == ast.UseSource {
			sourceExits(inner, exits)
		}
	}

	if x.Stdin != nil && x.Stdin.Kind == ast.InputCmd {
		sourceExits(x.Stdin.Value.(*ast.Cmd), exits)
	}

	*exits = append(*exits, x.Exit)
}

// readLines adds the loop of a for over a pipe to block
func (c *Compiler) readLines(block *Block, pipe *VarValue, st *ast.ForIn) {
	body := block.subBlock()
	line := block.nextTmp(stringType)
	block.registerVar(line)

	c.setLoopVar(body, st.Value, line)
	if st.Index != nil {
		counter := block.setTmp(&NumberValue{Value: 0})
		c.setLoopVar(body, st.Index, counter)
		body.set(counter, &Add{Left: counter, Right: &NumberValue{Value: 1}})
	}

	c.steps = append(c.steps, nil)
	c.evalBlock(body, st.Body)
	c.steps = c.steps[:len(c.steps)-1]

	block.add(&ReadLines{Pipe: pipe, Line: line, Block: body})
}
