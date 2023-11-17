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
	m.prog.Root.FuncDefs = append(m.prog.Root.FuncDefs, f)
}

type libTarget struct {
	prog *ast.Program
	lib  *ast.Library
}

func (l *libTarget) addImport(path string) {
	l.prog.AddImport(path)
}

func (l *libTarget) addMainStatement(s ast.Statement) {
	panic("lib must not have main statements")
}

func (l *libTarget) addFuncDef(f *ast.FuncDef) {
	l.lib.FuncDefs = append(l.lib.FuncDefs, f)
}
