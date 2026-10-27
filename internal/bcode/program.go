package bcode

import (
	"github.com/byte-wright/lush/internal/common"
)

type Program struct {
	Main   *Block
	Funcs  []*FuncDef
	varIdx int
}

func (p *Program) nextVar() int {
	p.varIdx++
	return p.varIdx
}

type FuncDef struct {
	Namespace string
	Name      string
	Params    []*Param
	Body      *Block
}

func (fd *FuncDef) FullName() string {
	return funcName(fd.Namespace, fd.Name)
}

type Param struct {
	Name string
	Type common.Type
}
