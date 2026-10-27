package parser

import (
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type mainTarget struct {
	prog *ast.Program
}

func (m *mainTarget) addImport(path string, loc *common.Location) error {
	return m.prog.File.Import(m.prog.AddImport(path, loc))
}

func (m *mainTarget) addMainStatement(s ast.Statement) {
	m.prog.Root.Statements = append(m.prog.Root.Statements, s)
}

func (m *mainTarget) addFuncDef(f *ast.FuncDef) {
	f.File = m.prog.File
	m.prog.Main.FuncDefs = append(m.prog.Main.FuncDefs, f)
}

func (m *mainTarget) addExternalFuncDef(f *ast.ExternalFuncDef) {
	panic("external func devs only in packages allowed")
}

type pkgTarget struct {
	prog *ast.Program
	file *ast.File
}

func (l *pkgTarget) addImport(path string, loc *common.Location) error {
	return l.file.Import(l.prog.AddImport(path, loc))
}

func (l *pkgTarget) addMainStatement(s ast.Statement) {
	panic("package must not have main statements")
}

func (l *pkgTarget) addFuncDef(f *ast.FuncDef) {
	f.File = l.file
	l.file.Package.FuncDefs = append(l.file.Package.FuncDefs, f)
}

func (l *pkgTarget) addExternalFuncDef(f *ast.ExternalFuncDef) {
	l.file.Package.ExternalFuncDefs = append(l.file.Package.ExternalFuncDefs, f)
}
