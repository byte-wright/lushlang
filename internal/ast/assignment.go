package ast

import "github.com/byte-wright/lush/internal/common"

type Assignment struct {
	Location    *common.Location
	Targets     []*AssignmentTarget
	Expressions []Expression
	Global      bool
	// Catch is nil without one
	Catch *Catch
}

type AssignmentTarget struct {
	Location *common.Location
	Name     string
	GlobalOf *Package
	// Element is the x[i] of x[i] = 1, Name is then the array. It's nil for a variable.
	Element *Index
}

// TargetValue is the current value of an element target, x[i] += 1 reads it with the
// index that the target evaluates, so the index is evaluated once.
type TargetValue struct {
	Target *AssignmentTarget
}

type GlobalDef struct {
	Assignment *Assignment
	File       *File
}
