package astpass

import (
	"github.com/byte-wright/lush/internal/ast"
)

type setParentPass struct {
	ast    *ast.Program
	errors []*ast.ASTError
}

func SetParent(ast *ast.Program) []*ast.ASTError {
	te := &setParentPass{
		ast: ast,
	}
	te.eval()

	return te.errors
}

func (t *setParentPass) eval() {
	t.evalBlock(t.ast.Root)
}

func (t *setParentPass) evalBlock(block *ast.Block) {
	for _, statement := range block.Statements {
		switch st := statement.(type) {
		case *ast.For:
			st.Body.Parent = block
			t.evalBlock(st.Body)

		case *ast.If:
			st.If.Parent = block
			t.evalBlock(st.If)
		}
	}
}
