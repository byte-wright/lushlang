package bcode

type Command interface{}

type Assignment struct {
	Var   *VarValue
	Value Atom
}

type Code struct {
	Code string
}
