package ast

type And struct {
	A Expression
	B Expression
}

type Or struct {
	A Expression
	B Expression
}

type Not struct {
	Expression
}
