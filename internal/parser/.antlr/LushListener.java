// Generated from /home/creichlin/workspaces/bytewright/lush/internal/parser/Lush.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link LushParser}.
 */
public interface LushListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link LushParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(LushParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(LushParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(LushParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(LushParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#if}.
	 * @param ctx the parse tree
	 */
	void enterIf(LushParser.IfContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#if}.
	 * @param ctx the parse tree
	 */
	void exitIf(LushParser.IfContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(LushParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(LushParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(LushParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(LushParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#funcStatement}.
	 * @param ctx the parse tree
	 */
	void enterFuncStatement(LushParser.FuncStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#funcStatement}.
	 * @param ctx the parse tree
	 */
	void exitFuncStatement(LushParser.FuncStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#indexExp}.
	 * @param ctx the parse tree
	 */
	void enterIndexExp(LushParser.IndexExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#indexExp}.
	 * @param ctx the parse tree
	 */
	void exitIndexExp(LushParser.IndexExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#unaryExp}.
	 * @param ctx the parse tree
	 */
	void enterUnaryExp(LushParser.UnaryExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#unaryExp}.
	 * @param ctx the parse tree
	 */
	void exitUnaryExp(LushParser.UnaryExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#mulExp}.
	 * @param ctx the parse tree
	 */
	void enterMulExp(LushParser.MulExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#mulExp}.
	 * @param ctx the parse tree
	 */
	void exitMulExp(LushParser.MulExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#addExp}.
	 * @param ctx the parse tree
	 */
	void enterAddExp(LushParser.AddExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#addExp}.
	 * @param ctx the parse tree
	 */
	void exitAddExp(LushParser.AddExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#relExp}.
	 * @param ctx the parse tree
	 */
	void enterRelExp(LushParser.RelExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#relExp}.
	 * @param ctx the parse tree
	 */
	void exitRelExp(LushParser.RelExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#andExp}.
	 * @param ctx the parse tree
	 */
	void enterAndExp(LushParser.AndExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#andExp}.
	 * @param ctx the parse tree
	 */
	void exitAndExp(LushParser.AndExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#orExp}.
	 * @param ctx the parse tree
	 */
	void enterOrExp(LushParser.OrExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#orExp}.
	 * @param ctx the parse tree
	 */
	void exitOrExp(LushParser.OrExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#condExp}.
	 * @param ctx the parse tree
	 */
	void enterCondExp(LushParser.CondExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#condExp}.
	 * @param ctx the parse tree
	 */
	void exitCondExp(LushParser.CondExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterAtom(LushParser.AtomContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitAtom(LushParser.AtomContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#func}.
	 * @param ctx the parse tree
	 */
	void enterFunc(LushParser.FuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#func}.
	 * @param ctx the parse tree
	 */
	void exitFunc(LushParser.FuncContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#var}.
	 * @param ctx the parse tree
	 */
	void enterVar(LushParser.VarContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#var}.
	 * @param ctx the parse tree
	 */
	void exitVar(LushParser.VarContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#envVar}.
	 * @param ctx the parse tree
	 */
	void enterEnvVar(LushParser.EnvVarContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#envVar}.
	 * @param ctx the parse tree
	 */
	void exitEnvVar(LushParser.EnvVarContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(LushParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(LushParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#string}.
	 * @param ctx the parse tree
	 */
	void enterString(LushParser.StringContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#string}.
	 * @param ctx the parse tree
	 */
	void exitString(LushParser.StringContext ctx);
	/**
	 * Enter a parse tree produced by {@link LushParser#number}.
	 * @param ctx the parse tree
	 */
	void enterNumber(LushParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link LushParser#number}.
	 * @param ctx the parse tree
	 */
	void exitNumber(LushParser.NumberContext ctx);
}