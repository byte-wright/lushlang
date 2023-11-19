package ast

type For struct {
	Initial   *Assignment
	Condition Expression
	Each      *Assignment
	Body      *Block
}
