package ast

import (
	"path/filepath"
)

type Program struct {
	Root     *Block
	FuncDefs []*FuncDef

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
