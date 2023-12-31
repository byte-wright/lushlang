package bcode

import (
	"github.com/byte-wright/lush/internal/common"
)

type Len struct {
	Parameter *VarValue
}

func (l *Len) Type() common.Type {
	return &common.BasicType{Type: common.Int}
}
