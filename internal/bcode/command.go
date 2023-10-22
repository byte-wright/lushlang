package bcode

type Command interface{}

type Assignment struct {
	Name  string
	Value Atom
}
