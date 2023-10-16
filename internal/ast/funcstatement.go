package ast

type FuncStatement struct {
	Func *Func
}

func (f *FuncStatement) print() []string {
	return []string{f.Func.print()}
}
