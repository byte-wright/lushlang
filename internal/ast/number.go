package ast

import "strconv"

type Number struct {
	Value float64
}

func (n *Number) print() string {
	return strconv.Itoa(int(n.Value))
}
