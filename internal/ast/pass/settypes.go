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
	statementCmd  *ast.Cmd
	// catchCall is the *ast.Func or *ast.Cmd whose error the current statement catches
	catchCall ast.Expression
	// pipeOver is the value of the current for in, the only place that can read a pipe
	pipeOver *ast.Var
	// pkgGlobal is set while checking the definition of a package's global
	pkgGlobal bool
	// loops is the number of loops around the current statement
	loops int
}

func SetTypes(ast *ast.Program) []*ast.ASTError {
	te := &setTypesPass{
		ast: ast,
	}

	te.declarations(ast.Main)
	for _, pkg := range ast.Packages() {
		te.declarations(pkg)
	}

	te.globalDefs()

	// Root is visited before the function bodies
	te.file = ast.File

	Visit(ast, te)

	return te.errors
}

func (t *setTypesPass) globalDefs() {
	c := &context{ast: t.ast, visitor: t}

	t.pkgGlobal = true
	for _, def := range t.ast.InitOrder {
		t.file, t.block = def.File, def.File.Package.Globals
		c.evalStatement(def.Assignment)
	}
	t.pkgGlobal = false
}

// declarations reports functions and parameters that are defined twice. Calls would be
// checked against the first one, but in bash the last one wins. An external function is
// defined once for each target, calls are checked against the first definition, so the
// others must have its types.
func (t *setTypesPass) declarations(pkg *ast.Package) {
	defined := map[string]*common.Location{}

	for _, fd := range pkg.FuncDefs {
		t.declare(defined, fd.Name, fd.Location)
		t.params(fd.Params)
	}

	type forTarget struct{ name, target string }
	targets := map[forTarget]*common.Location{}
	firsts := map[string]*ast.ExternalFuncDef{}

	for _, fd := range pkg.ExternalFuncDefs {
		t.params(fd.Params)

		key := forTarget{fd.Name, fd.Target}
		other := targets[key]
		if other != nil {
			t.errorf(fd.Location, "function %v for %v is already defined at %v", fd.Name, fd.Target, other)
			continue
		}
		targets[key] = fd.Location

		first := firsts[fd.Name]
		if first == nil {
			firsts[fd.Name] = fd
			t.declare(defined, fd.Name, fd.Location)
		} else if !sameSignature(first, fd) {
			t.errorf(fd.Location, "function %v for %v has another signature than at %v", fd.Name, fd.Target, first.Location)
		}
	}
}

func (t *setTypesPass) declare(defined map[string]*common.Location, name string, loc *common.Location) {
	first := defined[name]
	if first != nil {
		t.errorf(loc, "function %v is already defined at %v", name, first)
	} else {
		defined[name] = loc
	}
}

// the names of the parameters can differ, only the code sees them
func sameSignature(a, b *ast.ExternalFuncDef) bool {
	if len(a.Params) != len(b.Params) || len(a.Returns) != len(b.Returns) || a.CanFail != b.CanFail {
		return false
	}

	for i, p := range a.Params {
		if !p.Type.Is(b.Params[i].Type) {
			return false
		}
	}

	for i, r := range a.Returns {
		if !r.Is(b.Returns[i]) {
			return false
		}
	}

	return true
}

func (t *setTypesPass) params(params []*ast.Param) {
	seen := map[string]bool{}
	for _, p := range params {
		if p.Name != "_" && seen[p.Name] {
			t.errorf(p.Location, "parameter %v is already defined", p.Name)
		}
		seen[p.Name] = true
	}
}

