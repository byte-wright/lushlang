package parser

import (
	"os"
	"path/filepath"
	"strings"
)

type resolver struct {
	main string
	std  string
}

func (r *resolver) resolve(packagePath string) ([]string, error) {
	folder := ""
	if strings.HasPrefix(packagePath, "/") {
		folder = filepath.Join(r.main, packagePath[1:])
	} else {
		folder = filepath.Join(r.std, packagePath)
	}

	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	fs := []string{}
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".lsh") {
			continue
		}

		file := filepath.Join(folder, f.Name())
		fs = append(fs, file)
	}

	return fs, nil
}
