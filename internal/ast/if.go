package ast

type If struct {
	Condition Expression
	If        *Block
}

func (i *If) print() []string {
	bd := i.If.print()

	bd[0] = "if " + i.Condition.print() + " " + bd[0]

	return bd
}
