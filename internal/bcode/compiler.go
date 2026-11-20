package bcode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type Compiler struct {
	prog *Program
	ast  *ast.Program
	// namespaces keep the function names of the packages apart. Builtins have no package,
	// so they get the empty namespace.
	namespaces map[*ast.Package]string
	// steps holds the step of each loop around the current statement, nil for loops
	// without one
	steps []*ast.Assignment
	// elements are the evaluated element targets of the current assignment
	elements map[*ast.AssignmentTarget]*elementTarget
}

type elementTarget struct {
	array *VarValue
	index Atom
}

func Compile(program *ast.Program) (*Program, error) {
	a := &Compiler{
		prog:       &Program{},
		ast:        program,
		namespaces: map[*ast.Package]string{program.Main: "main"},
		elements:   map[*ast.AssignmentTarget]*elementTarget{},
	}

	// the number makes them unique, the name is only for reading the generated code
	for i, pkg := range program.Packages() {
		a.namespaces[pkg] = "p" + strconv.Itoa(i+1) + "_" + identifier(pkg.Name)
	}

	a.evalProgram()

	return a.prog, nil
}

func (c *Compiler) evalProgram() {
	c.prog.Main = &Block{
		tmpVar: c.prog,
		vars:   map[string]*VarValue{},
	}

	for _, def := range c.ast.InitOrder {
		c.evalAssignment(c.prog.Main, def.Assignment)
	}

	c.evalBlock(c.prog.Main, c.ast.Root)

	c.addFuncs(c.namespaces[c.ast.Main], c.ast.Main.FuncDefs)
	for _, pkg := range c.ast.Packages() {
		c.addFuncs(c.namespaces[pkg], pkg.FuncDefs)
		c.addExternalFuncs(c.namespaces[pkg], pkg.ExternalFuncDefs)
	}
}

// identifier replaces everything that can't be part of a bash function name.
func identifier(name string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}
		return '_'
	}, name)
}

func (c *Compiler) addFuncs(namespace string, funcDefs []*ast.FuncDef) {
	for _, fd := range funcDefs {
		block := &Block{
			tmpVar: c.prog,
			vars:   map[string]*VarValue{},
		}

		params := []*Param{}
		paramVars := []*VarValue{}

		for _, p := range fd.Params {
			params = append(params, &Param{Name: p.Name, Type: p.Type})
			paramVars = append(paramVars, block.paramVar(p.Name, p.Type))
		}

		if len(paramVars) > 0 {
			block.add(&Params{Vars: paramVars})
		}

		c.evalBlock(block, fd.Body)

		funcDef := &FuncDef{
			Namespace: namespace,
			Name:      fd.Name,
			Body:      block,
			Params:    params,
			CanFail:   fd.CanFail,
		}

		c.prog.Funcs = append(c.prog.Funcs, funcDef)
	}
}

func (c *Compiler) addExternalFuncs(namespace string, funcDefs []*ast.ExternalFuncDef) {
	for _, fd := range funcDefs {
		block := &Block{
			tmpVar: c.prog,
			vars:   map[string]*VarValue{},
		}

		params := []*Param{}
		paramVars := []*VarValue{}

		for _, p := range fd.Params {
			params = append(params, &Param{Name: p.Name, Type: p.Type})
			paramVars = append(paramVars, block.externalParamVar(p.Name, p.Type))
		}

		if len(paramVars) > 0 {
			block.add(&Params{Vars: paramVars})
		}

		block.add(&Code{Code: fd.Code})

		funcDef := &FuncDef{
			Namespace: namespace,
			Name:      fd.Name,
			Body:      block,
			Params:    params,
			CanFail:   fd.CanFail,
		}

		c.prog.Funcs = append(c.prog.Funcs, funcDef)
	}
}

func (c *Compiler) evalBlock(b *Block, block *ast.Block) {
	for _, s := range block.Statements {
		switch st := s.(type) {
		case *ast.Assignment:
			c.evalAssignment(b, st)

		case *ast.Block:
			c.evalBlock(b, st)

		case *ast.FuncStatement:
			c.evalFuncStatement(b, st)

		case *ast.If:
			c.evalIfStatement(b, st)

		case *ast.For:
			c.evalForStatement(b, st)

		case *ast.ForIn:
			c.evalForInStatement(b, st)

		case *ast.ReturnStatement:
			c.evalReturnStatement(b, st)

		case *ast.Fail:
			b.add(&Fail{Message: c.evalExpression(b, st.Message)[0]})

		case *ast.Break:
			b.add(&Break{})

		case *ast.Continue:
			// like in Go, continue runs the step before the next check
			step := c.steps[len(c.steps)-1]
			if step != nil {
				c.evalAssignment(b, step)
			}
			b.add(&Continue{})

		default:
			panic(fmt.Sprintf("no valid statement in basm transpiler %T", st))
		}
	}
}

