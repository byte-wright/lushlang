package bcode

import "github.com/byte-wright/lush/internal/common"

type VarValue struct {
	Name string
	T    common.Type
}

func (v *VarValue) Type() common.Type {
	return v.T
}

func (v *VarValue) Print() string {
	return v.Name + ":" + v.Type().Print()
}
