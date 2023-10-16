// Generated from /home/creichlin/workspaces/bytewright/lush/internal/parser/Lush.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class LushLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WHITESPACE=1, LAND=2, LOR=3, NOT=4, MINUS=5, PLUS=6, MUL=7, DIV=8, MOD=9, 
		LT=10, LTE=11, GT=12, GTE=13, EQ=14, NEQ=15, IF=16, LPAREN=17, RPAREN=18, 
		LCUR=19, RCUR=20, LSQ=21, RSQ=22, COMMA=23, QUESTION=24, COLON=25, ENVVAR=26, 
		ID=27, ASSIGN=28, STRING=29, NUMBER=30;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"WHITESPACE", "LAND", "LOR", "NOT", "MINUS", "PLUS", "MUL", "DIV", "MOD", 
			"LT", "LTE", "GT", "GTE", "EQ", "NEQ", "IF", "LPAREN", "RPAREN", "LCUR", 
			"RCUR", "LSQ", "RSQ", "COMMA", "QUESTION", "COLON", "ENVVAR", "ID", "ASSIGN", 
			"STRING", "NUMBER", "ESCAPED_VALUE"
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


	public LushLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lush.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u001e\u009e\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017"+
		"\u0002\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a"+
		"\u0002\u001b\u0007\u001b\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d"+
		"\u0002\u001e\u0007\u001e\u0001\u0000\u0004\u0000A\b\u0000\u000b\u0000"+
		"\f\u0000B\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001"+
		"\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001"+
		"\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018\u0001"+
		"\u0019\u0001\u0019\u0004\u0019\u0080\b\u0019\u000b\u0019\f\u0019\u0081"+
		"\u0001\u001a\u0001\u001a\u0005\u001a\u0086\b\u001a\n\u001a\f\u001a\u0089"+
		"\t\u001a\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0005"+
		"\u001c\u0090\b\u001c\n\u001c\f\u001c\u0093\t\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001d\u0004\u001d\u0098\b\u001d\u000b\u001d\f\u001d\u0099\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0000\u0000\u001f\u0001\u0001\u0003\u0002"+
		"\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013"+
		"\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011"+
		"#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017/\u00181\u00193\u001a5\u001b"+
		"7\u001c9\u001d;\u001e=\u0000\u0001\u0000\u0006\u0003\u0000\t\n\r\r  \u0004"+
		"\u000000AZ__az\u0002\u0000__az\u0002\u0000\"\"\\\\\u0001\u000009\u0005"+
		"\u0000\"\"\'\'\\\\nntt\u00a2\u0000\u0001\u0001\u0000\u0000\u0000\u0000"+
		"\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000"+
		"\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b"+
		"\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001"+
		"\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001"+
		"\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001"+
		"\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001"+
		"\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001"+
		"\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000"+
		"\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'\u0001\u0000\u0000\u0000"+
		"\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000\u0000\u0000\u0000-"+
		"\u0001\u0000\u0000\u0000\u0000/\u0001\u0000\u0000\u0000\u00001\u0001\u0000"+
		"\u0000\u0000\u00003\u0001\u0000\u0000\u0000\u00005\u0001\u0000\u0000\u0000"+
		"\u00007\u0001\u0000\u0000\u0000\u00009\u0001\u0000\u0000\u0000\u0000;"+
		"\u0001\u0000\u0000\u0000\u0001@\u0001\u0000\u0000\u0000\u0003F\u0001\u0000"+
		"\u0000\u0000\u0005I\u0001\u0000\u0000\u0000\u0007L\u0001\u0000\u0000\u0000"+
		"\tN\u0001\u0000\u0000\u0000\u000bP\u0001\u0000\u0000\u0000\rR\u0001\u0000"+
		"\u0000\u0000\u000fT\u0001\u0000\u0000\u0000\u0011V\u0001\u0000\u0000\u0000"+
		"\u0013X\u0001\u0000\u0000\u0000\u0015Z\u0001\u0000\u0000\u0000\u0017]"+
		"\u0001\u0000\u0000\u0000\u0019_\u0001\u0000\u0000\u0000\u001bb\u0001\u0000"+
		"\u0000\u0000\u001de\u0001\u0000\u0000\u0000\u001fh\u0001\u0000\u0000\u0000"+
		"!k\u0001\u0000\u0000\u0000#m\u0001\u0000\u0000\u0000%o\u0001\u0000\u0000"+
		"\u0000\'q\u0001\u0000\u0000\u0000)s\u0001\u0000\u0000\u0000+u\u0001\u0000"+
		"\u0000\u0000-w\u0001\u0000\u0000\u0000/y\u0001\u0000\u0000\u00001{\u0001"+
		"\u0000\u0000\u00003}\u0001\u0000\u0000\u00005\u0083\u0001\u0000\u0000"+
		"\u00007\u008a\u0001\u0000\u0000\u00009\u008c\u0001\u0000\u0000\u0000;"+
		"\u0097\u0001\u0000\u0000\u0000=\u009b\u0001\u0000\u0000\u0000?A\u0007"+
		"\u0000\u0000\u0000@?\u0001\u0000\u0000\u0000AB\u0001\u0000\u0000\u0000"+
		"B@\u0001\u0000\u0000\u0000BC\u0001\u0000\u0000\u0000CD\u0001\u0000\u0000"+
		"\u0000DE\u0006\u0000\u0000\u0000E\u0002\u0001\u0000\u0000\u0000FG\u0005"+
		"&\u0000\u0000GH\u0005&\u0000\u0000H\u0004\u0001\u0000\u0000\u0000IJ\u0005"+
		"|\u0000\u0000JK\u0005|\u0000\u0000K\u0006\u0001\u0000\u0000\u0000LM\u0005"+
		"!\u0000\u0000M\b\u0001\u0000\u0000\u0000NO\u0005-\u0000\u0000O\n\u0001"+
		"\u0000\u0000\u0000PQ\u0005+\u0000\u0000Q\f\u0001\u0000\u0000\u0000RS\u0005"+
		"*\u0000\u0000S\u000e\u0001\u0000\u0000\u0000TU\u0005/\u0000\u0000U\u0010"+
		"\u0001\u0000\u0000\u0000VW\u0005%\u0000\u0000W\u0012\u0001\u0000\u0000"+
		"\u0000XY\u0005<\u0000\u0000Y\u0014\u0001\u0000\u0000\u0000Z[\u0005<\u0000"+
		"\u0000[\\\u0005=\u0000\u0000\\\u0016\u0001\u0000\u0000\u0000]^\u0005>"+
		"\u0000\u0000^\u0018\u0001\u0000\u0000\u0000_`\u0005>\u0000\u0000`a\u0005"+
		"=\u0000\u0000a\u001a\u0001\u0000\u0000\u0000bc\u0005=\u0000\u0000cd\u0005"+
		"=\u0000\u0000d\u001c\u0001\u0000\u0000\u0000ef\u0005!\u0000\u0000fg\u0005"+
		"=\u0000\u0000g\u001e\u0001\u0000\u0000\u0000hi\u0005i\u0000\u0000ij\u0005"+
		"f\u0000\u0000j \u0001\u0000\u0000\u0000kl\u0005(\u0000\u0000l\"\u0001"+
		"\u0000\u0000\u0000mn\u0005)\u0000\u0000n$\u0001\u0000\u0000\u0000op\u0005"+
		"{\u0000\u0000p&\u0001\u0000\u0000\u0000qr\u0005}\u0000\u0000r(\u0001\u0000"+
		"\u0000\u0000st\u0005[\u0000\u0000t*\u0001\u0000\u0000\u0000uv\u0005]\u0000"+
		"\u0000v,\u0001\u0000\u0000\u0000wx\u0005,\u0000\u0000x.\u0001\u0000\u0000"+
		"\u0000yz\u0005?\u0000\u0000z0\u0001\u0000\u0000\u0000{|\u0005:\u0000\u0000"+
		"|2\u0001\u0000\u0000\u0000}\u007f\u0005$\u0000\u0000~\u0080\u0007\u0001"+
		"\u0000\u0000\u007f~\u0001\u0000\u0000\u0000\u0080\u0081\u0001\u0000\u0000"+
		"\u0000\u0081\u007f\u0001\u0000\u0000\u0000\u0081\u0082\u0001\u0000\u0000"+
		"\u0000\u00824\u0001\u0000\u0000\u0000\u0083\u0087\u0007\u0002\u0000\u0000"+
		"\u0084\u0086\u0007\u0001\u0000\u0000\u0085\u0084\u0001\u0000\u0000\u0000"+
		"\u0086\u0089\u0001\u0000\u0000\u0000\u0087\u0085\u0001\u0000\u0000\u0000"+
		"\u0087\u0088\u0001\u0000\u0000\u0000\u00886\u0001\u0000\u0000\u0000\u0089"+
		"\u0087\u0001\u0000\u0000\u0000\u008a\u008b\u0005=\u0000\u0000\u008b8\u0001"+
		"\u0000\u0000\u0000\u008c\u0091\u0005\"\u0000\u0000\u008d\u0090\b\u0003"+
		"\u0000\u0000\u008e\u0090\u0003=\u001e\u0000\u008f\u008d\u0001\u0000\u0000"+
		"\u0000\u008f\u008e\u0001\u0000\u0000\u0000\u0090\u0093\u0001\u0000\u0000"+
		"\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0091\u0092\u0001\u0000\u0000"+
		"\u0000\u0092\u0094\u0001\u0000\u0000\u0000\u0093\u0091\u0001\u0000\u0000"+
		"\u0000\u0094\u0095\u0005\"\u0000\u0000\u0095:\u0001\u0000\u0000\u0000"+
		"\u0096\u0098\u0007\u0004\u0000\u0000\u0097\u0096\u0001\u0000\u0000\u0000"+
		"\u0098\u0099\u0001\u0000\u0000\u0000\u0099\u0097\u0001\u0000\u0000\u0000"+
		"\u0099\u009a\u0001\u0000\u0000\u0000\u009a<\u0001\u0000\u0000\u0000\u009b"+
		"\u009c\u0005\\\u0000\u0000\u009c\u009d\u0007\u0005\u0000\u0000\u009d>"+
		"\u0001\u0000\u0000\u0000\u0007\u0000B\u0081\u0087\u008f\u0091\u0099\u0001"+
		"\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}