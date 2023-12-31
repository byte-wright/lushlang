package bcode

import (
	"strconv"

	"github.com/byte-wright/lush/internal/common"
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

func (b *Block) setTmps(vals ...Atom) []Value {
	vars := []Value{}

	for _, val := range vals {
		v := b.nextTmp(val.Type())
		b.registerVar(v)
		b.Commands = append(b.Commands, &Assignment{Var: v, Value: val})
		vars = append(vars, v)
	}

	return vars
}

// asVar takes a value and returns it if it's already a variable
// if not it assigns it to a variable and returns that variable.
func (b *Block) asVar(val Value) *VarValue {
	vVal, isVar := val.(*VarValue)
	if isVar {
		return vVal
	}
	return b.setTmp(val)
}

func (b *Block) nextTmp(t common.Type) *VarValue {
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
