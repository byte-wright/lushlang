package bcode

type Atom interface {
	Print() string
	Type() Type
}

type Op interface {
	Atom
}

type Add struct {
	Left  Value
	Right Value
}

func (a *Add) Type() Type {
	return a.Left.Type()
}

func (a *Add) Print() string {
	return a.Left.Print() + " + " + a.Right.Print()
}

type Sub struct {
	Left  Value
	Right Value
}

func (s *Sub) Type() Type {
	return s.Left.Type()
}

func (s *Sub) Print() string {
	return s.Left.Print() + " - " + s.Right.Print()
}

type Minus struct {
	Expression Value
}

func (m *Minus) Type() Type {
	return m.Expression.Type()
}

func (m *Minus) Print() string {
	return "-" + m.Expression.Print()
}

type Mul struct {
	Left  Value
	Right Value
}

func (m *Mul) Type() Type {
	return m.Left.Type()
}

func (m *Mul) Print() string {
	return m.Left.Print() + " * " + m.Right.Print()
}

type Div struct {
	Left  Value
	Right Value
}

func (d *Div) Type() Type {
	return d.Left.Type()
}

func (d *Div) Print() string {
	return d.Left.Print() + " / " + d.Right.Print()
}

type Mod struct {
	Left  Value
	Right Value
}

func (m *Mod) Type() Type {
	return m.Left.Type()
}

func (m *Mod) Print() string {
	return m.Left.Print() + " % " + m.Right.Print()
}
