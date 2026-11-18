package ast

import "github.com/byte-wright/lush/internal/common"

type FuncStatement struct {
	Call Expression // *Func or *Method
	// Catch is nil without one
	Catch *Catch
}

// Catch handles the error of the statement's call, Block runs when it fails. On a failure,
// an assignment doesn't assign its variables.
type Catch struct {
	Error    *AssignmentTarget
	Block    *Block
	Location *common.Location
}
