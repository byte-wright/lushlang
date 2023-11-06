package common

import "fmt"

type PrimitiveType int

const (
	None   PrimitiveType = 0
	String PrimitiveType = 1
	Int    PrimitiveType = 2
	Bool   PrimitiveType = 3
)

func PrimitiveTypeFor(t string) PrimitiveType {
	switch t {
	case "string":
		return String
	case "int":
		return Int
	case "bool":
		return Bool
	}
	panic(fmt.Sprintf("invalid type '%v'", t))
}

type Type interface {
	IsArray() bool
	Print() string
}

type BasicType struct {
	Type PrimitiveType
}

func (b *BasicType) IsArray() bool {
	return false
}

func (b *BasicType) Print() string {
	switch b.Type {
	case String:
		return "string"
	case Int:
		return "int"
	case Bool:
		return "bool"
	}
	return "none"
}

type ArrayType struct {
	ElementType Type
}

func (a *ArrayType) IsArray() bool {
	return true
}

func (a *ArrayType) Print() string {
	return "[]" + a.ElementType.Print()
}
