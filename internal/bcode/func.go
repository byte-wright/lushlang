package bcode

import "strings"

type Func struct {
	Name       string
	Parameters []Value
}

func (*Func) Type() Type {
	return &BasicType{Type: String}
}

func (f *Func) Print() string {
	ps := []string{}
	for _, p := range f.Parameters {
		ps = append(ps, p.Print())
	}

	return f.Name + "(" + strings.Join(ps, ", ") + ")"
}
