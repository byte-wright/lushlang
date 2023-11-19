package ast

type Ternary struct {
	Condition Expression
	True      Expression
	False     Expression
}
