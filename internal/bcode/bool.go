package bcode

import "github.com/byte-wright/lush/internal/common"

type And struct {
	Left  Value
	Right Value
}

func (a *And) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (a *And) Print() string {
	return a.Left.Print() + " && " + a.Right.Print()
}

type Not struct {
	Expression Value
}

func (n *Not) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (n *Not) Print() string {
	return "!" + n.Expression.Print()
}
