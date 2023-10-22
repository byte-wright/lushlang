package ast

import "strconv"

type Number struct {
	Value int
}

func (n *Number) print() string {
	return strconv.Itoa(int(n.Value))
}