func (c *Compiler) evalExpression(block *Block, exp ast.Expression) []Value {
	switch x := exp.(type) {
	case *ast.Number:
		return []Value{&NumberValue{Value: int(x.Value)}}

	case *ast.String:
		return []Value{&StringValue{Value: x.Value}}

	case *ast.Bool:
		return []Value{&BoolValue{Value: x.Value}}

	case *ast.Var:
		if x.GlobalOf != nil {
			global := &VarValue{Name: c.globalName(x.GlobalOf, x.Name), T: x.Type}

			// a call later in the expression can write the global, so a + f() has to read
			// a first. Copying a whole array for every read would be too slow.
			if x.Type.IsArray() {
				return []Value{global}
			}

			return []Value{block.setTmp(global)}
		}

		current := block.getVisibleVar(varName(x.Name))
		if current == nil {
			panic(fmt.Sprintf("var not defined %v", x.Name))
		}

		return []Value{current}

	case *ast.EnvVar:
		return []Value{&EnvVarValue{Name: x.Name}}

	case *ast.TargetValue:
		el := c.elements[x.Target]
		return []Value{block.setTmp(&Index{Value: el.array, Position: el.index, Location: x.Target.Element.Location})}

	case *ast.Add:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&Add{Left: a[0], Right: b[0]})

	case *ast.Sub:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&Sub{Left: a[0], Right: b[0]})

	case *ast.Minus:
		a := c.evalExpression(block, x.Expression)

		return block.setTmps(&Minus{Expression: a[0]})

	case *ast.Mul:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&Mul{Left: a[0], Right: b[0]})

	case *ast.Div:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&Div{Left: a[0], Right: b[0]})

	case *ast.Mod:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&Mod{Left: a[0], Right: b[0]})

	case *ast.And:
		return c.evalLogical(block, x.A, x.B, true)

	case *ast.Or:
		return c.evalLogical(block, x.A, x.B, false)

	case *ast.Not:
		a := c.evalExpression(block, x.Expression)

		return block.setTmps(&Not{Expression: a[0]})

	case *ast.Equal:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&Equal{Left: a[0], Right: b[0]})

	case *ast.LessThan:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&LessThan{Left: a[0], Right: b[0]})

	case *ast.LessThanEqual:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&LessThanEqual{Left: a[0], Right: b[0]})

	case *ast.GreaterThan:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&GreaterThan{Left: a[0], Right: b[0]})

	case *ast.GreaterThanEqual:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&GreaterThanEqual{Left: a[0], Right: b[0]})

	case *ast.NotEqual:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&NotEqual{Left: a[0], Right: b[0]})

	case *ast.Group:
		return c.evalExpression(block, x.Expression)

	case *ast.Slice:
		return []Value{c.evalSlice(block, x)}

	case *ast.Index:
		return []Value{c.evalIndex(block, x)}

	case *ast.Ternary:
		trueBlock := block.subBlock()
		trueValues := c.evalExpression(trueBlock, x.True)

		falseBlock := block.subBlock()
		falseValues := c.evalExpression(falseBlock, x.False)

		// the type check guarantees that both branches have the same number of values
		vars := []Value{}
		for i := range trueValues {
			v := block.nextTmp(trueValues[i].Type())
			trueBlock.add(&Assignment{Var: v, Value: trueValues[i]})
			falseBlock.add(&Assignment{Var: v, Value: falseValues[i]})
			vars = append(vars, v)
		}

		cond := block.setTmp(c.evalExpression(block, x.Condition)[0])
		notCond := block.setTmp(&Not{Expression: cond})

		block.add(&If{Condition: cond, Block: trueBlock})
		block.add(&If{Condition: notCond, Block: falseBlock})

		return vars

	case *ast.Func:
		if x.Package != nil {
			return c.evalCall(block, x)
		}

		switch x.Name {
		case "args":
			return []Value{block.setTmp(&ArgsValue{})}

		case "len":
			val := c.evalExpression(block, x.Parameters[0])
			return block.setTmps(&Len{
				Parameter: block.asVar(val[0]),
			})

		case "append":
			vals := c.evalExpressions(block, x.Parameters)

			ap := block.nextTmp(vals[0].Type())

			block.add(&Append{
				Array:    vals[0],
				Elements: vals[1:],
				Target:   ap,
			})

			return []Value{ap}
		}

		return c.evalCall(block, x)

	case *ast.Method:
		return c.evalCall(block, x.Func)

	case *ast.Cmd:
		return c.evalCmd(block, x, nil)

	case *ast.Array:
		a := &ArrayValue{
			Values: []Atom{},
			Tp:     x.Type,
		}
		for _, p := range x.Values {
			a.Values = append(a.Values, common.Map(c.evalExpression(block, p), func(x Value) Atom { return x })...)
		}

		return []Value{block.setTmp(a)}

	default:
		panic(fmt.Sprintf("no valid ast expression %T", x))
	}
}

