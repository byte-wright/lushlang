package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type Visitor struct {
	*BaseLushVisitor
	Target parseTarget
	Name   string
}

type parseTarget interface {
	addImport(path string)
	addMainStatement(ast.Statement)
	addFuncDef(*ast.FuncDef)
	addExternalFuncDef(*ast.ExternalFuncDef)
}

func newVisitor(name string, target parseTarget) *Visitor {
	return &Visitor{
		BaseLushVisitor: &BaseLushVisitor{},
		Target:          target,
		Name:            name,
	}
}

func (v *Visitor) VisitProgram(ctx *ProgramContext) any {
	for _, imp := range ctx.AllImportStatement() {
		path := imp.STRING().GetText()
		path = strings.Trim(path, "\" ")
		v.Target.addImport(path)
	}

	for _, st := range ctx.AllStatement() {
		statement := st.Accept(v)
		v.Target.addMainStatement(statement.(ast.Statement))
	}

	for _, fd := range ctx.AllFuncDef() {
		funcDef := fd.Accept(v)
		v.Target.addFuncDef(funcDef.(*ast.FuncDef))
	}

	return nil
}

func (v *Visitor) VisitLibrary(ctx *LibraryContext) any {
	for _, imp := range ctx.AllImportStatement() {
		path := imp.STRING().GetText()
		path = strings.Trim(path, "\" ")
		v.Target.addImport(path)
	}

	for _, fd := range ctx.AllFuncDef() {
		funcDef := fd.Accept(v)
		v.Target.addFuncDef(funcDef.(*ast.FuncDef))
	}

	for _, fd := range ctx.AllExternalFuncDef() {
		funcDef := fd.Accept(v)
		v.Target.addExternalFuncDef(funcDef.(*ast.ExternalFuncDef))
	}

	return nil
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

	if ctx.For_() != nil {
		return ctx.For_().Accept(v)
	}

	if ctx.ReturnStatement() != nil {
		return ctx.ReturnStatement().Accept(v)
	}

	panic("invalid statement in parser")
}

func (v *Visitor) VisitFuncDef(ctx *FuncDefContext) any {
	params := []*ast.Param{}

	for _, p := range ctx.AllParam() {
		params = append(params, p.Accept(v).(*ast.Param))
	}

	returns := []common.Type{}
	for _, tp := range ctx.AllType_() {
		returns = append(returns, tp.Accept(v).(common.Type))
	}

	return &ast.FuncDef{
		Name:    ctx.ID().GetText(),
		Params:  params,
		Returns: returns,
		Body:    ctx.Block().Accept(v).(*ast.Block),
	}
}

func (v *Visitor) VisitExternalFuncDef(ctx *ExternalFuncDefContext) any {
	params := []*ast.Param{}

	for _, p := range ctx.AllParam() {
		params = append(params, p.Accept(v).(*ast.Param))
	}

	returns := []common.Type{}
	for _, tp := range ctx.AllType_() {
		returns = append(returns, tp.Accept(v).(common.Type))
	}

	code := ctx.EXTERNAL_CODE().GetText()
	code = code[3 : len(code)-3]

	return &ast.ExternalFuncDef{
		Name:    ctx.ID().GetText(),
		Params:  params,
		Returns: returns,
		Code:    code,
	}
}

func (v *Visitor) VisitParam(ctx *ParamContext) any {
	return &ast.Param{
		Name: ctx.ID().GetText(),
		Type: ctx.Type_().Accept(v).(common.Type),
	}
}

func (v *Visitor) VisitIf(ctx *IfContext) any {
	return &ast.If{
		Condition: ctx.Expression().Accept(v).(ast.Expression),
		If:        ctx.Block().Accept(v).(*ast.Block),
	}
}

func (v *Visitor) VisitFor(ctx *ForContext) any {
	return &ast.For{
		Initial:   ctx.AllAssignment()[0].Accept(v).(*ast.Assignment),
		Condition: ctx.Expression().Accept(v).(ast.Expression),
		Each:      ctx.AllAssignment()[1].Accept(v).(*ast.Assignment),
		Body:      ctx.Block().Accept(v).(*ast.Block),
	}
}

func (v *Visitor) VisitReturnStatement(ctx *ReturnStatementContext) any {
	xps := []ast.Expression{}
	for _, x := range ctx.AllExpression() {
		xps = append(xps, x.Accept(v).(ast.Expression))
	}

	return &ast.ReturnStatement{
		Expressions: xps,
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
	names := []string{}
	for _, id := range ctx.AllID() {
		names = append(names, id.GetText())
	}

	return &ast.Assignment{
		Names: names,
		Expressions: common.Map(ctx.AllExpression(), func(x IExpressionContext) ast.Expression {
			return x.Accept(v).(ast.Expression)
		}),
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

	// method call
	if ctx.Func_() != nil {
		return &ast.Method{
			Expression: a,
			Func:       ctx.Func_().Accept(v).(*ast.Func),
		}
	}

	// unary ops
	if ctx.GetUnary_op() != nil {
		if ctx.MINUS() != nil {
			return &ast.Minus{Expression: a}
		}

		if ctx.NOT() != nil {
			return &ast.Not{Expression: a}
		}
	}

	// slice
	if ctx.GetSlice() != nil {
		var from ast.Expression
		if ctx.from != nil {
			from = ctx.from.Accept(v).(ast.Expression)
		}

		var to ast.Expression
		if ctx.to != nil {
			to = ctx.to.Accept(v).(ast.Expression)
		}
		return &ast.Slice{
			Value: a,
			From:  from,
			To:    to,
		}
	}

	// index
	if ctx.GetIndex() != nil {
		value := ctx.index.Accept(v).(ast.Expression)
		position := ctx.position.Accept(v).(ast.Expression)

		return &ast.Index{
			Value:    value,
			Position: position,
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

	panic(fmt.Sprint("invalid expression", ctx.GetText()))
}

func (v *Visitor) VisitAtom(ctx *AtomContext) any {
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

	if ctx.Bool_() != nil {
		v := ctx.Bool_().GetText() == "true"
		return &ast.Bool{
			Value: v,
		}
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

	if ctx.Array() != nil {
		return ctx.Array().Accept(v)
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

func (v *Visitor) VisitArray(ctx *ArrayContext) any {
	exps := []ast.Expression{}
	for _, exp := range ctx.AllExpression() {
		exps = append(exps, exp.Accept(v).(ast.Expression))
	}

	tp := common.None
	if ctx.PrimitiveType() != nil {
		tp = ctx.PrimitiveType().Accept(v).(common.PrimitiveType)
	}

	return &ast.Array{
		Values: exps,
		Type:   tp,
	}
}

func (v *Visitor) VisitType(ctx *TypeContext) any {
	if ctx.GetArrayType() != nil {
		return &common.ArrayType{
			ElementType: ctx.GetArrayType().Accept(v).(*common.BasicType),
		}
	}

	return &common.BasicType{Type: ctx.PrimitiveType().Accept(v).(common.PrimitiveType)}
}

func (v *Visitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) any {
	return common.PrimitiveTypeFor(ctx.GetText())
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
