package bcode

import "github.com/byte-wright/lush/internal/common"

type Equal struct {
	Left  Value
	Right Value
}

func (*Equal) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (e *Equal) Print() string {
	return e.Left.Print() + " == " + e.Right.Print()
}

type LessThan struct {
	Left  Value
	Right Value
}

func (*LessThan) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (l *LessThan) Print() string {
	return l.Left.Print() + " < " + l.Right.Print()
}

type LessThanEqual struct {
	Left  Value
	Right Value
}

func (*LessThanEqual) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (l *LessThanEqual) Print() string {
	return l.Left.Print() + " <= " + l.Right.Print()
}

type GreaterThan struct {
	Left  Value
	Right Value
}

func (*GreaterThan) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (g *GreaterThan) Print() string {
	return g.Left.Print() + " > " + g.Right.Print()
}

type GreaterThanEqual struct {
	Left  Value
	Right Value
}

func (*GreaterThanEqual) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (g *GreaterThanEqual) Print() string {
	return g.Left.Print() + " >= " + g.Right.Print()
}

type NotEqual struct {
	Left  Value
	Right Value
}

func (*NotEqual) Type() Type {
	return &BasicType{Type: common.Bool}
}

func (n *NotEqual) Print() string {
	return n.Left.Print() + " != " + n.Right.Print()
}
