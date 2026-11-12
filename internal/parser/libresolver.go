package parser

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/byte-wright/lush/internal/ast"
)

type resolver struct {
	std fs.FS
}

type sourceFile struct {
	name string
	code string
}

func isLocal(importPath string) bool {
	return strings.HasPrefix(importPath, "./") || strings.HasPrefix(importPath, "../")
}

// locate finds the folder of the package that a file of from imports as importPath. A local
// import starts at the folder of from, from a std package it stays in std.
func locate(from *ast.Package, importPath string) ast.Source {
	if !isLocal(importPath) {
		return ast.Source{Std: true, Dir: path.Clean(importPath)}
	}

	if from.Source.Std {
		return ast.Source{Std: true, Dir: path.Join(from.Source.Dir, importPath)}
	}

	return ast.Source{Dir: filepath.Join(from.Source.Dir, filepath.FromSlash(importPath))}
}

func (r *resolver) resolve(pkg *ast.Package) ([]sourceFile, error) {
	var files []sourceFile
	var err error
	if pkg.Source.Std {
		files, err = readPackage(r.std, "std", pkg.Source.Dir)
	} else {
		files, err = readPackage(os.DirFS(pkg.Source.Dir), pkg.Source.Dir, ".")
	}

	if err != nil {
		return nil, fmt.Errorf("package %q: %w", pkg.Path, err)
	}

	return files, nil
}

// readPackage reads the .lsh files of dir, prefix is put in front of their names.
func readPackage(fsys fs.FS, prefix, dir string) ([]sourceFile, error) {
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}

	files := []sourceFile{}
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".lsh") {
			continue
		}

		file := path.Join(dir, e.Name())
		code, err := fs.ReadFile(fsys, file)
		if err != nil {
			return nil, err
		}

		files = append(files, sourceFile{name: filepath.Join(prefix, file), code: string(code)})
	}

	return files, nil
}
