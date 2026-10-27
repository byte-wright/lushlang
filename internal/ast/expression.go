package ast

import "github.com/byte-wright/lush/internal/common"

type Expression interface {
	Types() []common.Type
}

type Minus struct {
	Expression
	Location *common.Location
}

type Mul struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Div struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Mod struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Add struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Sub struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type And struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Or struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Not struct {
	Expression
	Location *common.Location
}

type LessThan struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type LessThanEqual struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type GreaterThan struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type GreaterThanEqual struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Equal struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type NotEqual struct {
	A        Expression
	B        Expression
	Location *common.Location
}

type Slice struct {
	Value    Expression
	From     Expression
	To       Expression
	Location *common.Location
}

type Index struct {
	Value    Expression
	Position Expression
	Location *common.Location
}

type Ternary struct {
	Condition Expression
	True      Expression
	False     Expression
	Location  *common.Location
}

type Func struct {
	Name       string
	Parameters []Expression
	Location   *common.Location
	// calculated
	ReturnTypes []common.Type
	// Package defines the function, it's nil for builtins
	Package *Package
}

type Method struct {
	Namespace string
	Func      *Func
}

type Var struct {
	Name     string
	Location *common.Location
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
