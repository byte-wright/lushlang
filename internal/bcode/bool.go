package bcode

import "github.com/byte-wright/lush/internal/common"

type Not struct {
	Expression Value
}

func (n *Not) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}
