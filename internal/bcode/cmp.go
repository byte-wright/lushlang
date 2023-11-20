package bcode

import "github.com/byte-wright/lush/internal/common"

type Equal struct {
	Left  Value
	Right Value
}

func (*Equal) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

type LessThan struct {
	Left  Value
	Right Value
}

func (*LessThan) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

type LessThanEqual struct {
	Left  Value
	Right Value
}

func (*LessThanEqual) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

type GreaterThan struct {
	Left  Value
	Right Value
}

func (*GreaterThan) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

type GreaterThanEqual struct {
	Left  Value
	Right Value
}

func (*GreaterThanEqual) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}

type NotEqual struct {
	Left  Value
	Right Value
}

func (*NotEqual) Type() common.Type {
	return &common.BasicType{Type: common.Bool}
}