func (t *setTypesPass) BeforeFuncDef(fd *ast.FuncDef) {
	t.file = fd.File
	t.funcDef = fd

	if len(fd.Returns) > 0 && !terminates(fd.Body, false) {
		t.errorf(fd.Body.End, "missing return at the end of function %v", fd.Name)
	}

	for _, p := range fd.Params {
		// a duplicate is already reported by declarations
		if p.Name == "_" || fd.Body.GetVar(p.Name) != nil {
			continue
		}

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
	t.statementCall, t.statementCmd, t.pipeOver = nil, nil, nil
	t.catchCall = catchCall(statement)

	switch st := statement.(type) {
	case *ast.For:
		t.loops++
	case *ast.ForIn:
		t.loops++
		t.pipeOver, _ = st.Over.(*ast.Var)
	}

	st, isCall := statement.(*ast.FuncStatement)
	if !isCall {
		return
	}

	switch call := st.Call.(type) {
	case *ast.Func:
		t.statementCall = call
	case *ast.Method:
		t.statementCall = call.Func
	case *ast.Cmd:
		t.statementCmd = call
	}
}

func (t *setTypesPass) Statement(statement ast.Statement) {
	switch st := statement.(type) {
	case *ast.Assignment:
		if st.Catch != nil {
			t.catchAssignment(st)
		}
		t.assignment(st)

	case *ast.FuncStatement:
		if st.Catch != nil {
			t.catch(st, st.Catch)
		}

	case *ast.Fail:
		t.fail(st)

	case *ast.If:
		t.expectType(st.Location, "if condition", st.Condition, boolType)

	case *ast.For:
		t.loops--
		if st.Condition != nil {
			t.expectType(st.Location, "for condition", st.Condition, boolType)
		}

	case *ast.ForIn:
		t.loops--

	case *ast.Break:
		if t.loops == 0 {
			t.errorf(st.Location, "break is not in a loop")
		}

	case *ast.Continue:
		if t.loops == 0 {
			t.errorf(st.Location, "continue is not in a loop")
		}

	case *ast.ReturnStatement:
		t.returnStatement(st)
	}
}

// catchCall returns the *ast.Func or *ast.Cmd whose error the statement catches. It's nil
// without a catch, and when the statement has no single call to catch.
func catchCall(statement ast.Statement) ast.Expression {
	var call ast.Expression

	switch st := statement.(type) {
	case *ast.Assignment:
		if st.Catch == nil || len(st.Expressions) != 1 {
			return nil
		}
		call = st.Expressions[0]

	case *ast.FuncStatement:
		if st.Catch == nil {
			return nil
		}
		call = st.Call
	}

	switch f := call.(type) {
	case *ast.Func:
		return f
	case *ast.Method:
		return f.Func
	case *ast.Cmd:
		return f
	}

	return nil
}

// callInfo returns what the checks of errors need from a *ast.Func or *ast.Cmd
func callInfo(call ast.Expression) (types []common.Type, try, canFail bool) {
	if c, isCmd := call.(*ast.Cmd); isCmd {
		return c.ReturnTypes, c.Try, c.CanFail
	}

	f := call.(*ast.Func)
	return f.ReturnTypes, f.Try, f.CanFail
}

func (t *setTypesPass) BeforeCatch(catch *ast.Catch) {
	t.loopVar(catch.Block, catch.Error, stringType, "catch")
}

func (t *setTypesPass) AfterCatch(catch *ast.Catch) {
	t.block = catch.Block.Parent
}

// catch reports a catch without a call that can fail. It returns the call, nil when
// there is nothing to check.
func (t *setTypesPass) catch(statement ast.Statement, catch *ast.Catch) ast.Expression {
	call := catchCall(statement)
	if call == nil {
		t.errorf(catch.Location, "catch needs a single call as the value, like x = f() catch err { ... }")
		return nil
	}

	types, try, canFail := callInfo(call)
	if !known(types) || try {
		return nil
	}

	if !canFail {
		t.errorf(catch.Location, "%v can't fail, there is no error to catch", callName(statement))
		return nil
	}

	return call
}

// catchAssignment reports variables that would have no value after a failed call: on a
// failure, the assignment doesn't happen.
func (t *setTypesPass) catchAssignment(st *ast.Assignment) {
	call := t.catch(st, st.Catch)
	if call == nil || terminates(st.Catch.Block, true) {
		return
	}

	for _, target := range st.Targets {
		block := t.block
		if target.GlobalOf != nil {
			block = target.GlobalOf.Globals
		}

		// an element target needs an existing array, its index check reports a missing one
		if target.Name != "_" && target.Element == nil && block.GetVar(target.Name) == nil {
			t.errorf(target.Location, "%v isn't assigned when %v fails, assign it before or end the catch block with return or fail", target.Name, callName(st))
		}
	}
}

// callName is the name of the call that catchCall returns, as it's written
func callName(statement ast.Statement) string {
	var call ast.Expression

	switch st := statement.(type) {
	case *ast.Assignment:
		call = st.Expressions[0]
	case *ast.FuncStatement:
		call = st.Call
	}

	switch c := call.(type) {
	case *ast.Method:
		return c.Namespace + "." + c.Func.Name
	case *ast.Cmd:
		return "cmd"
	}

	return call.(*ast.Func).Name
}

func (t *setTypesPass) fail(st *ast.Fail) {
	t.expectType(st.Location, "fail message", st.Message, stringType)

	if t.funcDef != nil && !t.funcDef.CanFail {
		t.errorf(st.Location, "fail in function %v, which can't fail: add error to its results", t.funcDef.Name)
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
		if target.Element != nil {
			t.assignElement(target, types[i])
			continue
		}

		block := t.block
		if target.GlobalOf != nil {
			block = target.GlobalOf.Globals
		}

		t.assignTarget(block, target, types[i])
	}
}

// assignElement checks the value of x[i] = value, the index check of the target already
// reports an x that can't be indexed
func (t *setTypesPass) assignElement(target *ast.AssignmentTarget, tp common.Type) {
	array := target.Element.Value.Types()
	if !known(array) || !known([]common.Type{tp}) {
		return
	}

	if array[0].Is(stringType) {
		t.errorf(target.Element.Location, "cannot assign to a character of string '%v', strings can't be changed", target.Name)
		return
	}

	el := target.Element.Types()[0]
	if array[0].IsArray() && !tp.Is(el) {
		t.errorf(target.Element.Location, "an element of '%v' must be of type %v, not %v", target.Name, el.Print(), tp.Print())
	}
}

func (t *setTypesPass) assignTarget(block *ast.Block, target *ast.AssignmentTarget, tp common.Type) {
	if target.Name == "_" {
		return
	}

	v := block.GetVar(target.Name)
	if v != nil && !known([]common.Type{v.Type, tp}) {
		return
	}

	err := block.SetVar(target.Name, tp)
	if err != nil {
		t.errors = append(t.errors, &ast.ASTError{Location: target.Location, Message: err.Error()})
	}
}

func (t *setTypesPass) BeforeLoopBody(st *ast.ForIn) {
	t.loopVar(st.Body, st.Value, t.loopValue(st), "loop")

	if st.Index != nil {
		t.loopVar(st.Body, st.Index, intType, "loop")
	}
}

// loopVar declares a loop or catch variable in the body. Assigning a variable outside of
// it, or the one of an outer loop, would be easy to miss.
func (t *setTypesPass) loopVar(body *ast.Block, target *ast.AssignmentTarget, tp common.Type, kind string) {
	if target.Name == "_" {
		return
	}

	if body.GetVar(target.Name) != nil {
		t.errorf(target.Location, "%v variable %v is already defined", kind, target.Name)
	}

	// also after the error, so the body is checked with the loop's type
	body.Variables = append(body.Variables, &ast.Variable{Name: target.Name, Type: tp})
}

func (t *setTypesPass) loopValue(st *ast.ForIn) common.Type {
	if c, isCmd := st.Over.(*ast.Cmd); isCmd && c.Use == ast.UseLoop {
		return stringType
	}

	if st.From != nil {
		t.expectType(st.Location, "loop start", st.From, intType)
		t.expectType(st.Location, "loop end", st.Over, intType)
		t.noIndex(st)
		return intType
	}

	over := st.Over.Types()
	if !known(over) {
		return unknownTypes[0]
	}

	if len(over) == 1 && over[0].IsArray() {
		return over[0].(*common.ArrayType).ElementType
	}

	if len(over) == 1 && (over[0].Is(&common.PipeType{}) || over[0].Is(stringType)) {
		return stringType
	}

	if len(over) == 1 && over[0].Is(intType) {
		t.noIndex(st)
		return intType
	}

	t.errorf(st.Location, "for in needs an array, a string or an int, not %v", printTypes(over))
	return unknownTypes[0]
}

func (t *setTypesPass) noIndex(st *ast.ForIn) {
	if st.Index != nil {
		t.errorf(st.Index.Location, "only a loop over an array has an index")
	}
}

func (t *setTypesPass) Expression(expression ast.Expression) {
	switch exp := expression.(type) {
	case *ast.Func:
		t.call("", exp)

	case *ast.Cmd:
		t.cmd(exp)

	case *ast.Method:
		t.call(exp.Namespace, exp.Func)

	case *ast.Var:
		t.variable(exp)

	case *ast.Minus:
		t.unary(exp.Location, "-", exp.Expression, intType)
	case *ast.Not:
		t.unary(exp.Location, "!", exp.Expression, boolType)
	case *ast.Add:
		t.add(exp)
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

	case *ast.Array:
		t.array(exp)

	case *ast.EnvVar:
		t.envVar(exp)
	}
}

func (t *setTypesPass) variable(x *ast.Var) {
	switch {
	case x.Name == "_":
		t.errorf(x.Location, "cannot use _ as value")

	case x.GlobalOf != nil:
		x.Type = x.GlobalOf.Globals.GetVar(x.Name).Type

	case x.Namespace != "" && t.file.Imports[x.Namespace] == nil:
		t.errorf(x.Location, "namespace %v not found", x.Namespace)

	case x.Namespace != "":
		t.errorf(x.Location, "variable %v in namespace %v not found", x.Name, x.Namespace)

	default:
		v := t.block.GetVar(x.Name)
		if v == nil {
			t.errorf(x.Location, "variable '%v' is not defined", x.Name)
			return
		}

		x.Type = v.Type

		if x.Type.Is(&common.PipeType{}) && x != t.pipeOver {
			t.errorf(x.Location, "%v is a pipe, only for line in %v can read it", x.Name, x.Name)
			x.Type = unknownTypes[0]
		}
	}
}

func (t *setTypesPass) envVar(v *ast.EnvVar) {
	switch {
	case v.Name[0] >= '0' && v.Name[0] <= '9':
		t.errorf(v.Location, "$%v is not an environment variable, the script's arguments are in args()", v.Name)

	case v.Name == "_":
		t.errorf(v.Location, "$_ is not an environment variable")
	}
}

// array checks that the elements of a literal share one type that isn't an array, bash
// arrays can't nest.
func (t *setTypesPass) array(a *ast.Array) {
	if a.Type != common.None {
		return
	}

	if len(a.Values) == 0 {
		t.errorf(a.Location, "empty array needs an element type, like []int")
		return
	}

	elements := valueTypes(a.Values)
	if !known(elements) {
		return
	}

	for i, el := range elements {
		if el.IsArray() {
			t.errorf(a.Location, "array element %v can't be of type %v", i+1, el.Print())
			return
		}

		if !el.Is(elements[0]) {
			t.errorf(a.Location, "array element %v must be of type %v, not %v", i+1, elements[0].Print(), el.Print())
			return
		}
	}
}

func (t *setTypesPass) returnStatement(st *ast.ReturnStatement) {
	fd := t.funcDef
	// outside of functions it ends the script, like a return in Go's main
	if fd == nil {
		if len(st.Expressions) > 0 {
			t.errorf(st.Location, "return outside of a function can't have values")
		}
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
// statement counts. With jumps, a break or continue ends it too.
func terminates(b *ast.Block, jumps bool) bool {
	if len(b.Statements) == 0 {
		return false
	}

	switch st := b.Statements[len(b.Statements)-1].(type) {
	case *ast.ReturnStatement, *ast.Fail:
		return true

	case *ast.Break, *ast.Continue:
		return jumps

	case *ast.If:
		return st.Else != nil && terminates(st.If, jumps) && terminates(st.Else, jumps)

	case *ast.For:
		cond, isBool := st.Condition.(*ast.Bool)
		endless := st.Condition == nil || isBool && cond.Value
		return endless && !breaks(st.Body)
	}

	return false
}

// breaks reports whether a break leaves the loop of this body, a break in a nested loop
// only leaves that one
func breaks(b *ast.Block) bool {
	for _, statement := range b.Statements {
		switch st := statement.(type) {
		case *ast.Break:
			return true

		case *ast.If:
			if breaks(st.If) || st.Else != nil && breaks(st.Else) {
				return true
			}

		case *ast.Assignment:
			if st.Catch != nil && breaks(st.Catch.Block) || cmdBreaks(st.Expressions...) {
				return true
			}

		case *ast.FuncStatement:
			if st.Catch != nil && breaks(st.Catch.Block) || cmdBreaks(st.Call) {
				return true
			}
		}
	}

	return false
}

func (t *setTypesPass) call(namespace string, f *ast.Func) {
	name := f.Name
	if namespace != "" {
		name = namespace + "." + name
	}

	f.ReturnTypes = t.returnTypes(namespace, name, f)

	if known(f.ReturnTypes) {
		t.failingCall(name, f, f.Try, f.CanFail, f.Location)
	}

	// spreading would silently drop a call without values from the values around it
	if len(f.ReturnTypes) == 0 && f != t.statementCall {
		t.errorf(f.Location, "%v doesn't return a value", name)
		f.ReturnTypes = unknownTypes
	}
}

// failingCall checks that a call that can fail passes the error on with try or catches
// it, and that try is only used where the error can go on. call is a *ast.Func or *ast.Cmd.
func (t *setTypesPass) failingCall(name string, call ast.Expression, try, canFail bool, loc *common.Location) {
	switch {
	case try && call == t.catchCall:
		t.errorf(loc, "the error of %v can't go on with try and be caught, remove one", name)

	case canFail && t.pkgGlobal:
		t.errorf(loc, "%v can fail, a global of a package can't call it", name)

	case try && !canFail:
		t.errorf(loc, "%v can't fail, it needs no try", name)

	case try && t.funcDef != nil && !t.funcDef.CanFail:
		t.errorf(loc, "try in function %v, which can't fail: add error to its results or catch the error", t.funcDef.Name)

	case canFail && !try && call != t.catchCall:
		t.errorf(loc, "%v can fail, call it with try or catch the error", name)
	}
}

// returnTypes resolves the function and checks its arguments. Unqualified names are
// functions of the caller's own package or builtins.
func (t *setTypesPass) returnTypes(namespace, name string, f *ast.Func) []common.Type {
	args := valueTypes(f.Parameters)

	// like in Go, the functions of a package shadow the builtins inside of it
	isOwn := findFunc(t.file.Package, f.Name) != nil
	if namespace == "" && !isOwn {
		def, err := builtin(f.Name, args)
		if def != nil {
			if err != nil && known(args) {
				t.errorf(f.Location, "%v", err)
			}
			f.CanFail = def.canFail
			return def.returns
		}
	}

	def, pkg, err := t.getFunc(namespace, f.Name)
	if err != nil {
		t.errorf(f.Location, "%v", err)
		return unknownTypes
	}

	err = arguments(name, args, def.params)
	if err != nil && known(args) {
		t.errorf(f.Location, "%v", err)
	}

	f.Package = pkg
	f.CanFail = def.canFail

	return def.returns
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

// add lets a string on the left take an int or bool on the right as its text. Only the
// left side decides, so 1 + 2 + "x" is an error instead of "3x".
func (t *setTypesPass) add(exp *ast.Add) {
	at, bt := exp.A.Types(), exp.B.Types()
	if known(at) && known(bt) && len(at) == 1 && len(bt) == 1 && at[0].Is(stringType) && isOneOf(bt[0], []common.Type{intType, boolType}) {
		return
	}

	t.binary(exp.Location, "+", exp.A, exp.B, intType, stringType)
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

// signature is what a call needs to know of a function, of a FuncDef or ExternalFuncDef
type signature struct {
	params  []*ast.Param
	returns []common.Type
	canFail bool
}

// getFunc returns the signature of a function and the package that defines it.
func (t *setTypesPass) getFunc(namespace, name string) (*signature, *ast.Package, error) {
	pkg := t.file.Package

	if namespace != "" {
		pkg = t.file.Imports[namespace]
		if pkg == nil {
			return nil, nil, fmt.Errorf("namespace %v not found", namespace)
		}
	}

	def := findFunc(pkg, name)
	if def != nil {
		return def, pkg, nil
	}

	if namespace == "" {
		return nil, nil, fmt.Errorf("function %v not found", name)
	}

	return nil, nil, fmt.Errorf("function %v in namespace %v not found", name, namespace)
}

func findFunc(pkg *ast.Package, name string) *signature {
	for _, fd := range pkg.FuncDefs {
		if fd.Name == name {
			return &signature{params: fd.Params, returns: fd.Returns, canFail: fd.CanFail}
		}
	}

	for _, fd := range pkg.ExternalFuncDefs {
		if fd.Name == name {
			return &signature{params: fd.Params, returns: fd.Returns, canFail: fd.CanFail}
		}
	}

	return nil
}

// builtin returns the signature of a builtin function, nil if there is none with that name,
// and an error if the arguments don't fit. The signature has no params, the arguments are
// checked here.
func builtin(name string, args []common.Type) (*signature, error) {
	switch name {
	case "print", "println", "eprint", "eprintln":
		return &signature{returns: []common.Type{}}, noArrays(name, args)

	case "len":
		sig := &signature{returns: []common.Type{intType}}
		if len(args) != 1 {
			return sig, fmt.Errorf("len takes 1 argument, not %v", len(args))
		}

		if !args[0].IsArray() && !args[0].Is(stringType) {
			return sig, fmt.Errorf("argument 1 of len must be an array or string, not %v", args[0].Print())
		}

		return sig, nil

	case "parseInt":
		sig := &signature{returns: []common.Type{intType}, canFail: true}
		if len(args) != 1 {
			return sig, fmt.Errorf("parseInt takes 1 argument, not %v", len(args))
		}

		if !args[0].Is(stringType) {
			return sig, fmt.Errorf("argument 1 of parseInt must be of type string, not %v", args[0].Print())
		}

		return sig, nil

	case "args":
		sig := &signature{returns: []common.Type{stringArrayType}}
		if len(args) != 0 {
			return sig, fmt.Errorf("args takes %v, not %v", count(0, "argument"), len(args))
		}

		return sig, nil

	case "append":
		if len(args) == 0 {
			return &signature{returns: unknownTypes}, fmt.Errorf("append takes at least 1 argument, not 0")
		}

		array, isArray := args[0].(*common.ArrayType)
		if !isArray {
			return &signature{returns: unknownTypes}, fmt.Errorf("argument 1 of append must be an array, not %v", args[0].Print())
		}

		for i, arg := range args[1:] {
			if !arg.Is(array.ElementType) {
				return &signature{returns: args[:1]}, fmt.Errorf("argument %v of append must be of type %v, not %v", i+2, array.ElementType.Print(), arg.Print())
			}
		}

		return &signature{returns: args[:1]}, nil
	}

	return nil, nil
}

// noArrays is for the print functions, an array would show up as its length followed by its elements.
func noArrays(name string, args []common.Type) error {
	for i, arg := range args {
		if arg.IsArray() {
			return fmt.Errorf("argument %v of %v can't be of type %v", i+1, name, arg.Print())
		}
	}

	return nil
}
