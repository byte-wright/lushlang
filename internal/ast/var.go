package ast

type Var struct {
	Name string
}

func (v *Var) print() string {
	return v.Name
}
