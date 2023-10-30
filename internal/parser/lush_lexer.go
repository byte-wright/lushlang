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
		"", "", "'true'", "'false'", "'if'", "'for'", "'&&'", "'||'", "'!'",
		"'-'", "'+'", "'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='", "'=='",
		"'!='", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'?'", "':'",
		"';'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "TRUE", "FALSE", "IF", "FOR", "LAND", "LOR", "NOT",
		"MINUS", "PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ",
		"NEQ", "LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "SEMICOLON", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "TRUE", "FALSE", "IF", "FOR", "LAND", "LOR", "NOT", "MINUS",
		"PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ",
		"LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "SEMICOLON", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
		"ESCAPED_VALUE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 183, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 4, 0, 73, 8,
		0, 11, 0, 12, 0, 74, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 4, 29, 153, 8, 29, 11, 29, 12, 29, 154, 1, 30, 1, 30, 5, 30,
		159, 8, 30, 10, 30, 12, 30, 162, 9, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 5, 32, 169, 8, 32, 10, 32, 12, 32, 172, 9, 32, 1, 32, 1, 32, 1, 33,
		4, 33, 177, 8, 33, 11, 33, 12, 33, 178, 1, 34, 1, 34, 1, 34, 0, 0, 35,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 0, 1, 0, 6, 3, 0, 9, 10, 13,
		13, 32, 32, 4, 0, 48, 48, 65, 90, 95, 95, 97, 122, 2, 0, 95, 95, 97, 122,
		2, 0, 34, 34, 92, 92, 1, 0, 48, 57, 4, 0, 34, 34, 92, 92, 110, 110, 116,
		116, 187, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 1,
		72, 1, 0, 0, 0, 3, 78, 1, 0, 0, 0, 5, 83, 1, 0, 0, 0, 7, 89, 1, 0, 0, 0,
		9, 92, 1, 0, 0, 0, 11, 96, 1, 0, 0, 0, 13, 99, 1, 0, 0, 0, 15, 102, 1,
		0, 0, 0, 17, 104, 1, 0, 0, 0, 19, 106, 1, 0, 0, 0, 21, 108, 1, 0, 0, 0,
		23, 110, 1, 0, 0, 0, 25, 112, 1, 0, 0, 0, 27, 114, 1, 0, 0, 0, 29, 116,
		1, 0, 0, 0, 31, 119, 1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 124, 1, 0, 0,
		0, 37, 127, 1, 0, 0, 0, 39, 130, 1, 0, 0, 0, 41, 132, 1, 0, 0, 0, 43, 134,
		1, 0, 0, 0, 45, 136, 1, 0, 0, 0, 47, 138, 1, 0, 0, 0, 49, 140, 1, 0, 0,
		0, 51, 142, 1, 0, 0, 0, 53, 144, 1, 0, 0, 0, 55, 146, 1, 0, 0, 0, 57, 148,
		1, 0, 0, 0, 59, 150, 1, 0, 0, 0, 61, 156, 1, 0, 0, 0, 63, 163, 1, 0, 0,
		0, 65, 165, 1, 0, 0, 0, 67, 176, 1, 0, 0, 0, 69, 180, 1, 0, 0, 0, 71, 73,
		7, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 6, 0, 0, 0, 77, 2, 1, 0,
		0, 0, 78, 79, 5, 116, 0, 0, 79, 80, 5, 114, 0, 0, 80, 81, 5, 117, 0, 0,
		81, 82, 5, 101, 0, 0, 82, 4, 1, 0, 0, 0, 83, 84, 5, 102, 0, 0, 84, 85,
		5, 97, 0, 0, 85, 86, 5, 108, 0, 0, 86, 87, 5, 115, 0, 0, 87, 88, 5, 101,
		0, 0, 88, 6, 1, 0, 0, 0, 89, 90, 5, 105, 0, 0, 90, 91, 5, 102, 0, 0, 91,
		8, 1, 0, 0, 0, 92, 93, 5, 102, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 114,
		0, 0, 95, 10, 1, 0, 0, 0, 96, 97, 5, 38, 0, 0, 97, 98, 5, 38, 0, 0, 98,
		12, 1, 0, 0, 0, 99, 100, 5, 124, 0, 0, 100, 101, 5, 124, 0, 0, 101, 14,
		1, 0, 0, 0, 102, 103, 5, 33, 0, 0, 103, 16, 1, 0, 0, 0, 104, 105, 5, 45,
		0, 0, 105, 18, 1, 0, 0, 0, 106, 107, 5, 43, 0, 0, 107, 20, 1, 0, 0, 0,
		108, 109, 5, 42, 0, 0, 109, 22, 1, 0, 0, 0, 110, 111, 5, 47, 0, 0, 111,
		24, 1, 0, 0, 0, 112, 113, 5, 37, 0, 0, 113, 26, 1, 0, 0, 0, 114, 115, 5,
		60, 0, 0, 115, 28, 1, 0, 0, 0, 116, 117, 5, 60, 0, 0, 117, 118, 5, 61,
		0, 0, 118, 30, 1, 0, 0, 0, 119, 120, 5, 62, 0, 0, 120, 32, 1, 0, 0, 0,
		121, 122, 5, 62, 0, 0, 122, 123, 5, 61, 0, 0, 123, 34, 1, 0, 0, 0, 124,
		125, 5, 61, 0, 0, 125, 126, 5, 61, 0, 0, 126, 36, 1, 0, 0, 0, 127, 128,
		5, 33, 0, 0, 128, 129, 5, 61, 0, 0, 129, 38, 1, 0, 0, 0, 130, 131, 5, 40,
		0, 0, 131, 40, 1, 0, 0, 0, 132, 133, 5, 41, 0, 0, 133, 42, 1, 0, 0, 0,
		134, 135, 5, 123, 0, 0, 135, 44, 1, 0, 0, 0, 136, 137, 5, 125, 0, 0, 137,
		46, 1, 0, 0, 0, 138, 139, 5, 91, 0, 0, 139, 48, 1, 0, 0, 0, 140, 141, 5,
		93, 0, 0, 141, 50, 1, 0, 0, 0, 142, 143, 5, 44, 0, 0, 143, 52, 1, 0, 0,
		0, 144, 145, 5, 63, 0, 0, 145, 54, 1, 0, 0, 0, 146, 147, 5, 58, 0, 0, 147,
		56, 1, 0, 0, 0, 148, 149, 5, 59, 0, 0, 149, 58, 1, 0, 0, 0, 150, 152, 5,
		36, 0, 0, 151, 153, 7, 1, 0, 0, 152, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0,
		0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 60, 1, 0, 0, 0, 156,
		160, 7, 2, 0, 0, 157, 159, 7, 1, 0, 0, 158, 157, 1, 0, 0, 0, 159, 162,
		1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 62, 1, 0,
		0, 0, 162, 160, 1, 0, 0, 0, 163, 164, 5, 61, 0, 0, 164, 64, 1, 0, 0, 0,
		165, 170, 5, 34, 0, 0, 166, 169, 8, 3, 0, 0, 167, 169, 3, 69, 34, 0, 168,
		166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168,
		1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0, 0, 0, 172, 170, 1, 0,
		0, 0, 173, 174, 5, 34, 0, 0, 174, 66, 1, 0, 0, 0, 175, 177, 7, 4, 0, 0,
		176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178,
		179, 1, 0, 0, 0, 179, 68, 1, 0, 0, 0, 180, 181, 5, 92, 0, 0, 181, 182,
		7, 5, 0, 0, 182, 70, 1, 0, 0, 0, 7, 0, 74, 154, 160, 168, 170, 178, 1,
		6, 0, 0,
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
	LushLexerFOR        = 5
	LushLexerLAND       = 6
	LushLexerLOR        = 7
	LushLexerNOT        = 8
	LushLexerMINUS      = 9
	LushLexerPLUS       = 10
	LushLexerMUL        = 11
	LushLexerDIV        = 12
	LushLexerMOD        = 13
	LushLexerLT         = 14
	LushLexerLTE        = 15
	LushLexerGT         = 16
	LushLexerGTE        = 17
	LushLexerEQ         = 18
	LushLexerNEQ        = 19
	LushLexerLPAREN     = 20
	LushLexerRPAREN     = 21
	LushLexerLCUR       = 22
	LushLexerRCUR       = 23
	LushLexerLSQ        = 24
	LushLexerRSQ        = 25
	LushLexerCOMMA      = 26
	LushLexerQUESTION   = 27
	LushLexerCOLON      = 28
	LushLexerSEMICOLON  = 29
	LushLexerENVVAR     = 30
	LushLexerID         = 31
	LushLexerASSIGN     = 32
	LushLexerSTRING     = 33
	LushLexerNUMBER     = 34
)
