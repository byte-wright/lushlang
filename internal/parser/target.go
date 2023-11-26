package parser

import "github.com/byte-wright/lush/internal/ast"

type mainTarget struct {
	prog *ast.Program
}

func (m *mainTarget) addImport(path string) {
	m.prog.AddImport(path)
}

func (m *mainTarget) addMainStatement(s ast.Statement) {
	m.prog.Root.Statements = append(m.prog.Root.Statements, s)
}

func (m *mainTarget) addFuncDef(f *ast.FuncDef) {
	m.prog.FuncDefs = append(m.prog.FuncDefs, f)
}

func (m *mainTarget) addExternalFuncDef(f *ast.ExternalFuncDef) {
	panic("external func devs only in packages allowed")
}

type pkgTarget struct {
	prog *ast.Program
	pkg  *ast.Package
}

func (l *pkgTarget) addImport(path string) {
	l.prog.AddImport(path)
}

func (l *pkgTarget) addMainStatement(s ast.Statement) {
	panic("package must not have main statements")
}

func (l *pkgTarget) addFuncDef(f *ast.FuncDef) {
	l.pkg.FuncDefs = append(l.pkg.FuncDefs, f)
}

func (l *pkgTarget) addExternalFuncDef(f *ast.ExternalFuncDef) {
	l.pkg.ExternalFuncDefs = append(l.pkg.ExternalFuncDefs, f)
}
