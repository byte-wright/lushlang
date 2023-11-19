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

	for i := 0; i < len(prog.Packages()); i++ {
		pkg := prog.Packages()[i]

		files, err := resolver.resolve(pkg.Path)
		if err != nil {
			return nil, err
		}

		for _, f := range files {

			src, err := os.ReadFile(f)
			if err != nil {
				return nil, err
			}

			err = ParsePkgFile(prog, pkg, string(src), f)
			if err != nil {
				return nil, err
			}
		}
	}

	return prog, nil
}

func ParsePkgFile(p *ast.Program, pkg *ast.Package, code, name string) error {
	lexer := NewLushLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewLushParser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Package_()

	v := newVisitor(name, &pkgTarget{
		prog: p,
		pkg:  pkg,
	})

	tree.Accept(v)

	return nil
}
