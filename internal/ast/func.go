package ast

type Func struct {
	Name       string
	Parameters []Expression
}

type Method struct {
	Expression Expression
	Func       *Func
}
