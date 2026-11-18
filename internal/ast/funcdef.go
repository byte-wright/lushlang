package ast

import (
	"github.com/byte-wright/lush/internal/common"
)

type Param struct {
	Location *common.Location
	Name     string
	Type     common.Type
}

type FuncDef struct {
	Name     string
	Location *common.Location
	Params   []*Param
	// Returns doesn't hold the error of a function that can fail
	Returns []common.Type
	CanFail bool
	Body    *Block
	File    *File
}

type ExternalFuncDef struct {
	Name     string
	Location *common.Location
	Params   []*Param
	Returns  []common.Type
	CanFail  bool
	Code     string
}
