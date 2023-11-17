package ast

type Func struct {
	Name       string
	Parameters []Expression
}

func (f *Func) print() string {
	params := ""
	for _, x := range f.Parameters {
		params += ", " + x.print()
	}

	params = params[2:]
	return f.Name + "(" + params + ")"
}

type Method struct {
	Expression Expression
	Func       *Func
}

func (m *Method) print() string {
	return m.Expression.print() + "." + m.Func.print()
}
