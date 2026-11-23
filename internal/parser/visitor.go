package parser

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type Visitor struct {
	*BaseLushVisitor
	Target parseTarget
	Name   string
	errors []*ast.ASTError
	// checking a program with a broken import only finds follow-up errors, these stop it
	importErrors []*ast.ASTError
}

type parseTarget interface {
	// addImport is false when the file imports another package as name
	addImport(path, name string, loc *common.Location) bool
	addMainStatement(ast.Statement)
	addFuncDef(*ast.FuncDef)
	addExternalFuncDef(*ast.ExternalFuncDef)
	addGlobalDef(*ast.GlobalDef)
}

func newVisitor(name string, target parseTarget) *Visitor {
	return &Visitor{
		BaseLushVisitor: &BaseLushVisitor{},
		Target:          target,
		Name:            name,
	}
}

func (v *Visitor) loc(c antlr.ParserRuleContext) *common.Location {
	return &common.Location{
		File:     v.Name,
		LineFrom: c.GetStart().GetLine(),
		ColFrom:  c.GetStart().GetColumn() + 1,
		LineTo:   c.GetStop().GetLine(),
		ColTo:    c.GetStop().GetColumn() + 1,
	}
}

func (v *Visitor) locTk(c antlr.Token) *common.Location {
	return &common.Location{
		File:     v.Name,
		LineFrom: c.GetLine(),
		ColFrom:  c.GetColumn() + 1,
		LineTo:   c.GetLine(),
		ColTo:    c.GetColumn() + 1,
	}
}

