package bcode

import (
	"fmt"
	"strconv"
	"strings"
)

type tmpVar interface {
	nextVar() int
}

type Block struct {
	tmpVar   tmpVar
	Commands []Command
}

func (b *Block) set(name string, val Atom) {
	b.Commands = append(b.Commands, &Assignment{Name: name, Value: val})
}

func (b *Block) setTmp(val Atom) *VarValue {
	v := b.nextTmp()
	b.Commands = append(b.Commands, &Assignment{Name: v.Name, Value: val})
	return v
}

func (b *Block) nextTmp() *VarValue {
	return &VarValue{Name: "var_" + strconv.Itoa(b.tmpVar.nextVar())}
}

func (b *Block) add(cmd Command) {
	b.Commands = append(b.Commands, cmd)
}

func (b *Block) Print(indent int) string {
	r := ""

	ind := strings.Repeat("  ", indent)

	for _, c := range b.Commands {
		switch cmd := c.(type) {
		case *Assignment:
			r += ind + cmd.Name + " = " + cmd.Value.Print() + "\n"
		case *Func:
			r += ind + cmd.Print() + "\n"

		case *If:
			r += ind + "if " + cmd.Condition.Print() + " {\n"
			r += cmd.Block.Print(indent + 2)
			r += ind + "}\n"

		case *While:
			r += ind + "while " + cmd.Condition.Print() + " {\n"
			r += cmd.Block.Print(indent + 2)
			r += ind + "}\n"

		default:
			panic(fmt.Sprintf("unknown command %T", c))
		}
	}
	return r
}
