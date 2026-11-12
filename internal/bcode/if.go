package bcode

type If struct {
	Condition *VarValue
	Block     *Block
	// Else is nil without an else
	Else *Block
}

// While runs Check before each iteration, and stops when Condition isn't true. Both are
// nil for an endless loop.
type While struct {
	Check     *Block
	Condition *VarValue
	Block     *Block
}

// ForEach runs Block with Element set to each element of Array. The elements are read
// once, before the first iteration.
type ForEach struct {
	Element *VarValue
	Array   *VarValue
	Block   *Block
}

type Break struct{}

// Continue goes on with the next check of the loop
type Continue struct{}

type Return struct {
	Values []*VarValue
}