func (c *Compiler) evalExpressions(block *Block, exps []ast.Expression) []Value {
	vals := []Value{}

	for _, exp := range exps {
		vals = append(vals, c.evalExpression(block, exp)...)
	}

	return vals
}

func (c *Compiler) evalAssignment(b *Block, ass *ast.Assignment) {
	targets := map[string]bool{}
	for _, target := range ass.Targets {
		targets[c.targetName(target)] = true
	}

	c.evalElements(b, ass.Targets, targets)

	if ass.Catch != nil {
		assign := func(then *Block, results []Value) {
			c.assignTargets(b, then, ass.Targets, results)
		}

		if x, isCmd := ass.Expressions[0].(*ast.Cmd); isCmd {
			c.catchCmd(b, x, ass.Catch, assign)
		} else {
			c.evalCatch(b, ass.Expressions[0], ass.Catch, assign)
		}

		return
	}

	// every assignment and call copies an array, so nothing else can see this one grow
	if call := selfAppend(ass); call != nil {
		array := c.evalExpression(b, call.Parameters[0])[0].(*VarValue)
		b.add(&Append{Array: array, Elements: c.evalExpressions(b, call.Parameters[1:]), Target: array})

		return
	}

	results := []Value{}
	for _, exp := range ass.Expressions {
		for _, val := range c.evalExpression(b, exp) {
			// without the copy, a, b = b, a would read a after assigning b to it
			if v, isVar := val.(*VarValue); isVar && targets[v.Name] {
				val = b.setTmp(v)
			}
			results = append(results, val)
		}
	}

	c.assignTargets(b, b, ass.Targets, results)
}

// assignTargets adds the assignments to b, the new variables are visible in scope
func (c *Compiler) assignTargets(scope, b *Block, targets []*ast.AssignmentTarget, results []Value) {
	for i, target := range targets {
		if target.Name == "_" {
			continue
		}

		if target.Element != nil {
			el := c.elements[target]
			b.add(&SetElement{Array: el.array, Index: el.index, Value: results[i], Location: target.Element.Location})
			continue
		}

		v := &VarValue{Name: c.targetName(target), T: results[i].Type()}
		if target.GlobalOf != nil {
			b.add(&Assignment{Var: v, Value: results[i], Global: true})
			continue
		}

		if scope.getVisibleVar(v.Name) == nil {
			scope.registerVar(v)
		}
		b.add(&Assignment{Var: v, Value: results[i]})
	}
}

// evalElements evaluates the arrays and indexes of the element targets. Like in Go, that
// comes before the values, so i, a[i] = 1, 2 assigns the element at the old i.
func (c *Compiler) evalElements(b *Block, targets []*ast.AssignmentTarget, names map[string]bool) {
	for _, target := range targets {
		if target.Element == nil {
			continue
		}

		array := c.evalExpression(b, target.Element.Value)[0].(*VarValue)
		index := c.evalExpression(b, target.Element.Position)[0]
		if v, isVar := index.(*VarValue); isVar && names[v.Name] {
			index = b.setTmp(v)
		}

		c.elements[target] = &elementTarget{array: array, index: index}
	}
}

