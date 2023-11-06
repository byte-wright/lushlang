package bcode

import (
	"fmt"

	"github.com/byte-wright/lush/internal/common"
)

type Slice struct {
	Value *VarValue
	From  Atom
	To    Atom
}

func (s *Slice) Type() common.Type {
	return s.Value.Type()
}

func (s *Slice) Print() string {
	return s.Value.Print() + "[" + s.From.Print() + ":" + s.To.Print() + "]"
}

type Index struct {
	Value    *VarValue
	Position Atom
}

func (i *Index) Type() common.Type {
	t := i.Value.Type()

	switch tp := t.(type) {
	case *common.ArrayType:
		return tp.ElementType
	}

	panic(fmt.Sprintf("invalid type to index %T", t))
}

func (i *Index) Print() string {
	return i.Value.Print() + "[" + i.Position.Print() + "]"
}
