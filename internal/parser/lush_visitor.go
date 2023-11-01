// Code generated from internal/parser/Lush.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lush

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LushParser.
type LushVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LushParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by LushParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LushParser#if.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by LushParser#for.
	VisitFor(ctx *ForContext) interface{}

	// Visit a parse tree produced by LushParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LushParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by LushParser#funcStatement.
	VisitFuncStatement(ctx *FuncStatementContext) interface{}

	// Visit a parse tree produced by LushParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by LushParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by LushParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by LushParser#var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by LushParser#envVar.
	VisitEnvVar(ctx *EnvVarContext) interface{}

	// Visit a parse tree produced by LushParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by LushParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by LushParser#bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by LushParser#array.
	VisitArray(ctx *ArrayContext) interface{}
}
