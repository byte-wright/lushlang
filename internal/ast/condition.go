package ast

type LessThan struct {
	A Expression
	B Expression
}

type LessThanEqual struct {
	A Expression
	B Expression
}

type GreaterThan struct {
	A Expression
	B Expression
}

type GreaterThanEqual struct {
	A Expression
	B Expression
}

type Equal struct {
	A Expression
	B Expression
}

type NotEqual struct {
	A Expression
	B Expression
}
