package ast

import "github.com/byte-wright/lush/internal/common"

type If struct {
	Condition Expression
	If        *Block
	// Else is nil without an else, an else if is an Else block that only holds that if
	Else     *Block
	Location *common.Location
}
