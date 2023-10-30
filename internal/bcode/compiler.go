package bcode

import (
	"fmt"

	"github.com/byte-wright/lush/internal/ast"
)

type Compiler struct {
	prog *Program
	ast  *ast.Program
}

func Compile(program *ast.Program) (*Program, error) {
	a := &Compiler{
		prog: &Program{},
		ast:  program,
	}

	a.evalProgram()

	return a.prog, nil
}

func (c *Compiler) evalProgram() {
	c.prog.Main = &Block{tmpVar: c.prog}
	c.evalBlock(c.prog.Main, c.ast.Root)
}

func (c *Compiler) evalBlock(b *Block, block *ast.Block) {
	for _, s := range block.Statements {
		switch st := s.(type) {
		case *ast.Assignment:
			result := c.evalExpression(b, st.Expression)

			b.set(st.Name, result)

		case *ast.Block:
			c.evalBlock(b, st)

		case *ast.FuncStatement:
			c.evalFuncStatement(b, st)

		case *ast.If:
			c.evalIfStatement(b, st)

		case *ast.For:
			c.evalForStatement(b, st)

		default:
			panic(fmt.Sprintf("no valid statement %T", st))
		}
	}
}

func (c *Compiler) evalExpression(block *Block, exp ast.Expression) Value {
	switch x := exp.(type) {
	case *ast.Number:
		return &NumberValue{Value: int(x.Value)}

	case *ast.String:
		return &StringValue{Value: x.Value}

	case *ast.Bool:
		return &BoolValue{Value: x.Value}

	case *ast.Var:
		return &VarValue{Name: x.Name}

	case *ast.EnvVar:
		return &EnvVarValue{Name: x.Name}

	case *ast.Add:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&Add{Left: a, Right: b})

	case *ast.Sub:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&Sub{Left: a, Right: b})

	case *ast.Minus:
		a := c.evalExpression(block, x.Expression)

		return block.setTmp(&Minus{Expression: a})

	case *ast.Mul:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&Mul{Left: a, Right: b})

	case *ast.Div:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&Div{Left: a, Right: b})

	case *ast.Mod:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&Mod{Left: a, Right: b})

	case *ast.And:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&And{Left: a, Right: b})

	case *ast.Not:
		a := c.evalExpression(block, x.Expression)

		return block.setTmp(&Not{Expression: a})

	case *ast.Equal:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&Equal{Left: a, Right: b})

	case *ast.LessThan:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&LessThan{Left: a, Right: b})

	case *ast.LessThanEqual:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&LessThanEqual{Left: a, Right: b})

	case *ast.GreaterThan:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&GreaterThan{Left: a, Right: b})

	case *ast.GreaterThanEqual:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&GreaterThanEqual{Left: a, Right: b})

	case *ast.NotEqual:
		a := c.evalExpression(block, x.A)
		b := c.evalExpression(block, x.B)

		return block.setTmp(&NotEqual{Left: a, Right: b})

	case *ast.Group:
		return c.evalExpression(block, x.Expression)

	case *ast.Slice:
		return c.evalSlice(block, x)

	case *ast.Ternary:
		v := c.prog.Main.nextTmp()

		cond := block.setTmp(c.evalExpression(block, x.Condition))

		ifTrue := &If{
			Condition: cond,
			Block:     &Block{tmpVar: c.prog},
		}

		ifTrue.Block.add(&Assignment{
			Name:  v.Name,
			Value: c.evalExpression(ifTrue.Block, x.True),
		})

		block.add(ifTrue)

		notCond := block.setTmp(&Not{Expression: cond})

		ifFalse := &If{
			Condition: notCond,
			Block:     &Block{tmpVar: c.prog},
		}

		ifFalse.Block.add(&Assignment{
			Name:  v.Name,
			Value: c.evalExpression(ifFalse.Block, x.False),
		})

		block.add(ifFalse)

		return v

	case *ast.Func:
		f := &Func{
			Name: x.Name,
		}
		for _, p := range x.Parameters {
			f.Parameters = append(f.Parameters, c.evalExpression(block, p))
		}
		return block.setTmp(f)

	default:
		panic(fmt.Sprintf("no valid expression %T", x))
	}
}

func (c *Compiler) evalSlice(block *Block, s *ast.Slice) Atom {
	value := c.evalExpression(block, s.Value)

	var from Value
	if s.From != nil {
		from = c.evalExpression(block, s.From)
	} else {
		from = &NumberValue{Value: 0}
	}

	var to Value
	if s.To != nil {
		to = c.evalExpression(block, s.To)
	} else {
		to = block.setTmp(&Func{
			Name:       "len",
			Parameters: []Atom{value},
		})
	}

	return block.setTmp(&Slice{
		Value: value,
		From:  from,
		To:    to,
	})
}

func (c *Compiler) evalFuncStatement(block *Block, f *ast.FuncStatement) {
	fnc := &Func{
		Name: f.Func.Name,
	}
	for _, p := range f.Func.Parameters {
		fnc.Parameters = append(fnc.Parameters, c.evalExpression(block, p))
	}

	block.add(fnc)
}

func (c *Compiler) evalIfStatement(block *Block, ifst *ast.If) {
	condVar := block.setTmp(c.evalExpression(block, ifst.Condition))

	ifs := &If{
		Condition: condVar,
		Block:     &Block{tmpVar: c.prog},
	}

	c.evalBlock(ifs.Block, ifst.If)

	block.add(ifs)
}

func (c *Compiler) evalForStatement(block *Block, forst *ast.For) {
	result := c.evalExpression(block, forst.Initial.Expression)
	block.set(forst.Initial.Name, result)

	condVar := block.setTmp(c.evalExpression(block, forst.Condition))

	while := &While{
		Condition: condVar,
		Block:     &Block{tmpVar: c.prog},
	}

	c.evalBlock(while.Block, forst.Body)

	tCond := c.evalExpression(while.Block, forst.Each.Expression)
	while.Block.set(forst.Each.Name, tCond)

	while.Block.set(condVar.Name, c.evalExpression(while.Block, forst.Condition))

	block.add(while)
}
