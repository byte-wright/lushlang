package ast

import (
	"strings"

	"github.com/byte-wright/lush/internal/common"
)

type Assignment struct {
	Names       []string
	Expressions []Expression
}

func (a *Assignment) print() []string {
	xps := common.Map(a.Expressions, func(x Expression) string {
		return x.print()
	})

	return []string{strings.Join(a.Names, ", ") + " = " + strings.Join(xps, ", ")}
}
