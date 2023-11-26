package ast

import (
	"github.com/byte-wright/lush/internal/common"
)

var (
	boolT   = []common.Type{&common.BasicType{Type: common.Bool}}
	stringT = []common.Type{&common.BasicType{Type: common.String}}
	intT    = []common.Type{&common.BasicType{Type: common.Int}}
	noneT   = []common.Type{&common.BasicType{Type: common.None}}
)

func tl(t common.Type) []common.Type {
	return []common.Type{t}
}

func (n *Number) Types() []common.Type {
	return intT
}

func (s *String) Types() []common.Type {
	return stringT
}

func (b *Bool) Types() []common.Type {
	return boolT
}

func (v *Var) Types() []common.Type {
	if v.Type == nil {
		return noneT
	}
	return tl(v.Type)
}

func (v *EnvVar) Types() []common.Type {
	return tl(&common.BasicType{Type: common.String})
}

func (a *Add) Types() []common.Type {
	return a.A.Types()
}

func (a *Sub) Types() []common.Type {
	return a.A.Types()
}

func (a *Mul) Types() []common.Type {
	return a.A.Types()
}

func (a *Div) Types() []common.Type {
	return a.A.Types()
}

func (a *Mod) Types() []common.Type {
	return a.A.Types()
}

func (a *And) Types() []common.Type {
	return boolT
}

func (a *Or) Types() []common.Type {
	return boolT
}

func (a *LessThan) Types() []common.Type {
	return boolT
}

func (a *LessThanEqual) Types() []common.Type {
	return boolT
}

func (a *GreaterThan) Types() []common.Type {
	return boolT
}

func (a *GreaterThanEqual) Types() []common.Type {
	return boolT
}

func (a *Equal) Types() []common.Type {
	return boolT
}

func (a *NotEqual) Types() []common.Type {
	return boolT
}

func (s *Slice) Types() []common.Type {
	return s.Value.Types()
}

func (i *Index) Types() []common.Type {
	tps := i.Value.Types()
	f := tps[0]

	if f.IsArray() {
		return tl(f.(*common.ArrayType).ElementType)
	}

	return noneT
}

func (a *Ternary) Types() []common.Type {
	return a.True.Types()
}

func (s *Array) Types() []common.Type {
	if s.Type != common.None {
		return tl(&common.ArrayType{ElementType: &common.BasicType{Type: s.Type}})
	}

	return tl(&common.ArrayType{ElementType: s.Values[0].Types()[0]})
}

func (s *Func) Types() []common.Type {
	return s.ReturnTypes
}

func (s *Method) Types() []common.Type {
	return s.Func.Types()
}
