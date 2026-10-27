package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type errorListener struct {
	*antlr.DefaultErrorListener
	name   string
	errors []*ast.ASTError
}

func (l *errorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	l.errors = append(l.errors, &ast.ASTError{
		Location: &common.Location{
			File:     l.name,
			LineFrom: line,
			ColFrom:  column + 1,
			LineTo:   line,
			ColTo:    column + 1,
		},
		Message: msg,
	})
}

// listen replaces antlr's default listener, which prints to stderr and lets parsing go on.
func listen(lexer *LushLexer, parser *LushParser, name string) *errorListener {
	l := &errorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		name:                 name,
	}

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(l)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(l)

	return l
}
