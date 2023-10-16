package ast

type Minus struct {
	Expression
}

func (m *Minus) print() string {
	return "(" + "-" + m.Expression.print() + ")"
}

type Mul struct {
	A Expression
	B Expression
}

func (m *Mul) print() string {
	return "(" + m.A.print() + " * " + m.B.print() + ")"
}

type Div struct {
	A Expression
	B Expression
}

func (d *Div) print() string {
	return "(" + d.A.print() + " / " + d.B.print() + ")"
}

type Mod struct {
	A Expression
	B Expression
}

func (m *Mod) print() string {
	return "(" + m.A.print() + " % " + m.B.print() + ")"
}

type Add struct {
	A Expression
	B Expression
}

func (a *Add) print() string {
	return "(" + a.A.print() + " + " + a.B.print() + ")"
}

type Sub struct {
	A Expression
	B Expression
}

func (s *Sub) print() string {
	return "(" + s.A.print() + " - " + s.B.print() + ")"
}
