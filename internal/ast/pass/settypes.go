package astpass

import (
	"fmt"

	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

type setTypesPass struct {
	ast    *ast.Program
	errors []*ast.ASTError
	block  *ast.Block
}

func SetTypes(ast *ast.Program) []*ast.ASTError {
	te := &setTypesPass{
		ast: ast,
	}
	Visit(ast, te)

	return te.errors
}

func (t *setTypesPass) BeforeBlock(block *ast.Block) {
	t.block = block
}

func (t *setTypesPass) AfterBlock(block *ast.Block) {
}

func (t *setTypesPass) Statement(statement ast.Statement) {
	switch st := statement.(type) {
	case *ast.FuncStatement:
		t.Expression(st.Func)

	case *ast.Assignment:
		types := []common.Type{}
		for _, exp := range st.Expressions {
			types = append(types, exp.Types()...)
		}

		for i, target := range st.Targets {
			if i >= len(types) {
				continue
			}

			err := t.block.SetVar(target.Name, types[i])
			if err != nil {
				t.errors = append(t.errors, &ast.ASTError{Location: target.Location, Message: err.Error()})
			}
		}
	}
}

func (t *setTypesPass) Expression(expression ast.Expression) {
	switch exp := expression.(type) {
	case *ast.Func:
		tps, err := t.getFunc("", exp.Name, exp.Parameters)
		if err != nil {
			t.errors = append(t.errors, &ast.ASTError{Message: err.Error()})
		} else {
			exp.ReturnTypes = tps
		}

	case *ast.Method:
		tps, err := t.getFunc(exp.Namespace, exp.Func.Name, exp.Func.Parameters)
		if err != nil {
			t.errors = append(t.errors, &ast.ASTError{Message: err.Error()})
		} else {
			exp.Func.ReturnTypes = tps
		}

	case *ast.Var:
		v := t.block.GetVar(exp.Name)
		if v != nil {
			exp.Type = v.Type
		}
	}
}

func (t *setTypesPass) getFunc(namespace, name string, params []ast.Expression) ([]common.Type, error) {
	if name == "println" {
		return []common.Type{}, nil
	}
	if name == "len" {
		return []common.Type{&common.BasicType{Type: common.Int}}, nil
	}
	if name == "execVar" {
		return []common.Type{
			&common.BasicType{Type: common.String},
			&common.BasicType{Type: common.String},
			&common.BasicType{Type: common.Int},
		}, nil
	}
	if name == "append" {
		if len(params) == 0 {
			return nil, fmt.Errorf("append needs parameters")
		}

		ptps := params[0].Types()
		if len(ptps) != 1 {
			return nil, fmt.Errorf("first param in append must return one value")
		}

		if !ptps[0].IsArray() {
			return nil, fmt.Errorf("first param in append must be an array")
		}

		return ptps, nil

	}

	if namespace == "" {
		for _, fd := range t.ast.FuncDefs {
			if fd.Name == name {
				return fd.Returns, nil
			}
		}

		return nil, fmt.Errorf("function %v not found", name)
	}

	for _, pkg := range t.ast.Packages() {
		if pkg.Name == namespace {
			for _, fd := range pkg.ExternalFuncDefs {
				if fd.Name == name {
					return fd.Returns, nil
				}
			}

			for _, fd := range pkg.FuncDefs {
				if fd.Name == name {
					return fd.Returns, nil
				}
			}

			return nil, fmt.Errorf("function %v in namespace %v not found", name, namespace)
		}
	}

	return nil, fmt.Errorf("namespace %v not found", namespace)
}
