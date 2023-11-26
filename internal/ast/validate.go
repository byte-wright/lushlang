package ast

import "github.com/byte-wright/lush/internal/common"

type ASTError struct {
	Message string
}

func (a *ASTError) Error() string {
	return a.Message
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
			if len(st.Names) != len(rightTypes) {
				v.errors = append(v.errors, &ASTError{Message: "Wrong number of expressions"})
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
