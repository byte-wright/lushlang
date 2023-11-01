package ast

type Slice struct {
	Value Expression
	From  Expression
	To    Expression
}

func (s *Slice) print() string {
	f := ""
	if s.From != nil {
		f = s.From.print()
	}
	t := ""
	if s.To != nil {
		t = s.To.print()
	}
	return "(" + s.Value.print() + ")[" + f + ":" + t + "]"
}

type Index struct {
	Value    Expression
	Position Expression
}

func (i *Index) print() string {
	return "(" + i.Value.print() + ")[" + i.Position.print() + "]"
}
