package ast

type Block struct {
	Statements []Statement
	FuncDefs   []*FuncDef
}
