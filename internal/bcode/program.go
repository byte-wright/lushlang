package bcode

type Program struct {
	Main   *Block
	varIdx int
}

func (p *Program) nextVar() int {
	p.varIdx++
	return p.varIdx
}

func (p *Program) Print() string {
	return p.Main.Print(0)
}
