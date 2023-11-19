package ast

import (
	"path/filepath"
	"strings"
)

type Program struct {
	Root      *Block
	libs      []*Library
	addedLibs map[string]bool
}

type Library struct {
	Path             string
	Name             string
	FuncDefs         []*FuncDef
	ExternalFuncDefs []*ExternalFuncDef
}

func NewProgram() *Program {
	return &Program{
		Root:      &Block{},
		addedLibs: map[string]bool{},
	}
}

func (p *Program) Print() string {
	lines := []string{}

	for _, imp := range p.libs {
		lines = append(lines, "import \""+imp.Path+"\"")
	}

	if len(lines) > 0 {
		lines = append(lines, "")
	}

	for _, fd := range p.Root.FuncDefs {
		sls := fd.print()
		lines = append(lines, sls...)
		lines = append(lines, "")
	}

	for _, st := range p.Root.Statements {
		ls := st.print()
		lines = append(lines, ls...)
	}

	return strings.Join(lines, "\n")
}

func (p *Program) AddImport(path string) {
	if p.addedLibs[path] {
		return
	}
	p.libs = append(p.libs, &Library{
		Path: path,
		Name: filepath.Base(path),
	})

	p.addedLibs[path] = true
}

func (p *Program) Libraries() []*Library {
	return p.libs
}
