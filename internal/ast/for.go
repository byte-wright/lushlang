package ast

type For struct {
	Initial   *Assignment
	Condition Expression
	Each      *Assignment
	Body      *Block
}

func (f *For) print() []string {
	bd := f.Body.print()

	initial := f.Initial.print()
	if len(initial) != 1 {
		panic("assignment print must always return 1 line")
	}

	each := f.Each.print()
	if len(each) != 1 {
		panic("assignment print must always return 1 line")
	}

	bd[0] = "for " + initial[0] + "; " + f.Condition.print() + "; " + each[0] + " " + bd[0]

	return bd
}
