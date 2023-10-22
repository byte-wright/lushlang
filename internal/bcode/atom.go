package bcode

type Atom interface {
	Print() string
}

type Op interface {
	Atom
}

type Add struct {
	Left  Value
	Right Value
}

func (a *Add) Print() string {
	return a.Left.Print() + " + " + a.Right.Print()
}

type Minus struct {
	Expression Value
}

func (m *Minus) Print() string {
	return "-" + m.Expression.Print()
}

type Mul struct {
	Left  Value
	Right Value
}

func (m *Mul) Print() string {
	return m.Left.Print() + " * " + m.Right.Print()
}

type Mod struct {
	Left  Value
	Right Value
}

func (m *Mod) Print() string {
	return m.Left.Print() + " % " + m.Right.Print()
}
