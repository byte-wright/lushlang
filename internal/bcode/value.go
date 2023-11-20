package bcode

import (
	"github.com/byte-wright/lush/internal/common"
)

type Value interface {
	Atom
}

type EnvVarValue struct {
	Name string
}

func (e *EnvVarValue) Type() common.Type {
	return &common.BasicType{Type: common.String}
}

type NumberValue struct {
	Value int
}

func (n *NumberValue) Type() common.Type {
	return &common.BasicType{Type: common.Int}
}

type StringValue struct {
	Value string
}

func (s *StringValue) Type() common.Type {
	return &common.BasicType{Type: common.String}
}

type BoolValue struct {
	Value bool
}

func (b *BoolValue) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

type ArrayValue struct {
	Values []Atom
	Tp     common.PrimitiveType
}

func (a *ArrayValue) Type() common.Type {
	if a.Tp != common.None {
		return &common.ArrayType{ElementType: &common.BasicType{Type: a.Tp}}
	}

	return &common.ArrayType{ElementType: a.Values[0].Type()}
}

type MultiValue struct {
	Values []Value
}
