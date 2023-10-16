package ast

type Group struct {
	Expression
}

func (g *Group) print() string {
	return "(" + g.Expression.print() + ")"
}
