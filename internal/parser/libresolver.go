package parser

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type resolver struct {
	main string
	std  fs.FS
}

type sourceFile struct {
	name string
	code string
}

func (r *resolver) resolve(packagePath string) ([]sourceFile, error) {
	var files []sourceFile
	var err error
	if strings.HasPrefix(packagePath, "/") {
		files, err = readPackage(os.DirFS(r.main), r.main, packagePath[1:])
	} else {
		files, err = readPackage(r.std, "std", packagePath)
	}

	if err != nil {
		return nil, fmt.Errorf("package %q: %w", packagePath, err)
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
