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

func getSignature(prog *ast.Program, namespace, name string) *Signature {
	if namespace == "" {
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

	for _, lib := range prog.Libraries() {
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
