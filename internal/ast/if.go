package ast

type If struct {
	Condition Expression
	If        *Block
}
