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

	vars   map[string]*VarValue
	parent *Block
}

// subBlock creates a child block, will not be added to parent block, must be done manually
func (b *Block) subBlock() *Block {
	return &Block{
		tmpVar: b.tmpVar,
		vars:   map[string]*VarValue{},
		parent: b,
	}
}

func (b *Block) set(v *VarValue, val Atom) {
	current := b.getVisibleVar(v.Name)
	if current == nil {
		b.registerVar(v)
	}
	b.Commands = append(b.Commands, &Assignment{Var: v, Value: val})
}

func (b *Block) setTmp(val Atom) *VarValue {
	v := b.nextTmp(val.Type())
	b.registerVar(v)
	b.Commands = append(b.Commands, &Assignment{Var: v, Value: val})
	return v
}

func (b *Block) nextTmp(t Type) *VarValue {
	return &VarValue{
		Name: "var_" + strconv.Itoa(b.tmpVar.nextVar()),
		T:    t,
	}
}

func (b *Block) registerVar(v *VarValue) {
	b.vars[v.Name] = v
}

func (b *Block) getVisibleVar(name string) *VarValue {
	v, has := b.vars[name]
	if has {
		return v
	}

	if b.parent != nil {
		return b.parent.getVisibleVar(name)
	}

	return nil
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
			r += ind + cmd.Var.Name + " = " + cmd.Value.Print() + "\n"
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
