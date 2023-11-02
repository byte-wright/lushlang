package bcode

import (
	"strconv"
	"strings"
)

type Value interface {
	Atom
}

type EnvVarValue struct {
	Name string
}

func (e *EnvVarValue) Type() Type {
	return &BasicType{Type: String}
}

func (e *EnvVarValue) Print() string {
	return "$" + e.Name
}

type NumberValue struct {
	Value int
}

func (n *NumberValue) Type() Type {
	return &BasicType{Type: Int}
}

func (n *NumberValue) Print() string {
	return strconv.Itoa(n.Value)
}

type StringValue struct {
	Value string
}

func (s *StringValue) Type() Type {
	return &BasicType{Type: String}
}

func (s *StringValue) Print() string {
	return "\"" + s.Value + "\""
}

type BoolValue struct {
	Value bool
}

func (b *BoolValue) Type() Type {
	return &BasicType{Type: Bool}
}

func (b *BoolValue) Print() string {
	return strconv.FormatBool(b.Value)
}

type ArrayValue struct {
	Values []Atom
}

func (a *ArrayValue) Type() Type {
	return &ArrayType{ElementType: &BasicType{Type: String}}
}

func (a *ArrayValue) Print() string {
	ps := []string{}
	for _, p := range a.Values {
		ps = append(ps, p.Print())
	}

	return "[" + strings.Join(ps, ", ") + "]"
}
