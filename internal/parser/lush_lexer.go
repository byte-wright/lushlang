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
		"", "", "'true'", "'false'", "'if'", "'for'", "'string'", "'int'", "'bool'",
		"'func'", "'return'", "'&&'", "'||'", "'!'", "'-'", "'+'", "'*'", "'/'",
		"'%'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='", "'('", "')'", "'{'",
		"'}'", "'['", "']'", "','", "'?'", "':'", "';'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "TRUE", "FALSE", "IF", "FOR", "STRING_TYPE", "INT_TYPE",
		"BOOL_TYPE", "FUNC", "RETURN", "LAND", "LOR", "NOT", "MINUS", "PLUS",
		"MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ", "LPAREN",
		"RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION", "COLON",
		"SEMICOLON", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "TRUE", "FALSE", "IF", "FOR", "STRING_TYPE", "INT_TYPE",
		"BOOL_TYPE", "FUNC", "RETURN", "LAND", "LOR", "NOT", "MINUS", "PLUS",
		"MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ", "LPAREN",
		"RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION", "COLON",
		"SEMICOLON", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER", "ESCAPED_VALUE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 39, 221, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 4, 0, 83, 8, 0,
		11, 0, 12, 0, 84, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 4, 34, 191, 8, 34, 11, 34, 12,
		34, 192, 1, 35, 1, 35, 5, 35, 197, 8, 35, 10, 35, 12, 35, 200, 9, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 5, 37, 207, 8, 37, 10, 37, 12, 37, 210,
		9, 37, 1, 37, 1, 37, 1, 38, 4, 38, 215, 8, 38, 11, 38, 12, 38, 216, 1,
		39, 1, 39, 1, 39, 0, 0, 40, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69,
		35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 0, 1, 0, 6, 3, 0, 9, 10, 13, 13,
		32, 32, 4, 0, 48, 48, 65, 90, 95, 95, 97, 122, 2, 0, 95, 95, 97, 122, 2,
		0, 34, 34, 92, 92, 1, 0, 48, 57, 4, 0, 34, 34, 92, 92, 110, 110, 116, 116,
		225, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 1, 82, 1, 0, 0, 0, 3, 88, 1, 0, 0, 0, 5, 93, 1, 0, 0, 0,
		7, 99, 1, 0, 0, 0, 9, 102, 1, 0, 0, 0, 11, 106, 1, 0, 0, 0, 13, 113, 1,
		0, 0, 0, 15, 117, 1, 0, 0, 0, 17, 122, 1, 0, 0, 0, 19, 127, 1, 0, 0, 0,
		21, 134, 1, 0, 0, 0, 23, 137, 1, 0, 0, 0, 25, 140, 1, 0, 0, 0, 27, 142,
		1, 0, 0, 0, 29, 144, 1, 0, 0, 0, 31, 146, 1, 0, 0, 0, 33, 148, 1, 0, 0,
		0, 35, 150, 1, 0, 0, 0, 37, 152, 1, 0, 0, 0, 39, 154, 1, 0, 0, 0, 41, 157,
		1, 0, 0, 0, 43, 159, 1, 0, 0, 0, 45, 162, 1, 0, 0, 0, 47, 165, 1, 0, 0,
		0, 49, 168, 1, 0, 0, 0, 51, 170, 1, 0, 0, 0, 53, 172, 1, 0, 0, 0, 55, 174,
		1, 0, 0, 0, 57, 176, 1, 0, 0, 0, 59, 178, 1, 0, 0, 0, 61, 180, 1, 0, 0,
		0, 63, 182, 1, 0, 0, 0, 65, 184, 1, 0, 0, 0, 67, 186, 1, 0, 0, 0, 69, 188,
		1, 0, 0, 0, 71, 194, 1, 0, 0, 0, 73, 201, 1, 0, 0, 0, 75, 203, 1, 0, 0,
		0, 77, 214, 1, 0, 0, 0, 79, 218, 1, 0, 0, 0, 81, 83, 7, 0, 0, 0, 82, 81,
		1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0,
		85, 86, 1, 0, 0, 0, 86, 87, 6, 0, 0, 0, 87, 2, 1, 0, 0, 0, 88, 89, 5, 116,
		0, 0, 89, 90, 5, 114, 0, 0, 90, 91, 5, 117, 0, 0, 91, 92, 5, 101, 0, 0,
		92, 4, 1, 0, 0, 0, 93, 94, 5, 102, 0, 0, 94, 95, 5, 97, 0, 0, 95, 96, 5,
		108, 0, 0, 96, 97, 5, 115, 0, 0, 97, 98, 5, 101, 0, 0, 98, 6, 1, 0, 0,
		0, 99, 100, 5, 105, 0, 0, 100, 101, 5, 102, 0, 0, 101, 8, 1, 0, 0, 0, 102,
		103, 5, 102, 0, 0, 103, 104, 5, 111, 0, 0, 104, 105, 5, 114, 0, 0, 105,
		10, 1, 0, 0, 0, 106, 107, 5, 115, 0, 0, 107, 108, 5, 116, 0, 0, 108, 109,
		5, 114, 0, 0, 109, 110, 5, 105, 0, 0, 110, 111, 5, 110, 0, 0, 111, 112,
		5, 103, 0, 0, 112, 12, 1, 0, 0, 0, 113, 114, 5, 105, 0, 0, 114, 115, 5,
		110, 0, 0, 115, 116, 5, 116, 0, 0, 116, 14, 1, 0, 0, 0, 117, 118, 5, 98,
		0, 0, 118, 119, 5, 111, 0, 0, 119, 120, 5, 111, 0, 0, 120, 121, 5, 108,
		0, 0, 121, 16, 1, 0, 0, 0, 122, 123, 5, 102, 0, 0, 123, 124, 5, 117, 0,
		0, 124, 125, 5, 110, 0, 0, 125, 126, 5, 99, 0, 0, 126, 18, 1, 0, 0, 0,
		127, 128, 5, 114, 0, 0, 128, 129, 5, 101, 0, 0, 129, 130, 5, 116, 0, 0,
		130, 131, 5, 117, 0, 0, 131, 132, 5, 114, 0, 0, 132, 133, 5, 110, 0, 0,
		133, 20, 1, 0, 0, 0, 134, 135, 5, 38, 0, 0, 135, 136, 5, 38, 0, 0, 136,
		22, 1, 0, 0, 0, 137, 138, 5, 124, 0, 0, 138, 139, 5, 124, 0, 0, 139, 24,
		1, 0, 0, 0, 140, 141, 5, 33, 0, 0, 141, 26, 1, 0, 0, 0, 142, 143, 5, 45,
		0, 0, 143, 28, 1, 0, 0, 0, 144, 145, 5, 43, 0, 0, 145, 30, 1, 0, 0, 0,
		146, 147, 5, 42, 0, 0, 147, 32, 1, 0, 0, 0, 148, 149, 5, 47, 0, 0, 149,
		34, 1, 0, 0, 0, 150, 151, 5, 37, 0, 0, 151, 36, 1, 0, 0, 0, 152, 153, 5,
		60, 0, 0, 153, 38, 1, 0, 0, 0, 154, 155, 5, 60, 0, 0, 155, 156, 5, 61,
		0, 0, 156, 40, 1, 0, 0, 0, 157, 158, 5, 62, 0, 0, 158, 42, 1, 0, 0, 0,
		159, 160, 5, 62, 0, 0, 160, 161, 5, 61, 0, 0, 161, 44, 1, 0, 0, 0, 162,
		163, 5, 61, 0, 0, 163, 164, 5, 61, 0, 0, 164, 46, 1, 0, 0, 0, 165, 166,
		5, 33, 0, 0, 166, 167, 5, 61, 0, 0, 167, 48, 1, 0, 0, 0, 168, 169, 5, 40,
		0, 0, 169, 50, 1, 0, 0, 0, 170, 171, 5, 41, 0, 0, 171, 52, 1, 0, 0, 0,
		172, 173, 5, 123, 0, 0, 173, 54, 1, 0, 0, 0, 174, 175, 5, 125, 0, 0, 175,
		56, 1, 0, 0, 0, 176, 177, 5, 91, 0, 0, 177, 58, 1, 0, 0, 0, 178, 179, 5,
		93, 0, 0, 179, 60, 1, 0, 0, 0, 180, 181, 5, 44, 0, 0, 181, 62, 1, 0, 0,
		0, 182, 183, 5, 63, 0, 0, 183, 64, 1, 0, 0, 0, 184, 185, 5, 58, 0, 0, 185,
		66, 1, 0, 0, 0, 186, 187, 5, 59, 0, 0, 187, 68, 1, 0, 0, 0, 188, 190, 5,
		36, 0, 0, 189, 191, 7, 1, 0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0,
		0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 70, 1, 0, 0, 0, 194,
		198, 7, 2, 0, 0, 195, 197, 7, 1, 0, 0, 196, 195, 1, 0, 0, 0, 197, 200,
		1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 72, 1, 0,
		0, 0, 200, 198, 1, 0, 0, 0, 201, 202, 5, 61, 0, 0, 202, 74, 1, 0, 0, 0,
		203, 208, 5, 34, 0, 0, 204, 207, 8, 3, 0, 0, 205, 207, 3, 79, 39, 0, 206,
		204, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206,
		1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0, 210, 208, 1, 0,
		0, 0, 211, 212, 5, 34, 0, 0, 212, 76, 1, 0, 0, 0, 213, 215, 7, 4, 0, 0,
		214, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 78, 1, 0, 0, 0, 218, 219, 5, 92, 0, 0, 219, 220,
		7, 5, 0, 0, 220, 80, 1, 0, 0, 0, 7, 0, 84, 192, 198, 206, 208, 216, 1,
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
	LushLexerWHITESPACE  = 1
	LushLexerTRUE        = 2
	LushLexerFALSE       = 3
	LushLexerIF          = 4
	LushLexerFOR         = 5
	LushLexerSTRING_TYPE = 6
	LushLexerINT_TYPE    = 7
	LushLexerBOOL_TYPE   = 8
	LushLexerFUNC        = 9
	LushLexerRETURN      = 10
	LushLexerLAND        = 11
	LushLexerLOR         = 12
	LushLexerNOT         = 13
	LushLexerMINUS       = 14
	LushLexerPLUS        = 15
	LushLexerMUL         = 16
	LushLexerDIV         = 17
	LushLexerMOD         = 18
	LushLexerLT          = 19
	LushLexerLTE         = 20
	LushLexerGT          = 21
	LushLexerGTE         = 22
	LushLexerEQ          = 23
	LushLexerNEQ         = 24
	LushLexerLPAREN      = 25
	LushLexerRPAREN      = 26
	LushLexerLCUR        = 27
	LushLexerRCUR        = 28
	LushLexerLSQ         = 29
	LushLexerRSQ         = 30
	LushLexerCOMMA       = 31
	LushLexerQUESTION    = 32
	LushLexerCOLON       = 33
	LushLexerSEMICOLON   = 34
	LushLexerENVVAR      = 35
	LushLexerID          = 36
	LushLexerASSIGN      = 37
	LushLexerSTRING      = 38
	LushLexerNUMBER      = 39
)
