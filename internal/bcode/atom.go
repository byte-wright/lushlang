package bcode

import "github.com/byte-wright/lush/internal/common"

type Atom interface {
	Type() common.Type
}

type Op interface {
	Atom
}

type Add struct {
	Left  Value
	Right Value
}

func (a *Add) Type() common.Type {
	return a.Left.Type()
}

type Sub struct {
	Left  Value
	Right Value
}

func (s *Sub) Type() common.Type {
	return s.Left.Type()
}

type Minus struct {
	Expression Value
}

func (m *Minus) Type() common.Type {
	return m.Expression.Type()
}

type Mul struct {
	Left  Value
	Right Value
}

func (m *Mul) Type() common.Type {
	return m.Left.Type()
}

type Div struct {
	Left  Value
	Right Value
}

func (d *Div) Type() common.Type {
	return d.Left.Type()
}

type Mod struct {
	Left  Value
	Right Value
}

func (m *Mod) Type() common.Type {
	return m.Left.Type()
}
