package bcode

import (
	"github.com/byte-wright/lush/internal/common"
)

type Func struct {
	Namespace  string
	Name       string
	Parameters []Value
	Return     []*VarValue
}

func (f *Func) Type() common.Type {
	return &common.BasicType{Type: common.String}
}

func (f *Func) FullName() string {
	return funcName(f.Namespace, f.Name)
}

func funcName(namespace, name string) string {
	if namespace == "" {
		return name
	}

	return namespace + "__" + name
}
