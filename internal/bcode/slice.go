package bcode

type Slice struct {
	Value Atom
	From  Atom
	To    Atom
}

func (s *Slice) Print() string {
	return s.Value.Print() + "[" + s.From.Print() + ":" + s.To.Print() + "]"
}

type Index struct {
	Value    *VarValue
	Position Atom
}

func (i *Index) Print() string {
	return i.Value.Print() + "[" + i.Position.Print() + "]"
}
