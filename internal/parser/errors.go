package parser

import (
	"fmt"

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

// lowercase names can't clash with bash's variables like PATH and IFS, or with the
// generated L_ ones
func (l *errorListener) uppercaseNames(stream *antlr.CommonTokenStream) {
	for _, tk := range stream.GetAllTokens() {
		name := tk.GetText()
		if tk.GetTokenType() == LushLexerID && name[0] >= 'A' && name[0] <= 'Z' {
			l.SyntaxError(nil, nil, tk.GetLine(), tk.GetColumn(), fmt.Sprintf("name %v must start with a lowercase letter", name), nil)
		}
	}
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
