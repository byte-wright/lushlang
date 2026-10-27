package astpass

import (
	"fmt"
	"strings"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

var (
	intType         = &common.BasicType{Type: common.Int}
	stringType      = &common.BasicType{Type: common.String}
	boolType        = &common.BasicType{Type: common.Bool}
	stringArrayType = &common.ArrayType{ElementType: stringType}
	// unknownTypes are the types of an expression whose error is already reported
	unknownTypes = []common.Type{&common.BasicType{Type: common.None}}
)

type setTypesPass struct {
	ast    *ast.Program
	errors []*ast.ASTError
	block  *ast.Block
	file   *ast.File
	// funcDef is nil outside of functions
	funcDef *ast.FuncDef
	// statementCall is the call of the current function statement, the only call whose
	// value isn't used
	statementCall *ast.Func
}

func SetTypes(ast *ast.Program) []*ast.ASTError {
	te := &setTypesPass{
		ast: ast,
		// Root is visited before the function bodies
		file: ast.File,
	}
	Visit(ast, te)

	return te.errors
}

func (t *setTypesPass) BeforeFuncDef(fd *ast.FuncDef) {
	t.file = fd.File
	t.funcDef = fd

	if len(fd.Returns) > 0 && !terminates(fd.Body) {
		t.errorf(fd.Body.End, "missing return at the end of function %v", fd.Name)
	}

	for _, p := range fd.Params {
		err := fd.Body.SetVar(p.Name, p.Type)
		if err != nil {
			t.errors = append(t.errors, &ast.ASTError{Location: p.Location, Message: err.Error()})
		}
	}
}

func (t *setTypesPass) BeforeBlock(block *ast.Block) {
	t.block = block
}

func (t *setTypesPass) AfterBlock(block *ast.Block) {
}

func (t *setTypesPass) BeforeStatement(statement ast.Statement) {
	t.statementCall = nil

	st, isCall := statement.(*ast.FuncStatement)
	if !isCall {
		return
	}

	switch call := st.Call.(type) {
	case *ast.Func:
		t.statementCall = call
	case *ast.Method:
		t.statementCall = call.Func
	}
}

func (t *setTypesPass) Statement(statement ast.Statement) {
	switch st := statement.(type) {
	case *ast.Assignment:
		t.assignment(st)

	case *ast.If:
		t.expectType(st.Location, "if condition", st.Condition, boolType)

	case *ast.For:
		t.expectType(st.Location, "for condition", st.Condition, boolType)

	case *ast.ReturnStatement:
		t.returnStatement(st)
	}
}

func (t *setTypesPass) assignment(st *ast.Assignment) {
	types := valueTypes(st.Expressions)

	if len(types) != len(st.Targets) {
		if known(types) {
			t.errorf(st.Location, "Wrong number of expressions")
		}

		// the targets are still declared, so their uses aren't reported as undefined
		types = make([]common.Type, len(st.Targets))
		for i := range types {
			types[i] = unknownTypes[0]
		}
	}

	for i, target := range st.Targets {
		v := t.block.GetVar(target.Name)
		if v != nil && !known([]common.Type{v.Type, types[i]}) {
			continue
		}

		err := t.block.SetVar(target.Name, types[i])
		if err != nil {
			t.errors = append(t.errors, &ast.ASTError{Location: target.Location, Message: err.Error()})
		}
	}
}

func (t *setTypesPass) Expression(expression ast.Expression) {
	switch exp := expression.(type) {
	case *ast.Func:
		t.call("", exp)

	case *ast.Method:
		t.call(exp.Namespace, exp.Func)

	case *ast.Var:
		v := t.block.GetVar(exp.Name)
		if v == nil {
			t.errorf(exp.Location, "variable '%v' is not defined", exp.Name)
		} else {
			exp.Type = v.Type
		}

	case *ast.Minus:
		t.unary(exp.Location, "-", exp.Expression, intType)
	case *ast.Not:
		t.unary(exp.Location, "!", exp.Expression, boolType)
	case *ast.Add:
		t.binary(exp.Location, "+", exp.A, exp.B, intType, stringType)
	case *ast.Sub:
		t.binary(exp.Location, "-", exp.A, exp.B, intType)
	case *ast.Mul:
		t.binary(exp.Location, "*", exp.A, exp.B, intType)
	case *ast.Div:
		t.binary(exp.Location, "/", exp.A, exp.B, intType)
	case *ast.Mod:
		t.binary(exp.Location, "%", exp.A, exp.B, intType)
	case *ast.And:
		t.binary(exp.Location, "&&", exp.A, exp.B, boolType)
	case *ast.Or:
		t.binary(exp.Location, "||", exp.A, exp.B, boolType)
	case *ast.LessThan:
		t.binary(exp.Location, "<", exp.A, exp.B, intType)
	case *ast.LessThanEqual:
		t.binary(exp.Location, "<=", exp.A, exp.B, intType)
	case *ast.GreaterThan:
		t.binary(exp.Location, ">", exp.A, exp.B, intType)
	case *ast.GreaterThanEqual:
		t.binary(exp.Location, ">=", exp.A, exp.B, intType)
	case *ast.Equal:
		t.binary(exp.Location, "==", exp.A, exp.B, intType, stringType, boolType)
	case *ast.NotEqual:
		t.binary(exp.Location, "!=", exp.A, exp.B, intType, stringType, boolType)

	case *ast.Index:
		t.indexable(exp.Location, "index", exp.Value)
		t.expectType(exp.Location, "index", exp.Position, intType)

	case *ast.Slice:
		t.indexable(exp.Location, "slice", exp.Value)
		if exp.From != nil {
			t.expectType(exp.Location, "slice bound", exp.From, intType)
		}
		if exp.To != nil {
			t.expectType(exp.Location, "slice bound", exp.To, intType)
		}

	case *ast.Ternary:
		t.expectType(exp.Location, "ternary condition", exp.Condition, boolType)

		tt, ft := exp.True.Types(), exp.False.Types()
		if known(tt) && known(ft) && !sameTypes(tt, ft) {
			t.errorf(exp.Location, "ternary branches must have the same type, not %v and %v", printTypes(tt), printTypes(ft))
		}
	}
}

func (t *setTypesPass) returnStatement(st *ast.ReturnStatement) {
	fd := t.funcDef
	if fd == nil {
		t.errorf(st.Location, "return outside of a function")
		return
	}

	values := valueTypes(st.Expressions)
	if !known(values) {
		return
	}

	if len(values) != len(fd.Returns) {
		t.errorf(st.Location, "function %v returns %v, not %v", fd.Name, count(len(fd.Returns), "value"), len(values))
		return
	}

	for i, want := range fd.Returns {
		if !values[i].Is(want) {
			t.errorf(st.Location, "return value %v of %v must be of type %v, not %v", i+1, fd.Name, want.Print(), values[i].Print())
			return
		}
	}
}

// terminates reports whether the block can't reach its end. Like in Go, only its last
// statement counts.
func terminates(b *ast.Block) bool {
	if len(b.Statements) == 0 {
		return false
	}

	switch st := b.Statements[len(b.Statements)-1].(type) {
	case *ast.ReturnStatement:
		return true

	case *ast.If:
		return st.Else != nil && terminates(st.If) && terminates(st.Else)

	case *ast.For:
		// there's no break, so only a return can leave a loop whose condition is true
		cond, isBool := st.Condition.(*ast.Bool)
		return isBool && cond.Value
	}

	return false
}

func (t *setTypesPass) call(namespace string, f *ast.Func) {
	name := f.Name
	if namespace != "" {
		name = namespace + "." + name
	}

	f.ReturnTypes = t.returnTypes(namespace, name, f)

	// spreading would silently drop a call without values from the values around it
	if len(f.ReturnTypes) == 0 && f != t.statementCall {
		t.errorf(f.Location, "%v doesn't return a value", name)
		f.ReturnTypes = unknownTypes
	}
}

// returnTypes resolves the function and checks its arguments. Unqualified names are
// builtins or functions of the caller's own package.
func (t *setTypesPass) returnTypes(namespace, name string, f *ast.Func) []common.Type {
	args := valueTypes(f.Parameters)

	if namespace == "" {
		tps, isBuiltin, err := builtin(f.Name, args)
		if isBuiltin {
			if err != nil && known(args) {
				t.errorf(f.Location, "%v", err)
			}
			return tps
		}
	}

	params, returns, pkg, err := t.getFunc(namespace, f.Name)
	if err != nil {
		t.errorf(f.Location, "%v", err)
		return unknownTypes
	}

	err = arguments(name, args, params)
	if err != nil && known(args) {
		t.errorf(f.Location, "%v", err)
	}

	f.Package = pkg

	return returns
}

func arguments(name string, args []common.Type, params []*ast.Param) error {
	if len(args) != len(params) {
		return fmt.Errorf("%v takes %v, not %v", name, count(len(params), "argument"), len(args))
	}

	for i, p := range params {
		if !args[i].Is(p.Type) {
			return fmt.Errorf("argument %v of %v must be of type %v, not %v", i+1, name, p.Type.Print(), args[i].Print())
		}
	}

	return nil
}

func (t *setTypesPass) unary(loc *common.Location, op string, x ast.Expression, allowed ...common.Type) {
	xt := x.Types()
	if !known(xt) || len(xt) == 1 && isOneOf(xt[0], allowed) {
		return
	}

	t.errorf(loc, "operator '%v' can't be used with %v", op, printTypes(xt))
}

func (t *setTypesPass) binary(loc *common.Location, op string, a, b ast.Expression, allowed ...common.Type) {
	at, bt := a.Types(), b.Types()
	if !known(at) || !known(bt) {
		return
	}

	if len(at) == 1 && len(bt) == 1 && at[0].Is(bt[0]) && isOneOf(at[0], allowed) {
		return
	}

	t.errorf(loc, "operator '%v' can't be used with %v and %v", op, printTypes(at), printTypes(bt))
}

func (t *setTypesPass) expectType(loc *common.Location, what string, x ast.Expression, want common.Type) {
	xt := x.Types()
	if !known(xt) || len(xt) == 1 && xt[0].Is(want) {
		return
	}

	t.errorf(loc, "%v must be of type %v, not %v", what, want.Print(), printTypes(xt))
}

func (t *setTypesPass) indexable(loc *common.Location, op string, x ast.Expression) {
	xt := x.Types()
	if !known(xt) || len(xt) == 1 && (xt[0].IsArray() || xt[0].Is(stringType)) {
		return
	}

	t.errorf(loc, "cannot %v a value of type %v", op, printTypes(xt))
}

func (t *setTypesPass) errorf(loc *common.Location, format string, args ...any) {
	t.errors = append(t.errors, &ast.ASTError{Location: loc, Message: fmt.Sprintf(format, args...)})
}

// known is false when the types come from an expression with an error, checking
// them would only report follow-up errors.
func known(tps []common.Type) bool {
	for _, tp := range tps {
		if tp.Is(unknownTypes[0]) {
			return false
		}
	}

	return true
}

func isOneOf(tp common.Type, tps []common.Type) bool {
	for _, t := range tps {
		if tp.Is(t) {
			return true
		}
	}

	return false
}

func sameTypes(a, b []common.Type) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Is(b[i]) {
			return false
		}
	}

	return true
}

