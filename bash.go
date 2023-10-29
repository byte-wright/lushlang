package lush

import (
	"github.com/byte-wright/lush/internal/bash"
	"github.com/byte-wright/lush/internal/bcode"
	"github.com/byte-wright/lush/internal/parser"
)

func ToBashDebug(source, name string) (string, string, string, error) {
	prog, err := parser.Parse(string(source), name)
	if err != nil {
		return "", "", "", err
	}

	ast := prog.Print()

	bc, err := bcode.Compile(prog)
	if err != nil {
		return "", "", "", err
	}

	bCode := bc.Print()

	return ast, bCode, bash.Translate(bc), nil
}
