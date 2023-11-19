package bcode

import (
	"fmt"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type Signature struct {
	Namespace  string
	Name       string
	Parameters []common.Type
	Return     []common.Type
}

var internalFuncs = func() map[string]*Signature {
	return map[string]*Signature{
		"join": {
			Name: "join",
			Parameters: []common.Type{
				&common.ArrayType{ElementType: &common.BasicType{Type: common.String}},
				&common.BasicType{Type: common.String},
			},
			Return: []common.Type{
				&common.BasicType{Type: common.String},
			},
		},
		"append": {
			Name: "append",
			Parameters: []common.Type{
				&common.ArrayType{ElementType: &common.BasicType{Type: common.String}},
				&common.BasicType{Type: common.String},
			},
			Return: []common.Type{
				&common.ArrayType{ElementType: &common.BasicType{Type: common.String}},
			},
		},
		"hasPrefix": {
			Name: "hasPrefix",
			Parameters: []common.Type{
				&common.BasicType{Type: common.String},
				&common.BasicType{Type: common.String},
			},
			Return: []common.Type{
				&common.BasicType{Type: common.Bool},
			},
		},
		"hasSuffix": {
			Name: "hasSuffix",
			Parameters: []common.Type{
				&common.BasicType{Type: common.String},
				&common.BasicType{Type: common.String},
			},
			Return: []common.Type{
				&common.BasicType{Type: common.Bool},
			},
		},
		"trimPrefix": {
			Name: "trimPrefix",
			Parameters: []common.Type{
				&common.BasicType{Type: common.String},
				&common.BasicType{Type: common.String},
			},
			Return: []common.Type{
				&common.BasicType{Type: common.String},
			},
		},
		"trimSuffix": {
			Name: "trimSuffix",
			Parameters: []common.Type{
				&common.BasicType{Type: common.String},
				&common.BasicType{Type: common.String},
			},
			Return: []common.Type{
				&common.BasicType{Type: common.String},
			},
		},
	}
}()

func getSignature(prog *ast.Program, namespace, name string) *Signature {
	if namespace == "" {
		internal, isInternal := internalFuncs[name]
		if isInternal {
			return internal
		}

		for _, fd := range prog.Root.FuncDefs {
			if fd.Name == name {
				return &Signature{
					Name: name,
					Parameters: common.Map(fd.Params, func(p *ast.Param) common.Type {
						return p.Type
					}),
					Return: fd.Returns,
				}
			}
		}
	}

	for _, lib := range prog.Libs {
		if lib.Name == namespace {
			for _, fd := range lib.FuncDefs {
				if fd.Name == name {
					return &Signature{
						Namespace: lib.Path,
						Name:      name,
						Parameters: common.Map(fd.Params, func(p *ast.Param) common.Type {
							return p.Type
						}),
						Return: fd.Returns,
					}
				}
			}
			for _, fd := range lib.ExternalFuncDefs {
				if fd.Name == name {
					return &Signature{
						Namespace: lib.Path,
						Name:      name,
						Parameters: common.Map(fd.Params, func(p *ast.Param) common.Type {
							return p.Type
						}),
						Return: fd.Returns,
					}
				}
			}
		}
	}

	panic(fmt.Sprintf("func signature for '%s' not found", name))
}
