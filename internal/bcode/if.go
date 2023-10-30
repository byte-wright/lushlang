package bcode

type If struct {
	Condition *VarValue
	Block     *Block
}

type While struct {
	Condition *VarValue
	Block     *Block
}
