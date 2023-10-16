package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/byte-wright/lush/internal/ast"
)

func Parse(code, name string) (*ast.Program, error) {
	lexer := NewLushLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewLushParser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Program()

	v := newVisitor(name)

	program := tree.Accept(v)
	return program.(*ast.Program), nil
}
