package ast

import "strconv"

type Number struct {
	Value int
}

func (n *Number) print() string {
	return strconv.Itoa(int(n.Value))
}

type Bool struct {
	Value bool
}

func (b *Bool) print() string {
	return strconv.FormatBool(b.Value)
}

type String struct {
	Value string
}

func (s *String) print() string {
	return "\"" + s.Value + "\""
}

type Array struct {
	Values []Expression
}

func (a *Array) print() string {
	r := "["
	for i, xp := range a.Values {
		if i > 0 {
			r += ", "
		}
		r += xp.print()
	}
	return r + "]"
}
