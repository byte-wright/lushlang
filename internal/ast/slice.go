package ast

type Slice struct {
	Value Expression
	From  Expression
	To    Expression
}

type Index struct {
	Value    Expression
	Position Expression
}
