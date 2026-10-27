package lush

import (
	"embed"
	"errors"
	"io/fs"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/bash"
	"github.com/byte-wright/lush/internal/bcode"
	"github.com/byte-wright/lush/internal/parser"
)

//go:embed std
var stdFiles embed.FS

// ToBashDebug compiles source to bash, it also returns the printed AST and bcode. Imports
// starting with / are read from the root folder, std is built in.
func ToBashDebug(source, name, root string) (string, string, string, error) {
	std, err := fs.Sub(stdFiles, "std")
	if err != nil {
		return "", "", "", err
	}

	prog, astErrs, err := parser.Parse(source, name, root, std)
	if err != nil {
		return "", "", "", err
	}

	if len(astErrs) > 0 {
		errs := []error{}
		for _, e := range astErrs {
			errs = append(errs, e)
		}

		return "", "", "", errors.Join(errs...)
	}

	ast := ast.Print(prog)

	bc, err := bcode.Compile(prog)
	if err != nil {
		return "", "", "", err
	}

	bCode := bcode.Print(bc)

	return ast, bCode, bash.Translate(bc), nil
}
