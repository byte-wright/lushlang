package astpass

import (
	"fmt"
	"slices"
	"strings"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type nodeKind int

const (
	funcNode nodeKind = iota
	// the globals of packages get their values before main runs, sorted by their references
	defNode
	// the globals of main get their values at the top level, in source order
	mainGlobalNode
	statementNode
)

type initNode struct {
	kind  nodeKind
	label string
	def   *ast.GlobalDef
	// assignedAt is the index of the top level statement that assigns a global of main
	// first, assignment is its location
	assignedAt    int
	assignment    *common.Location
	refs          []initRef
	sorted        bool
	cycleReported bool
}

type initRef struct {
	node     *initNode
	location *common.Location
}

type initOrderPass struct {
	prog     *ast.Program
	errors   []*ast.ASTError
	vars     map[*ast.Package]map[string]*initNode
	funcs    map[*ast.Package]map[string]*initNode
	defNodes []*initNode
	// statements are the top level statements of main, in source order
	statements []*initNode
	file       *ast.File
	current    *initNode
	// root is the top level statement that is visited, it's nil in functions
	root ast.Statement
	path []*initNode
}

func SetInitOrder(prog *ast.Program) []*ast.ASTError {
	p := &initOrderPass{
		prog:  prog,
		vars:  map[*ast.Package]map[string]*initNode{},
		funcs: map[*ast.Package]map[string]*initNode{},
	}

	for _, pkg := range prog.Packages() {
		p.declare(pkg)
	}
	p.declareMain()

	p.collectReferences()

	// a local named like a global would show up as a false cycle or early use
	if len(p.errors) > 0 {
		return p.errors
	}

	for _, node := range p.defNodes {
		p.sort(node)
	}

	p.checkMain()

	return p.errors
}

func (p *initOrderPass) declare(pkg *ast.Package) {
	p.vars[pkg] = map[string]*initNode{}
	defined := map[string]*common.Location{}

	for _, def := range pkg.GlobalDefs {
		node := &initNode{
			kind: defNode,
			def:  def,
			label: strings.Join(common.Map(def.Assignment.Targets, func(t *ast.AssignmentTarget) string {
				return p.qualified(pkg, t.Name)
			}), ", "),
		}
		p.defNodes = append(p.defNodes, node)

		for _, target := range def.Assignment.Targets {
			target.GlobalOf = pkg
			first := defined[target.Name]

			switch {
			case target.Name == "_":

			case first != nil:
				p.errorf(target.Location, "global variable %v is already defined at %v", target.Name, first)

			default:
				defined[target.Name] = target.Location
				p.vars[pkg][target.Name] = node
			}
		}
	}

	p.declareFuncs(pkg)
}

// declareMain declares a global of main where the top level assigns it first. A global
// assignment in a block doesn't declare, the block might not run.
func (p *initOrderPass) declareMain() {
	main := p.prog.Main
	p.vars[main] = map[string]*initNode{}

	for i, statement := range p.prog.Root.Statements {
		assignment, isAssignment := statement.(*ast.Assignment)
		if !isAssignment || !assignment.Global {
			continue
		}

		for _, target := range assignment.Targets {
			if target.Name != "_" && p.vars[main][target.Name] == nil {
				p.vars[main][target.Name] = &initNode{
					kind:       mainGlobalNode,
					label:      target.Name,
					assignedAt: i,
					assignment: target.Location,
				}
			}
		}
	}

	p.declareFuncs(main)
}

func (p *initOrderPass) declareFuncs(pkg *ast.Package) {
	p.funcs[pkg] = map[string]*initNode{}

	for _, fd := range pkg.FuncDefs {
		p.funcs[pkg][fd.Name] = &initNode{kind: funcNode, label: p.qualified(pkg, fd.Name)}
	}
}

func (p *initOrderPass) qualified(pkg *ast.Package, name string) string {
	if pkg == p.prog.Main {
		return name
	}

	return pkg.Name + "." + name
}

func (p *initOrderPass) collectReferences() {
	c := &context{ast: p.prog, visitor: p}

	for _, node := range p.defNodes {
		p.file, p.current = node.def.File, node

		for _, exp := range node.def.Assignment.Expressions {
			c.evalExpression(exp)
		}
	}

	p.file = p.prog.File
	for _, statement := range p.prog.Root.Statements {
		p.root, p.current = statement, &initNode{kind: statementNode}
		p.statements = append(p.statements, p.current)
		c.evalStatement(statement)
	}
	p.root = nil

	for _, fd := range funcDefs(p.prog) {
		p.BeforeFuncDef(fd)
		c.evalBlock(fd.Body)
	}
}

func (p *initOrderPass) BeforeFuncDef(fd *ast.FuncDef) {
	pkg := fd.File.Package
	p.file, p.current = fd.File, p.funcs[pkg][fd.Name]

	for _, param := range fd.Params {
		if p.vars[pkg][param.Name] != nil {
			p.errorf(param.Location, "parameter %v has the name of a global variable", param.Name)
		}
	}
}

func (p *initOrderPass) BeforeStatement(ast.Statement) {}

func (p *initOrderPass) BeforeLoopBody(st *ast.ForIn) {
	for _, target := range []*ast.AssignmentTarget{st.Value, st.Index} {
		if target != nil && p.vars[p.file.Package][target.Name] != nil {
			p.errorf(target.Location, "loop variable %v has the name of a global variable", target.Name)
		}
	}
}

func (p *initOrderPass) BeforeCatch(catch *ast.Catch) {
	if p.vars[p.file.Package][catch.Error.Name] != nil {
		p.errorf(catch.Error.Location, "catch variable %v has the name of a global variable", catch.Error.Name)
	}
}

func (p *initOrderPass) AfterCatch(*ast.Catch) {}

func (p *initOrderPass) BeforeBlock(*ast.Block) {}

func (p *initOrderPass) AfterBlock(*ast.Block) {}

func (p *initOrderPass) Expression(expression ast.Expression) {
	switch exp := expression.(type) {
	case *ast.Var:
		pkg := p.namespace(exp.Namespace)

		node := p.vars[pkg][exp.Name]
		if node != nil {
			exp.GlobalOf = pkg
			p.refer(node, exp.Location)
		}

	case *ast.Func:
		p.refer(p.funcs[p.file.Package][exp.Name], exp.Location)

	case *ast.Method:
		p.refer(p.funcs[p.namespace(exp.Namespace)][exp.Func.Name], exp.Func.Location)
	}
}

func (p *initOrderPass) namespace(name string) *ast.Package {
	if name == "" {
		return p.file.Package
	}

	return p.file.Imports[name]
}

func (p *initOrderPass) Statement(statement ast.Statement) {
	assignment, isAssignment := statement.(*ast.Assignment)
	if !isAssignment {
		return
	}

	pkg := p.file.Package

	for _, target := range assignment.Targets {
		node := p.vars[pkg][target.Name]

		switch {
		case target.Name == "_":

		case assignment.Global && node == nil:
			p.errorf(target.Location, "global variable %v is not defined", target.Name)

		case assignment.Global:
			target.GlobalOf = pkg
			// a write counts as a use: a package's initial value would overwrite it, and in
			// main it would come before the definition. A top level write of main is never
			// early, it's the definition or after it.
			if statement != p.root {
				p.refer(node, target.Location)
			}

		case node != nil:
			p.errorf(target.Location, "%v is a global variable, write it with global %v = ... or use another name", target.Name, target.Name)
		}
	}
}

func (p *initOrderPass) refer(node *initNode, loc *common.Location) {
	if node == nil || slices.ContainsFunc(p.current.refs, func(r initRef) bool { return r.node == node }) {
		return
	}

	p.current.refs = append(p.current.refs, initRef{node: node, location: loc})
}

func (p *initOrderPass) sort(node *initNode) {
	if node.sorted {
		return
	}

	start := slices.Index(p.path, node)
	if start >= 0 {
		p.reportCycle(node, p.path[start:])
		return
	}

	p.path = append(p.path, node)
	// a function already followed for another variable can still close a cycle for this
	// one, so every variable starts with nothing seen
	p.follow(node, map[*initNode]bool{})
	p.path = p.path[:len(p.path)-1]

	node.sorted = true
	p.prog.InitOrder = append(p.prog.InitOrder, node.def)
}

func (p *initOrderPass) follow(node *initNode, seen map[*initNode]bool) {
	for _, ref := range node.refs {
		switch {
		case ref.node.kind == defNode:
			p.sort(ref.node)

		case ref.node.kind == funcNode && !seen[ref.node]:
			seen[ref.node] = true
			p.path = append(p.path, ref.node)
			p.follow(ref.node, seen)
			p.path = p.path[:len(p.path)-1]
		}
	}
}

func (p *initOrderPass) reportCycle(node *initNode, cycle []*initNode) {
	if node.cycleReported {
		return
	}

	node.cycleReported = true

	labels := common.Map(append(slices.Clone(cycle), node), func(n *initNode) string { return n.label })
	p.errorf(node.def.Assignment.Location, "initialization cycle: %v", strings.Join(labels, " -> "))
}

// checkMain reports the globals of main that a top level statement uses before they are
// assigned, also through the functions it calls. Like a cycle in a package, a use counts
// even in a branch that never runs.
func (p *initOrderPass) checkMain() {
	for i, statement := range p.statements {
		// one error per global and statement is enough
		reported := map[*initNode]bool{}

		for _, ref := range statement.refs {
			switch {
			case ref.node.kind == mainGlobalNode && ref.node.assignedAt >= i && !reported[ref.node]:
				reported[ref.node] = true
				p.errorf(ref.location, "%v is used before it is assigned at %v", ref.node.label, ref.node.assignment)

			case ref.node.kind == funcNode:
				p.checkCall(ref.location, i, reported, []*initNode{ref.node}, map[*initNode]bool{ref.node: true})
			}
		}
	}
}

func (p *initOrderPass) checkCall(call *common.Location, statement int, reported map[*initNode]bool, path []*initNode, seen map[*initNode]bool) {
	for _, ref := range path[len(path)-1].refs {
		switch {
		case ref.node.kind == mainGlobalNode && ref.node.assignedAt >= statement && !reported[ref.node]:
			reported[ref.node] = true
			labels := common.Map(path, func(n *initNode) string { return n.label })
			p.errorf(call, "%v uses %v before it is assigned at %v", strings.Join(labels, " -> "), ref.node.label, ref.node.assignment)

		case ref.node.kind == funcNode && !seen[ref.node]:
			seen[ref.node] = true
			p.checkCall(call, statement, reported, append(path, ref.node), seen)
		}
	}
}

func (p *initOrderPass) errorf(loc *common.Location, format string, args ...any) {
	p.errors = append(p.errors, &ast.ASTError{Location: loc, Message: fmt.Sprintf(format, args...)})
}
