package ast

import "strings"

type Program struct {
	Root *Block
}

func (p *Program) Print() string {
	lines := []string{}

	for _, fd := range p.Root.FuncDefs {
		sls := fd.print()
		lines = append(lines, sls...)
		lines = append(lines, "")
	}

	for _, st := range p.Root.Statements {
		ls := st.print()
		lines = append(lines, ls...)
	}

	return strings.Join(lines, "\n")
}
