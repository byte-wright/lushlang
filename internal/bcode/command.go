package bcode

type Command interface{}

type Assignment struct {
	Var   *VarValue
	Value Atom
}
