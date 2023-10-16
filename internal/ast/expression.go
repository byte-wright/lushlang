package ast

type Expression interface {
	print() string
}

type Statement interface {
	print() []string
}
