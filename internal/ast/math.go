package ast

type Minus struct {
	Expression
}

type Mul struct {
	A Expression
	B Expression
}

type Div struct {
	A Expression
	B Expression
}

type Mod struct {
	A Expression
	B Expression
}

type Add struct {
	A Expression
	B Expression
}

type Sub struct {
	A Expression
	B Expression
}
