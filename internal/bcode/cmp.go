package bcode

type Equal struct {
	Left  Value
	Right Value
}

func (e *Equal) Print() string {
	return e.Left.Print() + " == " + e.Right.Print()
}

type NotEqual struct {
	Left  Value
	Right Value
}

func (n *NotEqual) Print() string {
	return n.Left.Print() + " != " + n.Right.Print()
}
