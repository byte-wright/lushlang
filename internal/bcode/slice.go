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
