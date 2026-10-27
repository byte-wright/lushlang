package bcode

type If struct {
	Condition *VarValue
	Block     *Block
	// Else is nil without an else
	Else *Block
}

type While struct {
	Condition *VarValue
	Block     *Block
}

type Return struct {
	Values []*VarValue
}
