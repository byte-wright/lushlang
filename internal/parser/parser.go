package parser

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/antlr4-go/antlr/v4"
	"github.com/byte-wright/lush/internal/ast"
	astpass "github.com/byte-wright/lush/internal/ast/pass"
)

// Parse parses the main file and the packages it imports. Imports starting with / are
// read from the mainRoot folder, all others from std.
func Parse(code, name, mainRoot string, std fs.FS) (*ast.Program, []*ast.ASTError, error) {
	lexer := NewLushLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewLushParser(stream)
	listener := listen(lexer, parser, name)

	parser.BuildParseTrees = true
	tree := parser.Program()

	// the visitor can't handle the incomplete tree of a source with syntax errors
	if len(listener.errors) > 0 {
		return nil, listener.errors, nil
	}

	prog := ast.NewProgram()

	v := newVisitor(name, &mainTarget{prog: prog})

	tree.Accept(v)

	resolver := &resolver{
		main: mainRoot,
		std:  std,
	}

	pkgErrs := []*ast.ASTError{}
	for i := 0; i < len(prog.Packages()); i++ {
		pkg := prog.Packages()[i]

		files, err := resolver.resolve(pkg.Path)
		// an invalid path is one that leaves the root, like /../x
		if errors.Is(err, fs.ErrNotExist) || errors.Is(err, fs.ErrInvalid) {
			pkgErrs = append(pkgErrs, &ast.ASTError{Location: pkg.Location, Message: fmt.Sprintf("package %q not found", pkg.Path)})
			continue
		}
		if err != nil {
			return nil, nil, err
		}

		for _, f := range files {
			pkgErrs = append(pkgErrs, ParsePkgFile(prog, pkg, f.code, f.name)...)
		}
	}

	// without all packages, checking the types would only find follow-up errors
	if len(pkgErrs) > 0 {
		return nil, append(v.errors, pkgErrs...), nil
	}

	errs := append(v.errors, astpass.SetParent(prog)...)
	errs = append(errs, astpass.SetTypes(prog)...)

	return prog, errs, nil
}

func ParsePkgFile(p *ast.Program, pkg *ast.Package, code, name string) []*ast.ASTError {
	lexer := NewLushLexer(antlr.NewInputStream(code))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewLushParser(stream)
	listener := listen(lexer, parser, name)

	parser.BuildParseTrees = true
	tree := parser.Package_()

	if len(listener.errors) > 0 {
		return listener.errors
	}

	v := newVisitor(name, &pkgTarget{
		prog: p,
		file: ast.NewFile(pkg),
	})

	tree.Accept(v)

	return v.errors
}
