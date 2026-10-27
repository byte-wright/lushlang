package bcode

import (
	"fmt"

	"github.com/byte-wright/lush/internal/common"
)

type Slice struct {
	Value    *VarValue
	From     Atom
	To       Atom
	Location *common.Location
}

func (s *Slice) Type() common.Type {
	return s.Value.Type()
}

type Index struct {
	Value    *VarValue
	Position Atom
	Location *common.Location
}

func (i *Index) Type() common.Type {
	t := i.Value.Type()

	switch tp := t.(type) {
	case *common.ArrayType:
		return tp.ElementType
	case *common.BasicType:
		if tp.Type == common.String {
			return tp
		}
	}

	panic(fmt.Sprintf("invalid type to index %T", t))
}