func (v *Visitor) VisitProgram(ctx *ProgramContext) any {
	v.imports(ctx.AllImportStatement())

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

func (v *Visitor) VisitPackage(ctx *PackageContext) any {
	v.imports(ctx.AllImportStatement())

	for _, ga := range ctx.AllGlobalAssignment() {
		v.Target.addGlobalDef(&ast.GlobalDef{Assignment: ga.Accept(v).(*ast.Assignment)})
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

func (v *Visitor) imports(imps []IImportStatementContext) {
	paths := map[string]string{}

	for _, imp := range imps {
		pkgPath := strings.Trim(imp.STRING().GetText(), "\" ")
		loc := v.loc(imp)

		if strings.HasPrefix(pkgPath, "/") {
			v.importError(loc, "package %q can't start with /, a local package starts with ./ or ../", pkgPath)
			continue
		}

		name := path.Base(pkgPath)
		if imp.GetAlias() != nil {
			name = imp.GetAlias().GetText()
		} else {
			// code can't refer to it, the error is here instead of misleading ones at a use
			invalid := invalidPackageName(name)
			if invalid != "" {
				v.importError(loc, "package %q needs a name with as, the folder name %v %v", pkgPath, name, invalid)
				continue
			}
		}

		if !v.Target.addImport(pkgPath, name, loc) {
			v.importError(loc, "package %v is already imported from %q", name, paths[name])
			continue
		}

		if paths[name] == "" {
			paths[name] = pkgPath
		}
	}
}

func (v *Visitor) importError(loc *common.Location, format string, a ...any) {
	v.importErrors = append(v.importErrors, &ast.ASTError{Location: loc, Message: fmt.Sprintf(format, a...)})
}

func (v *Visitor) VisitStatement(ctx *StatementContext) any {
	if ctx.Assignment() != nil {
		assignment := ctx.Assignment().Accept(v).(*ast.Assignment)
		v.cmdBody(assignment, ctx.GetCmdBody())
		assignment.Catch = v.catch(ctx.CatchClause())
		return assignment
	}

	if ctx.GlobalAssignment() != nil {
		assignment := ctx.GlobalAssignment().Accept(v).(*ast.Assignment)
		v.cmdBody(assignment, ctx.GetCmdBody())
		assignment.Catch = v.catch(ctx.CatchClause())
		return assignment
	}

	if ctx.FuncStatement() != nil {
		st := ctx.FuncStatement().Accept(v).(*ast.FuncStatement)
		st.Catch = v.catch(ctx.CatchClause())
		return st
	}

	if ctx.CmdStatement() != nil {
		st := ctx.CmdStatement().Accept(v).(*ast.FuncStatement)
		st.Catch = v.catch(ctx.CatchClause())
		return st
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

	if ctx.FailStatement() != nil {
		return ctx.FailStatement().Accept(v)
	}

	if ctx.BREAK() != nil {
		return &ast.Break{Location: v.loc(ctx)}
	}

	if ctx.CONTINUE() != nil {
		return &ast.Continue{Location: v.loc(ctx)}
	}

	panic("invalid statement in parser")
}

func (v *Visitor) VisitFuncDef(ctx *FuncDefContext) any {
	params := []*ast.Param{}

	for _, p := range ctx.AllParam() {
		params = append(params, p.Accept(v).(*ast.Param))
	}

	returns, canFail := v.results(ctx.Results())

	return &ast.FuncDef{
		Name:     ctx.ID().GetText(),
		Location: v.locTk(ctx.ID().GetSymbol()),
		Params:   params,
		Returns:  returns,
		CanFail:  canFail,
		Body:     ctx.Block().Accept(v).(*ast.Block),
	}
}

func (v *Visitor) VisitExternalFuncDef(ctx *ExternalFuncDefContext) any {
	params := []*ast.Param{}

	for _, p := range ctx.AllParam() {
		params = append(params, p.Accept(v).(*ast.Param))
	}

	returns, canFail := v.results(ctx.Results())

	name := ctx.ID().GetText()
	tk := ctx.EXTERNAL_CODE().GetSymbol()
	code := tk.GetText()
	code = code[3 : len(code)-3]

	end := strings.IndexByte(code, '\n')
	if end < 0 {
		end = len(code)
	}

	target := strings.TrimSpace(code[:end])
	if !isTargetName(target) {
		v.errors = append(v.errors, &ast.ASTError{Location: v.locTk(tk), Message: fmt.Sprintf("the code of %v must start with its target, like ```bash", name)})
	}

	return &ast.ExternalFuncDef{
		Name:     name,
		Location: v.locTk(ctx.ID().GetSymbol()),
		Params:   params,
		Returns:  returns,
		CanFail:  canFail,
		Target:   target,
		Code:     code[end:],
	}
}

func isTargetName(s string) bool {
	if s == "" || s[0] < 'a' || s[0] > 'z' {
		return false
	}

	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') {
			return false
		}
	}

	return true
}

// results returns the types of the values, and whether the last result is an error
func (v *Visitor) results(ctx IResultsContext) ([]common.Type, bool) {
	returns := []common.Type{}
	if ctx == nil {
		return returns, false
	}

	results := ctx.AllResult()
	for i, r := range results {
		if r.ERROR_TYPE() == nil {
			returns = append(returns, r.Type_().Accept(v).(common.Type))
			continue
		}

		if i < len(results)-1 {
			v.errors = append(v.errors, &ast.ASTError{Location: v.loc(r), Message: "error must be the last result"})
		}
	}

	return returns, results[len(results)-1].ERROR_TYPE() != nil
}

func (v *Visitor) VisitParam(ctx *ParamContext) any {
	return &ast.Param{
		Location: v.loc(ctx),
		Name:     ctx.ID().GetText(),
		Type:     ctx.Type_().Accept(v).(common.Type),
	}
}

func (v *Visitor) VisitIf(ctx *IfContext) any {
	st := &ast.If{
		Condition: ctx.Expression().Accept(v).(ast.Expression),
		If:        ctx.GetThen().Accept(v).(*ast.Block),
		Location:  v.loc(ctx),
	}

	if ctx.GetElseIf() != nil {
		st.Else = &ast.Block{Statements: []ast.Statement{ctx.GetElseIf().Accept(v)}}
	}
	if ctx.GetElseBlock() != nil {
		st.Else = ctx.GetElseBlock().Accept(v).(*ast.Block)
	}

	return st
}

func (v *Visitor) VisitFor(ctx *ForContext) any {
	if ctx.GetOver() != nil {
		return v.forIn(ctx)
	}

	st := &ast.For{
		Body:     ctx.Block().Accept(v).(*ast.Block),
		Location: v.loc(ctx),
	}

	if ctx.GetInit() != nil {
		st.Initial = ctx.GetInit().Accept(v).(*ast.Assignment)
		st.Each = ctx.GetStep().Accept(v).(*ast.Assignment)
	}

	if ctx.GetCond() != nil {
		st.Condition = ctx.GetCond().Accept(v).(ast.Expression)
	}

	return st
}

func (v *Visitor) forIn(ctx *ForContext) *ast.ForIn {
	st := &ast.ForIn{
		Value:    v.target(ctx.GetValue()),
		Over:     ctx.GetOver().Accept(v).(ast.Expression),
		Body:     ctx.Block().Accept(v).(*ast.Block),
		Location: v.loc(ctx),
	}

	// without a pipe, it's a loop over the lines of stdout: []string
	if c, isCmd := st.Over.(*ast.Cmd); isCmd && c.Pipe() != "" {
		c.Use = ast.UseLoop
	}

	if ctx.GetIndex() != nil {
		st.Index = v.target(ctx.GetIndex())
	}

	if ctx.GetFrom() != nil {
		st.From = ctx.GetFrom().Accept(v).(ast.Expression)
	}

	return st
}

func (v *Visitor) target(tk antlr.Token) *ast.AssignmentTarget {
	return &ast.AssignmentTarget{Location: v.locTk(tk), Name: tk.GetText()}
}

func (v *Visitor) catch(ctx ICatchClauseContext) *ast.Catch {
	if ctx == nil {
		return nil
	}

	return &ast.Catch{
		Error:    v.target(ctx.GetName()),
		Block:    ctx.Block().Accept(v).(*ast.Block),
		Location: v.locTk(ctx.CATCH().GetSymbol()),
	}
}

func (v *Visitor) VisitFailStatement(ctx *FailStatementContext) any {
	return &ast.Fail{
		Message:  ctx.Expression().Accept(v).(ast.Expression),
		Location: v.loc(ctx),
	}
}

func (v *Visitor) VisitReturnStatement(ctx *ReturnStatementContext) any {
	xps := []ast.Expression{}
	for _, x := range ctx.AllExpression() {
		xps = append(xps, x.Accept(v).(ast.Expression))
	}

	return &ast.ReturnStatement{
		Expressions: xps,
		Location:    v.loc(ctx),
	}
}

func (v *Visitor) VisitBlock(ctx *BlockContext) any {
	b := &ast.Block{End: v.locTk(ctx.RCUR().GetSymbol())}

	for _, s := range ctx.AllStatement() {
		b.Statements = append(b.Statements, s.Accept(v).(ast.Statement))
	}

	return b
}

func (v *Visitor) VisitAssignment(ctx *AssignmentContext) any {
	targets := common.Map(ctx.AllAssignTarget(), func(t IAssignTargetContext) *ast.AssignmentTarget {
		return t.Accept(v).(*ast.AssignmentTarget)
	})

	expressions := common.Map(ctx.AllExpression(), func(x IExpressionContext) ast.Expression {
		return x.Accept(v).(ast.Expression)
	})

	op := ctx.GetOp()
	if op.GetTokenType() != LushParserASSIGN {
		if len(targets) != 1 || len(expressions) != 1 {
			v.errors = append(v.errors, &ast.ASTError{Location: v.locTk(op), Message: fmt.Sprintf("%v needs one variable and one value", op.GetText())})
			// only the first pair is checked, so there are no follow-up errors
			targets, expressions = targets[:1], expressions[:1]
		}

		expressions = []ast.Expression{v.compound(ctx, targets[0], expressions[0])}
	}

	return &ast.Assignment{
		Location:    v.loc(ctx),
		Targets:     targets,
		Expressions: expressions,
	}
}

func (v *Visitor) VisitAssignTarget(ctx *AssignTargetContext) any {
	target := v.target(ctx.ID().GetSymbol())

	if ctx.GetPosition() != nil {
		target.Element = &ast.Index{
			Value:    &ast.Var{Name: target.Name, Location: target.Location},
			Position: ctx.GetPosition().Accept(v).(ast.Expression),
			Location: v.loc(ctx),
		}
	}

	return target
}

// compound turns x += y into x = x + (y)
func (v *Visitor) compound(ctx *AssignmentContext, target *ast.AssignmentTarget, value ast.Expression) ast.Expression {
	var a ast.Expression = &ast.Var{Name: target.Name, Location: target.Location}
	if target.Element != nil {
		a = &ast.TargetValue{Target: target}
	}
	b := &ast.Group{Expression: value}
	loc := v.loc(ctx)

	switch ctx.GetOp().GetTokenType() {
	case LushParserPLUS_ASSIGN:
		return &ast.Add{A: a, B: b, Location: loc}
	case LushParserMINUS_ASSIGN:
		return &ast.Sub{A: a, B: b, Location: loc}
	case LushParserMUL_ASSIGN:
		return &ast.Mul{A: a, B: b, Location: loc}
	case LushParserDIV_ASSIGN:
		return &ast.Div{A: a, B: b, Location: loc}
	case LushParserMOD_ASSIGN:
		return &ast.Mod{A: a, B: b, Location: loc}
	}

	panic("invalid assignment operator in parser")
}

func (v *Visitor) VisitGlobalAssignment(ctx *GlobalAssignmentContext) any {
	assignment := ctx.Assignment().Accept(v).(*ast.Assignment)
	assignment.Global = true

	return assignment
}

func (v *Visitor) VisitFuncStatement(ctx *FuncStatementContext) any {
	return &ast.FuncStatement{Call: v.call(ctx.GetNamespace(), ctx.Func_(), ctx.TRY())}
}

func (v *Visitor) call(namespace antlr.Token, ctx IFuncContext, try antlr.TerminalNode) ast.Expression {
	f := ctx.Accept(v).(*ast.Func)
	f.Try = try != nil

	if namespace != nil {
		return &ast.Method{Namespace: namespace.GetText(), Func: f}
	}

	return f
}

func (v *Visitor) VisitExpression(ctx *ExpressionContext) any {
	if ctx.Atom() != nil {
		return ctx.Atom().Accept(v)
	}

	a := ctx.Expression(0).Accept(v).(ast.Expression)
	loc := v.loc(ctx)

	// unary ops
	if ctx.GetUnary_op() != nil {
		if ctx.MINUS() != nil {
			return &ast.Minus{Expression: a, Location: loc}
		}

		if ctx.NOT() != nil {
			return &ast.Not{Expression: a, Location: loc}
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
			Value:    a,
			From:     from,
			To:       to,
			Location: loc,
		}
	}

	// index
	if ctx.GetIndex() != nil {
		position := ctx.position.Accept(v).(ast.Expression)

		return &ast.Index{
			Value:    a,
			Position: position,
			Location: loc,
		}
	}

	b := ctx.Expression(1).Accept(v).(ast.Expression)

	// mul ops
	if ctx.MUL() != nil {
		return &ast.Mul{A: a, B: b, Location: loc}
	}

	if ctx.DIV() != nil {
		return &ast.Div{A: a, B: b, Location: loc}
	}

	if ctx.MOD() != nil {
		return &ast.Mod{A: a, B: b, Location: loc}
	}

	// add ops
	if ctx.PLUS() != nil {
		return &ast.Add{A: a, B: b, Location: loc}
	}

	if ctx.MINUS() != nil {
		return &ast.Sub{A: a, B: b, Location: loc}
	}

	// rel ops
	if ctx.LT() != nil {
		return &ast.LessThan{A: a, B: b, Location: loc}
	}
	if ctx.LTE() != nil {
		return &ast.LessThanEqual{A: a, B: b, Location: loc}
	}
	if ctx.GT() != nil {
		return &ast.GreaterThan{A: a, B: b, Location: loc}
	}
	if ctx.GTE() != nil {
		return &ast.GreaterThanEqual{A: a, B: b, Location: loc}
	}
	if ctx.EQ() != nil {
		return &ast.Equal{A: a, B: b, Location: loc}
	}
	if ctx.NEQ() != nil {
		return &ast.NotEqual{A: a, B: b, Location: loc}
	}

	// logical and
	if ctx.LAND() != nil {
		return &ast.And{A: a, B: b, Location: loc}
	}

	// logical or
	if ctx.LOR() != nil {
		return &ast.Or{A: a, B: b, Location: loc}
	}

	// ternary
	if ctx.GetTernary() != nil {
		c := ctx.Expression(2).Accept(v).(ast.Expression)
		return &ast.Ternary{
			Condition: a,
			True:      b,
			False:     c,
			Location:  loc,
		}
	}

	panic(fmt.Sprint("invalid expression", ctx.GetText()))
}

func (v *Visitor) VisitAtom(ctx *AtomContext) any {
	if ctx.String_() != nil {
		return &ast.String{
			Value: v.unquote(ctx.String_().GetStart()),
		}
	}

	if ctx.Number() != nil {
		txt := ctx.Number().GetText()
		n, err := strconv.Atoi(txt)
		// the lexer only lets digits through, so it can only be too big
		if err != nil {
			v.errors = append(v.errors, &ast.ASTError{Location: v.loc(ctx), Message: fmt.Sprintf("number %v is too big for an int", txt)})
		}

		return &ast.Number{
			Value: n,
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
		return v.call(ctx.GetNamespace(), ctx.Func_(), ctx.TRY())
	}

	if ctx.Cmd() != nil {
		return ctx.Cmd().Accept(v)
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
		Name:     ctx.ID().GetText(),
		Location: v.loc(ctx),
	}
	for _, p := range ctx.AllExpression() {
		f.Parameters = append(f.Parameters, p.Accept(v).(ast.Expression))
	}
	return f
}

func (v *Visitor) VisitVar(ctx *VarContext) any {
	namespace := ""
	if ctx.GetNamespace() != nil {
		namespace = ctx.GetNamespace().GetText()
	}

	return &ast.Var{
		Namespace: namespace,
		Name:      ctx.GetName().GetText(),
		Location:  v.loc(ctx),
	}
}

func (v *Visitor) VisitEnvVar(ctx *EnvVarContext) any {
	id := ctx.ENVVAR().GetText()
	id = id[1:]
	return &ast.EnvVar{
		Name:     id,
		Location: v.loc(ctx),
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
		Values:   exps,
		Type:     tp,
		Location: v.loc(ctx),
	}
}

func (v *Visitor) VisitType(ctx *TypeContext) any {
	if ctx.GetArrayType() != nil {
		return &common.ArrayType{
			ElementType: &common.BasicType{Type: ctx.GetArrayType().Accept(v).(common.PrimitiveType)},
		}
	}

	return &common.BasicType{Type: ctx.PrimitiveType().Accept(v).(common.PrimitiveType)}
}

func (v *Visitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) any {
	return common.PrimitiveTypeFor(ctx.GetText())
}

// unquote decodes the escapes of a Go string. The lexer lets any character follow a \,
// so a wrong escape is one error here instead of several from the lexer.
func (v *Visitor) unquote(tk antlr.Token) string {
	txt := tk.GetText()
	txt = txt[1 : len(txt)-1]

	var b strings.Builder
	for i := 0; i < len(txt); {
		if txt[i] != '\\' {
			b.WriteByte(txt[i])
			i++
			continue
		}

		r, multibyte, rest, err := strconv.UnquoteChar(txt[i:], '"')
		if err != nil {
			_, size := utf8.DecodeRuneInString(txt[i+1:])
			v.errors = append(v.errors, &ast.ASTError{Location: escapeLoc(v.locTk(tk), txt[:i]), Message: escapeError(txt[i:])})
			i += 1 + size
			continue
		}

		if r == 0 {
			// bash strings end at a NUL
			v.errors = append(v.errors, &ast.ASTError{Location: escapeLoc(v.locTk(tk), txt[:i]), Message: "strings can't hold NUL"})
		}

		// like strconv.Unquote: \xff and octal escapes are single bytes, not UTF-8
		if r < utf8.RuneSelf || !multibyte {
			b.WriteByte(byte(r))
		} else {
			b.WriteRune(r)
		}
		i = len(txt) - len(rest)
	}

	return b.String()
}

// before is the text between the opening quote and the escape
func escapeLoc(loc *common.Location, before string) *common.Location {
	nl := strings.LastIndexByte(before, '\n')
	if nl < 0 {
		loc.ColFrom += 1 + utf8.RuneCountInString(before)
	} else {
		loc.LineFrom += strings.Count(before, "\n")
		loc.ColFrom = 1 + utf8.RuneCountInString(before[nl+1:])
	}
	loc.LineTo = loc.LineFrom
	loc.ColTo = loc.ColFrom

	return loc
}

func escapeError(s string) string {
	c, _ := utf8.DecodeRuneInString(s[1:])
	switch {
	case c == 'x':
		return `escape \x needs 2 hex digits`
	case c == 'u':
		return `escape \u needs 4 hex digits of a valid code point`
	case c == 'U':
		return `escape \U needs 8 hex digits of a valid code point`
	case c >= '0' && c <= '7':
		return `octal escape needs 3 digits from \000 to \377`
	case unicode.IsPrint(c):
		return fmt.Sprintf(`unknown escape \%c`, c)
	}

	return fmt.Sprintf(`unknown escape, \ is followed by %U`, c)
}
