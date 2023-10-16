package ast

type Slice struct {
	Value Expression
	From  Expression
	To    Expression
}

func (s *Slice) print() string {
	return "(" + s.Value.print() + ")[" + s.From.print() + ":" + s.To.print() + "]"
}
