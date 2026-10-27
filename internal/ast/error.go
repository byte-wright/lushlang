package ast

import (
	"strconv"

	"github.com/byte-wright/lush/internal/common"
)

type ASTError struct {
	Location *common.Location
	Message  string
}

func (a *ASTError) Error() string {
	l := ""
	if a.Location != nil {
		l = a.Location.File + " Ln:" + strconv.Itoa(a.Location.LineFrom) + " Col:" + strconv.Itoa(a.Location.ColFrom) + " "
	}
	return l + a.Message
}
