package ast

type ReturnStatement struct {
	Expression Expression
}

func (r *ReturnStatement) print() []string {
	return []string{"return " + r.Expression.print()}
}
