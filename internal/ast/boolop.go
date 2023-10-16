package ast

type And struct {
	A Expression
	B Expression
}

func (a *And) print() string {
	return a.A.print() + " && " + a.B.print()
}

type Or struct {
	A Expression
	B Expression
}

func (o *Or) print() string {
	return o.A.print() + " || " + o.B.print()
}

type Not struct {
	Expression
}

func (n *Not) print() string {
	return "!" + n.Expression.print()
}
