package bcode

import "strconv"

type Value interface {
	Atom
}

type VarValue struct {
	Name string
}

func (v *VarValue) Print() string {
	return v.Name
}

type EnvVarValue struct {
	Name string
}

func (v *EnvVarValue) Print() string {
	return "$" + v.Name
}

type NumberValue struct {
	Value int
}

func (n *NumberValue) Print() string {
	return strconv.Itoa(n.Value)
}

type StringValue struct {
	Value string
}

func (s *StringValue) Print() string {
	return "\"" + s.Value + "\""
}

type BoolValue struct {
	Value bool
}

func (b *BoolValue) Print() string {
	return strconv.FormatBool(b.Value)
}
