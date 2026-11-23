package parser

import (
	"fmt"
	"regexp"

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

// a convention so names are written one way, not a technical need: the backend prefixes
// the names it emits
func (l *errorListener) uppercaseNames(stream *antlr.CommonTokenStream) {
	for _, tk := range stream.GetAllTokens() {
		name := tk.GetText()
		if tk.GetTokenType() == LushLexerID && name[0] >= 'A' && name[0] <= 'Z' {
			l.SyntaxError(nil, nil, tk.GetLine(), tk.GetColumn(), fmt.Sprintf("name %v must start with a lowercase letter", name), nil)
		}
	}
}

var nameShape = regexp.MustCompile(`^[a-z_][a-zA-Z_0-9]*$`)

// invalidPackageName tells why code can't refer to a package by its folder name, like
// my-pkg or if. It's empty for a valid name.
func invalidPackageName(name string) string {
	if !nameShape.MatchString(name) {
		return "must start with a lowercase letter or _ and contain only letters, digits and _"
	}

	lexer := NewLushLexer(antlr.NewInputStream(name))
	lexer.RemoveErrorListeners()
	if lexer.NextToken().GetTokenType() != LushLexerID {
		return "is a keyword"
	}

	return ""
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
