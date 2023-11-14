package bcode

import (
	"strconv"
	"strings"

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

func (e *EnvVarValue) Print() string {
	return "$" + e.Name
}

type NumberValue struct {
	Value int
}

func (n *NumberValue) Type() common.Type {
	return &common.BasicType{Type: common.Int}
}

func (n *NumberValue) Print() string {
	return strconv.Itoa(n.Value)
}

type StringValue struct {
	Value string
}

func (s *StringValue) Type() common.Type {
	return &common.BasicType{Type: common.String}
}

func (s *StringValue) Print() string {
	return "\"" + s.Value + "\""
}

type BoolValue struct {
	Value bool
}

func (b *BoolValue) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

func (b *BoolValue) Print() string {
	return strconv.FormatBool(b.Value)
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

func (a *ArrayValue) Print() string {
	ps := []string{}
	for _, p := range a.Values {
		ps = append(ps, p.Print())
	}

	return "[" + strings.Join(ps, ", ") + "]"
}

type MultiValue struct {
	Values []Value
}
