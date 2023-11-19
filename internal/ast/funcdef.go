package ast

import (
	"github.com/byte-wright/lush/internal/common"
)

type Param struct {
	Name string
	Type common.Type
}

type FuncDef struct {
	Name    string
	Params  []*Param
	Returns []common.Type
	Body    *Block
}

type ExternalFuncDef struct {
	Name    string
	Params  []*Param
	Returns []common.Type
	Code    string
}
