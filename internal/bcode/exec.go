package bcode

import "strings"

type ExecVar struct {
	Command    Value
	Parameters []Value
	Stdout     *VarValue
	Stderr     *VarValue
	Err        *VarValue
}

func (e *ExecVar) Print() string {
	params := []string{e.Command.Print()}

	for _, p := range e.Parameters {
		params = append(params, p.Print())
	}

	return "execVar(" + strings.Join(params, ", ") + ")"
}
