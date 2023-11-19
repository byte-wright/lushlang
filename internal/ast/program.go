package ast

import (
	"path/filepath"
	"strings"
)

type Program struct {
	Root          *Block
	packages      []*Package
	addedPackages map[string]bool
}

type Package struct {
	Path             string
	Name             string
	FuncDefs         []*FuncDef
	ExternalFuncDefs []*ExternalFuncDef
}

func NewProgram() *Program {
	return &Program{
		Root:          &Block{},
		addedPackages: map[string]bool{},
	}
}

func (p *Program) Print() string {
	lines := []string{}

	for _, imp := range p.packages {
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
	if p.addedPackages[path] {
		return
	}
	p.packages = append(p.packages, &Package{
		Path: path,
		Name: filepath.Base(path),
	})

	p.addedPackages[path] = true
}

func (p *Program) Packages() []*Package {
	return p.packages
}
