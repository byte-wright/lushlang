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
		"", "", "'import'", "'true'", "'false'", "'if'", "'else'", "'for'",
		"'string'", "'int'", "'bool'", "'func'", "'return'", "'&&'", "'||'",
		"'!'", "'-'", "'+'", "'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='",
		"'=='", "'!='", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'?'",
		"':'", "';'", "'.'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "IMPORT", "TRUE", "FALSE", "IF", "ELSE", "FOR", "STRING_TYPE",
		"INT_TYPE", "BOOL_TYPE", "FUNC", "RETURN", "LAND", "LOR", "NOT", "MINUS",
		"PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ",
		"LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "SEMICOLON", "DOT", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
		"EXTERNAL_CODE",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "IMPORT", "TRUE", "FALSE", "IF", "ELSE", "FOR", "STRING_TYPE",
		"INT_TYPE", "BOOL_TYPE", "FUNC", "RETURN", "LAND", "LOR", "NOT", "MINUS",
		"PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ",
		"LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA", "QUESTION",
		"COLON", "SEMICOLON", "DOT", "ENVVAR", "ID", "ASSIGN", "STRING", "NUMBER",
		"EXTERNAL_CODE", "ESCAPED_VALUE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 258, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 4, 0, 91, 8, 0, 11, 0, 12, 0, 92,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 4, 37, 213, 8, 37, 11,
		37, 12, 37, 214, 1, 38, 1, 38, 5, 38, 219, 8, 38, 10, 38, 12, 38, 222,
		9, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40, 229, 8, 40, 10, 40, 12,
		40, 232, 9, 40, 1, 40, 1, 40, 1, 41, 4, 41, 237, 8, 41, 11, 41, 12, 41,
		238, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 247, 8, 42, 10, 42,
		12, 42, 250, 9, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1,
		248, 0, 44, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 0, 1, 0, 7, 3,
		0, 9, 10, 13, 13, 32, 32, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0,
		95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 1, 0, 48, 57, 2, 0, 10, 10, 13,
		13, 4, 0, 34, 34, 92, 92, 110, 110, 116, 116, 264, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0,
		0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0,
		0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0,
		0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 1, 90, 1,
		0, 0, 0, 3, 96, 1, 0, 0, 0, 5, 103, 1, 0, 0, 0, 7, 108, 1, 0, 0, 0, 9,
		114, 1, 0, 0, 0, 11, 117, 1, 0, 0, 0, 13, 122, 1, 0, 0, 0, 15, 126, 1,
		0, 0, 0, 17, 133, 1, 0, 0, 0, 19, 137, 1, 0, 0, 0, 21, 142, 1, 0, 0, 0,
		23, 147, 1, 0, 0, 0, 25, 154, 1, 0, 0, 0, 27, 157, 1, 0, 0, 0, 29, 160,
		1, 0, 0, 0, 31, 162, 1, 0, 0, 0, 33, 164, 1, 0, 0, 0, 35, 166, 1, 0, 0,
		0, 37, 168, 1, 0, 0, 0, 39, 170, 1, 0, 0, 0, 41, 172, 1, 0, 0, 0, 43, 174,
		1, 0, 0, 0, 45, 177, 1, 0, 0, 0, 47, 179, 1, 0, 0, 0, 49, 182, 1, 0, 0,
		0, 51, 185, 1, 0, 0, 0, 53, 188, 1, 0, 0, 0, 55, 190, 1, 0, 0, 0, 57, 192,
		1, 0, 0, 0, 59, 194, 1, 0, 0, 0, 61, 196, 1, 0, 0, 0, 63, 198, 1, 0, 0,
		0, 65, 200, 1, 0, 0, 0, 67, 202, 1, 0, 0, 0, 69, 204, 1, 0, 0, 0, 71, 206,
		1, 0, 0, 0, 73, 208, 1, 0, 0, 0, 75, 210, 1, 0, 0, 0, 77, 216, 1, 0, 0,
		0, 79, 223, 1, 0, 0, 0, 81, 225, 1, 0, 0, 0, 83, 236, 1, 0, 0, 0, 85, 240,
		1, 0, 0, 0, 87, 255, 1, 0, 0, 0, 89, 91, 7, 0, 0, 0, 90, 89, 1, 0, 0, 0,
		91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1,
		0, 0, 0, 94, 95, 6, 0, 0, 0, 95, 2, 1, 0, 0, 0, 96, 97, 5, 105, 0, 0, 97,
		98, 5, 109, 0, 0, 98, 99, 5, 112, 0, 0, 99, 100, 5, 111, 0, 0, 100, 101,
		5, 114, 0, 0, 101, 102, 5, 116, 0, 0, 102, 4, 1, 0, 0, 0, 103, 104, 5,
		116, 0, 0, 104, 105, 5, 114, 0, 0, 105, 106, 5, 117, 0, 0, 106, 107, 5,
		101, 0, 0, 107, 6, 1, 0, 0, 0, 108, 109, 5, 102, 0, 0, 109, 110, 5, 97,
		0, 0, 110, 111, 5, 108, 0, 0, 111, 112, 5, 115, 0, 0, 112, 113, 5, 101,
		0, 0, 113, 8, 1, 0, 0, 0, 114, 115, 5, 105, 0, 0, 115, 116, 5, 102, 0,
		0, 116, 10, 1, 0, 0, 0, 117, 118, 5, 101, 0, 0, 118, 119, 5, 108, 0, 0,
		119, 120, 5, 115, 0, 0, 120, 121, 5, 101, 0, 0, 121, 12, 1, 0, 0, 0, 122,
		123, 5, 102, 0, 0, 123, 124, 5, 111, 0, 0, 124, 125, 5, 114, 0, 0, 125,
		14, 1, 0, 0, 0, 126, 127, 5, 115, 0, 0, 127, 128, 5, 116, 0, 0, 128, 129,
		5, 114, 0, 0, 129, 130, 5, 105, 0, 0, 130, 131, 5, 110, 0, 0, 131, 132,
		5, 103, 0, 0, 132, 16, 1, 0, 0, 0, 133, 134, 5, 105, 0, 0, 134, 135, 5,
		110, 0, 0, 135, 136, 5, 116, 0, 0, 136, 18, 1, 0, 0, 0, 137, 138, 5, 98,
		0, 0, 138, 139, 5, 111, 0, 0, 139, 140, 5, 111, 0, 0, 140, 141, 5, 108,
		0, 0, 141, 20, 1, 0, 0, 0, 142, 143, 5, 102, 0, 0, 143, 144, 5, 117, 0,
		0, 144, 145, 5, 110, 0, 0, 145, 146, 5, 99, 0, 0, 146, 22, 1, 0, 0, 0,
		147, 148, 5, 114, 0, 0, 148, 149, 5, 101, 0, 0, 149, 150, 5, 116, 0, 0,
		150, 151, 5, 117, 0, 0, 151, 152, 5, 114, 0, 0, 152, 153, 5, 110, 0, 0,
		153, 24, 1, 0, 0, 0, 154, 155, 5, 38, 0, 0, 155, 156, 5, 38, 0, 0, 156,
		26, 1, 0, 0, 0, 157, 158, 5, 124, 0, 0, 158, 159, 5, 124, 0, 0, 159, 28,
		1, 0, 0, 0, 160, 161, 5, 33, 0, 0, 161, 30, 1, 0, 0, 0, 162, 163, 5, 45,
		0, 0, 163, 32, 1, 0, 0, 0, 164, 165, 5, 43, 0, 0, 165, 34, 1, 0, 0, 0,
		166, 167, 5, 42, 0, 0, 167, 36, 1, 0, 0, 0, 168, 169, 5, 47, 0, 0, 169,
		38, 1, 0, 0, 0, 170, 171, 5, 37, 0, 0, 171, 40, 1, 0, 0, 0, 172, 173, 5,
		60, 0, 0, 173, 42, 1, 0, 0, 0, 174, 175, 5, 60, 0, 0, 175, 176, 5, 61,
		0, 0, 176, 44, 1, 0, 0, 0, 177, 178, 5, 62, 0, 0, 178, 46, 1, 0, 0, 0,
		179, 180, 5, 62, 0, 0, 180, 181, 5, 61, 0, 0, 181, 48, 1, 0, 0, 0, 182,
		183, 5, 61, 0, 0, 183, 184, 5, 61, 0, 0, 184, 50, 1, 0, 0, 0, 185, 186,
		5, 33, 0, 0, 186, 187, 5, 61, 0, 0, 187, 52, 1, 0, 0, 0, 188, 189, 5, 40,
		0, 0, 189, 54, 1, 0, 0, 0, 190, 191, 5, 41, 0, 0, 191, 56, 1, 0, 0, 0,
		192, 193, 5, 123, 0, 0, 193, 58, 1, 0, 0, 0, 194, 195, 5, 125, 0, 0, 195,
		60, 1, 0, 0, 0, 196, 197, 5, 91, 0, 0, 197, 62, 1, 0, 0, 0, 198, 199, 5,
		93, 0, 0, 199, 64, 1, 0, 0, 0, 200, 201, 5, 44, 0, 0, 201, 66, 1, 0, 0,
		0, 202, 203, 5, 63, 0, 0, 203, 68, 1, 0, 0, 0, 204, 205, 5, 58, 0, 0, 205,
		70, 1, 0, 0, 0, 206, 207, 5, 59, 0, 0, 207, 72, 1, 0, 0, 0, 208, 209, 5,
		46, 0, 0, 209, 74, 1, 0, 0, 0, 210, 212, 5, 36, 0, 0, 211, 213, 7, 1, 0,
		0, 212, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214,
		215, 1, 0, 0, 0, 215, 76, 1, 0, 0, 0, 216, 220, 7, 2, 0, 0, 217, 219, 7,
		1, 0, 0, 218, 217, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218, 1, 0, 0,
		0, 220, 221, 1, 0, 0, 0, 221, 78, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 223,
		224, 5, 61, 0, 0, 224, 80, 1, 0, 0, 0, 225, 230, 5, 34, 0, 0, 226, 229,
		8, 3, 0, 0, 227, 229, 3, 87, 43, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1,
		0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0,
		0, 231, 233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 234, 5, 34, 0, 0, 234,
		82, 1, 0, 0, 0, 235, 237, 7, 4, 0, 0, 236, 235, 1, 0, 0, 0, 237, 238, 1,
		0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 84, 1, 0, 0,
		0, 240, 241, 5, 96, 0, 0, 241, 242, 5, 96, 0, 0, 242, 243, 5, 96, 0, 0,
		243, 248, 1, 0, 0, 0, 244, 247, 9, 0, 0, 0, 245, 247, 7, 5, 0, 0, 246,
		244, 1, 0, 0, 0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0,
		0, 0, 251, 252, 5, 96, 0, 0, 252, 253, 5, 96, 0, 0, 253, 254, 5, 96, 0,
		0, 254, 86, 1, 0, 0, 0, 255, 256, 5, 92, 0, 0, 256, 257, 7, 6, 0, 0, 257,
		88, 1, 0, 0, 0, 9, 0, 92, 214, 220, 228, 230, 238, 246, 248, 1, 6, 0, 0,
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
	LushLexerELSE          = 6
	LushLexerFOR           = 7
	LushLexerSTRING_TYPE   = 8
	LushLexerINT_TYPE      = 9
	LushLexerBOOL_TYPE     = 10
	LushLexerFUNC          = 11
	LushLexerRETURN        = 12
	LushLexerLAND          = 13
	LushLexerLOR           = 14
	LushLexerNOT           = 15
	LushLexerMINUS         = 16
	LushLexerPLUS          = 17
	LushLexerMUL           = 18
	LushLexerDIV           = 19
	LushLexerMOD           = 20
	LushLexerLT            = 21
	LushLexerLTE           = 22
	LushLexerGT            = 23
	LushLexerGTE           = 24
	LushLexerEQ            = 25
	LushLexerNEQ           = 26
	LushLexerLPAREN        = 27
	LushLexerRPAREN        = 28
	LushLexerLCUR          = 29
	LushLexerRCUR          = 30
	LushLexerLSQ           = 31
	LushLexerRSQ           = 32
	LushLexerCOMMA         = 33
	LushLexerQUESTION      = 34
	LushLexerCOLON         = 35
	LushLexerSEMICOLON     = 36
	LushLexerDOT           = 37
	LushLexerENVVAR        = 38
	LushLexerID            = 39
	LushLexerASSIGN        = 40
	LushLexerSTRING        = 41
	LushLexerNUMBER        = 42
	LushLexerEXTERNAL_CODE = 43
)
