package parser

import (
	"fmt"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

// cmdBody sets the block after an assignment as the body of its cmd
func (v *Visitor) cmdBody(assignment *ast.Assignment, ctx IBlockContext) {
	if ctx == nil {
		return
	}

	body := ctx.Accept(v).(*ast.Block)

	c, isCmd := assignment.Expressions[0].(*ast.Cmd)
	if len(assignment.Expressions) != 1 || !isCmd {
		v.errorf(v.loc(ctx), "only a cmd has a block after it, as the only value of the assignment")
		return
	}

	c.Body = body
}

func (v *Visitor) VisitCmdStatement(ctx *CmdStatementContext) any {
	c := ctx.Cmd().Accept(v).(*ast.Cmd)
	if ctx.Block() != nil {
		c.Body = ctx.Block().Accept(v).(*ast.Block)
	}

	return &ast.FuncStatement{Call: c}
}

// resultOrder is the position of the options that are results, they must come in this order
var resultOrder = map[string]int{"stdout": 0, "stderr": 1, "exit": 2}

func (v *Visitor) VisitCmd(ctx *CmdContext) any {
	// like the one of a call, the location starts after the try
	loc := v.loc(ctx)
	start := v.locTk(ctx.CMD().GetSymbol())
	loc.LineFrom, loc.ColFrom = start.LineFrom, start.ColFrom

	c := &ast.Cmd{Try: ctx.TRY() != nil, Location: loc}

	seen := map[string]bool{}
	lastResult := ""

	for _, a := range ctx.AllCmdArg() {
		arg := a.(*CmdArgContext)

		if arg.GetName() == nil && arg.GetEnv() == nil {
			if len(seen) > 0 {
				v.errorf(v.loc(arg), "the arguments of cmd come before its options")
			}

			c.Args = append(c.Args, &ast.CmdArg{
				Value:    arg.Expression().Accept(v).(ast.Expression),
				Spread:   arg.GetSpread() != nil,
				Location: v.loc(arg),
			})

			continue
		}

		name := ""
		if arg.GetEnv() != nil {
			name = arg.GetEnv().GetText()
		} else {
			name = arg.GetName().GetText()
		}

		loc := v.loc(arg)
		if seen[name] {
			v.errorf(loc, "option %v: is given twice", name)
		}
		seen[name] = true

		pos, isResult := resultOrder[name]
		if isResult {
			if lastResult != "" && resultOrder[lastResult] > pos {
				v.errorf(loc, "%v: must come before %v:, the results are in the order stdout, stderr, exit", name, lastResult)
			}
			lastResult = name
		}

		v.cmdOption(c, name, arg.CmdValue().(*CmdValueContext), loc)
	}

	if len(c.Args) == 0 {
		v.errorf(c.Location, "cmd needs the name of the command")
	}

	// a cmd with a pipe as an argument or stdin is read by the outer one, with another
	// stdout it's a value like any other
	for _, arg := range c.Args {
		inner, isCmd := arg.Value.(*ast.Cmd)
		if isCmd && isPipe(inner.Stdout) {
			inner.Use = ast.UseSource
		}
	}
	if c.Stdin != nil && c.Stdin.Kind == ast.InputCmd {
		c.Stdin.Value.(*ast.Cmd).Use = ast.UseSource
	}

	return c
}

func isPipe(t *ast.Target) bool {
	return t != nil && t.Kind == ast.PipeTarget
}

// cmdValue is a type like string, or an expression
type cmdValue struct {
	primitive common.PrimitiveType
	exp       ast.Expression
}

func (v *Visitor) cmdValue(ctx *CmdValueContext) cmdValue {
	if ctx.PrimitiveType() != nil {
		return cmdValue{primitive: ctx.PrimitiveType().Accept(v).(common.PrimitiveType)}
	}

	return cmdValue{exp: ctx.Expression().Accept(v).(ast.Expression)}
}

func (v *Visitor) cmdOption(c *ast.Cmd, name string, ctx *CmdValueContext, loc *common.Location) {
	val := v.cmdValue(ctx)

	switch name {
	case "stdout":
		c.Stdout = v.cmdTarget(name, val, loc)

	case "stderr":
		c.Stderr = v.cmdTarget(name, val, loc)

	case "exit":
		switch {
		case val.primitive == common.Int:
			c.Exit = ast.ExitInt
		case val.primitive == common.Bool:
			c.Exit = ast.ExitBool
		case isName(val.exp, "_"):
			c.Exit = ast.ExitIgnore
		default:
			v.errorf(loc, "exit: takes int, bool or _")
		}

	case "stdin":
		c.Stdin = v.cmdInput(val, loc)

	case "dir":
		if val.exp == nil {
			v.errorf(loc, "dir: takes the path of a folder, not a type")
			return
		}
		c.Dir = val.exp

	default:
		if name[0] != '$' {
			v.errorf(loc, "cmd has no option %v:, it has stdout:, stderr:, exit:, stdin:, dir: and $NAME: for env vars", name)
			return
		}

		if val.exp == nil {
			v.errorf(loc, "%v: takes the value of the env var, not a type", name)
			return
		}
		c.Env = append(c.Env, &ast.CmdEnv{Name: name[1:], Value: val.exp, Location: loc})
	}
}

func (v *Visitor) cmdTarget(name string, val cmdValue, loc *common.Location) *ast.Target {
	t := &ast.Target{Location: loc}
	array, isArray := val.exp.(*ast.Array)

	switch {
	case val.primitive == common.String:
		t.Kind = ast.CollectString

	case isArray && len(array.Values) == 0 && array.Type == common.String:
		t.Kind = ast.CollectLines

	case isName(val.exp, "pipe"):
		t.Kind = ast.PipeTarget

	case isName(val.exp, "_"):
		t.Kind = ast.DiscardTarget

	case name == "stderr" && isName(val.exp, "stdout"):
		t.Kind = ast.StdoutTarget

	case v.isFile(val.exp, "file", loc):
		t.Kind = ast.FileTarget
		t.Path = val.exp.(*ast.Func).Parameters[0]

	case v.isFile(val.exp, "appendFile", loc):
		t.Kind = ast.AppendTarget
		t.Path = val.exp.(*ast.Func).Parameters[0]

	default:
		also := ""
		if name == "stderr" {
			also = ", stdout"
		}
		v.errorf(loc, "%v: takes string, []string, pipe, file(path), appendFile(path)%v or _", name, also)
		return nil
	}

	return t
}

func (v *Visitor) cmdInput(val cmdValue, loc *common.Location) *ast.Input {
	if val.exp == nil {
		v.errorf(loc, "stdin: takes a string, file(path) or a cmd with stdout: pipe, not a type")
		return nil
	}

	if v.isFile(val.exp, "file", loc) {
		return &ast.Input{Kind: ast.InputFile, Value: val.exp.(*ast.Func).Parameters[0]}
	}

	if c, isCmd := val.exp.(*ast.Cmd); isCmd && isPipe(c.Stdout) {
		return &ast.Input{Kind: ast.InputCmd, Value: c}
	}

	return &ast.Input{Kind: ast.InputString, Value: val.exp}
}

// isFile reports whether x is file(path) or appendFile(path), as name. A wrong number of
// arguments is reported here.
func (v *Visitor) isFile(x ast.Expression, name string, loc *common.Location) bool {
	f, isFunc := x.(*ast.Func)
	if !isFunc || f.Name != name || f.Try {
		return false
	}

	if len(f.Parameters) != 1 {
		v.errorf(loc, "%v takes 1 argument, the path, not %v", name, len(f.Parameters))
		f.Parameters = []ast.Expression{&ast.String{}}
	}

	return true
}

func isName(x ast.Expression, name string) bool {
	vr, isVar := x.(*ast.Var)
	return isVar && vr.Namespace == "" && vr.Name == name
}

func (v *Visitor) errorf(loc *common.Location, format string, a ...any) {
	v.errors = append(v.errors, &ast.ASTError{Location: loc, Message: fmt.Sprintf(format, a...)})
}
