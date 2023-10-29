package ast

type Ternary struct {
	Condition Expression
	True      Expression
	False     Expression
}

func (t *Ternary) print() string {
	return "(" + t.Condition.print() + " ? " + t.True.print() + " : " + t.False.print() + ")"
}