func selfAppend(ass *ast.Assignment) *ast.Func {
	// in a, b = append(a, "x"), a the second value has to be the old a
	if len(ass.Targets) != 1 || len(ass.Expressions) != 1 {
		return nil
	}

	// without elements there would be no bash line, and bash rejects an empty if body
	call, isCall := ass.Expressions[0].(*ast.Func)
	if !isCall || call.Package != nil || call.Name != "append" || len(call.Parameters) < 2 {
		return nil
	}

	array, isVar := call.Parameters[0].(*ast.Var)
	target := ass.Targets[0]
	if !isVar || target.Element != nil || array.Name != target.Name || array.GlobalOf != target.GlobalOf {
		return nil
	}

	return call
}

func (c *Compiler) targetName(target *ast.AssignmentTarget) string {
	if target.GlobalOf != nil {
		return c.globalName(target.GlobalOf, target.Name)
	}

	return varName(target.Name)
}

func (c *Compiler) globalName(pkg *ast.Package, name string) string {
	return "G_" + c.namespaces[pkg] + "_" + name
}

// evalLogical evaluates b only when a doesn't decide the result already,
// so && and || short-circuit like in Go.
func (c *Compiler) evalLogical(block *Block, a, b ast.Expression, isAnd bool) []Value {
	result := block.setTmp(c.evalExpression(block, a)[0])

	cond := result
	if !isAnd {
		cond = block.setTmp(&Not{Expression: result})
	}

	right := block.subBlock()
	right.set(result, c.evalExpression(right, b)[0])

	block.add(&If{Condition: cond, Block: right})

	return []Value{result}
}

func (c *Compiler) evalSlice(block *Block, s *ast.Slice) Atom {
	value := c.evalExpression(block, s.Value)

	// we want a varValue for slice operation
	varValue, isVar := value[0].(*VarValue)
	if !isVar {
		varValue = block.setTmp(value[0])
	}

	var from Value
	if s.From != nil {
		from = c.evalExpression(block, s.From)[0]
	} else {
		from = &NumberValue{Value: 0}
	}

	var to Value
	if s.To != nil {
		to = c.evalExpression(block, s.To)[0]
	} else {
		to = block.setTmp(&Len{
			Parameter: block.asVar(value[0]),
		})
	}

	return block.setTmp(&Slice{
		Value:    varValue,
		From:     from,
		To:       to,
		Location: s.Location,
	})
}

func (c *Compiler) evalIndex(block *Block, s *ast.Index) Atom {
	value := block.asVar(c.evalExpression(block, s.Value)[0])
	position := c.evalExpression(block, s.Position)[0]

	return block.setTmp(&Index{
		Value:    value,
		Position: position,
		Location: s.Location,
	})
}

func (c *Compiler) evalFuncStatement(block *Block, f *ast.FuncStatement) {
	if x, isCmd := f.Call.(*ast.Cmd); isCmd && f.Catch != nil {
		c.catchCmd(block, x, f.Catch, nil)
		return
	}

	if f.Catch != nil {
		c.evalCatch(block, f.Call, f.Catch, nil)
		return
	}

	c.evalExpression(block, f.Call)
}

// evalCatch calls the function of call, assign gets its results. When the call fails, the
// block of catch runs instead.
func (c *Compiler) evalCatch(block *Block, call ast.Expression, catch *ast.Catch, assign func(then *Block, results []Value)) {
	x, isFunc := call.(*ast.Func)
	if !isFunc {
		x = call.(*ast.Method).Func
	}

	st := &Catch{
		Call:  c.newCall(block, x),
		Block: block.subBlock(),
		Then:  block.subBlock(),
	}

	if catch.Error.Name != "_" {
		st.Block.set(&VarValue{Name: varName(catch.Error.Name), T: &common.BasicType{Type: common.String}}, &ErrorValue{})
	}

	c.evalBlock(st.Block, catch.Block)

	if assign != nil {
		assign(st.Then, common.Map(st.Call.Return, func(v *VarValue) Value { return v }))
	}

	block.add(st)
}

func (c *Compiler) evalCall(block *Block, x *ast.Func) []Value {
	f := c.newCall(block, x)
	block.add(f)

	return common.Map(f.Return, func(i *VarValue) Value { return i })
}

