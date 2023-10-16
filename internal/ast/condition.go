package ast

type LessThan struct {
	A Expression
	B Expression
}

func (l *LessThan) print() string {
	return l.A.print() + " < " + l.B.print()
}

type LessThanEqual struct {
	A Expression
	B Expression
}

func (l *LessThanEqual) print() string {
	return l.A.print() + " <= " + l.B.print()
}

type GreaterThan struct {
	A Expression
	B Expression
}

func (g *GreaterThan) print() string {
	return g.A.print() + " > " + g.B.print()
}

type GreaterThanEqual struct {
	A Expression
	B Expression
}

func (g *GreaterThanEqual) print() string {
	return g.A.print() + " >= " + g.B.print()
}

type Equal struct {
	A Expression
	B Expression
}

func (e *Equal) print() string {
	return e.A.print() + " == " + e.B.print()
}

type NotEqual struct {
	A Expression
	B Expression
}

func (n *NotEqual) print() string {
	return n.A.print() + " != " + n.B.print()
}
