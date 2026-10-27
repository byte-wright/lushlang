package ast

import "github.com/byte-wright/lush/internal/common"

type ReturnStatement struct {
	Expressions []Expression
	Location    *common.Location
}
