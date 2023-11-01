package bcode

type PrimitiveType int

const (
	None   PrimitiveType = 0
	String PrimitiveType = 1
	Int    PrimitiveType = 2
	Bool   PrimitiveType = 3
)

type Type interface {
	IsArray() bool
}

type BasicType struct {
	Type PrimitiveType
}

func (b *BasicType) IsArray() bool {
	return false
}

type ArrayType struct {
	Type
}

func (a *ArrayType) IsArray() bool {
	return true
}

type VarValue struct {
	Name string
	T    Type
}

func (v *VarValue) Type() Type {
	return v.T
}

func (v *VarValue) Print() string {
	return v.Name
}
