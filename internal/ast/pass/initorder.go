package astpass

import (
	"fmt"
	"slices"
	"strings"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type initNode struct {
	label         string
	varDef        *ast.VarDef
	refs          []*initNode
	sorted        bool
	cycleReported bool
}

func (n *initNode) isFunc() bool {
	return n.varDef == nil
}

type initOrderPass struct {
	prog     *ast.Program
	errors   []*ast.ASTError
	vars     map[*ast.Package]map[string]*initNode
	funcs    map[*ast.Package]map[string]*initNode
	varNodes []*initNode
	file     *ast.File
	current  *initNode
	path     []*initNode
}

func SetInitOrder(prog *ast.Program) []*ast.ASTError {
	p := &initOrderPass{
		prog:  prog,
		vars:  map[*ast.Package]map[string]*initNode{},
		funcs: map[*ast.Package]map[string]*initNode{},
	}

	for _, pkg := range slices.Concat(prog.Packages(), []*ast.Package{prog.Main}) {
		p.declare(pkg)
	}

	p.collectReferences()

	// a local named like a global would show up as a false cycle
	if len(p.errors) > 0 {
		return p.errors
	}

	for _, node := range p.varNodes {
		p.sort(node)
	}

	return p.errors
}

func (p *initOrderPass) declare(pkg *ast.Package) {
	p.vars[pkg] = map[string]*initNode{}
	p.funcs[pkg] = map[string]*initNode{}
	defined := map[string]*common.Location{}

	for _, vd := range pkg.VarDefs {
		node := &initNode{
			varDef: vd,
			label: strings.Join(common.Map(vd.Assignment.Targets, func(t *ast.AssignmentTarget) string {
				return p.qualified(pkg, t.Name)
			}), ", "),
		}
		p.varNodes = append(p.varNodes, node)

		for _, target := range vd.Assignment.Targets {
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

	for _, fd := range pkg.FuncDefs {
		p.funcs[pkg][fd.Name] = &initNode{label: p.qualified(pkg, fd.Name)}
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

	for _, node := range p.varNodes {
		p.file, p.current = node.varDef.File, node

		for _, exp := range node.varDef.Assignment.Expressions {
			c.evalExpression(exp)
		}
	}

	// the top level runs after all initial values, so nothing depends on its references
	p.file, p.current = p.prog.File, &initNode{}

	Visit(p.prog, p)
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

func (p *initOrderPass) BeforeBlock(*ast.Block) {}

func (p *initOrderPass) AfterBlock(*ast.Block) {}

func (p *initOrderPass) Expression(expression ast.Expression) {
	switch exp := expression.(type) {
	case *ast.Var:
		pkg := p.namespace(exp.Namespace)

		node := p.vars[pkg][exp.Name]
		if node != nil {
			exp.GlobalOf = pkg
			p.refer(node)
		}

	case *ast.Func:
		p.refer(p.funcs[p.file.Package][exp.Name])

	case *ast.Method:
		p.refer(p.funcs[p.namespace(exp.Namespace)][exp.Func.Name])
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
			// a write before the variable's own initial value would be overwritten by it
			p.refer(node)

		case node != nil:
			p.errorf(target.Location, "%v is a global variable, write it with global %v = ... or use another name", target.Name, target.Name)
		}
	}
}

func (p *initOrderPass) refer(node *initNode) {
	if node != nil && !slices.Contains(p.current.refs, node) {
		p.current.refs = append(p.current.refs, node)
	}
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
	p.prog.InitOrder = append(p.prog.InitOrder, node.varDef)
}

func (p *initOrderPass) follow(node *initNode, seen map[*initNode]bool) {
	for _, ref := range node.refs {
		switch {
		case !ref.isFunc():
			p.sort(ref)

		case !seen[ref]:
			seen[ref] = true
			p.path = append(p.path, ref)
			p.follow(ref, seen)
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
	p.errorf(node.varDef.Assignment.Location, "initialization cycle: %v", strings.Join(labels, " -> "))
}

func (p *initOrderPass) errorf(loc *common.Location, format string, args ...any) {
	p.errors = append(p.errors, &ast.ASTError{Location: loc, Message: fmt.Sprintf(format, args...)})
}
