package parser

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/byte-wright/lush/internal/ast"
)

func Parse(code, name, libRoot string) (*ast.Program, error) {
	lexer := NewLushLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewLushParser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Program()

	prog := &ast.Program{
		Root: &ast.Block{},
	}

	v := newVisitor(name, &mainTarget{prog: prog})

	tree.Accept(v)

	imported := map[string]bool{}

	for len(imported) < len(prog.Libs) {
		for _, lib := range prog.Libs {
			libDir := filepath.Join(libRoot, lib.Path)
			files, err := os.ReadDir(libDir)
			if err != nil {
				return nil, err
			}

			for _, f := range files {
				if !strings.HasSuffix(f.Name(), ".lsh") {
					continue
				}

				src, err := os.ReadFile(filepath.Join(libDir, f.Name()))
				if err != nil {
					return nil, err
				}

				name := path.Join(lib.Path, f.Name())

				err = ParseLibFile(prog, lib, string(src), name)
				if err != nil {
					return nil, err
				}
			}

			imported[lib.Path] = true
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
