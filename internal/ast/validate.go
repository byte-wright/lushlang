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

type Validator struct {
	ast    *Program
	errors []*ASTError
}

func Validate(p *Program) []*ASTError {
	validator := &Validator{
		ast:    p,
		errors: []*ASTError{},
	}

	validator.validate()

	return validator.errors
}

func (v *Validator) validate() {
	v.validateBlock(v.ast.Root)
}

func (v *Validator) validateBlock(block *Block) {
	for _, statement := range block.Statements {
		switch st := statement.(type) {
		case *Assignment:
			rightTypes := typesForExpressions(st.Expressions)
			if len(st.Targets) != len(rightTypes) {
				v.errors = append(v.errors, &ASTError{Location: st.Location, Message: "Wrong number of expressions"})
				continue
			}
		}
	}
}

func typesForExpressions(xps []Expression) []common.Type {
	tps := []common.Type{}

	for _, xp := range xps {
		tps = append(tps, xp.Types()...)
	}

	return tps
}
