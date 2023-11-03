package bcode

import "github.com/byte-wright/lush/internal/common"

type Type interface {
	IsArray() bool
	Print() string
}

type BasicType struct {
	Type common.PrimitiveType
}

func (b *BasicType) IsArray() bool {
	return false
}

func (b *BasicType) Print() string {
	switch b.Type {
	case common.String:
		return "string"
	case common.Int:
		return "int"
	case common.Bool:
		return "bool"
	}
	return "none"
}

type ArrayType struct {
	ElementType Type
}

func (a *ArrayType) IsArray() bool {
	return true
}

func (a *ArrayType) Print() string {
	return "[]" + a.ElementType.Print()
}

type VarValue struct {
	Name string
	T    Type
}

func (v *VarValue) Type() Type {
	return v.T
}

func (v *VarValue) Print() string {
	return v.Name + ":" + v.Type().Print()
}
