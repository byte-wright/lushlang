package ast

import "github.com/byte-wright/lush/internal/common"

type Assignment struct {
	Location    *common.Location
	Names       []string
	Expressions []Expression
}
