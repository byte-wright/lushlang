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
		"", "", "'import'", "'true'", "'false'", "'if'", "'for'", "'string'",
		"'int'", "'bool'", "'func'", "'return'", "'&&'", "'||'", "'!'", "'-'",
		"'+'", "'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='",
		"'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'?'", "':'", "';'",
		"'.'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "IMPORT", "TRUE", "FALSE", "IF", "FOR", "STRING_TYPE",
		"INT_TYPE", "BOOL_TYPE", "FUNC", "RETURN", "LAND", "LOR", "NOT", "MINUS",
		"PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ",
		"LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "SEMICOLON", "DOT", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
		"EXTERNAL_CODE",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "IMPORT", "TRUE", "FALSE", "IF", "FOR", "STRING_TYPE",
		"INT_TYPE", "BOOL_TYPE", "FUNC", "RETURN", "LAND", "LOR", "NOT", "MINUS",
		"PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ",
		"LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "SEMICOLON", "DOT", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
		"EXTERNAL_CODE", "ESCAPED_VALUE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 251, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 4, 0, 89, 8, 0, 11, 0, 12, 0, 90, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		4, 36, 206, 8, 36, 11, 36, 12, 36, 207, 1, 37, 1, 37, 5, 37, 212, 8, 37,
		10, 37, 12, 37, 215, 9, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 5, 39, 222,
		8, 39, 10, 39, 12, 39, 225, 9, 39, 1, 39, 1, 39, 1, 40, 4, 40, 230, 8,
		40, 11, 40, 12, 40, 231, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41,
		240, 8, 41, 10, 41, 12, 41, 243, 9, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 42, 1, 241, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 0,
		1, 0, 7, 3, 0, 9, 10, 13, 13, 32, 32, 4, 0, 48, 48, 65, 90, 95, 95, 97,
		122, 2, 0, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 1, 0, 48, 57, 2, 0, 10,
		10, 13, 13, 4, 0, 34, 34, 92, 92, 110, 110, 116, 116, 257, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 88, 1, 0, 0, 0,
		3, 94, 1, 0, 0, 0, 5, 101, 1, 0, 0, 0, 7, 106, 1, 0, 0, 0, 9, 112, 1, 0,
		0, 0, 11, 115, 1, 0, 0, 0, 13, 119, 1, 0, 0, 0, 15, 126, 1, 0, 0, 0, 17,
		130, 1, 0, 0, 0, 19, 135, 1, 0, 0, 0, 21, 140, 1, 0, 0, 0, 23, 147, 1,
		0, 0, 0, 25, 150, 1, 0, 0, 0, 27, 153, 1, 0, 0, 0, 29, 155, 1, 0, 0, 0,
		31, 157, 1, 0, 0, 0, 33, 159, 1, 0, 0, 0, 35, 161, 1, 0, 0, 0, 37, 163,
		1, 0, 0, 0, 39, 165, 1, 0, 0, 0, 41, 167, 1, 0, 0, 0, 43, 170, 1, 0, 0,
		0, 45, 172, 1, 0, 0, 0, 47, 175, 1, 0, 0, 0, 49, 178, 1, 0, 0, 0, 51, 181,
		1, 0, 0, 0, 53, 183, 1, 0, 0, 0, 55, 185, 1, 0, 0, 0, 57, 187, 1, 0, 0,
		0, 59, 189, 1, 0, 0, 0, 61, 191, 1, 0, 0, 0, 63, 193, 1, 0, 0, 0, 65, 195,
		1, 0, 0, 0, 67, 197, 1, 0, 0, 0, 69, 199, 1, 0, 0, 0, 71, 201, 1, 0, 0,
		0, 73, 203, 1, 0, 0, 0, 75, 209, 1, 0, 0, 0, 77, 216, 1, 0, 0, 0, 79, 218,
		1, 0, 0, 0, 81, 229, 1, 0, 0, 0, 83, 233, 1, 0, 0, 0, 85, 248, 1, 0, 0,
		0, 87, 89, 7, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88,
		1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 6, 0, 0, 0,
		93, 2, 1, 0, 0, 0, 94, 95, 5, 105, 0, 0, 95, 96, 5, 109, 0, 0, 96, 97,
		5, 112, 0, 0, 97, 98, 5, 111, 0, 0, 98, 99, 5, 114, 0, 0, 99, 100, 5, 116,
		0, 0, 100, 4, 1, 0, 0, 0, 101, 102, 5, 116, 0, 0, 102, 103, 5, 114, 0,
		0, 103, 104, 5, 117, 0, 0, 104, 105, 5, 101, 0, 0, 105, 6, 1, 0, 0, 0,
		106, 107, 5, 102, 0, 0, 107, 108, 5, 97, 0, 0, 108, 109, 5, 108, 0, 0,
		109, 110, 5, 115, 0, 0, 110, 111, 5, 101, 0, 0, 111, 8, 1, 0, 0, 0, 112,
		113, 5, 105, 0, 0, 113, 114, 5, 102, 0, 0, 114, 10, 1, 0, 0, 0, 115, 116,
		5, 102, 0, 0, 116, 117, 5, 111, 0, 0, 117, 118, 5, 114, 0, 0, 118, 12,
		1, 0, 0, 0, 119, 120, 5, 115, 0, 0, 120, 121, 5, 116, 0, 0, 121, 122, 5,
		114, 0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5, 110, 0, 0, 124, 125, 5,
		103, 0, 0, 125, 14, 1, 0, 0, 0, 126, 127, 5, 105, 0, 0, 127, 128, 5, 110,
		0, 0, 128, 129, 5, 116, 0, 0, 129, 16, 1, 0, 0, 0, 130, 131, 5, 98, 0,
		0, 131, 132, 5, 111, 0, 0, 132, 133, 5, 111, 0, 0, 133, 134, 5, 108, 0,
		0, 134, 18, 1, 0, 0, 0, 135, 136, 5, 102, 0, 0, 136, 137, 5, 117, 0, 0,
		137, 138, 5, 110, 0, 0, 138, 139, 5, 99, 0, 0, 139, 20, 1, 0, 0, 0, 140,
		141, 5, 114, 0, 0, 141, 142, 5, 101, 0, 0, 142, 143, 5, 116, 0, 0, 143,
		144, 5, 117, 0, 0, 144, 145, 5, 114, 0, 0, 145, 146, 5, 110, 0, 0, 146,
		22, 1, 0, 0, 0, 147, 148, 5, 38, 0, 0, 148, 149, 5, 38, 0, 0, 149, 24,
		1, 0, 0, 0, 150, 151, 5, 124, 0, 0, 151, 152, 5, 124, 0, 0, 152, 26, 1,
		0, 0, 0, 153, 154, 5, 33, 0, 0, 154, 28, 1, 0, 0, 0, 155, 156, 5, 45, 0,
		0, 156, 30, 1, 0, 0, 0, 157, 158, 5, 43, 0, 0, 158, 32, 1, 0, 0, 0, 159,
		160, 5, 42, 0, 0, 160, 34, 1, 0, 0, 0, 161, 162, 5, 47, 0, 0, 162, 36,
		1, 0, 0, 0, 163, 164, 5, 37, 0, 0, 164, 38, 1, 0, 0, 0, 165, 166, 5, 60,
		0, 0, 166, 40, 1, 0, 0, 0, 167, 168, 5, 60, 0, 0, 168, 169, 5, 61, 0, 0,
		169, 42, 1, 0, 0, 0, 170, 171, 5, 62, 0, 0, 171, 44, 1, 0, 0, 0, 172, 173,
		5, 62, 0, 0, 173, 174, 5, 61, 0, 0, 174, 46, 1, 0, 0, 0, 175, 176, 5, 61,
		0, 0, 176, 177, 5, 61, 0, 0, 177, 48, 1, 0, 0, 0, 178, 179, 5, 33, 0, 0,
		179, 180, 5, 61, 0, 0, 180, 50, 1, 0, 0, 0, 181, 182, 5, 40, 0, 0, 182,
		52, 1, 0, 0, 0, 183, 184, 5, 41, 0, 0, 184, 54, 1, 0, 0, 0, 185, 186, 5,
		123, 0, 0, 186, 56, 1, 0, 0, 0, 187, 188, 5, 125, 0, 0, 188, 58, 1, 0,
		0, 0, 189, 190, 5, 91, 0, 0, 190, 60, 1, 0, 0, 0, 191, 192, 5, 93, 0, 0,
		192, 62, 1, 0, 0, 0, 193, 194, 5, 44, 0, 0, 194, 64, 1, 0, 0, 0, 195, 196,
		5, 63, 0, 0, 196, 66, 1, 0, 0, 0, 197, 198, 5, 58, 0, 0, 198, 68, 1, 0,
		0, 0, 199, 200, 5, 59, 0, 0, 200, 70, 1, 0, 0, 0, 201, 202, 5, 46, 0, 0,
		202, 72, 1, 0, 0, 0, 203, 205, 5, 36, 0, 0, 204, 206, 7, 1, 0, 0, 205,
		204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208,
		1, 0, 0, 0, 208, 74, 1, 0, 0, 0, 209, 213, 7, 2, 0, 0, 210, 212, 7, 1,
		0, 0, 211, 210, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0,
		213, 214, 1, 0, 0, 0, 214, 76, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217,
		5, 61, 0, 0, 217, 78, 1, 0, 0, 0, 218, 223, 5, 34, 0, 0, 219, 222, 8, 3,
		0, 0, 220, 222, 3, 85, 42, 0, 221, 219, 1, 0, 0, 0, 221, 220, 1, 0, 0,
		0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		226, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 5, 34, 0, 0, 227, 80,
		1, 0, 0, 0, 228, 230, 7, 4, 0, 0, 229, 228, 1, 0, 0, 0, 230, 231, 1, 0,
		0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 82, 1, 0, 0, 0,
		233, 234, 5, 96, 0, 0, 234, 235, 5, 96, 0, 0, 235, 236, 5, 96, 0, 0, 236,
		241, 1, 0, 0, 0, 237, 240, 9, 0, 0, 0, 238, 240, 7, 5, 0, 0, 239, 237,
		1, 0, 0, 0, 239, 238, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0, 241, 242, 1, 0,
		0, 0, 241, 239, 1, 0, 0, 0, 242, 244, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0,
		244, 245, 5, 96, 0, 0, 245, 246, 5, 96, 0, 0, 246, 247, 5, 96, 0, 0, 247,
		84, 1, 0, 0, 0, 248, 249, 5, 92, 0, 0, 249, 250, 7, 6, 0, 0, 250, 86, 1,
		0, 0, 0, 9, 0, 90, 207, 213, 221, 223, 231, 239, 241, 1, 6, 0, 0,
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
	LushLexerWHITESPACE    = 1
	LushLexerIMPORT        = 2
	LushLexerTRUE          = 3
	LushLexerFALSE         = 4
	LushLexerIF            = 5
	LushLexerFOR           = 6
	LushLexerSTRING_TYPE   = 7
	LushLexerINT_TYPE      = 8
	LushLexerBOOL_TYPE     = 9
	LushLexerFUNC          = 10
	LushLexerRETURN        = 11
	LushLexerLAND          = 12
	LushLexerLOR           = 13
	LushLexerNOT           = 14
	LushLexerMINUS         = 15
	LushLexerPLUS          = 16
	LushLexerMUL           = 17
	LushLexerDIV           = 18
	LushLexerMOD           = 19
	LushLexerLT            = 20
	LushLexerLTE           = 21
	LushLexerGT            = 22
	LushLexerGTE           = 23
	LushLexerEQ            = 24
	LushLexerNEQ           = 25
	LushLexerLPAREN        = 26
	LushLexerRPAREN        = 27
	LushLexerLCUR          = 28
	LushLexerRCUR          = 29
	LushLexerLSQ           = 30
	LushLexerRSQ           = 31
	LushLexerCOMMA         = 32
	LushLexerQUESTION      = 33
	LushLexerCOLON         = 34
	LushLexerSEMICOLON     = 35
	LushLexerDOT           = 36
	LushLexerENVVAR        = 37
	LushLexerID            = 38
	LushLexerASSIGN        = 39
	LushLexerSTRING        = 40
	LushLexerNUMBER        = 41
	LushLexerEXTERNAL_CODE = 42
)
