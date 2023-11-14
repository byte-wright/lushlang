package ast

import "strings"

type ReturnStatement struct {
	Expressions []Expression
}

func (r *ReturnStatement) print() []string {
	xps := []string{}
	for _, x := range r.Expressions {
		xps = append(xps, x.print())
	}
	return []string{"return " + strings.Join(xps, ", ")}
}
