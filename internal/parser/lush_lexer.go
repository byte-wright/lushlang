// Code generated from internal/parser/Lush.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LushLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LushLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lushlexerLexerInit() {
	staticData := &LushLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'&&'", "'||'", "'!'", "'-'", "'+'", "'*'", "'/'", "'%'", "'<'",
		"'<='", "'>'", "'>='", "'=='", "'!='", "'if'", "'('", "')'", "'{'",
		"'}'", "'['", "']'", "','", "'?'", "':'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "LAND", "LOR", "NOT", "MINUS", "PLUS", "MUL", "DIV",
		"MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ", "IF", "LPAREN", "RPAREN",
		"LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION", "COLON", "ENVVAR",
		"ID", "ASSIGN", "STRING", "NUMBER",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "LAND", "LOR", "NOT", "MINUS", "PLUS", "MUL", "DIV", "MOD",
		"LT", "LTE", "GT", "GTE", "EQ", "NEQ", "IF", "LPAREN", "RPAREN", "LCUR",
		"RCUR", "LSQ", "RSQ", "COMMA", "QUESTION", "COLON", "ENVVAR", "ID",
		"ASSIGN", "STRING", "NUMBER", "ESCAPED_VALUE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 158, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 4, 0, 65, 8, 0, 11, 0, 12, 0, 66, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 4, 25, 128,
		8, 25, 11, 25, 12, 25, 129, 1, 26, 1, 26, 5, 26, 134, 8, 26, 10, 26, 12,
		26, 137, 9, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 5, 28, 144, 8, 28, 10,
		28, 12, 28, 147, 9, 28, 1, 28, 1, 28, 1, 29, 4, 29, 152, 8, 29, 11, 29,
		12, 29, 153, 1, 30, 1, 30, 1, 30, 0, 0, 31, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 0, 1, 0, 6,
		3, 0, 9, 10, 13, 13, 32, 32, 4, 0, 48, 48, 65, 90, 95, 95, 97, 122, 2,
		0, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 1, 0, 48, 57, 5, 0, 34, 34, 39,
		39, 92, 92, 110, 110, 116, 116, 162, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43,
		1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0,
		51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0,
		0, 59, 1, 0, 0, 0, 1, 64, 1, 0, 0, 0, 3, 70, 1, 0, 0, 0, 5, 73, 1, 0, 0,
		0, 7, 76, 1, 0, 0, 0, 9, 78, 1, 0, 0, 0, 11, 80, 1, 0, 0, 0, 13, 82, 1,
		0, 0, 0, 15, 84, 1, 0, 0, 0, 17, 86, 1, 0, 0, 0, 19, 88, 1, 0, 0, 0, 21,
		90, 1, 0, 0, 0, 23, 93, 1, 0, 0, 0, 25, 95, 1, 0, 0, 0, 27, 98, 1, 0, 0,
		0, 29, 101, 1, 0, 0, 0, 31, 104, 1, 0, 0, 0, 33, 107, 1, 0, 0, 0, 35, 109,
		1, 0, 0, 0, 37, 111, 1, 0, 0, 0, 39, 113, 1, 0, 0, 0, 41, 115, 1, 0, 0,
		0, 43, 117, 1, 0, 0, 0, 45, 119, 1, 0, 0, 0, 47, 121, 1, 0, 0, 0, 49, 123,
		1, 0, 0, 0, 51, 125, 1, 0, 0, 0, 53, 131, 1, 0, 0, 0, 55, 138, 1, 0, 0,
		0, 57, 140, 1, 0, 0, 0, 59, 151, 1, 0, 0, 0, 61, 155, 1, 0, 0, 0, 63, 65,
		7, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0,
		66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 6, 0, 0, 0, 69, 2, 1, 0,
		0, 0, 70, 71, 5, 38, 0, 0, 71, 72, 5, 38, 0, 0, 72, 4, 1, 0, 0, 0, 73,
		74, 5, 124, 0, 0, 74, 75, 5, 124, 0, 0, 75, 6, 1, 0, 0, 0, 76, 77, 5, 33,
		0, 0, 77, 8, 1, 0, 0, 0, 78, 79, 5, 45, 0, 0, 79, 10, 1, 0, 0, 0, 80, 81,
		5, 43, 0, 0, 81, 12, 1, 0, 0, 0, 82, 83, 5, 42, 0, 0, 83, 14, 1, 0, 0,
		0, 84, 85, 5, 47, 0, 0, 85, 16, 1, 0, 0, 0, 86, 87, 5, 37, 0, 0, 87, 18,
		1, 0, 0, 0, 88, 89, 5, 60, 0, 0, 89, 20, 1, 0, 0, 0, 90, 91, 5, 60, 0,
		0, 91, 92, 5, 61, 0, 0, 92, 22, 1, 0, 0, 0, 93, 94, 5, 62, 0, 0, 94, 24,
		1, 0, 0, 0, 95, 96, 5, 62, 0, 0, 96, 97, 5, 61, 0, 0, 97, 26, 1, 0, 0,
		0, 98, 99, 5, 61, 0, 0, 99, 100, 5, 61, 0, 0, 100, 28, 1, 0, 0, 0, 101,
		102, 5, 33, 0, 0, 102, 103, 5, 61, 0, 0, 103, 30, 1, 0, 0, 0, 104, 105,
		5, 105, 0, 0, 105, 106, 5, 102, 0, 0, 106, 32, 1, 0, 0, 0, 107, 108, 5,
		40, 0, 0, 108, 34, 1, 0, 0, 0, 109, 110, 5, 41, 0, 0, 110, 36, 1, 0, 0,
		0, 111, 112, 5, 123, 0, 0, 112, 38, 1, 0, 0, 0, 113, 114, 5, 125, 0, 0,
		114, 40, 1, 0, 0, 0, 115, 116, 5, 91, 0, 0, 116, 42, 1, 0, 0, 0, 117, 118,
		5, 93, 0, 0, 118, 44, 1, 0, 0, 0, 119, 120, 5, 44, 0, 0, 120, 46, 1, 0,
		0, 0, 121, 122, 5, 63, 0, 0, 122, 48, 1, 0, 0, 0, 123, 124, 5, 58, 0, 0,
		124, 50, 1, 0, 0, 0, 125, 127, 5, 36, 0, 0, 126, 128, 7, 1, 0, 0, 127,
		126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130,
		1, 0, 0, 0, 130, 52, 1, 0, 0, 0, 131, 135, 7, 2, 0, 0, 132, 134, 7, 1,
		0, 0, 133, 132, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0,
		135, 136, 1, 0, 0, 0, 136, 54, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139,
		5, 61, 0, 0, 139, 56, 1, 0, 0, 0, 140, 145, 5, 34, 0, 0, 141, 144, 8, 3,
		0, 0, 142, 144, 3, 61, 30, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0,
		0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		148, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 34, 0, 0, 149, 58,
		1, 0, 0, 0, 150, 152, 7, 4, 0, 0, 151, 150, 1, 0, 0, 0, 152, 153, 1, 0,
		0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 60, 1, 0, 0, 0,
		155, 156, 5, 92, 0, 0, 156, 157, 7, 5, 0, 0, 157, 62, 1, 0, 0, 0, 7, 0,
		66, 129, 135, 143, 145, 153, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LushLexerInit initializes any static state used to implement LushLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLushLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LushLexerInit() {
	staticData := &LushLexerLexerStaticData
	staticData.once.Do(lushlexerLexerInit)
}

// NewLushLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLushLexer(input antlr.CharStream) *LushLexer {
	LushLexerInit()
	l := new(LushLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LushLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lush.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LushLexer tokens.
const (
	LushLexerWHITESPACE = 1
	LushLexerLAND       = 2
	LushLexerLOR        = 3
	LushLexerNOT        = 4
	LushLexerMINUS      = 5
	LushLexerPLUS       = 6
	LushLexerMUL        = 7
	LushLexerDIV        = 8
	LushLexerMOD        = 9
	LushLexerLT         = 10
	LushLexerLTE        = 11
	LushLexerGT         = 12
	LushLexerGTE        = 13
	LushLexerEQ         = 14
	LushLexerNEQ        = 15
	LushLexerIF         = 16
	LushLexerLPAREN     = 17
	LushLexerRPAREN     = 18
	LushLexerLCUR       = 19
	LushLexerRCUR       = 20
	LushLexerLSQ        = 21
	LushLexerRSQ        = 22
	LushLexerCOMMA      = 23
	LushLexerQUESTION   = 24
	LushLexerCOLON      = 25
	LushLexerENVVAR     = 26
	LushLexerID         = 27
	LushLexerASSIGN     = 28
	LushLexerSTRING     = 29
	LushLexerNUMBER     = 30
)
