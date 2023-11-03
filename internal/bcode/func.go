package bcode

import (
	"strings"

	"github.com/byte-wright/lush/internal/common"
)

type Func struct {
	Name       string
	Parameters []Value
}

func (f *Func) Type() Type {
	if f.Name == "append" {
		return f.Parameters[0].Type()
	}

	return &BasicType{Type: common.String}
}

func (f *Func) Print() string {
	ps := []string{}
	for _, p := range f.Parameters {
		ps = append(ps, p.Print())
	}

	return f.Name + "(" + strings.Join(ps, ", ") + ")"
}
