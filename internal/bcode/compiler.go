package bcode

import (
	"fmt"
	"strconv"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type Compiler struct {
	prog *Program
	ast  *ast.Program
}

func Compile(program *ast.Program) (*Program, error) {
	a := &Compiler{
		prog: &Program{
			FuncsByName: map[string]*FuncDef{},
		},
		ast: program,
	}

	a.evalProgram()

	return a.prog, nil
}

func (c *Compiler) evalProgram() {
	c.prog.Main = &Block{
		tmpVar: c.prog,
		vars:   map[string]*VarValue{},
	}
	c.evalBlock(c.prog.Main, c.ast.Root)

	c.addFuncs("", c.ast.Root.FuncDefs)
	for _, lib := range c.ast.Libraries() {
		c.addFuncs(lib.Path, lib.FuncDefs)
		c.addExternalFuncs(lib.Path, lib.ExternalFuncDefs)
	}
}

func (c *Compiler) addFuncs(namespace string, funcDefs []*ast.FuncDef) {
	for _, fd := range funcDefs {
		block := &Block{
			parent: c.prog.Main,
			tmpVar: c.prog,
			vars:   map[string]*VarValue{},
		}

		params := []*Param{}

		paramIndex := 1
		for _, p := range fd.Params {
			params = append(params, &Param{Name: p.Name, Type: p.Type})
			block.set(&VarValue{Name: p.Name, T: p.Type}, &EnvVarValue{Name: strconv.Itoa(paramIndex)})

			paramIndex++
		}

		c.evalBlock(block, fd.Body)

		funcDef := &FuncDef{
			Namespace: namespace,
			Name:      fd.Name,
			Body:      block,
			Params:    params,
		}

		c.prog.Funcs = append(c.prog.Funcs, funcDef)
		c.prog.FuncsByName[funcDef.FullName()] = funcDef
	}
}

func (c *Compiler) addExternalFuncs(namespace string, funcDefs []*ast.ExternalFuncDef) {
	for _, fd := range funcDefs {
		block := &Block{
			parent: c.prog.Main,
			tmpVar: c.prog,
			vars:   map[string]*VarValue{},
		}

		params := []*Param{}

		paramIndex := 1
		for _, p := range fd.Params {
			params = append(params, &Param{Name: p.Name, Type: p.Type})
			block.set(&VarValue{Name: p.Name, T: p.Type}, &EnvVarValue{Name: strconv.Itoa(paramIndex)})

			paramIndex++
		}

		block.add(&Code{Code: fd.Code})

		funcDef := &FuncDef{
			Namespace: namespace,
			Name:      fd.Name,
			Body:      block,
			Params:    params,
		}

		c.prog.Funcs = append(c.prog.Funcs, funcDef)
		c.prog.FuncsByName[funcDef.FullName()] = funcDef
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

		case *ast.ReturnStatement:
			c.evalReturnStatement(b, st)

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
		current := block.getVisibleVar(x.Name)
		if current == nil {
			panic(fmt.Sprintf("var not defined %v", x.Name))
		}

		return []Value{current}

	case *ast.EnvVar:
		return []Value{&EnvVarValue{Name: x.Name}}

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
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmps(&And{Left: a[0], Right: b[0]})

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
		trueValue := c.evalExpression(trueBlock, x.True)

		falseBlock := block.subBlock()
		falseValue := c.evalExpression(falseBlock, x.False)

		v := c.prog.Main.nextTmp(trueValue[0].Type())

		cond := block.setTmp(c.evalExpression(block, x.Condition)[0])
		notCond := block.setTmp(&Not{Expression: cond})

		ifTrue := &If{
			Condition: cond,
			Block:     trueBlock,
		}

		ifTrue.Block.add(&Assignment{
			Var:   v,
			Value: trueValue[0],
		})

		block.add(ifTrue)

		ifFalse := &If{
			Condition: notCond,
			Block:     falseBlock,
		}

		ifFalse.Block.add(&Assignment{
			Var:   v,
			Value: falseValue[0],
		})

		block.add(ifFalse)

		return []Value{v}

	case *ast.Func:
		switch x.Name {
		case "len":
			val := c.evalExpression(block, x.Parameters[0])
			return []Value{&Len{
				Parameter: block.asVar(val[0]),
			}}

		case "execVar":
			vals := c.evalExpressions(block, x.Parameters)

			stdout := block.nextTmp(&common.ArrayType{ElementType: &common.BasicType{Type: common.String}})
			stderr := block.nextTmp(&common.ArrayType{ElementType: &common.BasicType{Type: common.String}})
			xerr := block.nextTmp(&common.BasicType{Type: common.Int})

			block.add(&ExecVar{
				Command:    vals[0],
				Parameters: vals[1:],
				Stdout:     stdout,
				Stderr:     stderr,
				Err:        xerr,
			})

			return []Value{stdout, stderr, xerr}

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

		fSig := getSignature(c.ast, "", x.Name)

		ret := common.Map(fSig.Return, func(t common.Type) *VarValue {
			return block.nextTmp(t)
		})

		f := &Func{
			Name:   x.Name,
			Return: ret,
		}
		for _, p := range x.Parameters {
			f.Parameters = append(f.Parameters, c.evalExpression(block, p)...)
		}

		block.add(f)

		return common.Map(f.Return, func(i *VarValue) Value { return i })

	case *ast.Method:
		vr, is := x.Expression.(*ast.Var)
		if is {
			fSig := getSignature(c.ast, vr.Name, x.Func.Name)

			ret := common.Map(fSig.Return, func(t common.Type) *VarValue {
				return block.nextTmp(t)
			})
			f := &Func{
				Namespace: fSig.Namespace,
				Name:      x.Func.Name,
				Return:    ret,
			}
			for _, p := range x.Func.Parameters {
				f.Parameters = append(f.Parameters, c.evalExpression(block, p)...)
			}

			block.add(f)
			return common.Map(f.Return, func(i *VarValue) Value { return i })
		}

		panic(fmt.Sprint("unknown package name", vr.Name))

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
	results := []Value{}
	for _, exp := range ass.Expressions {
		results = append(results, c.evalExpression(b, exp)...)
	}

	for i, name := range ass.Names {
		b.set(&VarValue{Name: name, T: results[i].Type()}, results[i])
	}
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
		Value: varValue,
		From:  from,
		To:    to,
	})
}

func (c *Compiler) evalIndex(block *Block, s *ast.Index) Atom {
	return block.setTmp(&Index{
		Value:    c.evalExpression(block, s.Value)[0].(*VarValue),
		Position: c.evalExpression(block, s.Position)[0],
	})
}

func (c *Compiler) evalFuncStatement(block *Block, f *ast.FuncStatement) {
	fnc := &Func{
		Name: f.Func.Name,
	}
	for _, p := range f.Func.Parameters {
		fnc.Parameters = append(fnc.Parameters, c.evalExpression(block, p)[0])
	}

	block.add(fnc)
}

func (c *Compiler) evalIfStatement(block *Block, ifst *ast.If) {
	condVar := block.setTmp(c.evalExpression(block, ifst.Condition)[0])

	ifs := &If{
		Condition: condVar,
		Block:     block.subBlock(),
	}

	c.evalBlock(ifs.Block, ifst.If)

	block.add(ifs)
}

func (c *Compiler) evalForStatement(block *Block, forst *ast.For) {
	c.evalAssignment(block, forst.Initial)

	condVar := block.setTmp(c.evalExpression(block, forst.Condition)[0])

	while := &While{
		Condition: condVar,
		Block:     block.subBlock(),
	}

	c.evalBlock(while.Block, forst.Body)

	c.evalAssignment(while.Block, forst.Each)

	while.Block.set(condVar, c.evalExpression(while.Block, forst.Condition)[0])

	block.add(while)
}

func (c *Compiler) evalReturnStatement(block *Block, ret *ast.ReturnStatement) {
	values := []*VarValue{}

	for _, x := range ret.Expressions {
		values = append(values, block.setTmp(c.evalExpression(block, x)[0]))
	}

	block.add(&Return{
		Values: values,
	})
}
