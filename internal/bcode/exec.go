package bcode

type ExecVar struct {
	Command    Value
	Parameters []Value
	Stdout     *VarValue
	Stderr     *VarValue
	Err        *VarValue
}
