package bcode

type Equal struct {
	Left  Value
	Right Value
}

func (e *Equal) Print() string {
	return e.Left.Print() + " == " + e.Right.Print()
}

type LessThan struct {
	Left  Value
	Right Value
}

func (l *LessThan) Print() string {
	return l.Left.Print() + " < " + l.Right.Print()
}

type LessThanEqual struct {
	Left  Value
	Right Value
}

func (l *LessThanEqual) Print() string {
	return l.Left.Print() + " <= " + l.Right.Print()
}

type GreaterThan struct {
	Left  Value
	Right Value
}

func (g *GreaterThan) Print() string {
	return g.Left.Print() + " > " + g.Right.Print()
}

type GreaterThanEqual struct {
	Left  Value
	Right Value
}

func (g *GreaterThanEqual) Print() string {
	return g.Left.Print() + " >= " + g.Right.Print()
}

type NotEqual struct {
	Left  Value
	Right Value
}

func (n *NotEqual) Print() string {
	return n.Left.Print() + " != " + n.Right.Print()
}
