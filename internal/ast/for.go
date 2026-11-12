package ast

import "github.com/byte-wright/lush/internal/common"

// For is a loop with an optional Initial and Each, and an endless loop without a Condition
type For struct {
	Initial   *Assignment
	Condition Expression
	Each      *Assignment
	Body      *Block
	Location  *common.Location
}

type Break struct {
	Location *common.Location
}

type Continue struct {
	Location *common.Location
}

// ForIn runs Body for each element of the array Over, or for each int from From (0 without
// it) to just before Over. Index is nil without one.
type ForIn struct {
	Value    *AssignmentTarget
	Index    *AssignmentTarget
	From     Expression
	Over     Expression
	Body     *Block
	Location *common.Location
}