// valueTypes returns the types of all values, an expression can have several.
func valueTypes(exps []ast.Expression) []common.Type {
	tps := []common.Type{}
	for _, exp := range exps {
		tps = append(tps, exp.Types()...)
	}

	return tps
}

func count(n int, noun string) string {
	if n == 1 {
		return "1 " + noun
	}

	return fmt.Sprintf("%v %vs", n, noun)
}

func printTypes(tps []common.Type) string {
	if len(tps) == 1 {
		return tps[0].Print()
	}

	return "(" + strings.Join(common.Map(tps, func(t common.Type) string { return t.Print() }), ", ") + ")"
}

// getFunc returns the parameters and return types of a function and the package that
// defines it.
func (t *setTypesPass) getFunc(namespace, name string) ([]*ast.Param, []common.Type, *ast.Package, error) {
	pkg := t.file.Package

	if namespace != "" {
		pkg = t.file.Imports[namespace]
		if pkg == nil {
			return nil, nil, nil, fmt.Errorf("namespace %v not found", namespace)
		}
	}

	for _, fd := range pkg.FuncDefs {
		if fd.Name == name {
			return fd.Params, fd.Returns, pkg, nil
		}
	}

	for _, fd := range pkg.ExternalFuncDefs {
		if fd.Name == name {
			return fd.Params, fd.Returns, pkg, nil
		}
	}

	if namespace == "" {
		return nil, nil, nil, fmt.Errorf("function %v not found", name)
	}

	return nil, nil, nil, fmt.Errorf("function %v in namespace %v not found", name, namespace)
}

