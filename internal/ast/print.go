package ast

import (
	"gitlab.com/akabio/gogen"
	"gitlab.com/akabio/gogen/parser"

	_ "embed"
)

//go:embed print.gg
var template string

var ggAST = func() *gogen.AST {
	ast, err := parser.Parse(template, "print.gg", func(i string) (string, string, error) {
		panic("imports not implemented")
	})
	if err != nil {
		panic(err)
	}

	return ast
}()

func Print(p *Program) string {
	r, err := gogen.Execute(ggAST, p)
	if err != nil {
		panic(err)
	}

	return r
}
