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
		"", "", "'true'", "'false'", "'if'", "'&&'", "'||'", "'!'", "'-'", "'+'",
		"'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='", "'('",
		"')'", "'{'", "'}'", "'['", "']'", "','", "'?'", "':'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "TRUE", "FALSE", "IF", "LAND", "LOR", "NOT", "MINUS",
		"PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ",
		"LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "TRUE", "FALSE", "IF", "LAND", "LOR", "NOT", "MINUS",
		"PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ",
		"LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER", "ESCAPED_VALUE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 173, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 4, 0, 69, 8, 0, 11, 0, 12, 0, 70, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 4, 27, 143, 8, 27, 11, 27, 12, 27, 144, 1, 28, 1, 28,
		5, 28, 149, 8, 28, 10, 28, 12, 28, 152, 9, 28, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 30, 5, 30, 159, 8, 30, 10, 30, 12, 30, 162, 9, 30, 1, 30, 1, 30,
		1, 31, 4, 31, 167, 8, 31, 11, 31, 12, 31, 168, 1, 32, 1, 32, 1, 32, 0,
		0, 33, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 0, 1, 0, 6, 3, 0, 9, 10, 13, 13, 32,
		32, 4, 0, 48, 48, 65, 90, 95, 95, 97, 122, 2, 0, 95, 95, 97, 122, 2, 0,
		34, 34, 92, 92, 1, 0, 48, 57, 4, 0, 34, 34, 92, 92, 110, 110, 116, 116,
		177, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 68, 1, 0, 0, 0, 3, 74, 1, 0, 0, 0, 5, 79,
		1, 0, 0, 0, 7, 85, 1, 0, 0, 0, 9, 88, 1, 0, 0, 0, 11, 91, 1, 0, 0, 0, 13,
		94, 1, 0, 0, 0, 15, 96, 1, 0, 0, 0, 17, 98, 1, 0, 0, 0, 19, 100, 1, 0,
		0, 0, 21, 102, 1, 0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 106, 1, 0, 0, 0, 27,
		108, 1, 0, 0, 0, 29, 111, 1, 0, 0, 0, 31, 113, 1, 0, 0, 0, 33, 116, 1,
		0, 0, 0, 35, 119, 1, 0, 0, 0, 37, 122, 1, 0, 0, 0, 39, 124, 1, 0, 0, 0,
		41, 126, 1, 0, 0, 0, 43, 128, 1, 0, 0, 0, 45, 130, 1, 0, 0, 0, 47, 132,
		1, 0, 0, 0, 49, 134, 1, 0, 0, 0, 51, 136, 1, 0, 0, 0, 53, 138, 1, 0, 0,
		0, 55, 140, 1, 0, 0, 0, 57, 146, 1, 0, 0, 0, 59, 153, 1, 0, 0, 0, 61, 155,
		1, 0, 0, 0, 63, 166, 1, 0, 0, 0, 65, 170, 1, 0, 0, 0, 67, 69, 7, 0, 0,
		0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71,
		1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 6, 0, 0, 0, 73, 2, 1, 0, 0, 0,
		74, 75, 5, 116, 0, 0, 75, 76, 5, 114, 0, 0, 76, 77, 5, 117, 0, 0, 77, 78,
		5, 101, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 102, 0, 0, 80, 81, 5, 97, 0,
		0, 81, 82, 5, 108, 0, 0, 82, 83, 5, 115, 0, 0, 83, 84, 5, 101, 0, 0, 84,
		6, 1, 0, 0, 0, 85, 86, 5, 105, 0, 0, 86, 87, 5, 102, 0, 0, 87, 8, 1, 0,
		0, 0, 88, 89, 5, 38, 0, 0, 89, 90, 5, 38, 0, 0, 90, 10, 1, 0, 0, 0, 91,
		92, 5, 124, 0, 0, 92, 93, 5, 124, 0, 0, 93, 12, 1, 0, 0, 0, 94, 95, 5,
		33, 0, 0, 95, 14, 1, 0, 0, 0, 96, 97, 5, 45, 0, 0, 97, 16, 1, 0, 0, 0,
		98, 99, 5, 43, 0, 0, 99, 18, 1, 0, 0, 0, 100, 101, 5, 42, 0, 0, 101, 20,
		1, 0, 0, 0, 102, 103, 5, 47, 0, 0, 103, 22, 1, 0, 0, 0, 104, 105, 5, 37,
		0, 0, 105, 24, 1, 0, 0, 0, 106, 107, 5, 60, 0, 0, 107, 26, 1, 0, 0, 0,
		108, 109, 5, 60, 0, 0, 109, 110, 5, 61, 0, 0, 110, 28, 1, 0, 0, 0, 111,
		112, 5, 62, 0, 0, 112, 30, 1, 0, 0, 0, 113, 114, 5, 62, 0, 0, 114, 115,
		5, 61, 0, 0, 115, 32, 1, 0, 0, 0, 116, 117, 5, 61, 0, 0, 117, 118, 5, 61,
		0, 0, 118, 34, 1, 0, 0, 0, 119, 120, 5, 33, 0, 0, 120, 121, 5, 61, 0, 0,
		121, 36, 1, 0, 0, 0, 122, 123, 5, 40, 0, 0, 123, 38, 1, 0, 0, 0, 124, 125,
		5, 41, 0, 0, 125, 40, 1, 0, 0, 0, 126, 127, 5, 123, 0, 0, 127, 42, 1, 0,
		0, 0, 128, 129, 5, 125, 0, 0, 129, 44, 1, 0, 0, 0, 130, 131, 5, 91, 0,
		0, 131, 46, 1, 0, 0, 0, 132, 133, 5, 93, 0, 0, 133, 48, 1, 0, 0, 0, 134,
		135, 5, 44, 0, 0, 135, 50, 1, 0, 0, 0, 136, 137, 5, 63, 0, 0, 137, 52,
		1, 0, 0, 0, 138, 139, 5, 58, 0, 0, 139, 54, 1, 0, 0, 0, 140, 142, 5, 36,
		0, 0, 141, 143, 7, 1, 0, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0,
		144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 56, 1, 0, 0, 0, 146, 150,
		7, 2, 0, 0, 147, 149, 7, 1, 0, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0,
		0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 58, 1, 0, 0, 0,
		152, 150, 1, 0, 0, 0, 153, 154, 5, 61, 0, 0, 154, 60, 1, 0, 0, 0, 155,
		160, 5, 34, 0, 0, 156, 159, 8, 3, 0, 0, 157, 159, 3, 65, 32, 0, 158, 156,
		1, 0, 0, 0, 158, 157, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0,
		0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0,
		163, 164, 5, 34, 0, 0, 164, 62, 1, 0, 0, 0, 165, 167, 7, 4, 0, 0, 166,
		165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169,
		1, 0, 0, 0, 169, 64, 1, 0, 0, 0, 170, 171, 5, 92, 0, 0, 171, 172, 7, 5,
		0, 0, 172, 66, 1, 0, 0, 0, 7, 0, 70, 144, 150, 158, 160, 168, 1, 6, 0,
		0,
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
	LushLexerTRUE       = 2
	LushLexerFALSE      = 3
	LushLexerIF         = 4
	LushLexerLAND       = 5
	LushLexerLOR        = 6
	LushLexerNOT        = 7
	LushLexerMINUS      = 8
	LushLexerPLUS       = 9
	LushLexerMUL        = 10
	LushLexerDIV        = 11
	LushLexerMOD        = 12
	LushLexerLT         = 13
	LushLexerLTE        = 14
	LushLexerGT         = 15
	LushLexerGTE        = 16
	LushLexerEQ         = 17
	LushLexerNEQ        = 18
	LushLexerLPAREN     = 19
	LushLexerRPAREN     = 20
	LushLexerLCUR       = 21
	LushLexerRCUR       = 22
	LushLexerLSQ        = 23
	LushLexerRSQ        = 24
	LushLexerCOMMA      = 25
	LushLexerQUESTION   = 26
	LushLexerCOLON      = 27
	LushLexerENVVAR     = 28
	LushLexerID         = 29
	LushLexerASSIGN     = 30
	LushLexerSTRING     = 31
	LushLexerNUMBER     = 32
)
