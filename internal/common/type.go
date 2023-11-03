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
