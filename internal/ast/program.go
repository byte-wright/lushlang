package ast

import (
	"fmt"
	"path/filepath"

	"github.com/byte-wright/lush/internal/common"
)

type Program struct {
	Root *Block
	Main *Package
	// File is the main file, Root is part of it
	File *File

	packages      []*Package
	addedPackages map[string]*Package
}

type Package struct {
	Path             string
	Name             string
	FuncDefs         []*FuncDef
	ExternalFuncDefs []*ExternalFuncDef
	// Location is the import that added the package, it's nil for main
	Location *common.Location
}

// File holds the imports of a source file, they are only visible inside that file. Its
// functions are visible in the whole package.
type File struct {
	Package *Package
	Imports map[string]*Package
}

func NewProgram() *Program {
	main := &Package{Name: "main"}

	return &Program{
		Root:          &Block{},
		Main:          main,
		File:          NewFile(main),
		addedPackages: map[string]*Package{},
	}
}

func NewFile(pkg *Package) *File {
	return &File{
		Package: pkg,
		Imports: map[string]*Package{},
	}
}

// AddImport returns the package of the path, it's only loaded once per program.
func (p *Program) AddImport(path string, loc *common.Location) *Package {
	pkg := p.addedPackages[path]
	if pkg != nil {
		return pkg
	}

	pkg = &Package{
		Path:     path,
		Name:     filepath.Base(path),
		Location: loc,
	}
	p.packages = append(p.packages, pkg)
	p.addedPackages[path] = pkg

	return pkg
}

// Packages returns the imported packages, without the main package.
func (p *Program) Packages() []*Package {
	return p.packages
}

func (f *File) Import(pkg *Package) error {
	other := f.Imports[pkg.Name]
	if other != nil && other != pkg {
		return fmt.Errorf("package %v is already imported from %q", pkg.Name, other.Path)
	}

	f.Imports[pkg.Name] = pkg

	return nil
}
