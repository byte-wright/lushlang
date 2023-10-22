package bcode

type And struct {
	Left  Value
	Right Value
}

func (a *And) Print() string {
	return a.Left.Print() + " && " + a.Right.Print()
}

type Not struct {
	Expression Value
}

func (n *Not) Print() string {
	return "!" + n.Expression.Print()
}
