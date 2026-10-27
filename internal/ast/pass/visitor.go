package astpass

import (
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type Visitor interface {
	Expression(ast.Expression)
	BeforeStatement(ast.Statement)
	Statement(ast.Statement)
	BeforeFuncDef(*ast.FuncDef)
	BeforeBlock(*ast.Block)
	AfterBlock(*ast.Block)
}

type context struct {
	ast     *ast.Program
	visitor Visitor
}

func Visit(ast *ast.Program, v Visitor) {
	c := &context{
		ast:     ast,
		visitor: v,
	}

	c.visit()
}

func (c *context) visit() {
	c.evalBlock(c.ast.Root)

	for _, fd := range funcDefs(c.ast) {
		c.visitor.BeforeFuncDef(fd)
		c.evalBlock(fd.Body)
	}
}

func funcDefs(prog *ast.Program) []*ast.FuncDef {
	fds := append([]*ast.FuncDef{}, prog.Main.FuncDefs...)
	for _, pkg := range prog.Packages() {
		fds = append(fds, pkg.FuncDefs...)
	}

	return fds
}

func (c context) evalBlock(block *ast.Block) {
	for _, statement := range block.Statements {
		c.visitor.BeforeBlock(block)
		c.evalStatement(statement)
		c.visitor.AfterBlock(block)
	}
}

func (c context) evalStatement(statement ast.Statement) {
	c.visitor.BeforeStatement(statement)

	switch st := statement.(type) {
	case *ast.Assignment:
		for _, exp := range st.Expressions {
			c.evalExpression(exp)
		}
	case *ast.For:
		c.evalStatement(st.Initial)
		c.evalExpression(st.Condition)
		c.evalStatement(st.Each)
		c.evalBlock(st.Body)
	case *ast.If:
		c.evalExpression(st.Condition)
		c.evalBlock(st.If)
		if st.Else != nil {
			c.evalBlock(st.Else)
		}

	case *ast.FuncStatement:
		c.evalExpression(st.Call)

	case *ast.ReturnStatement:
		for _, exp := range st.Expressions {
			c.evalExpression(exp)
		}

	default:
		common.IncompleteSwitch(st)
	}

	c.visitor.Statement(statement)
}

func (c context) evalExpression(expression ast.Expression) {
	switch exp := expression.(type) {
	case *ast.Func:
		for _, p := range exp.Parameters {
			c.evalExpression(p)
		}
	case *ast.Method:
		for _, p := range exp.Func.Parameters {
			c.evalExpression(p)
		}
	case *ast.Add:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.Sub:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.Minus:
		c.evalExpression(exp.Expression)
	case *ast.Not:
		c.evalExpression(exp.Expression)
	case *ast.Mul:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.Div:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.Mod:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.And:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.Or:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.LessThan:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.LessThanEqual:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.GreaterThan:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.GreaterThanEqual:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.Equal:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.NotEqual:
		c.evalExpression(exp.A)
		c.evalExpression(exp.B)
	case *ast.Array:
		for _, v := range exp.Values {
			c.evalExpression(v)
		}
	case *ast.Slice:
		c.evalExpression(exp.Value)
		if exp.From != nil {
			c.evalExpression(exp.From)
		}
		if exp.To != nil {
			c.evalExpression(exp.To)
		}
	case *ast.Index:
		c.evalExpression(exp.Value)
		c.evalExpression(exp.Position)
	case *ast.Ternary:
		c.evalExpression(exp.Condition)
		c.evalExpression(exp.True)
		c.evalExpression(exp.False)
	case *ast.Group:
		c.evalExpression(exp.Expression)
	case *ast.String, *ast.Number, *ast.Bool, *ast.Var, *ast.EnvVar:

	default:
		common.IncompleteSwitch(exp)
	}

	c.visitor.Expression(expression)
}
