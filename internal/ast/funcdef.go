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
	Name    string
	Params  []*Param
	Returns []common.Type
	Body    *Block
}

func (f *FuncDef) print() []string {
	params := []string{}
	for _, p := range f.Params {
		params = append(params, p.print())
	}

	retParams := []string{}
	for _, ret := range f.Returns {
		retParams = append(retParams, ret.Print())
	}

	retParam := ""
	if len(retParams) > 0 {
		retParam = " " + strings.Join(retParams, ", ")
	}

	ls := []string{}
	ls = append(ls, "func "+f.Name+"("+strings.Join(params, ", ")+")"+retParam+" {")

	for _, l := range f.Body.printBare() {
		ls = append(ls, "  "+l)
	}

	ls = append(ls, "}")

	return ls
}
