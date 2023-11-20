package bcode

import "github.com/byte-wright/lush/internal/common"

type And struct {
	Left  Value
	Right Value
}

func (a *And) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

type Not struct {
	Expression Value
}

func (n *Not) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}
