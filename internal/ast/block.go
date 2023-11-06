package ast

type Block struct {
	Statements []Statement
	FuncDefs   []*FuncDef
}

func (b *Block) printBare() []string {
	ls := []string{}

	for _, st := range b.Statements {
		sls := st.print()
		for _, l := range sls {
			ls = append(ls, l)
		}
	}

	return ls
}

func (b *Block) print() []string {
	ls := []string{}
	ls = append(ls, "{")

	ls = append(ls, b.printBare()...)

	ls = append(ls, "}")
	return ls
}
