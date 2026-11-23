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

// ToBashDebug compiles source to bash, it also returns the printed AST and bcode. dir is the
// folder of source, imports starting with ./ or ../ start there, the others are from the
// built in std.
func ToBashDebug(source, name, dir string) (string, string, string, error) {
	std, err := fs.Sub(stdFiles, "std")
	if err != nil {
		return "", "", "", err
	}

	prog, astErrs, err := parser.Parse(source, name, dir, std)
	if err != nil {
		return "", "", "", err
	}

	if len(astErrs) > 0 {
		return "", "", "", joinErrors(astErrs)
	}

	printed := ast.Print(prog)

	bc, err := bcode.Compile(prog)
	if err != nil {
		return "", "", "", err
	}

	bCode := bcode.Print(bc)

	sh, astErrs := bash.Translate(bc)
	if len(astErrs) > 0 {
		return "", "", "", joinErrors(astErrs)
	}

	return printed, bCode, sh, nil
}

func joinErrors(astErrs []*ast.ASTError) error {
	errs := []error{}
	for _, e := range astErrs {
		errs = append(errs, e)
	}

	return errors.Join(errs...)
}
