// Generated from /home/creichlin/workspaces/bytewright/lush/internal/parser/Lush.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class LushParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WHITESPACE=1, LAND=2, LOR=3, NOT=4, MINUS=5, PLUS=6, MUL=7, DIV=8, MOD=9, 
		LT=10, LTE=11, GT=12, GTE=13, EQ=14, NEQ=15, IF=16, LPAREN=17, RPAREN=18, 
		LCUR=19, RCUR=20, LSQ=21, RSQ=22, COMMA=23, QUESTION=24, COLON=25, ENVVAR=26, 
		ID=27, ASSIGN=28, STRING=29, NUMBER=30;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_if = 2, RULE_block = 3, RULE_assignment = 4, 
		RULE_funcStatement = 5, RULE_expression = 6, RULE_atom = 7, RULE_func = 8, 
		RULE_var = 9, RULE_envVar = 10, RULE_value = 11, RULE_string = 12, RULE_number = 13;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "if", "block", "assignment", "funcStatement", 
			"expression", "atom", "func", "var", "envVar", "value", "string", "number"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'&&'", "'||'", "'!'", "'-'", "'+'", "'*'", "'/'", "'%'", 
			"'<'", "'<='", "'>'", "'>='", "'=='", "'!='", "'if'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "','", "'?'", "':'", null, null, "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WHITESPACE", "LAND", "LOR", "NOT", "MINUS", "PLUS", "MUL", "DIV", 
			"MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ", "IF", "LPAREN", "RPAREN", 
			"LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION", "COLON", "ENVVAR", 
			"ID", "ASSIGN", "STRING", "NUMBER"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Lush.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public LushParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(LushParser.EOF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IF || _la==ID) {
				{
				{
				setState(28);
				statement();
				}
				}
				setState(33);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(34);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public FuncStatementContext funcStatement() {
			return getRuleContext(FuncStatementContext.class,0);
		}
		public IfContext if_() {
			return getRuleContext(IfContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(39);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				assignment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(37);
				funcStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(38);
				if_();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(LushParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if; }
	}

	public final IfContext if_() throws RecognitionException {
		IfContext _localctx = new IfContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_if);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(IF);
			setState(42);
			expression(0);
			setState(43);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LCUR() { return getToken(LushParser.LCUR, 0); }
		public TerminalNode RCUR() { return getToken(LushParser.RCUR, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			match(LCUR);
			setState(49);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IF || _la==ID) {
				{
				{
				setState(46);
				statement();
				}
				}
				setState(51);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(52);
			match(RCUR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(LushParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(LushParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(ID);
			setState(55);
			match(ASSIGN);
			setState(56);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncStatementContext extends ParserRuleContext {
		public FuncContext func() {
			return getRuleContext(FuncContext.class,0);
		}
		public FuncStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcStatement; }
	}

	public final FuncStatementContext funcStatement() throws RecognitionException {
		FuncStatementContext _localctx = new FuncStatementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_funcStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			func();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext ternary;
		public ExpressionContext slice;
		public Token unary_op;
		public Token mul_op;
		public Token add_op;
		public Token rel_op;
		public AtomContext atom() {
			return getRuleContext(AtomContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MINUS() { return getToken(LushParser.MINUS, 0); }
		public TerminalNode NOT() { return getToken(LushParser.NOT, 0); }
		public TerminalNode MUL() { return getToken(LushParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(LushParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(LushParser.MOD, 0); }
		public TerminalNode PLUS() { return getToken(LushParser.PLUS, 0); }
		public TerminalNode EQ() { return getToken(LushParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(LushParser.NEQ, 0); }
		public TerminalNode LT() { return getToken(LushParser.LT, 0); }
		public TerminalNode LTE() { return getToken(LushParser.LTE, 0); }
		public TerminalNode GT() { return getToken(LushParser.GT, 0); }
		public TerminalNode GTE() { return getToken(LushParser.GTE, 0); }
		public TerminalNode LAND() { return getToken(LushParser.LAND, 0); }
		public TerminalNode LOR() { return getToken(LushParser.LOR, 0); }
		public TerminalNode QUESTION() { return getToken(LushParser.QUESTION, 0); }
		public TerminalNode COLON() { return getToken(LushParser.COLON, 0); }
		public TerminalNode LSQ() { return getToken(LushParser.LSQ, 0); }
		public TerminalNode RSQ() { return getToken(LushParser.RSQ, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
			case ENVVAR:
			case ID:
			case STRING:
			case NUMBER:
				{
				setState(61);
				atom();
				}
				break;
			case NOT:
			case MINUS:
				{
				setState(62);
				((ExpressionContext)_localctx).unary_op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==NOT || _la==MINUS) ) {
					((ExpressionContext)_localctx).unary_op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(63);
				expression(8);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(96);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(94);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(66);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(67);
						((ExpressionContext)_localctx).mul_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 896L) != 0)) ) {
							((ExpressionContext)_localctx).mul_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(68);
						expression(8);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(69);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(70);
						((ExpressionContext)_localctx).add_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MINUS || _la==PLUS) ) {
							((ExpressionContext)_localctx).add_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(71);
						expression(7);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(72);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(73);
						((ExpressionContext)_localctx).rel_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 64512L) != 0)) ) {
							((ExpressionContext)_localctx).rel_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(74);
						expression(6);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(75);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(76);
						match(LAND);
						setState(77);
						expression(5);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(78);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(79);
						match(LOR);
						setState(80);
						expression(4);
						}
						break;
					case 6:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.ternary = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(81);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(82);
						match(QUESTION);
						setState(83);
						expression(0);
						setState(84);
						match(COLON);
						setState(85);
						expression(3);
						}
						break;
					case 7:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						_localctx.slice = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(87);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(88);
						match(LSQ);
						setState(89);
						expression(0);
						setState(90);
						match(COLON);
						setState(91);
						expression(0);
						setState(92);
						match(RSQ);
						}
						break;
					}
					} 
				}
				setState(98);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AtomContext extends ParserRuleContext {
		public ExpressionContext group;
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public EnvVarContext envVar() {
			return getRuleContext(EnvVarContext.class,0);
		}
		public VarContext var() {
			return getRuleContext(VarContext.class,0);
		}
		public FuncContext func() {
			return getRuleContext(FuncContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(LushParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(LushParser.RPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_atom);
		try {
			setState(107);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(99);
				value();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(100);
				envVar();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(101);
				var();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(102);
				func();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(103);
				match(LPAREN);
				setState(104);
				((AtomContext)_localctx).group = expression(0);
				setState(105);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(LushParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(LushParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(LushParser.RPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(LushParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(LushParser.COMMA, i);
		}
		public FuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func; }
	}

	public final FuncContext func() throws RecognitionException {
		FuncContext _localctx = new FuncContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_func);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(ID);
			setState(110);
			match(LPAREN);
			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1812070448L) != 0)) {
				{
				setState(111);
				expression(0);
				setState(116);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(112);
					match(COMMA);
					setState(113);
					expression(0);
					}
					}
					setState(118);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(121);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(LushParser.ID, 0); }
		public VarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var; }
	}

	public final VarContext var() throws RecognitionException {
		VarContext _localctx = new VarContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnvVarContext extends ParserRuleContext {
		public TerminalNode ENVVAR() { return getToken(LushParser.ENVVAR, 0); }
		public EnvVarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_envVar; }
	}

	public final EnvVarContext envVar() throws RecognitionException {
		EnvVarContext _localctx = new EnvVarContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_envVar);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(ENVVAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValueContext extends ParserRuleContext {
		public StringContext string() {
			return getRuleContext(StringContext.class,0);
		}
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_value);
		try {
			setState(129);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(127);
				string();
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(128);
				number();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(LushParser.STRING, 0); }
		public StringContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string; }
	}

	public final StringContext string() throws RecognitionException {
		StringContext _localctx = new StringContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_string);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NumberContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(LushParser.NUMBER, 0); }
		public NumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number; }
	}

	public final NumberContext number() throws RecognitionException {
		NumberContext _localctx = new NumberContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_number);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			match(NUMBER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 6:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 7);
		case 1:
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		case 4:
			return precpred(_ctx, 3);
		case 5:
			return precpred(_ctx, 2);
		case 6:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u001e\u0088\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004"+
		"\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007"+
		"\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b"+
		"\u0002\f\u0007\f\u0002\r\u0007\r\u0001\u0000\u0005\u0000\u001e\b\u0000"+
		"\n\u0000\f\u0000!\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0003\u0001(\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0005\u00030\b\u0003\n\u0003\f\u0003"+
		"3\t\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0003\u0006A\b\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0005\u0006_\b\u0006\n\u0006\f\u0006b\t\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0003\u0007l\b\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0005\bs\b\b\n\b\f\bv\t\b\u0003\bx\b\b\u0001\b\u0001\b\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0003\u000b\u0082\b\u000b\u0001"+
		"\f\u0001\f\u0001\r\u0001\r\u0001\r\u0000\u0001\f\u000e\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u0000\u0004\u0001"+
		"\u0000\u0004\u0005\u0001\u0000\u0007\t\u0001\u0000\u0005\u0006\u0001\u0000"+
		"\n\u000f\u008c\u0000\u001f\u0001\u0000\u0000\u0000\u0002\'\u0001\u0000"+
		"\u0000\u0000\u0004)\u0001\u0000\u0000\u0000\u0006-\u0001\u0000\u0000\u0000"+
		"\b6\u0001\u0000\u0000\u0000\n:\u0001\u0000\u0000\u0000\f@\u0001\u0000"+
		"\u0000\u0000\u000ek\u0001\u0000\u0000\u0000\u0010m\u0001\u0000\u0000\u0000"+
		"\u0012{\u0001\u0000\u0000\u0000\u0014}\u0001\u0000\u0000\u0000\u0016\u0081"+
		"\u0001\u0000\u0000\u0000\u0018\u0083\u0001\u0000\u0000\u0000\u001a\u0085"+
		"\u0001\u0000\u0000\u0000\u001c\u001e\u0003\u0002\u0001\u0000\u001d\u001c"+
		"\u0001\u0000\u0000\u0000\u001e!\u0001\u0000\u0000\u0000\u001f\u001d\u0001"+
		"\u0000\u0000\u0000\u001f \u0001\u0000\u0000\u0000 \"\u0001\u0000\u0000"+
		"\u0000!\u001f\u0001\u0000\u0000\u0000\"#\u0005\u0000\u0000\u0001#\u0001"+
		"\u0001\u0000\u0000\u0000$(\u0003\b\u0004\u0000%(\u0003\n\u0005\u0000&"+
		"(\u0003\u0004\u0002\u0000\'$\u0001\u0000\u0000\u0000\'%\u0001\u0000\u0000"+
		"\u0000\'&\u0001\u0000\u0000\u0000(\u0003\u0001\u0000\u0000\u0000)*\u0005"+
		"\u0010\u0000\u0000*+\u0003\f\u0006\u0000+,\u0003\u0006\u0003\u0000,\u0005"+
		"\u0001\u0000\u0000\u0000-1\u0005\u0013\u0000\u0000.0\u0003\u0002\u0001"+
		"\u0000/.\u0001\u0000\u0000\u000003\u0001\u0000\u0000\u00001/\u0001\u0000"+
		"\u0000\u000012\u0001\u0000\u0000\u000024\u0001\u0000\u0000\u000031\u0001"+
		"\u0000\u0000\u000045\u0005\u0014\u0000\u00005\u0007\u0001\u0000\u0000"+
		"\u000067\u0005\u001b\u0000\u000078\u0005\u001c\u0000\u000089\u0003\f\u0006"+
		"\u00009\t\u0001\u0000\u0000\u0000:;\u0003\u0010\b\u0000;\u000b\u0001\u0000"+
		"\u0000\u0000<=\u0006\u0006\uffff\uffff\u0000=A\u0003\u000e\u0007\u0000"+
		">?\u0007\u0000\u0000\u0000?A\u0003\f\u0006\b@<\u0001\u0000\u0000\u0000"+
		"@>\u0001\u0000\u0000\u0000A`\u0001\u0000\u0000\u0000BC\n\u0007\u0000\u0000"+
		"CD\u0007\u0001\u0000\u0000D_\u0003\f\u0006\bEF\n\u0006\u0000\u0000FG\u0007"+
		"\u0002\u0000\u0000G_\u0003\f\u0006\u0007HI\n\u0005\u0000\u0000IJ\u0007"+
		"\u0003\u0000\u0000J_\u0003\f\u0006\u0006KL\n\u0004\u0000\u0000LM\u0005"+
		"\u0002\u0000\u0000M_\u0003\f\u0006\u0005NO\n\u0003\u0000\u0000OP\u0005"+
		"\u0003\u0000\u0000P_\u0003\f\u0006\u0004QR\n\u0002\u0000\u0000RS\u0005"+
		"\u0018\u0000\u0000ST\u0003\f\u0006\u0000TU\u0005\u0019\u0000\u0000UV\u0003"+
		"\f\u0006\u0003V_\u0001\u0000\u0000\u0000WX\n\u0001\u0000\u0000XY\u0005"+
		"\u0015\u0000\u0000YZ\u0003\f\u0006\u0000Z[\u0005\u0019\u0000\u0000[\\"+
		"\u0003\f\u0006\u0000\\]\u0005\u0016\u0000\u0000]_\u0001\u0000\u0000\u0000"+
		"^B\u0001\u0000\u0000\u0000^E\u0001\u0000\u0000\u0000^H\u0001\u0000\u0000"+
		"\u0000^K\u0001\u0000\u0000\u0000^N\u0001\u0000\u0000\u0000^Q\u0001\u0000"+
		"\u0000\u0000^W\u0001\u0000\u0000\u0000_b\u0001\u0000\u0000\u0000`^\u0001"+
		"\u0000\u0000\u0000`a\u0001\u0000\u0000\u0000a\r\u0001\u0000\u0000\u0000"+
		"b`\u0001\u0000\u0000\u0000cl\u0003\u0016\u000b\u0000dl\u0003\u0014\n\u0000"+
		"el\u0003\u0012\t\u0000fl\u0003\u0010\b\u0000gh\u0005\u0011\u0000\u0000"+
		"hi\u0003\f\u0006\u0000ij\u0005\u0012\u0000\u0000jl\u0001\u0000\u0000\u0000"+
		"kc\u0001\u0000\u0000\u0000kd\u0001\u0000\u0000\u0000ke\u0001\u0000\u0000"+
		"\u0000kf\u0001\u0000\u0000\u0000kg\u0001\u0000\u0000\u0000l\u000f\u0001"+
		"\u0000\u0000\u0000mn\u0005\u001b\u0000\u0000nw\u0005\u0011\u0000\u0000"+
		"ot\u0003\f\u0006\u0000pq\u0005\u0017\u0000\u0000qs\u0003\f\u0006\u0000"+
		"rp\u0001\u0000\u0000\u0000sv\u0001\u0000\u0000\u0000tr\u0001\u0000\u0000"+
		"\u0000tu\u0001\u0000\u0000\u0000ux\u0001\u0000\u0000\u0000vt\u0001\u0000"+
		"\u0000\u0000wo\u0001\u0000\u0000\u0000wx\u0001\u0000\u0000\u0000xy\u0001"+
		"\u0000\u0000\u0000yz\u0005\u0012\u0000\u0000z\u0011\u0001\u0000\u0000"+
		"\u0000{|\u0005\u001b\u0000\u0000|\u0013\u0001\u0000\u0000\u0000}~\u0005"+
		"\u001a\u0000\u0000~\u0015\u0001\u0000\u0000\u0000\u007f\u0082\u0003\u0018"+
		"\f\u0000\u0080\u0082\u0003\u001a\r\u0000\u0081\u007f\u0001\u0000\u0000"+
		"\u0000\u0081\u0080\u0001\u0000\u0000\u0000\u0082\u0017\u0001\u0000\u0000"+
		"\u0000\u0083\u0084\u0005\u001d\u0000\u0000\u0084\u0019\u0001\u0000\u0000"+
		"\u0000\u0085\u0086\u0005\u001e\u0000\u0000\u0086\u001b\u0001\u0000\u0000"+
		"\u0000\n\u001f\'1@^`ktw\u0081";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}