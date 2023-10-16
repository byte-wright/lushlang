package ast

type String struct {
	Value string
}

func (s *String) print() string {
	return "\"" + s.Value + "\""
}
