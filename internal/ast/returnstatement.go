package ast

import "github.com/byte-wright/lush/internal/common"

type ReturnStatement struct {
	Expressions []Expression
	Location    *common.Location
}

// Fail returns the error Message to the caller, at the top level it ends the script.
type Fail struct {
	Message  Expression
	Location *common.Location
}
