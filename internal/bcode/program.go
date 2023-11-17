package bcode

import "github.com/byte-wright/lush/internal/common"

type Program struct {
	Main  *Block
	Funcs []*FuncDef
	// Funcs by name contains teh same funcs like the Funcs slice
	FuncsByName map[string]*FuncDef
	varIdx      int
}

func (p *Program) nextVar() int {
	p.varIdx++
	return p.varIdx
}

func (p *Program) Print() string {
	str := ""
	for _, fd := range p.Funcs {
		str += "func " + fd.FullName() + "() {\n"
		str += fd.Body.Print(2)
		str += "}\n"
	}

	return str + p.Main.Print(0)
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
