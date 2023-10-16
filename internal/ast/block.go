package ast

type Block struct {
	Statements []Statement
}

func (b *Block) print() []string {
	ls := []string{"{"}

	for _, st := range b.Statements {
		sls := st.print()
		for _, l := range sls {
			ls = append(ls, "  "+l)
		}
	}

	ls = append(ls, "}")
	return ls
}
