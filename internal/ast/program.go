package ast

import (
	"path/filepath"
	"strings"
)

type Program struct {
	Root *Block
	Libs []*Library
}

type Library struct {
	Path     string
	Name     string
	FuncDefs []*FuncDef
}

func (p *Program) Print() string {
	lines := []string{}

	for _, imp := range p.Libs {
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
	p.Libs = append(p.Libs, &Library{
		Path: path,
		Name: filepath.Base(path),
	})
}
