package bcode

import "fmt"

type Slice struct {
	Value Atom
	From  Atom
	To    Atom
}

func (s *Slice) Type() Type {
	t := s.Value.Type()

	switch tp := t.(type) {
	case *BasicType:
		if tp.Type == String {
			return &BasicType{Type: String}
		}
	}

	panic(fmt.Sprintf("invalid type to slice %T", t))
}

func (s *Slice) Print() string {
	return s.Value.Print() + "[" + s.From.Print() + ":" + s.To.Print() + "]"
}

type Index struct {
	Value    *VarValue
	Position Atom
}

func (i *Index) Type() Type {
	t := i.Value.Type()

	switch tp := t.(type) {
	case *ArrayType:
		return tp.Type
	}

	panic(fmt.Sprintf("invalid type to index %T", t))
}

func (i *Index) Print() string {
	return i.Value.Print() + "[" + i.Position.Print() + "]"
}
