package ast

import (
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
	addedPackages map[Source]*Package
}

// Source is the folder of a package's files. It is a path in std, or on disk relative to the
// working directory.
type Source struct {
	Std bool
	Dir string
}

type Package struct {
	// Path is how the import that added the package wrote it
	Path string
	// Name is the folder name, files can import the package with another one
	Name   string
	Source Source
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

// NewProgram creates a program whose main file is in the folder dir.
func NewProgram(dir string) *Program {
	main := &Package{Name: "main", Source: Source{Dir: dir}, Globals: &Block{}}

	return &Program{
		Root:          &Block{},
		Main:          main,
		File:          NewFile(main),
		addedPackages: map[Source]*Package{},
	}
}

func NewFile(pkg *Package) *File {
	return &File{
		Package: pkg,
		Imports: map[string]*Package{},
	}
}

// AddImport returns the package in src, it's only loaded once per program, however the
// imports write its path.
func (p *Program) AddImport(src Source, path string, loc *common.Location) *Package {
	pkg := p.addedPackages[src]
	if pkg != nil {
		return pkg
	}

	pkg = &Package{
		Path:     path,
		Name:     filepath.Base(src.Dir),
		Source:   src,
		Globals:  &Block{},
		Location: loc,
	}
	p.packages = append(p.packages, pkg)
	p.addedPackages[src] = pkg

	return pkg
}

// Packages returns the imported packages, without the main package.
func (p *Program) Packages() []*Package {
	return p.packages
}

// Import makes pkg visible as name in the file, it's false when name is another package.
func (f *File) Import(name string, pkg *Package) bool {
	other := f.Imports[name]
	if other != nil && other != pkg {
		return false
	}

	f.Imports[name] = pkg

	return true
}
