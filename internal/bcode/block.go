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

// identifiers can't start with an uppercase letter, so they never collide with a temp
func (b *Block) nextTmp(t common.Type) *VarValue {
	return &VarValue{
		Name: "L_t" + strconv.Itoa(b.tmpVar.nextVar()),
		T:    t,
	}
}

// without the prefix, a variable named like an environment variable (http_proxy) would
// hide it from the functions it calls and from child processes
func varName(identifier string) string {
	return "V_" + identifier
}

// a blank parameter still takes its position in the arguments, so it gets a temp
func (b *Block) paramVar(name string, t common.Type) *VarValue {
	if name == "_" {
		return b.nextTmp(t)
	}

	v := &VarValue{Name: varName(name), T: t}
	b.registerVar(v)
	return v
}

// the bash code of an external function reads its parameters by their names
func (b *Block) externalParamVar(name string, t common.Type) *VarValue {
	if name == "_" {
		return b.nextTmp(t)
	}

	return &VarValue{Name: name, T: t}
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
