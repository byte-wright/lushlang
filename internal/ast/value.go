package ast

import (
	"github.com/byte-wright/lush/internal/common"
)

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
	Type   common.PrimitiveType
}
