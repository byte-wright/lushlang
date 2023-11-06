package ast

import (
	"strings"

	"github.com/byte-wright/lush/internal/common"
)

type Param struct {
	Name string
	Type common.Type
}

func (p *Param) print() string {
	return p.Name + " " + p.Type.Print()
}

type FuncDef struct {
	Params []*Param
	Name   string
	Body   *Block
}

func (f *FuncDef) print() []string {
	ls := []string{}
	params := []string{}

	for _, p := range f.Params {
		params = append(params, p.print())
	}
	ls = append(ls, "func "+f.Name+"("+strings.Join(params, ", ")+") {")

	for _, l := range f.Body.printBare() {
		ls = append(ls, "  "+l)
	}

	ls = append(ls, "}")

	return ls
}
