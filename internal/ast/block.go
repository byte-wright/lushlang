package ast

import (
	"fmt"

	"github.com/byte-wright/lush/internal/common"
)

type Statement interface{}

type Block struct {
	Statements []Statement

	// calculated
	Parent    *Block
	Variables []*Variable
}

type Variable struct {
	Name string
	Type common.Type
}

func (b *Block) GetVar(name string) *Variable {
	for _, v := range b.Variables {
		if v.Name == name {
			return v
		}
	}

	if b.Parent == nil {
		return nil
	}
	return b.Parent.GetVar(name)
}

func (b *Block) SetVar(name string, tp common.Type) error {
	v := b.GetVar(name)
	if v == nil {
		b.Variables = append(b.Variables, &Variable{
			Name: name,
			Type: tp,
		})

		return nil
	}

	if v.Type.Is(tp) {
		return nil
	}

	return fmt.Errorf("you cannot assign an %v value to the variable '%v' of type %v", tp.Print(), name, v.Type.Print())
}
