// Code generated from internal/parser/Lush.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lush

import "github.com/antlr4-go/antlr/v4"

type BaseLushVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLushVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitPackage(ctx *PackageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitFuncDef(ctx *FuncDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitExternalFuncDef(ctx *ExternalFuncDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitIf(ctx *IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitFor(ctx *ForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitFuncStatement(ctx *FuncStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitEnvVar(ctx *EnvVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLushVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
