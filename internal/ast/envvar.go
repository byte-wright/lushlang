package ast

type EnvVar struct {
	Name string
}

func (e *EnvVar) print() string {
	return "$" + e.Name
}
