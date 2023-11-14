package bcode

import (
	"strings"

	"github.com/byte-wright/lush/internal/common"
)

type Func struct {
	Name       string
	Parameters []Value
	Return     []*VarValue
}

func (f *Func) Type() common.Type {
	if f.Name == "append" {
		return f.Parameters[0].Type()
	}

	return &common.BasicType{Type: common.String}
}

func (f *Func) Print() string {
	ps := []string{}
	for _, p := range f.Parameters {
		ps = append(ps, p.Print())
	}

	assign := ""
	if len(f.Return) > 0 {
		rps := []string{}
		for _, r := range f.Return {
			rps = append(rps, r.Print())
		}
		assign = strings.Join(rps, ", ") + " = "
	}

	return assign + f.Name + "(" + strings.Join(ps, ", ") + ")"
}
