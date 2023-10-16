package ast

type Assignment struct {
	Name       string
	Expression Expression
}

func (a *Assignment) print() []string {
	return []string{a.Name + " = " + a.Expression.print()}
}