// newCall evaluates the arguments, the call isn't added to the block
func (c *Compiler) newCall(block *Block, x *ast.Func) *Func {
	ret := common.Map(x.ReturnTypes, func(t common.Type) *VarValue {
		return block.nextTmp(t)
	})

	f := &Func{
		Namespace: c.namespaces[x.Package],
		Name:      x.Name,
		Return:    ret,
		Try:       x.Try,
	}
	for _, p := range x.Parameters {
		f.Parameters = append(f.Parameters, c.evalExpression(block, p)...)
	}

	return f
}

func (c *Compiler) evalIfStatement(block *Block, ifst *ast.If) {
	condVar := block.setTmp(c.evalExpression(block, ifst.Condition)[0])

	ifs := &If{
		Condition: condVar,
		Block:     block.subBlock(),
	}

	c.evalBlock(ifs.Block, ifst.If)

	if ifst.Else != nil {
		ifs.Else = block.subBlock()
		c.evalBlock(ifs.Else, ifst.Else)
	}

	block.add(ifs)
}

func (c *Compiler) evalForStatement(block *Block, forst *ast.For) {
	if forst.Initial != nil {
		c.evalAssignment(block, forst.Initial)
	}

	while := &While{Block: block.subBlock()}

	if forst.Condition != nil {
		while.Check = block.subBlock()
		while.Condition = while.Check.setTmp(c.evalExpression(while.Check, forst.Condition)[0])
	}

	c.steps = append(c.steps, forst.Each)
	c.evalBlock(while.Block, forst.Body)
	c.steps = c.steps[:len(c.steps)-1]

	if forst.Each != nil {
		c.evalAssignment(while.Block, forst.Each)
	}

	block.add(while)
}

// evalForInStatement sets the loop variables and moves the counter on at the start of the
// body, so a continue needs no step and changing them doesn't change the loop
func (c *Compiler) evalForInStatement(block *Block, st *ast.ForIn) {
	if x, isCmd := st.Over.(*ast.Cmd); isCmd && x.Use == ast.UseLoop {
		c.evalCmd(block, x, st)
		return
	}

	if st.Over.Types()[0].Is(&common.PipeType{}) {
		c.readLines(block, c.evalExpression(block, st.Over)[0].(*VarValue), st)
		return
	}

	body := block.subBlock()
	var loop Command

	over := st.Over.Types()[0]
	if st.From == nil && (over.IsArray() || over.Is(&common.BasicType{Type: common.String})) {
		array := block.asVar(c.evalExpression(block, st.Over)[0])
		if !over.IsArray() {
			array = block.setTmp(&Chars{Value: array})
		}
		element := block.nextTmp(array.T.(*common.ArrayType).ElementType)
		block.registerVar(element)

		c.setLoopVar(body, st.Value, element)
		if st.Index != nil {
			counter := block.setTmp(&NumberValue{Value: 0})
			c.setLoopVar(body, st.Index, counter)
			body.set(counter, &Add{Left: counter, Right: &NumberValue{Value: 1}})
		}

		loop = &ForEach{Element: element, Array: array, Block: body}
	} else {
		var from Atom = &NumberValue{Value: 0}
		if st.From != nil {
			from = c.evalExpression(block, st.From)[0]
		}
		counter := block.setTmp(from)
		end := block.setTmp(c.evalExpression(block, st.Over)[0])

		check := block.subBlock()
		cond := check.setTmp(&LessThan{Left: counter, Right: end})

		c.setLoopVar(body, st.Value, counter)
		body.set(counter, &Add{Left: counter, Right: &NumberValue{Value: 1}})

		loop = &While{Check: check, Condition: cond, Block: body}
	}

	c.steps = append(c.steps, nil)
	c.evalBlock(body, st.Body)
	c.steps = c.steps[:len(c.steps)-1]

	block.add(loop)
}

func (c *Compiler) setLoopVar(body *Block, target *ast.AssignmentTarget, value *VarValue) {
	if target.Name != "_" {
		body.set(&VarValue{Name: varName(target.Name), T: value.T}, value)
	}
}

func (c *Compiler) evalReturnStatement(block *Block, ret *ast.ReturnStatement) {
	values := []*VarValue{}

	for _, x := range ret.Expressions {
		for _, val := range c.evalExpression(block, x) {
			values = append(values, block.setTmp(val))
		}
	}

	block.add(&Return{
		Values: values,
	})
}
