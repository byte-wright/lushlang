package bcode

import (
	"github.com/byte-wright/lush/internal/common"
)

type Func struct {
	Namespace  string
	Name       string
	Parameters []Value
	Return     []*VarValue
	// Try makes the current function fail with the error of a failing call, main ends the
	// script with it
	Try bool
}

func (f *Func) Type() common.Type {
	return &common.BasicType{Type: common.String}
}

func (f *Func) FullName() string {
	return funcName(f.Namespace, f.Name)
}

func funcName(namespace, name string) string {
	if namespace == "" {
		return name
	}

	return namespace + "_" + name
}

// Catch calls a function that can fail. When it fails, Block runs, otherwise Then runs
// after the results are read.
type Catch struct {
	Call  *Func
	Block *Block
	Then  *Block
}

// Fail makes the current function fail with Message, main ends the script with it.
type Fail struct {
	Message Value
}

// ErrorValue is the message of the call that failed, Block of a Catch can read it.
type ErrorValue struct{}

func (e *ErrorValue) Type() common.Type {
	return &common.BasicType{Type: common.String}
}

// Params binds the arguments of a function call to its parameters.
type Params struct {
	Vars []*VarValue
}
