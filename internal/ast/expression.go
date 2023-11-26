package ast

import "github.com/byte-wright/lush/internal/common"

type Expression interface {
	Types() []common.Type
}

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

type Slice struct {
	Value Expression
	From  Expression
	To    Expression
}

type Index struct {
	Value    Expression
	Position Expression
}

type Ternary struct {
	Condition Expression
	True      Expression
	False     Expression
}

type Func struct {
	Name       string
	Parameters []Expression
	// calculated
	ReturnTypes []common.Type
}

type Method struct {
	Namespace string
	Func      *Func
}

type Var struct {
	Name string
	// calculated
	Type common.Type
}

type Number struct {
	Value int
}

type Bool struct {
	Value bool
}

type String struct {
	Value string
}

type Array struct {
	Values []Expression
	// Type is set only when declaring an empty array.
	// Otherwise it needs to be read from the arrays values.
	Type common.PrimitiveType
}
