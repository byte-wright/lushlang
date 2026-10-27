package ast

import "github.com/byte-wright/lush/internal/common"

type For struct {
	Initial   *Assignment
	Condition Expression
	Each      *Assignment
	Body      *Block
	Location  *common.Location
}
