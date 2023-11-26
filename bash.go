package lush

import (
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/bash"
	"github.com/byte-wright/lush/internal/bcode"
	"github.com/byte-wright/lush/internal/parser"
)

func ToBashDebug(source, name string) (string, string, string, error) {
	prog, _, err := parser.Parse(string(source), name, "./std", "./std")
	if err != nil {
		return "", "", "", err
	}

	ast := ast.Print(prog)

	bc, err := bcode.Compile(prog)
	if err != nil {
		return "", "", "", err
	}

	bCode := bcode.Print(bc)

	return ast, bCode, bash.Translate(bc), nil
}
