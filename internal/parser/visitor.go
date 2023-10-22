package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/byte-wright/lush/internal/ast"
)

type Visitor struct {
	*BaseLushVisitor
	Program *ast.Program
	Name    string
}

func newVisitor(name string) *Visitor {
	return &Visitor{
		BaseLushVisitor: &BaseLushVisitor{},
		Program: &ast.Program{
			Root: &ast.Block{},
		},
		Name: name,
	}
}

func (v *Visitor) VisitProgram(ctx *ProgramContext) any {
	for _, st := range ctx.AllStatement() {
		statement := st.Accept(v)
		v.Program.Root.Statements = append(v.Program.Root.Statements, statement.(ast.Statement))
	}

	return v.Program
}

func (v *Visitor) VisitStatement(ctx *StatementContext) any {
	if ctx.Assignment() != nil {
		return ctx.Assignment().Accept(v)
	}
	if ctx.FuncStatement() != nil {
		return ctx.FuncStatement().Accept(v)
	}
	if ctx.If_() != nil {
		return ctx.If_().Accept(v)
	}

	panic("invalid statement")
}

func (v *Visitor) VisitIf(ctx *IfContext) any {
	return &ast.If{
		Condition: ctx.Expression().Accept(v).(ast.Expression),
		If:        ctx.Block().Accept(v).(*ast.Block),
	}
}

func (v *Visitor) VisitBlock(ctx *BlockContext) any {
	b := &ast.Block{}

	for _, s := range ctx.AllStatement() {
		b.Statements = append(b.Statements, s.Accept(v).(ast.Statement))
	}

	return b
}

func (v *Visitor) VisitAssignment(ctx *AssignmentContext) any {
	name := ctx.ID().GetText()
	exp := ctx.Expression().Accept(v)
	return &ast.Assignment{
		Name:       name,
		Expression: exp.(ast.Expression),
	}
}

func (v *Visitor) VisitFuncStatement(ctx *FuncStatementContext) any {
	return &ast.FuncStatement{
		Func: ctx.Func_().Accept(v).(*ast.Func),
	}
}

func (v *Visitor) VisitExpression(ctx *ExpressionContext) any {
	if ctx.Atom() != nil {
		return ctx.Atom().Accept(v)
	}

	a := ctx.Expression(0).Accept(v).(ast.Expression)

	// unary ops
	if ctx.GetUnary_op() != nil {
		if ctx.MINUS() != nil {
			return &ast.Minus{Expression: a}
		}

		if ctx.NOT() != nil {
			return &ast.Not{Expression: a}
		}
	}

	b := ctx.Expression(1).Accept(v).(ast.Expression)

	// mul ops
	if ctx.MUL() != nil {
		return &ast.Mul{A: a, B: b}
	}

	if ctx.DIV() != nil {
		return &ast.Div{A: a, B: b}
	}

	if ctx.MOD() != nil {
		return &ast.Mod{A: a, B: b}
	}

	// add ops
	if ctx.PLUS() != nil {
		return &ast.Add{A: a, B: b}
	}

	if ctx.MINUS() != nil {
		return &ast.Sub{A: a, B: b}
	}

	// rel ops
	if ctx.LT() != nil {
		return &ast.LessThan{A: a, B: b}
	}
	if ctx.LTE() != nil {
		return &ast.LessThanEqual{A: a, B: b}
	}
	if ctx.GT() != nil {
		return &ast.GreaterThan{A: a, B: b}
	}
	if ctx.GTE() != nil {
		return &ast.GreaterThanEqual{A: a, B: b}
	}
	if ctx.EQ() != nil {
		return &ast.Equal{A: a, B: b}
	}
	if ctx.NEQ() != nil {
		return &ast.NotEqual{A: a, B: b}
	}

	// logical and
	if ctx.LAND() != nil {
		return &ast.And{A: a, B: b}
	}

	// logical or
	if ctx.LOR() != nil {
		return &ast.Or{A: a, B: b}
	}

	// ternary
	if ctx.GetTernary() != nil {
		c := ctx.Expression(2).Accept(v).(ast.Expression)
		return &ast.Ternary{
			Condition: a,
			True:      b,
			False:     c,
		}
	}

	// slice
	if ctx.GetSlice() != nil {
		c := ctx.Expression(2).Accept(v).(ast.Expression)
		return &ast.Slice{
			Value: a,
			From:  b,
			To:    c,
		}
	}

	panic(fmt.Sprint("invalid expression", ctx.GetText()))
}

func (v *Visitor) VisitAtom(ctx *AtomContext) any {
	if ctx.Value() != nil {
		return ctx.Value().Accept(v)
	}

	if ctx.EnvVar() != nil {
		return ctx.EnvVar().Accept(v)
	}

	if ctx.Var_() != nil {
		return ctx.Var_().Accept(v)
	}

	if ctx.Func_() != nil {
		return ctx.Func_().Accept(v)
	}

	if ctx.GetGroup() != nil {
		return &ast.Group{Expression: ctx.GetGroup().Accept(v).(ast.Expression)}
	}

	panic("invalid atom")
}

func (v *Visitor) VisitFunc(ctx *FuncContext) any {
	f := &ast.Func{
		Name: ctx.ID().GetText(),
	}
	for _, p := range ctx.AllExpression() {
		f.Parameters = append(f.Parameters, p.Accept(v).(ast.Expression))
	}
	return f
}

func (v *Visitor) VisitVar(ctx *VarContext) any {
	return &ast.Var{
		Name: ctx.ID().GetText(),
	}
}

func (v *Visitor) VisitEnvVar(ctx *EnvVarContext) any {
	id := ctx.ENVVAR().GetText()
	id = id[1:]
	return &ast.EnvVar{
		Name: id,
	}
}

var strEscapes = []struct {
	k string
	v string
}{{
	k: "\\n",
	v: "\n",
}, {
	k: "\\t",
	v: "\t",
}, {
	k: "\\\\",
	v: "\\",
}, {
	k: "\\\"",
	v: "\"",
}}

func (v *Visitor) VisitValue(ctx *ValueContext) any {
	if ctx.String_() != nil {
		txt := ctx.String_().GetText()
		txt = txt[1 : len(txt)-1]

		for _, e := range strEscapes {
			txt = strings.ReplaceAll(txt, e.k, e.v)
		}

		return &ast.String{
			Value: txt,
		}
	}

	if ctx.Number() != nil {
		v, err := strconv.Atoi(ctx.Number().GetText())
		if err != nil {
			panic(err)
		}

		return &ast.Number{
			Value: v,
		}
	}

	panic("invalid value")
}
