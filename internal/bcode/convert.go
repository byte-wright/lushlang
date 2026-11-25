package bcode

import "github.com/byte-wright/lush/internal/common"

// Convert is a conversion that can't fail: a value to a string, an array element by element
// to a []string, or an array to its text, like [1 2 3]. The ones that parse a string are
// calls of builtins, they can fail.
type Convert struct {
	Value *VarValue
	To    common.Type
}

func (c *Convert) Type() common.Type {
	return c.To
}
