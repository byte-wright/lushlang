package ast

import "strings"

type Program struct {
	Root *Block
}

func (p *Program) Print() string {
	lines := []string{}

	for _, st := range p.Root.Statements {
		ls := st.print()
		lines = append(lines, ls...)
	}

	return strings.Join(lines, "\n")
}
