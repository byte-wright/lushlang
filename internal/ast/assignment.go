package ast

import "github.com/byte-wright/lush/internal/common"

type Assignment struct {
	Location    *common.Location
	Targets     []*AssignmentTarget
	Expressions []Expression
}

type AssignmentTarget struct {
	Location *common.Location
	Name     string
}
