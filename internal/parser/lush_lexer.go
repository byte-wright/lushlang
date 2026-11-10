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
		"'string'", "'int'", "'bool'", "'func'", "'return'", "'global'", "'&&'",
		"'||'", "'!'", "'-'", "'+'", "'*'", "'/'", "'%'", "'<'", "'<='", "'>'",
		"'>='", "'=='", "'!='", "'('", "')'", "'{'", "'}'", "'['", "']'", "','",
		"'?'", "':'", "';'", "'.'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "IMPORT", "TRUE", "FALSE", "IF", "ELSE", "FOR", "STRING_TYPE",
		"INT_TYPE", "BOOL_TYPE", "FUNC", "RETURN", "GLOBAL", "LAND", "LOR",
		"NOT", "MINUS", "PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE",
		"EQ", "NEQ", "LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA",
		"QUESTION", "COLON", "SEMICOLON", "DOT", "ENVVAR", "ID", "ASSIGN", "STRING",
		"NUMBER", "EXTERNAL_CODE",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "IMPORT", "TRUE", "FALSE", "IF", "ELSE", "FOR", "STRING_TYPE",
		"INT_TYPE", "BOOL_TYPE", "FUNC", "RETURN", "GLOBAL", "LAND", "LOR",
		"NOT", "MINUS", "PLUS", "MUL", "DIV", "MOD", "LT", "LTE", "GT", "GTE",
		"EQ", "NEQ", "LPAREN", "RPAREN", "LCUR", "RCUR", "LSQ", "RSQ", "COMMA",
		"QUESTION", "COLON", "SEMICOLON", "DOT", "ENVVAR", "ID", "ASSIGN", "STRING",
		"NUMBER", "EXTERNAL_CODE", "ESCAPED_VALUE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 267, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 4, 0, 93, 8, 0, 11,
		0, 12, 0, 94, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 4, 38, 222, 8, 38, 11, 38, 12,
		38, 223, 1, 39, 1, 39, 5, 39, 228, 8, 39, 10, 39, 12, 39, 231, 9, 39, 1,
		40, 1, 40, 1, 41, 1, 41, 1, 41, 5, 41, 238, 8, 41, 10, 41, 12, 41, 241,
		9, 41, 1, 41, 1, 41, 1, 42, 4, 42, 246, 8, 42, 11, 42, 12, 42, 247, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 256, 8, 43, 10, 43, 12, 43,
		259, 9, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 257, 0,
		45, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75,
		38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 0, 1, 0, 7, 3,
		0, 9, 10, 13, 13, 32, 32, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0,
		65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 1, 0, 48, 57, 2, 0, 10,
		10, 13, 13, 4, 0, 34, 34, 92, 92, 110, 110, 116, 116, 273, 0, 1, 1, 0,
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
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0,
		0, 87, 1, 0, 0, 0, 1, 92, 1, 0, 0, 0, 3, 98, 1, 0, 0, 0, 5, 105, 1, 0,
		0, 0, 7, 110, 1, 0, 0, 0, 9, 116, 1, 0, 0, 0, 11, 119, 1, 0, 0, 0, 13,
		124, 1, 0, 0, 0, 15, 128, 1, 0, 0, 0, 17, 135, 1, 0, 0, 0, 19, 139, 1,
		0, 0, 0, 21, 144, 1, 0, 0, 0, 23, 149, 1, 0, 0, 0, 25, 156, 1, 0, 0, 0,
		27, 163, 1, 0, 0, 0, 29, 166, 1, 0, 0, 0, 31, 169, 1, 0, 0, 0, 33, 171,
		1, 0, 0, 0, 35, 173, 1, 0, 0, 0, 37, 175, 1, 0, 0, 0, 39, 177, 1, 0, 0,
		0, 41, 179, 1, 0, 0, 0, 43, 181, 1, 0, 0, 0, 45, 183, 1, 0, 0, 0, 47, 186,
		1, 0, 0, 0, 49, 188, 1, 0, 0, 0, 51, 191, 1, 0, 0, 0, 53, 194, 1, 0, 0,
		0, 55, 197, 1, 0, 0, 0, 57, 199, 1, 0, 0, 0, 59, 201, 1, 0, 0, 0, 61, 203,
		1, 0, 0, 0, 63, 205, 1, 0, 0, 0, 65, 207, 1, 0, 0, 0, 67, 209, 1, 0, 0,
		0, 69, 211, 1, 0, 0, 0, 71, 213, 1, 0, 0, 0, 73, 215, 1, 0, 0, 0, 75, 217,
		1, 0, 0, 0, 77, 219, 1, 0, 0, 0, 79, 225, 1, 0, 0, 0, 81, 232, 1, 0, 0,
		0, 83, 234, 1, 0, 0, 0, 85, 245, 1, 0, 0, 0, 87, 249, 1, 0, 0, 0, 89, 264,
		1, 0, 0, 0, 91, 93, 7, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 6,
		0, 0, 0, 97, 2, 1, 0, 0, 0, 98, 99, 5, 105, 0, 0, 99, 100, 5, 109, 0, 0,
		100, 101, 5, 112, 0, 0, 101, 102, 5, 111, 0, 0, 102, 103, 5, 114, 0, 0,
		103, 104, 5, 116, 0, 0, 104, 4, 1, 0, 0, 0, 105, 106, 5, 116, 0, 0, 106,
		107, 5, 114, 0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 101, 0, 0, 109,
		6, 1, 0, 0, 0, 110, 111, 5, 102, 0, 0, 111, 112, 5, 97, 0, 0, 112, 113,
		5, 108, 0, 0, 113, 114, 5, 115, 0, 0, 114, 115, 5, 101, 0, 0, 115, 8, 1,
		0, 0, 0, 116, 117, 5, 105, 0, 0, 117, 118, 5, 102, 0, 0, 118, 10, 1, 0,
		0, 0, 119, 120, 5, 101, 0, 0, 120, 121, 5, 108, 0, 0, 121, 122, 5, 115,
		0, 0, 122, 123, 5, 101, 0, 0, 123, 12, 1, 0, 0, 0, 124, 125, 5, 102, 0,
		0, 125, 126, 5, 111, 0, 0, 126, 127, 5, 114, 0, 0, 127, 14, 1, 0, 0, 0,
		128, 129, 5, 115, 0, 0, 129, 130, 5, 116, 0, 0, 130, 131, 5, 114, 0, 0,
		131, 132, 5, 105, 0, 0, 132, 133, 5, 110, 0, 0, 133, 134, 5, 103, 0, 0,
		134, 16, 1, 0, 0, 0, 135, 136, 5, 105, 0, 0, 136, 137, 5, 110, 0, 0, 137,
		138, 5, 116, 0, 0, 138, 18, 1, 0, 0, 0, 139, 140, 5, 98, 0, 0, 140, 141,
		5, 111, 0, 0, 141, 142, 5, 111, 0, 0, 142, 143, 5, 108, 0, 0, 143, 20,
		1, 0, 0, 0, 144, 145, 5, 102, 0, 0, 145, 146, 5, 117, 0, 0, 146, 147, 5,
		110, 0, 0, 147, 148, 5, 99, 0, 0, 148, 22, 1, 0, 0, 0, 149, 150, 5, 114,
		0, 0, 150, 151, 5, 101, 0, 0, 151, 152, 5, 116, 0, 0, 152, 153, 5, 117,
		0, 0, 153, 154, 5, 114, 0, 0, 154, 155, 5, 110, 0, 0, 155, 24, 1, 0, 0,
		0, 156, 157, 5, 103, 0, 0, 157, 158, 5, 108, 0, 0, 158, 159, 5, 111, 0,
		0, 159, 160, 5, 98, 0, 0, 160, 161, 5, 97, 0, 0, 161, 162, 5, 108, 0, 0,
		162, 26, 1, 0, 0, 0, 163, 164, 5, 38, 0, 0, 164, 165, 5, 38, 0, 0, 165,
		28, 1, 0, 0, 0, 166, 167, 5, 124, 0, 0, 167, 168, 5, 124, 0, 0, 168, 30,
		1, 0, 0, 0, 169, 170, 5, 33, 0, 0, 170, 32, 1, 0, 0, 0, 171, 172, 5, 45,
		0, 0, 172, 34, 1, 0, 0, 0, 173, 174, 5, 43, 0, 0, 174, 36, 1, 0, 0, 0,
		175, 176, 5, 42, 0, 0, 176, 38, 1, 0, 0, 0, 177, 178, 5, 47, 0, 0, 178,
		40, 1, 0, 0, 0, 179, 180, 5, 37, 0, 0, 180, 42, 1, 0, 0, 0, 181, 182, 5,
		60, 0, 0, 182, 44, 1, 0, 0, 0, 183, 184, 5, 60, 0, 0, 184, 185, 5, 61,
		0, 0, 185, 46, 1, 0, 0, 0, 186, 187, 5, 62, 0, 0, 187, 48, 1, 0, 0, 0,
		188, 189, 5, 62, 0, 0, 189, 190, 5, 61, 0, 0, 190, 50, 1, 0, 0, 0, 191,
		192, 5, 61, 0, 0, 192, 193, 5, 61, 0, 0, 193, 52, 1, 0, 0, 0, 194, 195,
		5, 33, 0, 0, 195, 196, 5, 61, 0, 0, 196, 54, 1, 0, 0, 0, 197, 198, 5, 40,
		0, 0, 198, 56, 1, 0, 0, 0, 199, 200, 5, 41, 0, 0, 200, 58, 1, 0, 0, 0,
		201, 202, 5, 123, 0, 0, 202, 60, 1, 0, 0, 0, 203, 204, 5, 125, 0, 0, 204,
		62, 1, 0, 0, 0, 205, 206, 5, 91, 0, 0, 206, 64, 1, 0, 0, 0, 207, 208, 5,
		93, 0, 0, 208, 66, 1, 0, 0, 0, 209, 210, 5, 44, 0, 0, 210, 68, 1, 0, 0,
		0, 211, 212, 5, 63, 0, 0, 212, 70, 1, 0, 0, 0, 213, 214, 5, 58, 0, 0, 214,
		72, 1, 0, 0, 0, 215, 216, 5, 59, 0, 0, 216, 74, 1, 0, 0, 0, 217, 218, 5,
		46, 0, 0, 218, 76, 1, 0, 0, 0, 219, 221, 5, 36, 0, 0, 220, 222, 7, 1, 0,
		0, 221, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223,
		224, 1, 0, 0, 0, 224, 78, 1, 0, 0, 0, 225, 229, 7, 2, 0, 0, 226, 228, 7,
		1, 0, 0, 227, 226, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0,
		0, 229, 230, 1, 0, 0, 0, 230, 80, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232,
		233, 5, 61, 0, 0, 233, 82, 1, 0, 0, 0, 234, 239, 5, 34, 0, 0, 235, 238,
		8, 3, 0, 0, 236, 238, 3, 89, 44, 0, 237, 235, 1, 0, 0, 0, 237, 236, 1,
		0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0,
		0, 240, 242, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 243, 5, 34, 0, 0, 243,
		84, 1, 0, 0, 0, 244, 246, 7, 4, 0, 0, 245, 244, 1, 0, 0, 0, 246, 247, 1,
		0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 86, 1, 0, 0,
		0, 249, 250, 5, 96, 0, 0, 250, 251, 5, 96, 0, 0, 251, 252, 5, 96, 0, 0,
		252, 257, 1, 0, 0, 0, 253, 256, 9, 0, 0, 0, 254, 256, 7, 5, 0, 0, 255,
		253, 1, 0, 0, 0, 255, 254, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 258,
		1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 260, 1, 0, 0, 0, 259, 257, 1, 0,
		0, 0, 260, 261, 5, 96, 0, 0, 261, 262, 5, 96, 0, 0, 262, 263, 5, 96, 0,
		0, 263, 88, 1, 0, 0, 0, 264, 265, 5, 92, 0, 0, 265, 266, 7, 6, 0, 0, 266,
		90, 1, 0, 0, 0, 9, 0, 94, 223, 229, 237, 239, 247, 255, 257, 1, 6, 0, 0,
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
	LushLexerGLOBAL        = 13
	LushLexerLAND          = 14
	LushLexerLOR           = 15
	LushLexerNOT           = 16
	LushLexerMINUS         = 17
	LushLexerPLUS          = 18
	LushLexerMUL           = 19
	LushLexerDIV           = 20
	LushLexerMOD           = 21
	LushLexerLT            = 22
	LushLexerLTE           = 23
	LushLexerGT            = 24
	LushLexerGTE           = 25
	LushLexerEQ            = 26
	LushLexerNEQ           = 27
	LushLexerLPAREN        = 28
	LushLexerRPAREN        = 29
	LushLexerLCUR          = 30
	LushLexerRCUR          = 31
	LushLexerLSQ           = 32
	LushLexerRSQ           = 33
	LushLexerCOMMA         = 34
	LushLexerQUESTION      = 35
	LushLexerCOLON         = 36
	LushLexerSEMICOLON     = 37
	LushLexerDOT           = 38
	LushLexerENVVAR        = 39
	LushLexerID            = 40
	LushLexerASSIGN        = 41
	LushLexerSTRING        = 42
	LushLexerNUMBER        = 43
	LushLexerEXTERNAL_CODE = 44
)