// builtin returns the return types of a builtin function and an error if the arguments
// don't fit.
func builtin(name string, args []common.Type) ([]common.Type, bool, error) {
	switch name {
	case "println":
		return []common.Type{}, true, noArrays(name, args)

	case "len":
		tps := []common.Type{intType}
		if len(args) != 1 {
			return tps, true, fmt.Errorf("len takes 1 argument, not %v", len(args))
		}

		if !args[0].IsArray() && !args[0].Is(stringType) {
			return tps, true, fmt.Errorf("argument 1 of len must be an array or string, not %v", args[0].Print())
		}

		return tps, true, nil

	case "execVar":
		tps := []common.Type{stringArrayType, stringArrayType, intType}
		if len(args) == 0 {
			return tps, true, fmt.Errorf("execVar takes at least 1 argument, not 0")
		}

		return tps, true, noArrays(name, args)

	case "append":
		if len(args) == 0 {
			return unknownTypes, true, fmt.Errorf("append takes at least 1 argument, not 0")
		}

		array, isArray := args[0].(*common.ArrayType)
		if !isArray {
			return unknownTypes, true, fmt.Errorf("argument 1 of append must be an array, not %v", args[0].Print())
		}

		for i, arg := range args[1:] {
			if !arg.Is(array.ElementType) {
				return args[:1], true, fmt.Errorf("argument %v of append must be of type %v, not %v", i+2, array.ElementType.Print(), arg.Print())
			}
		}

		return args[:1], true, nil
	}

	return nil, false, nil
}

// noArrays is for builtins whose arguments go to a command, an array would show up as
// its length followed by its elements in println and be a syntax error in execVar.
func noArrays(name string, args []common.Type) error {
	for i, arg := range args {
		if arg.IsArray() {
			return fmt.Errorf("argument %v of %v can't be of type %v", i+1, name, arg.Print())
		}
	}

	return nil
}
