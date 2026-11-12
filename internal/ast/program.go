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

	// InitOrder holds the GlobalDefs of all packages, in the order they get their values
	InitOrder []*GlobalDef

	packages      []*Package
	addedPackages map[string]*Package
}

type Package struct {
	Path string
	Name string
	// GlobalDefs is empty for main, its globals are defined by global assignments in Root
	GlobalDefs       []*GlobalDef
	Globals          *Block
	FuncDefs         []*FuncDef
	ExternalFuncDefs []*ExternalFuncDef
	// Location is the import that added the package, it's nil for main
	Location *common.Location
}

// File holds the imports of a source file, they are only visible inside that file. Its
// functions and globals are visible in the whole package.
type File struct {
	Package *Package
	Imports map[string]*Package
}

func NewProgram() *Program {
	main := &Package{Name: "main", Globals: &Block{}}

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
		Globals:  &Block{},
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

// Import makes pkg visible as name in the file, that's its folder name or the one after as.
func (f *File) Import(name string, pkg *Package) error {
	other := f.Imports[name]
	if other != nil && other != pkg {
		return fmt.Errorf("package %v is already imported from %q", name, other.Path)
	}

	f.Imports[name] = pkg

	return nil
}
