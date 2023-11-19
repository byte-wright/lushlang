package parser

import (
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/byte-wright/lush/internal/ast"
)

func Parse(code, name, mainRoot, stdRoot string) (*ast.Program, error) {
	lexer := NewLushLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewLushParser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Program()

	prog := ast.NewProgram()

	v := newVisitor(name, &mainTarget{prog: prog})

	tree.Accept(v)

	resolver := &resolver{
		main: mainRoot,
		std:  stdRoot,
	}

	for i := 0; i < len(prog.Libraries()); i++ {
		lib := prog.Libraries()[i]

		files, err := resolver.resolve(lib.Path)
		if err != nil {
			return nil, err
		}

		for _, f := range files {

			src, err := os.ReadFile(f)
			if err != nil {
				return nil, err
			}

			err = ParseLibFile(prog, lib, string(src), f)
			if err != nil {
				return nil, err
			}
		}
	}

	return prog, nil
}

func ParseLibFile(p *ast.Program, lib *ast.Library, code, name string) error {
	lexer := NewLushLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewLushParser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Library()

	v := newVisitor(name, &libTarget{
		prog: p,
		lib:  lib,
	})

	tree.Accept(v)

	return nil
}
