// Code generated from internal/parser/Lush.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lush

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LushParser struct {
	*antlr.BaseParser
}

var LushParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lushParserInit() {
	staticData := &LushParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "'import'", "'as'", "'true'", "'false'", "'if'", "'else'",
		"'for'", "'string'", "'int'", "'bool'", "'func'", "'return'", "'break'",
		"'continue'", "'in'", "'until'", "'global'", "'try'", "'catch'", "'fail'",
		"'error'", "'cmd'", "'&&'", "'||'", "'!'", "'-'", "'+'", "'*'", "'/'",
		"'%'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='", "'('", "')'", "'{'",
		"'}'", "'['", "']'", "','", "'?'", "':'", "';'", "'...'", "'.'", "",
		"", "'='", "'+='", "'-='", "'*='", "'/='", "'%='",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "LINE_COMMENT", "IMPORT", "AS", "TRUE", "FALSE", "IF",
		"ELSE", "FOR", "STRING_TYPE", "INT_TYPE", "BOOL_TYPE", "FUNC", "RETURN",
		"BREAK", "CONTINUE", "IN", "UNTIL", "GLOBAL", "TRY", "CATCH", "FAIL",
		"ERROR_TYPE", "CMD", "LAND", "LOR", "NOT", "MINUS", "PLUS", "MUL", "DIV",
		"MOD", "LT", "LTE", "GT", "GTE", "EQ", "NEQ", "LPAREN", "RPAREN", "LCUR",
		"RCUR", "LSQ", "RSQ", "COMMA", "QUESTION", "COLON", "SEMICOLON", "ELLIPSIS",
		"DOT", "ENVVAR", "ID", "ASSIGN", "PLUS_ASSIGN", "MINUS_ASSIGN", "MUL_ASSIGN",
		"DIV_ASSIGN", "MOD_ASSIGN", "STRING", "NUMBER", "EXTERNAL_CODE",
	}
	staticData.RuleNames = []string{
		"program", "package", "importStatement", "statement", "funcDef", "externalFuncDef",
		"param", "results", "result", "if", "for", "returnStatement", "failStatement",
		"catchClause", "block", "assignment", "assignTarget", "globalAssignment",
		"funcStatement", "cmdStatement", "cmd", "cmdArg", "cmdValue", "expression",
		"atom", "func", "var", "envVar", "string", "number", "bool", "array",
		"type", "primitiveType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 460, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 5, 0, 70, 8, 0, 10, 0, 12, 0, 73,
		9, 0, 1, 0, 1, 0, 5, 0, 77, 8, 0, 10, 0, 12, 0, 80, 9, 0, 1, 0, 1, 0, 1,
		1, 5, 1, 85, 8, 1, 10, 1, 12, 1, 88, 9, 1, 1, 1, 1, 1, 1, 1, 5, 1, 93,
		8, 1, 10, 1, 12, 1, 96, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		104, 8, 2, 1, 3, 1, 3, 3, 3, 108, 8, 3, 1, 3, 3, 3, 111, 8, 3, 1, 3, 1,
		3, 3, 3, 115, 8, 3, 1, 3, 3, 3, 118, 8, 3, 1, 3, 1, 3, 3, 3, 122, 8, 3,
		1, 3, 1, 3, 3, 3, 126, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		134, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 142, 8, 4, 10, 4,
		12, 4, 145, 9, 4, 3, 4, 147, 8, 4, 1, 4, 1, 4, 3, 4, 151, 8, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 161, 8, 5, 10, 5, 12, 5, 164,
		9, 5, 3, 5, 166, 8, 5, 1, 5, 1, 5, 3, 5, 170, 8, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 180, 8, 7, 10, 7, 12, 7, 183, 9, 7, 1,
		8, 1, 8, 3, 8, 187, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 195,
		8, 9, 3, 9, 197, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 217, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 223, 8, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 228, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		5, 11, 235, 8, 11, 10, 11, 12, 11, 238, 9, 11, 3, 11, 240, 8, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 5,
		14, 253, 8, 14, 10, 14, 12, 14, 256, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 5, 15, 263, 8, 15, 10, 15, 12, 15, 266, 9, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 5, 15, 272, 8, 15, 10, 15, 12, 15, 275, 9, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 282, 8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 3,
		18, 288, 8, 18, 1, 18, 1, 18, 3, 18, 292, 8, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 3, 19, 298, 8, 19, 1, 20, 3, 20, 301, 8, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 5, 20, 308, 8, 20, 10, 20, 12, 20, 311, 9, 20, 1, 20, 3, 20,
		314, 8, 20, 3, 20, 316, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 328, 8, 21, 3, 21, 330, 8, 21, 1, 22,
		1, 22, 3, 22, 334, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 340, 8, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 366, 8, 23, 1, 23, 1, 23, 3, 23, 370, 8, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 378, 8, 23, 10, 23, 12,
		23, 381, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 389, 8,
		24, 1, 24, 1, 24, 3, 24, 393, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 3, 24, 402, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5,
		25, 409, 8, 25, 10, 25, 12, 25, 412, 9, 25, 1, 25, 3, 25, 415, 8, 25, 3,
		25, 417, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 423, 8, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 31, 5, 31, 439, 8, 31, 10, 31, 12, 31, 442, 9, 31, 3, 31, 444,
		8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 450, 8, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 456, 8, 32, 1, 33, 1, 33, 1, 33, 0, 1, 46, 34, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 0, 7, 1, 0, 53, 58,
		1, 0, 27, 28, 1, 0, 30, 32, 1, 0, 28, 29, 1, 0, 33, 38, 1, 0, 5, 6, 1,
		0, 10, 12, 509, 0, 71, 1, 0, 0, 0, 2, 86, 1, 0, 0, 0, 4, 99, 1, 0, 0, 0,
		6, 133, 1, 0, 0, 0, 8, 135, 1, 0, 0, 0, 10, 154, 1, 0, 0, 0, 12, 173, 1,
		0, 0, 0, 14, 176, 1, 0, 0, 0, 16, 186, 1, 0, 0, 0, 18, 188, 1, 0, 0, 0,
		20, 227, 1, 0, 0, 0, 22, 229, 1, 0, 0, 0, 24, 241, 1, 0, 0, 0, 26, 246,
		1, 0, 0, 0, 28, 250, 1, 0, 0, 0, 30, 259, 1, 0, 0, 0, 32, 276, 1, 0, 0,
		0, 34, 283, 1, 0, 0, 0, 36, 287, 1, 0, 0, 0, 38, 295, 1, 0, 0, 0, 40, 300,
		1, 0, 0, 0, 42, 329, 1, 0, 0, 0, 44, 333, 1, 0, 0, 0, 46, 339, 1, 0, 0,
		0, 48, 401, 1, 0, 0, 0, 50, 403, 1, 0, 0, 0, 52, 422, 1, 0, 0, 0, 54, 426,
		1, 0, 0, 0, 56, 428, 1, 0, 0, 0, 58, 430, 1, 0, 0, 0, 60, 432, 1, 0, 0,
		0, 62, 449, 1, 0, 0, 0, 64, 455, 1, 0, 0, 0, 66, 457, 1, 0, 0, 0, 68, 70,
		3, 4, 2, 0, 69, 68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		71, 72, 1, 0, 0, 0, 72, 78, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 77, 3,
		6, 3, 0, 75, 77, 3, 8, 4, 0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77,
		80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 1, 0, 0,
		0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 0, 0, 1, 82, 1, 1, 0, 0, 0, 83, 85, 3,
		4, 2, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86,
		87, 1, 0, 0, 0, 87, 94, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 93, 3, 8, 4,
		0, 90, 93, 3, 10, 5, 0, 91, 93, 3, 34, 17, 0, 92, 89, 1, 0, 0, 0, 92, 90,
		1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 98, 5,
		0, 0, 1, 98, 3, 1, 0, 0, 0, 99, 100, 5, 3, 0, 0, 100, 103, 5, 59, 0, 0,
		101, 102, 5, 4, 0, 0, 102, 104, 5, 52, 0, 0, 103, 101, 1, 0, 0, 0, 103,
		104, 1, 0, 0, 0, 104, 5, 1, 0, 0, 0, 105, 107, 3, 30, 15, 0, 106, 108,
		3, 28, 14, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 1,
		0, 0, 0, 109, 111, 3, 26, 13, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0,
		0, 0, 111, 134, 1, 0, 0, 0, 112, 114, 3, 34, 17, 0, 113, 115, 3, 28, 14,
		0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116,
		118, 3, 26, 13, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 134,
		1, 0, 0, 0, 119, 121, 3, 36, 18, 0, 120, 122, 3, 26, 13, 0, 121, 120, 1,
		0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 134, 1, 0, 0, 0, 123, 125, 3, 38, 19,
		0, 124, 126, 3, 26, 13, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0,
		126, 134, 1, 0, 0, 0, 127, 134, 3, 18, 9, 0, 128, 134, 3, 20, 10, 0, 129,
		134, 3, 22, 11, 0, 130, 134, 3, 24, 12, 0, 131, 134, 5, 15, 0, 0, 132,
		134, 5, 16, 0, 0, 133, 105, 1, 0, 0, 0, 133, 112, 1, 0, 0, 0, 133, 119,
		1, 0, 0, 0, 133, 123, 1, 0, 0, 0, 133, 127, 1, 0, 0, 0, 133, 128, 1, 0,
		0, 0, 133, 129, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0,
		133, 132, 1, 0, 0, 0, 134, 7, 1, 0, 0, 0, 135, 136, 5, 13, 0, 0, 136, 137,
		5, 52, 0, 0, 137, 146, 5, 39, 0, 0, 138, 143, 3, 12, 6, 0, 139, 140, 5,
		45, 0, 0, 140, 142, 3, 12, 6, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1, 0,
		0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0,
		145, 143, 1, 0, 0, 0, 146, 138, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 150, 5, 40, 0, 0, 149, 151, 3, 14, 7, 0, 150, 149,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 3, 28,
		14, 0, 153, 9, 1, 0, 0, 0, 154, 155, 5, 13, 0, 0, 155, 156, 5, 52, 0, 0,
		156, 165, 5, 39, 0, 0, 157, 162, 3, 12, 6, 0, 158, 159, 5, 45, 0, 0, 159,
		161, 3, 12, 6, 0, 160, 158, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0,
		0, 0, 165, 157, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0,
		167, 169, 5, 40, 0, 0, 168, 170, 3, 14, 7, 0, 169, 168, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 5, 61, 0, 0, 172, 11,
		1, 0, 0, 0, 173, 174, 5, 52, 0, 0, 174, 175, 3, 64, 32, 0, 175, 13, 1,
		0, 0, 0, 176, 181, 3, 16, 8, 0, 177, 178, 5, 45, 0, 0, 178, 180, 3, 16,
		8, 0, 179, 177, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0,
		181, 182, 1, 0, 0, 0, 182, 15, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 187,
		3, 64, 32, 0, 185, 187, 5, 23, 0, 0, 186, 184, 1, 0, 0, 0, 186, 185, 1,
		0, 0, 0, 187, 17, 1, 0, 0, 0, 188, 189, 5, 7, 0, 0, 189, 190, 3, 46, 23,
		0, 190, 196, 3, 28, 14, 0, 191, 194, 5, 8, 0, 0, 192, 195, 3, 18, 9, 0,
		193, 195, 3, 28, 14, 0, 194, 192, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195,
		197, 1, 0, 0, 0, 196, 191, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 19, 1,
		0, 0, 0, 198, 199, 5, 9, 0, 0, 199, 228, 3, 28, 14, 0, 200, 201, 5, 9,
		0, 0, 201, 202, 3, 46, 23, 0, 202, 203, 3, 28, 14, 0, 203, 228, 1, 0, 0,
		0, 204, 205, 5, 9, 0, 0, 205, 206, 3, 30, 15, 0, 206, 207, 5, 48, 0, 0,
		207, 208, 3, 46, 23, 0, 208, 209, 5, 48, 0, 0, 209, 210, 3, 30, 15, 0,
		210, 211, 3, 28, 14, 0, 211, 228, 1, 0, 0, 0, 212, 213, 5, 9, 0, 0, 213,
		216, 5, 52, 0, 0, 214, 215, 5, 45, 0, 0, 215, 217, 5, 52, 0, 0, 216, 214,
		1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 222, 5, 17,
		0, 0, 219, 220, 3, 46, 23, 0, 220, 221, 5, 18, 0, 0, 221, 223, 1, 0, 0,
		0, 222, 219, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		225, 3, 46, 23, 0, 225, 226, 3, 28, 14, 0, 226, 228, 1, 0, 0, 0, 227, 198,
		1, 0, 0, 0, 227, 200, 1, 0, 0, 0, 227, 204, 1, 0, 0, 0, 227, 212, 1, 0,
		0, 0, 228, 21, 1, 0, 0, 0, 229, 239, 5, 14, 0, 0, 230, 231, 4, 11, 0, 0,
		231, 236, 3, 46, 23, 0, 232, 233, 5, 45, 0, 0, 233, 235, 3, 46, 23, 0,
		234, 232, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236,
		237, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 230,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 23, 1, 0, 0, 0, 241, 242, 5, 22,
		0, 0, 242, 243, 5, 39, 0, 0, 243, 244, 3, 46, 23, 0, 244, 245, 5, 40, 0,
		0, 245, 25, 1, 0, 0, 0, 246, 247, 5, 21, 0, 0, 247, 248, 5, 52, 0, 0, 248,
		249, 3, 28, 14, 0, 249, 27, 1, 0, 0, 0, 250, 254, 5, 41, 0, 0, 251, 253,
		3, 6, 3, 0, 252, 251, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0,
		0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0,
		257, 258, 5, 42, 0, 0, 258, 29, 1, 0, 0, 0, 259, 264, 3, 32, 16, 0, 260,
		261, 5, 45, 0, 0, 261, 263, 3, 32, 16, 0, 262, 260, 1, 0, 0, 0, 263, 266,
		1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267, 1, 0,
		0, 0, 266, 264, 1, 0, 0, 0, 267, 268, 7, 0, 0, 0, 268, 273, 3, 46, 23,
		0, 269, 270, 5, 45, 0, 0, 270, 272, 3, 46, 23, 0, 271, 269, 1, 0, 0, 0,
		272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274,
		31, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 281, 5, 52, 0, 0, 277, 278,
		5, 43, 0, 0, 278, 279, 3, 46, 23, 0, 279, 280, 5, 44, 0, 0, 280, 282, 1,
		0, 0, 0, 281, 277, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 33, 1, 0, 0,
		0, 283, 284, 5, 19, 0, 0, 284, 285, 3, 30, 15, 0, 285, 35, 1, 0, 0, 0,
		286, 288, 5, 20, 0, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288,
		291, 1, 0, 0, 0, 289, 290, 5, 52, 0, 0, 290, 292, 5, 50, 0, 0, 291, 289,
		1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 294, 3, 50,
		25, 0, 294, 37, 1, 0, 0, 0, 295, 297, 3, 40, 20, 0, 296, 298, 3, 28, 14,
		0, 297, 296, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 39, 1, 0, 0, 0, 299,
		301, 5, 20, 0, 0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302,
		1, 0, 0, 0, 302, 303, 5, 24, 0, 0, 303, 315, 5, 39, 0, 0, 304, 309, 3,
		42, 21, 0, 305, 306, 5, 45, 0, 0, 306, 308, 3, 42, 21, 0, 307, 305, 1,
		0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0,
		0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 314, 5, 45, 0, 0, 313,
		312, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 304,
		1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 318, 5, 40,
		0, 0, 318, 41, 1, 0, 0, 0, 319, 320, 5, 52, 0, 0, 320, 321, 5, 47, 0, 0,
		321, 330, 3, 44, 22, 0, 322, 323, 5, 51, 0, 0, 323, 324, 5, 47, 0, 0, 324,
		330, 3, 44, 22, 0, 325, 327, 3, 46, 23, 0, 326, 328, 5, 49, 0, 0, 327,
		326, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 330, 1, 0, 0, 0, 329, 319,
		1, 0, 0, 0, 329, 322, 1, 0, 0, 0, 329, 325, 1, 0, 0, 0, 330, 43, 1, 0,
		0, 0, 331, 334, 3, 66, 33, 0, 332, 334, 3, 46, 23, 0, 333, 331, 1, 0, 0,
		0, 333, 332, 1, 0, 0, 0, 334, 45, 1, 0, 0, 0, 335, 336, 6, 23, -1, 0, 336,
		340, 3, 48, 24, 0, 337, 338, 7, 1, 0, 0, 338, 340, 3, 46, 23, 7, 339, 335,
		1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 340, 379, 1, 0, 0, 0, 341, 342, 10, 6,
		0, 0, 342, 343, 7, 2, 0, 0, 343, 378, 3, 46, 23, 7, 344, 345, 10, 5, 0,
		0, 345, 346, 7, 3, 0, 0, 346, 378, 3, 46, 23, 6, 347, 348, 10, 4, 0, 0,
		348, 349, 7, 4, 0, 0, 349, 378, 3, 46, 23, 5, 350, 351, 10, 3, 0, 0, 351,
		352, 5, 25, 0, 0, 352, 378, 3, 46, 23, 4, 353, 354, 10, 2, 0, 0, 354, 355,
		5, 26, 0, 0, 355, 378, 3, 46, 23, 3, 356, 357, 10, 1, 0, 0, 357, 358, 5,
		46, 0, 0, 358, 359, 3, 46, 23, 0, 359, 360, 5, 47, 0, 0, 360, 361, 3, 46,
		23, 1, 361, 378, 1, 0, 0, 0, 362, 363, 10, 9, 0, 0, 363, 365, 5, 43, 0,
		0, 364, 366, 3, 46, 23, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0,
		366, 367, 1, 0, 0, 0, 367, 369, 5, 47, 0, 0, 368, 370, 3, 46, 23, 0, 369,
		368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 378,
		5, 44, 0, 0, 372, 373, 10, 8, 0, 0, 373, 374, 5, 43, 0, 0, 374, 375, 3,
		46, 23, 0, 375, 376, 5, 44, 0, 0, 376, 378, 1, 0, 0, 0, 377, 341, 1, 0,
		0, 0, 377, 344, 1, 0, 0, 0, 377, 347, 1, 0, 0, 0, 377, 350, 1, 0, 0, 0,
		377, 353, 1, 0, 0, 0, 377, 356, 1, 0, 0, 0, 377, 362, 1, 0, 0, 0, 377,
		372, 1, 0, 0, 0, 378, 381, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380,
		1, 0, 0, 0, 380, 47, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 382, 402, 3, 56,
		28, 0, 383, 402, 3, 58, 29, 0, 384, 402, 3, 60, 30, 0, 385, 402, 3, 54,
		27, 0, 386, 402, 3, 52, 26, 0, 387, 389, 5, 20, 0, 0, 388, 387, 1, 0, 0,
		0, 388, 389, 1, 0, 0, 0, 389, 392, 1, 0, 0, 0, 390, 391, 5, 52, 0, 0, 391,
		393, 5, 50, 0, 0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 394,
		1, 0, 0, 0, 394, 402, 3, 50, 25, 0, 395, 402, 3, 40, 20, 0, 396, 397, 5,
		39, 0, 0, 397, 398, 3, 46, 23, 0, 398, 399, 5, 40, 0, 0, 399, 402, 1, 0,
		0, 0, 400, 402, 3, 62, 31, 0, 401, 382, 1, 0, 0, 0, 401, 383, 1, 0, 0,
		0, 401, 384, 1, 0, 0, 0, 401, 385, 1, 0, 0, 0, 401, 386, 1, 0, 0, 0, 401,
		388, 1, 0, 0, 0, 401, 395, 1, 0, 0, 0, 401, 396, 1, 0, 0, 0, 401, 400,
		1, 0, 0, 0, 402, 49, 1, 0, 0, 0, 403, 404, 5, 52, 0, 0, 404, 416, 5, 39,
		0, 0, 405, 410, 3, 46, 23, 0, 406, 407, 5, 45, 0, 0, 407, 409, 3, 46, 23,
		0, 408, 406, 1, 0, 0, 0, 409, 412, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 410,
		411, 1, 0, 0, 0, 411, 414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 415,
		5, 45, 0, 0, 414, 413, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 417, 1, 0,
		0, 0, 416, 405, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0,
		418, 419, 5, 40, 0, 0, 419, 51, 1, 0, 0, 0, 420, 421, 5, 52, 0, 0, 421,
		423, 5, 50, 0, 0, 422, 420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424,
		1, 0, 0, 0, 424, 425, 5, 52, 0, 0, 425, 53, 1, 0, 0, 0, 426, 427, 5, 51,
		0, 0, 427, 55, 1, 0, 0, 0, 428, 429, 5, 59, 0, 0, 429, 57, 1, 0, 0, 0,
		430, 431, 5, 60, 0, 0, 431, 59, 1, 0, 0, 0, 432, 433, 7, 5, 0, 0, 433,
		61, 1, 0, 0, 0, 434, 443, 5, 43, 0, 0, 435, 440, 3, 46, 23, 0, 436, 437,
		5, 45, 0, 0, 437, 439, 3, 46, 23, 0, 438, 436, 1, 0, 0, 0, 439, 442, 1,
		0, 0, 0, 440, 438, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 444, 1, 0, 0,
		0, 442, 440, 1, 0, 0, 0, 443, 435, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444,
		445, 1, 0, 0, 0, 445, 450, 5, 44, 0, 0, 446, 447, 5, 43, 0, 0, 447, 448,
		5, 44, 0, 0, 448, 450, 3, 66, 33, 0, 449, 434, 1, 0, 0, 0, 449, 446, 1,
		0, 0, 0, 450, 63, 1, 0, 0, 0, 451, 456, 3, 66, 33, 0, 452, 453, 5, 43,
		0, 0, 453, 454, 5, 44, 0, 0, 454, 456, 3, 66, 33, 0, 455, 451, 1, 0, 0,
		0, 455, 452, 1, 0, 0, 0, 456, 65, 1, 0, 0, 0, 457, 458, 7, 6, 0, 0, 458,
		67, 1, 0, 0, 0, 59, 71, 76, 78, 86, 92, 94, 103, 107, 110, 114, 117, 121,
		125, 133, 143, 146, 150, 162, 165, 169, 181, 186, 194, 196, 216, 222, 227,
		236, 239, 254, 264, 273, 281, 287, 291, 297, 300, 309, 313, 315, 327, 329,
		333, 339, 365, 369, 377, 379, 388, 392, 401, 410, 414, 416, 422, 440, 443,
		449, 455,
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

// LushParserInit initializes any static state used to implement LushParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLushParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LushParserInit() {
	staticData := &LushParserStaticData
	staticData.once.Do(lushParserInit)
}

// NewLushParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLushParser(input antlr.TokenStream) *LushParser {
	LushParserInit()
	this := new(LushParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LushParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lush.g4"

	return this
}

// LushParser tokens.
const (
	LushParserEOF           = antlr.TokenEOF
	LushParserWHITESPACE    = 1
	LushParserLINE_COMMENT  = 2
	LushParserIMPORT        = 3
	LushParserAS            = 4
	LushParserTRUE          = 5
	LushParserFALSE         = 6
	LushParserIF            = 7
	LushParserELSE          = 8
	LushParserFOR           = 9
	LushParserSTRING_TYPE   = 10
	LushParserINT_TYPE      = 11
	LushParserBOOL_TYPE     = 12
	LushParserFUNC          = 13
	LushParserRETURN        = 14
	LushParserBREAK         = 15
	LushParserCONTINUE      = 16
	LushParserIN            = 17
	LushParserUNTIL         = 18
	LushParserGLOBAL        = 19
	LushParserTRY           = 20
	LushParserCATCH         = 21
	LushParserFAIL          = 22
	LushParserERROR_TYPE    = 23
	LushParserCMD           = 24
	LushParserLAND          = 25
	LushParserLOR           = 26
	LushParserNOT           = 27
	LushParserMINUS         = 28
	LushParserPLUS          = 29
	LushParserMUL           = 30
	LushParserDIV           = 31
	LushParserMOD           = 32
	LushParserLT            = 33
	LushParserLTE           = 34
	LushParserGT            = 35
	LushParserGTE           = 36
	LushParserEQ            = 37
	LushParserNEQ           = 38
	LushParserLPAREN        = 39
	LushParserRPAREN        = 40
	LushParserLCUR          = 41
	LushParserRCUR          = 42
	LushParserLSQ           = 43
	LushParserRSQ           = 44
	LushParserCOMMA         = 45
	LushParserQUESTION      = 46
	LushParserCOLON         = 47
	LushParserSEMICOLON     = 48
	LushParserELLIPSIS      = 49
	LushParserDOT           = 50
	LushParserENVVAR        = 51
	LushParserID            = 52
	LushParserASSIGN        = 53
	LushParserPLUS_ASSIGN   = 54
	LushParserMINUS_ASSIGN  = 55
	LushParserMUL_ASSIGN    = 56
	LushParserDIV_ASSIGN    = 57
	LushParserMOD_ASSIGN    = 58
	LushParserSTRING        = 59
	LushParserNUMBER        = 60
	LushParserEXTERNAL_CODE = 61
)

// LushParser rules.
const (
	LushParserRULE_program          = 0
	LushParserRULE_package          = 1
	LushParserRULE_importStatement  = 2
	LushParserRULE_statement        = 3
	LushParserRULE_funcDef          = 4
	LushParserRULE_externalFuncDef  = 5
	LushParserRULE_param            = 6
	LushParserRULE_results          = 7
	LushParserRULE_result           = 8
	LushParserRULE_if               = 9
	LushParserRULE_for              = 10
	LushParserRULE_returnStatement  = 11
	LushParserRULE_failStatement    = 12
	LushParserRULE_catchClause      = 13
	LushParserRULE_block            = 14
	LushParserRULE_assignment       = 15
	LushParserRULE_assignTarget     = 16
	LushParserRULE_globalAssignment = 17
	LushParserRULE_funcStatement    = 18
	LushParserRULE_cmdStatement     = 19
	LushParserRULE_cmd              = 20
	LushParserRULE_cmdArg           = 21
	LushParserRULE_cmdValue         = 22
	LushParserRULE_expression       = 23
	LushParserRULE_atom             = 24
	LushParserRULE_func             = 25
	LushParserRULE_var              = 26
	LushParserRULE_envVar           = 27
	LushParserRULE_string           = 28
	LushParserRULE_number           = 29
	LushParserRULE_bool             = 30
	LushParserRULE_array            = 31
	LushParserRULE_type             = 32
	LushParserRULE_primitiveType    = 33
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllFuncDef() []IFuncDefContext
	FuncDef(i int) IFuncDefContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(LushParserEOF, 0)
}

func (s *ProgramContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) AllFuncDef() []IFuncDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDefContext); ok {
			len++
		}
	}

	tst := make([]IFuncDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDefContext); ok {
			tst[i] = t.(IFuncDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncDef(i int) IFuncDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LushParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LushParserIMPORT {
		{
			p.SetState(68)
			p.ImportStatement()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503599650038400) != 0 {
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LushParserIF, LushParserFOR, LushParserRETURN, LushParserBREAK, LushParserCONTINUE, LushParserGLOBAL, LushParserTRY, LushParserFAIL, LushParserCMD, LushParserID:
			{
				p.SetState(74)
				p.Statement()
			}

		case LushParserFUNC:
			{
				p.SetState(75)
				p.FuncDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(LushParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageContext is an interface to support dynamic dispatch.
type IPackageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
	AllFuncDef() []IFuncDefContext
	FuncDef(i int) IFuncDefContext
	AllExternalFuncDef() []IExternalFuncDefContext
	ExternalFuncDef(i int) IExternalFuncDefContext
	AllGlobalAssignment() []IGlobalAssignmentContext
	GlobalAssignment(i int) IGlobalAssignmentContext

	// IsPackageContext differentiates from other interfaces.
	IsPackageContext()
}

type PackageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageContext() *PackageContext {
	var p = new(PackageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_package
	return p
}

func InitEmptyPackageContext(p *PackageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_package
}

func (*PackageContext) IsPackageContext() {}

func NewPackageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageContext {
	var p = new(PackageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_package

	return p
}

func (s *PackageContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageContext) EOF() antlr.TerminalNode {
	return s.GetToken(LushParserEOF, 0)
}

func (s *PackageContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *PackageContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *PackageContext) AllFuncDef() []IFuncDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDefContext); ok {
			len++
		}
	}

	tst := make([]IFuncDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDefContext); ok {
			tst[i] = t.(IFuncDefContext)
			i++
		}
	}

	return tst
}

func (s *PackageContext) FuncDef(i int) IFuncDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *PackageContext) AllExternalFuncDef() []IExternalFuncDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExternalFuncDefContext); ok {
			len++
		}
	}

	tst := make([]IExternalFuncDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExternalFuncDefContext); ok {
			tst[i] = t.(IExternalFuncDefContext)
			i++
		}
	}

	return tst
}

func (s *PackageContext) ExternalFuncDef(i int) IExternalFuncDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternalFuncDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternalFuncDefContext)
}

func (s *PackageContext) AllGlobalAssignment() []IGlobalAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobalAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IGlobalAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobalAssignmentContext); ok {
			tst[i] = t.(IGlobalAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *PackageContext) GlobalAssignment(i int) IGlobalAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalAssignmentContext)
}

func (s *PackageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitPackage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Package_() (localctx IPackageContext) {
	localctx = NewPackageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LushParserRULE_package)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LushParserIMPORT {
		{
			p.SetState(83)
			p.ImportStatement()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LushParserFUNC || _la == LushParserGLOBAL {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(89)
				p.FuncDef()
			}

		case 2:
			{
				p.SetState(90)
				p.ExternalFuncDef()
			}

		case 3:
			{
				p.SetState(91)
				p.GlobalAssignment()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Match(LushParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// Getter signatures
	IMPORT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	AS() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) GetAlias() antlr.Token { return s.alias }

func (s *ImportStatementContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(LushParserIMPORT, 0)
}

func (s *ImportStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(LushParserSTRING, 0)
}

func (s *ImportStatementContext) AS() antlr.TerminalNode {
	return s.GetToken(LushParserAS, 0)
}

func (s *ImportStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LushParserRULE_importStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(LushParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(LushParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserAS {
		{
			p.SetState(101)
			p.Match(LushParserAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)

			var _m = p.Match(LushParserID)

			localctx.(*ImportStatementContext).alias = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCmdBody returns the cmdBody rule contexts.
	GetCmdBody() IBlockContext

	// SetCmdBody sets the cmdBody rule contexts.
	SetCmdBody(IBlockContext)

	// Getter signatures
	Assignment() IAssignmentContext
	CatchClause() ICatchClauseContext
	Block() IBlockContext
	GlobalAssignment() IGlobalAssignmentContext
	FuncStatement() IFuncStatementContext
	CmdStatement() ICmdStatementContext
	If_() IIfContext
	For_() IForContext
	ReturnStatement() IReturnStatementContext
	FailStatement() IFailStatementContext
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	cmdBody IBlockContext
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GetCmdBody() IBlockContext { return s.cmdBody }

func (s *StatementContext) SetCmdBody(v IBlockContext) { s.cmdBody = v }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) CatchClause() ICatchClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatchClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatchClauseContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) GlobalAssignment() IGlobalAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalAssignmentContext)
}

func (s *StatementContext) FuncStatement() IFuncStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncStatementContext)
}

func (s *StatementContext) CmdStatement() ICmdStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdStatementContext)
}

func (s *StatementContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *StatementContext) For_() IForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) FailStatement() IFailStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFailStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFailStatementContext)
}

func (s *StatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(LushParserBREAK, 0)
}

func (s *StatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(LushParserCONTINUE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LushParserRULE_statement)
	var _la int

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Assignment()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserLCUR {
			{
				p.SetState(106)

				var _x = p.Block()

				localctx.(*StatementContext).cmdBody = _x
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCATCH {
			{
				p.SetState(109)
				p.CatchClause()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.GlobalAssignment()
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserLCUR {
			{
				p.SetState(113)

				var _x = p.Block()

				localctx.(*StatementContext).cmdBody = _x
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCATCH {
			{
				p.SetState(116)
				p.CatchClause()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.FuncStatement()
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCATCH {
			{
				p.SetState(120)
				p.CatchClause()
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.CmdStatement()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCATCH {
			{
				p.SetState(124)
				p.CatchClause()
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.If_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.For_()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(129)
			p.ReturnStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(130)
			p.FailStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(131)
			p.Match(LushParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(132)
			p.Match(LushParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	AllParam() []IParamContext
	Param(i int) IParamContext
	Results() IResultsContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_funcDef
	return p
}

func InitEmptyFuncDefContext(p *FuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_funcDef
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(LushParserFUNC, 0)
}

func (s *FuncDefContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *FuncDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserLPAREN, 0)
}

func (s *FuncDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserRPAREN, 0)
}

func (s *FuncDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDefContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *FuncDefContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *FuncDefContext) Results() IResultsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultsContext)
}

func (s *FuncDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *FuncDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitFuncDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LushParserRULE_funcDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(LushParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(LushParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserID {
		{
			p.SetState(138)
			p.Param()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LushParserCOMMA {
			{
				p.SetState(139)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(140)
				p.Param()
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(148)
		p.Match(LushParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796101417984) != 0 {
		{
			p.SetState(149)
			p.Results()
		}

	}
	{
		p.SetState(152)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternalFuncDefContext is an interface to support dynamic dispatch.
type IExternalFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	EXTERNAL_CODE() antlr.TerminalNode
	AllParam() []IParamContext
	Param(i int) IParamContext
	Results() IResultsContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExternalFuncDefContext differentiates from other interfaces.
	IsExternalFuncDefContext()
}

type ExternalFuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternalFuncDefContext() *ExternalFuncDefContext {
	var p = new(ExternalFuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_externalFuncDef
	return p
}

func InitEmptyExternalFuncDefContext(p *ExternalFuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_externalFuncDef
}

func (*ExternalFuncDefContext) IsExternalFuncDefContext() {}

func NewExternalFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternalFuncDefContext {
	var p = new(ExternalFuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_externalFuncDef

	return p
}

func (s *ExternalFuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternalFuncDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(LushParserFUNC, 0)
}

func (s *ExternalFuncDefContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *ExternalFuncDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserLPAREN, 0)
}

func (s *ExternalFuncDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserRPAREN, 0)
}

func (s *ExternalFuncDefContext) EXTERNAL_CODE() antlr.TerminalNode {
	return s.GetToken(LushParserEXTERNAL_CODE, 0)
}

func (s *ExternalFuncDefContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ExternalFuncDefContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ExternalFuncDefContext) Results() IResultsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultsContext)
}

func (s *ExternalFuncDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *ExternalFuncDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *ExternalFuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternalFuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternalFuncDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitExternalFuncDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) ExternalFuncDef() (localctx IExternalFuncDefContext) {
	localctx = NewExternalFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LushParserRULE_externalFuncDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(LushParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(LushParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserID {
		{
			p.SetState(157)
			p.Param()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LushParserCOMMA {
			{
				p.SetState(158)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(159)
				p.Param()
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(167)
		p.Match(LushParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796101417984) != 0 {
		{
			p.SetState(168)
			p.Results()
		}

	}
	{
		p.SetState(171)
		p.Match(LushParserEXTERNAL_CODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_() ITypeContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LushParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResultsContext is an interface to support dynamic dispatch.
type IResultsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllResult() []IResultContext
	Result(i int) IResultContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsResultsContext differentiates from other interfaces.
	IsResultsContext()
}

type ResultsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultsContext() *ResultsContext {
	var p = new(ResultsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_results
	return p
}

func InitEmptyResultsContext(p *ResultsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_results
}

func (*ResultsContext) IsResultsContext() {}

func NewResultsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultsContext {
	var p = new(ResultsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_results

	return p
}

func (s *ResultsContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultsContext) AllResult() []IResultContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResultContext); ok {
			len++
		}
	}

	tst := make([]IResultContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResultContext); ok {
			tst[i] = t.(IResultContext)
			i++
		}
	}

	return tst
}

func (s *ResultsContext) Result(i int) IResultContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultContext)
}

func (s *ResultsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *ResultsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *ResultsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitResults(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Results() (localctx IResultsContext) {
	localctx = NewResultsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LushParserRULE_results)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Result()
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LushParserCOMMA {
		{
			p.SetState(177)
			p.Match(LushParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Result()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResultContext is an interface to support dynamic dispatch.
type IResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ERROR_TYPE() antlr.TerminalNode

	// IsResultContext differentiates from other interfaces.
	IsResultContext()
}

type ResultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultContext() *ResultContext {
	var p = new(ResultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_result
	return p
}

func InitEmptyResultContext(p *ResultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_result
}

func (*ResultContext) IsResultContext() {}

func NewResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultContext {
	var p = new(ResultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_result

	return p
}

func (s *ResultContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ResultContext) ERROR_TYPE() antlr.TerminalNode {
	return s.GetToken(LushParserERROR_TYPE, 0)
}

func (s *ResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitResult(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Result() (localctx IResultContext) {
	localctx = NewResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LushParserRULE_result)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LushParserSTRING_TYPE, LushParserINT_TYPE, LushParserBOOL_TYPE, LushParserLSQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Type_()
		}

	case LushParserERROR_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(LushParserERROR_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThen returns the then rule contexts.
	GetThen() IBlockContext

	// GetElseIf returns the elseIf rule contexts.
	GetElseIf() IIfContext

	// GetElseBlock returns the elseBlock rule contexts.
	GetElseBlock() IBlockContext

	// SetThen sets the then rule contexts.
	SetThen(IBlockContext)

	// SetElseIf sets the elseIf rule contexts.
	SetElseIf(IIfContext)

	// SetElseBlock sets the elseBlock rule contexts.
	SetElseBlock(IBlockContext)

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	If_() IIfContext

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	then      IBlockContext
	elseIf    IIfContext
	elseBlock IBlockContext
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) GetThen() IBlockContext { return s.then }

func (s *IfContext) GetElseIf() IIfContext { return s.elseIf }

func (s *IfContext) GetElseBlock() IBlockContext { return s.elseBlock }

func (s *IfContext) SetThen(v IBlockContext) { s.then = v }

func (s *IfContext) SetElseIf(v IIfContext) { s.elseIf = v }

func (s *IfContext) SetElseBlock(v IBlockContext) { s.elseBlock = v }

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(LushParserIF, 0)
}

func (s *IfContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(LushParserELSE, 0)
}

func (s *IfContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LushParserRULE_if)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(LushParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.expression(0)
	}
	{
		p.SetState(190)

		var _x = p.Block()

		localctx.(*IfContext).then = _x
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserELSE {
		{
			p.SetState(191)
			p.Match(LushParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LushParserIF:
			{
				p.SetState(192)

				var _x = p.If_()

				localctx.(*IfContext).elseIf = _x
			}

		case LushParserLCUR:
			{
				p.SetState(193)

				var _x = p.Block()

				localctx.(*IfContext).elseBlock = _x
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForContext is an interface to support dynamic dispatch.
type IForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// GetIndex returns the index token.
	GetIndex() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// SetIndex sets the index token.
	SetIndex(antlr.Token)

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetInit returns the init rule contexts.
	GetInit() IAssignmentContext

	// GetStep returns the step rule contexts.
	GetStep() IAssignmentContext

	// GetFrom returns the from rule contexts.
	GetFrom() IExpressionContext

	// GetOver returns the over rule contexts.
	GetOver() IExpressionContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetInit sets the init rule contexts.
	SetInit(IAssignmentContext)

	// SetStep sets the step rule contexts.
	SetStep(IAssignmentContext)

	// SetFrom sets the from rule contexts.
	SetFrom(IExpressionContext)

	// SetOver sets the over rule contexts.
	SetOver(IExpressionContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	Block() IBlockContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	IN() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	COMMA() antlr.TerminalNode
	UNTIL() antlr.TerminalNode

	// IsForContext differentiates from other interfaces.
	IsForContext()
}

type ForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExpressionContext
	init   IAssignmentContext
	step   IAssignmentContext
	value  antlr.Token
	index  antlr.Token
	from   IExpressionContext
	over   IExpressionContext
}

func NewEmptyForContext() *ForContext {
	var p = new(ForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_for
	return p
}

func InitEmptyForContext(p *ForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_for
}

func (*ForContext) IsForContext() {}

func NewForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForContext {
	var p = new(ForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_for

	return p
}

func (s *ForContext) GetParser() antlr.Parser { return s.parser }

func (s *ForContext) GetValue() antlr.Token { return s.value }

func (s *ForContext) GetIndex() antlr.Token { return s.index }

func (s *ForContext) SetValue(v antlr.Token) { s.value = v }

func (s *ForContext) SetIndex(v antlr.Token) { s.index = v }

func (s *ForContext) GetCond() IExpressionContext { return s.cond }

func (s *ForContext) GetInit() IAssignmentContext { return s.init }

func (s *ForContext) GetStep() IAssignmentContext { return s.step }

func (s *ForContext) GetFrom() IExpressionContext { return s.from }

func (s *ForContext) GetOver() IExpressionContext { return s.over }

func (s *ForContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *ForContext) SetInit(v IAssignmentContext) { s.init = v }

func (s *ForContext) SetStep(v IAssignmentContext) { s.step = v }

func (s *ForContext) SetFrom(v IExpressionContext) { s.from = v }

func (s *ForContext) SetOver(v IExpressionContext) { s.over = v }

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(LushParserFOR, 0)
}

func (s *ForContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ForContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(LushParserSEMICOLON)
}

func (s *ForContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(LushParserSEMICOLON, i)
}

func (s *ForContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ForContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForContext) IN() antlr.TerminalNode {
	return s.GetToken(LushParserIN, 0)
}

func (s *ForContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LushParserID)
}

func (s *ForContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LushParserID, i)
}

func (s *ForContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, 0)
}

func (s *ForContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(LushParserUNTIL, 0)
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) For_() (localctx IForContext) {
	localctx = NewForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LushParserRULE_for)
	var _la int

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Match(LushParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.Match(LushParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)

			var _x = p.expression(0)

			localctx.(*ForContext).cond = _x
		}
		{
			p.SetState(202)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(204)
			p.Match(LushParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)

			var _x = p.Assignment()

			localctx.(*ForContext).init = _x
		}
		{
			p.SetState(206)
			p.Match(LushParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)

			var _x = p.expression(0)

			localctx.(*ForContext).cond = _x
		}
		{
			p.SetState(208)
			p.Match(LushParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)

			var _x = p.Assignment()

			localctx.(*ForContext).step = _x
		}
		{
			p.SetState(210)
			p.Block()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(212)
			p.Match(LushParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)

			var _m = p.Match(LushParserID)

			localctx.(*ForContext).value = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCOMMA {
			{
				p.SetState(214)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(215)

				var _m = p.Match(LushParserID)

				localctx.(*ForContext).index = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(218)
			p.Match(LushParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(219)

				var _x = p.expression(0)

				localctx.(*ForContext).from = _x
			}
			{
				p.SetState(220)
				p.Match(LushParserUNTIL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(224)

			var _x = p.expression(0)

			localctx.(*ForContext).over = _x
		}
		{
			p.SetState(225)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(LushParserRETURN, 0)
}

func (s *ReturnStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ReturnStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *ReturnStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LushParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(LushParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		p.SetState(230)

		if !(p.GetTokenStream().LT(-1).GetLine() == p.GetTokenStream().LT(1).GetLine()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(-1).GetLine() == p.GetTokenStream().LT(1).GetLine()", ""))
			goto errorExit
		}
		{
			p.SetState(231)
			p.expression(0)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LushParserCOMMA {
			{
				p.SetState(232)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(233)
				p.expression(0)
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFailStatementContext is an interface to support dynamic dispatch.
type IFailStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FAIL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsFailStatementContext differentiates from other interfaces.
	IsFailStatementContext()
}

type FailStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFailStatementContext() *FailStatementContext {
	var p = new(FailStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_failStatement
	return p
}

func InitEmptyFailStatementContext(p *FailStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_failStatement
}

func (*FailStatementContext) IsFailStatementContext() {}

func NewFailStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FailStatementContext {
	var p = new(FailStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_failStatement

	return p
}

func (s *FailStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FailStatementContext) FAIL() antlr.TerminalNode {
	return s.GetToken(LushParserFAIL, 0)
}

func (s *FailStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserLPAREN, 0)
}

func (s *FailStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FailStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserRPAREN, 0)
}

func (s *FailStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FailStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FailStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitFailStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) FailStatement() (localctx IFailStatementContext) {
	localctx = NewFailStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LushParserRULE_failStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(LushParserFAIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(LushParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expression(0)
	}
	{
		p.SetState(244)
		p.Match(LushParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatchClauseContext is an interface to support dynamic dispatch.
type ICatchClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	CATCH() antlr.TerminalNode
	Block() IBlockContext
	ID() antlr.TerminalNode

	// IsCatchClauseContext differentiates from other interfaces.
	IsCatchClauseContext()
}

type CatchClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyCatchClauseContext() *CatchClauseContext {
	var p = new(CatchClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_catchClause
	return p
}

func InitEmptyCatchClauseContext(p *CatchClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_catchClause
}

func (*CatchClauseContext) IsCatchClauseContext() {}

func NewCatchClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchClauseContext {
	var p = new(CatchClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_catchClause

	return p
}

func (s *CatchClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchClauseContext) GetName() antlr.Token { return s.name }

func (s *CatchClauseContext) SetName(v antlr.Token) { s.name = v }

func (s *CatchClauseContext) CATCH() antlr.TerminalNode {
	return s.GetToken(LushParserCATCH, 0)
}

func (s *CatchClauseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CatchClauseContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *CatchClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitCatchClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) CatchClause() (localctx ICatchClauseContext) {
	localctx = NewCatchClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LushParserRULE_catchClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(LushParserCATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)

		var _m = p.Match(LushParserID)

		localctx.(*CatchClauseContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCUR() antlr.TerminalNode
	RCUR() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LCUR() antlr.TerminalNode {
	return s.GetToken(LushParserLCUR, 0)
}

func (s *BlockContext) RCUR() antlr.TerminalNode {
	return s.GetToken(LushParserRCUR, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LushParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(LushParserLCUR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503599650030208) != 0 {
		{
			p.SetState(251)
			p.Statement()
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
		p.Match(LushParserRCUR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllAssignTarget() []IAssignTargetContext
	AssignTarget(i int) IAssignTargetContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	ASSIGN() antlr.TerminalNode
	PLUS_ASSIGN() antlr.TerminalNode
	MINUS_ASSIGN() antlr.TerminalNode
	MUL_ASSIGN() antlr.TerminalNode
	DIV_ASSIGN() antlr.TerminalNode
	MOD_ASSIGN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) GetOp() antlr.Token { return s.op }

func (s *AssignmentContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignmentContext) AllAssignTarget() []IAssignTargetContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignTargetContext); ok {
			len++
		}
	}

	tst := make([]IAssignTargetContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignTargetContext); ok {
			tst[i] = t.(IAssignTargetContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) AssignTarget(i int) IAssignTargetContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignTargetContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignTargetContext)
}

func (s *AssignmentContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(LushParserASSIGN, 0)
}

func (s *AssignmentContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(LushParserPLUS_ASSIGN, 0)
}

func (s *AssignmentContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(LushParserMINUS_ASSIGN, 0)
}

func (s *AssignmentContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(LushParserMUL_ASSIGN, 0)
}

func (s *AssignmentContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(LushParserDIV_ASSIGN, 0)
}

func (s *AssignmentContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(LushParserMOD_ASSIGN, 0)
}

func (s *AssignmentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *AssignmentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LushParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.AssignTarget()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LushParserCOMMA {
		{
			p.SetState(260)
			p.Match(LushParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.AssignTarget()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(267)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&567453553048682496) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(268)
		p.expression(0)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LushParserCOMMA {
		{
			p.SetState(269)
			p.Match(LushParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.expression(0)
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignTargetContext is an interface to support dynamic dispatch.
type IAssignTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPosition returns the position rule contexts.
	GetPosition() IExpressionContext

	// SetPosition sets the position rule contexts.
	SetPosition(IExpressionContext)

	// Getter signatures
	ID() antlr.TerminalNode
	LSQ() antlr.TerminalNode
	RSQ() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignTargetContext differentiates from other interfaces.
	IsAssignTargetContext()
}

type AssignTargetContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	position IExpressionContext
}

func NewEmptyAssignTargetContext() *AssignTargetContext {
	var p = new(AssignTargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_assignTarget
	return p
}

func InitEmptyAssignTargetContext(p *AssignTargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_assignTarget
}

func (*AssignTargetContext) IsAssignTargetContext() {}

func NewAssignTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignTargetContext {
	var p = new(AssignTargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_assignTarget

	return p
}

func (s *AssignTargetContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignTargetContext) GetPosition() IExpressionContext { return s.position }

func (s *AssignTargetContext) SetPosition(v IExpressionContext) { s.position = v }

func (s *AssignTargetContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *AssignTargetContext) LSQ() antlr.TerminalNode {
	return s.GetToken(LushParserLSQ, 0)
}

func (s *AssignTargetContext) RSQ() antlr.TerminalNode {
	return s.GetToken(LushParserRSQ, 0)
}

func (s *AssignTargetContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignTargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignTargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignTargetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitAssignTarget(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) AssignTarget() (localctx IAssignTargetContext) {
	localctx = NewAssignTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LushParserRULE_assignTarget)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserLSQ {
		{
			p.SetState(277)
			p.Match(LushParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)

			var _x = p.expression(0)

			localctx.(*AssignTargetContext).position = _x
		}
		{
			p.SetState(279)
			p.Match(LushParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGlobalAssignmentContext is an interface to support dynamic dispatch.
type IGlobalAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GLOBAL() antlr.TerminalNode
	Assignment() IAssignmentContext

	// IsGlobalAssignmentContext differentiates from other interfaces.
	IsGlobalAssignmentContext()
}

type GlobalAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalAssignmentContext() *GlobalAssignmentContext {
	var p = new(GlobalAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_globalAssignment
	return p
}

func InitEmptyGlobalAssignmentContext(p *GlobalAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_globalAssignment
}

func (*GlobalAssignmentContext) IsGlobalAssignmentContext() {}

func NewGlobalAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalAssignmentContext {
	var p = new(GlobalAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_globalAssignment

	return p
}

func (s *GlobalAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalAssignmentContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(LushParserGLOBAL, 0)
}

func (s *GlobalAssignmentContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *GlobalAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitGlobalAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) GlobalAssignment() (localctx IGlobalAssignmentContext) {
	localctx = NewGlobalAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LushParserRULE_globalAssignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(LushParserGLOBAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Assignment()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncStatementContext is an interface to support dynamic dispatch.
type IFuncStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNamespace returns the namespace token.
	GetNamespace() antlr.Token

	// SetNamespace sets the namespace token.
	SetNamespace(antlr.Token)

	// Getter signatures
	Func_() IFuncContext
	TRY() antlr.TerminalNode
	DOT() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsFuncStatementContext differentiates from other interfaces.
	IsFuncStatementContext()
}

type FuncStatementContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	namespace antlr.Token
}

func NewEmptyFuncStatementContext() *FuncStatementContext {
	var p = new(FuncStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_funcStatement
	return p
}

func InitEmptyFuncStatementContext(p *FuncStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_funcStatement
}

func (*FuncStatementContext) IsFuncStatementContext() {}

func NewFuncStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncStatementContext {
	var p = new(FuncStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_funcStatement

	return p
}

func (s *FuncStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncStatementContext) GetNamespace() antlr.Token { return s.namespace }

func (s *FuncStatementContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *FuncStatementContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *FuncStatementContext) TRY() antlr.TerminalNode {
	return s.GetToken(LushParserTRY, 0)
}

func (s *FuncStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(LushParserDOT, 0)
}

func (s *FuncStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *FuncStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitFuncStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) FuncStatement() (localctx IFuncStatementContext) {
	localctx = NewFuncStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LushParserRULE_funcStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserTRY {
		{
			p.SetState(286)
			p.Match(LushParserTRY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(289)

			var _m = p.Match(LushParserID)

			localctx.(*FuncStatementContext).namespace = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(LushParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(293)
		p.Func_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdStatementContext is an interface to support dynamic dispatch.
type ICmdStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Cmd() ICmdContext
	Block() IBlockContext

	// IsCmdStatementContext differentiates from other interfaces.
	IsCmdStatementContext()
}

type CmdStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdStatementContext() *CmdStatementContext {
	var p = new(CmdStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmdStatement
	return p
}

func InitEmptyCmdStatementContext(p *CmdStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmdStatement
}

func (*CmdStatementContext) IsCmdStatementContext() {}

func NewCmdStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdStatementContext {
	var p = new(CmdStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_cmdStatement

	return p
}

func (s *CmdStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdStatementContext) Cmd() ICmdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CmdStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitCmdStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) CmdStatement() (localctx ICmdStatementContext) {
	localctx = NewCmdStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LushParserRULE_cmdStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Cmd()
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserLCUR {
		{
			p.SetState(296)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdContext is an interface to support dynamic dispatch.
type ICmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CMD() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	TRY() antlr.TerminalNode
	AllCmdArg() []ICmdArgContext
	CmdArg(i int) ICmdArgContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCmdContext differentiates from other interfaces.
	IsCmdContext()
}

type CmdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdContext() *CmdContext {
	var p = new(CmdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmd
	return p
}

func InitEmptyCmdContext(p *CmdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmd
}

func (*CmdContext) IsCmdContext() {}

func NewCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdContext {
	var p = new(CmdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_cmd

	return p
}

func (s *CmdContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdContext) CMD() antlr.TerminalNode {
	return s.GetToken(LushParserCMD, 0)
}

func (s *CmdContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserLPAREN, 0)
}

func (s *CmdContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserRPAREN, 0)
}

func (s *CmdContext) TRY() antlr.TerminalNode {
	return s.GetToken(LushParserTRY, 0)
}

func (s *CmdContext) AllCmdArg() []ICmdArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdArgContext); ok {
			len++
		}
	}

	tst := make([]ICmdArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdArgContext); ok {
			tst[i] = t.(ICmdArgContext)
			i++
		}
	}

	return tst
}

func (s *CmdContext) CmdArg(i int) ICmdArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdArgContext)
}

func (s *CmdContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *CmdContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *CmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitCmd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Cmd() (localctx ICmdContext) {
	localctx = NewCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LushParserRULE_cmd)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserTRY {
		{
			p.SetState(299)
			p.Match(LushParserTRY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(302)
		p.Match(LushParserCMD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(LushParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1736147002620641376) != 0 {
		{
			p.SetState(304)
			p.CmdArg()
		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(305)
					p.Match(LushParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(306)
					p.CmdArg()
				}

			}
			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCOMMA {
			{
				p.SetState(312)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(317)
		p.Match(LushParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdArgContext is an interface to support dynamic dispatch.
type ICmdArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetEnv returns the env token.
	GetEnv() antlr.Token

	// GetSpread returns the spread token.
	GetSpread() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetEnv sets the env token.
	SetEnv(antlr.Token)

	// SetSpread sets the spread token.
	SetSpread(antlr.Token)

	// Getter signatures
	COLON() antlr.TerminalNode
	CmdValue() ICmdValueContext
	ID() antlr.TerminalNode
	ENVVAR() antlr.TerminalNode
	Expression() IExpressionContext
	ELLIPSIS() antlr.TerminalNode

	// IsCmdArgContext differentiates from other interfaces.
	IsCmdArgContext()
}

type CmdArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	env    antlr.Token
	spread antlr.Token
}

func NewEmptyCmdArgContext() *CmdArgContext {
	var p = new(CmdArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmdArg
	return p
}

func InitEmptyCmdArgContext(p *CmdArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmdArg
}

func (*CmdArgContext) IsCmdArgContext() {}

func NewCmdArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdArgContext {
	var p = new(CmdArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_cmdArg

	return p
}

func (s *CmdArgContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdArgContext) GetName() antlr.Token { return s.name }

func (s *CmdArgContext) GetEnv() antlr.Token { return s.env }

func (s *CmdArgContext) GetSpread() antlr.Token { return s.spread }

func (s *CmdArgContext) SetName(v antlr.Token) { s.name = v }

func (s *CmdArgContext) SetEnv(v antlr.Token) { s.env = v }

func (s *CmdArgContext) SetSpread(v antlr.Token) { s.spread = v }

func (s *CmdArgContext) COLON() antlr.TerminalNode {
	return s.GetToken(LushParserCOLON, 0)
}

func (s *CmdArgContext) CmdValue() ICmdValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdValueContext)
}

func (s *CmdArgContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *CmdArgContext) ENVVAR() antlr.TerminalNode {
	return s.GetToken(LushParserENVVAR, 0)
}

func (s *CmdArgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CmdArgContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(LushParserELLIPSIS, 0)
}

func (s *CmdArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitCmdArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) CmdArg() (localctx ICmdArgContext) {
	localctx = NewCmdArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LushParserRULE_cmdArg)
	var _la int

	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)

			var _m = p.Match(LushParserID)

			localctx.(*CmdArgContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Match(LushParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.CmdValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)

			var _m = p.Match(LushParserENVVAR)

			localctx.(*CmdArgContext).env = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(LushParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.CmdValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(325)
			p.expression(0)
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserELLIPSIS {
			{
				p.SetState(326)

				var _m = p.Match(LushParserELLIPSIS)

				localctx.(*CmdArgContext).spread = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdValueContext is an interface to support dynamic dispatch.
type ICmdValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	Expression() IExpressionContext

	// IsCmdValueContext differentiates from other interfaces.
	IsCmdValueContext()
}

type CmdValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdValueContext() *CmdValueContext {
	var p = new(CmdValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmdValue
	return p
}

func InitEmptyCmdValueContext(p *CmdValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_cmdValue
}

func (*CmdValueContext) IsCmdValueContext() {}

func NewCmdValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdValueContext {
	var p = new(CmdValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_cmdValue

	return p
}

func (s *CmdValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdValueContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *CmdValueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CmdValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitCmdValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) CmdValue() (localctx ICmdValueContext) {
	localctx = NewCmdValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LushParserRULE_cmdValue)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LushParserSTRING_TYPE, LushParserINT_TYPE, LushParserBOOL_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.PrimitiveType()
		}

	case LushParserTRUE, LushParserFALSE, LushParserTRY, LushParserCMD, LushParserNOT, LushParserMINUS, LushParserLPAREN, LushParserLSQ, LushParserENVVAR, LushParserID, LushParserSTRING, LushParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.expression(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUnary_op returns the unary_op token.
	GetUnary_op() antlr.Token

	// GetMul_op returns the mul_op token.
	GetMul_op() antlr.Token

	// GetAdd_op returns the add_op token.
	GetAdd_op() antlr.Token

	// GetRel_op returns the rel_op token.
	GetRel_op() antlr.Token

	// SetUnary_op sets the unary_op token.
	SetUnary_op(antlr.Token)

	// SetMul_op sets the mul_op token.
	SetMul_op(antlr.Token)

	// SetAdd_op sets the add_op token.
	SetAdd_op(antlr.Token)

	// SetRel_op sets the rel_op token.
	SetRel_op(antlr.Token)

	// GetSlice returns the slice rule contexts.
	GetSlice() IExpressionContext

	// GetIndex returns the index rule contexts.
	GetIndex() IExpressionContext

	// GetTernary returns the ternary rule contexts.
	GetTernary() IExpressionContext

	// GetFrom returns the from rule contexts.
	GetFrom() IExpressionContext

	// GetTo returns the to rule contexts.
	GetTo() IExpressionContext

	// GetPosition returns the position rule contexts.
	GetPosition() IExpressionContext

	// SetSlice sets the slice rule contexts.
	SetSlice(IExpressionContext)

	// SetIndex sets the index rule contexts.
	SetIndex(IExpressionContext)

	// SetTernary sets the ternary rule contexts.
	SetTernary(IExpressionContext)

	// SetFrom sets the from rule contexts.
	SetFrom(IExpressionContext)

	// SetTo sets the to rule contexts.
	SetTo(IExpressionContext)

	// SetPosition sets the position rule contexts.
	SetPosition(IExpressionContext)

	// Getter signatures
	Atom() IAtomContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LAND() antlr.TerminalNode
	LOR() antlr.TerminalNode
	QUESTION() antlr.TerminalNode
	COLON() antlr.TerminalNode
	LSQ() antlr.TerminalNode
	RSQ() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	slice    IExpressionContext
	index    IExpressionContext
	ternary  IExpressionContext
	unary_op antlr.Token
	mul_op   antlr.Token
	add_op   antlr.Token
	rel_op   antlr.Token
	from     IExpressionContext
	to       IExpressionContext
	position IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetUnary_op() antlr.Token { return s.unary_op }

func (s *ExpressionContext) GetMul_op() antlr.Token { return s.mul_op }

func (s *ExpressionContext) GetAdd_op() antlr.Token { return s.add_op }

func (s *ExpressionContext) GetRel_op() antlr.Token { return s.rel_op }

func (s *ExpressionContext) SetUnary_op(v antlr.Token) { s.unary_op = v }

func (s *ExpressionContext) SetMul_op(v antlr.Token) { s.mul_op = v }

func (s *ExpressionContext) SetAdd_op(v antlr.Token) { s.add_op = v }

func (s *ExpressionContext) SetRel_op(v antlr.Token) { s.rel_op = v }

func (s *ExpressionContext) GetSlice() IExpressionContext { return s.slice }

func (s *ExpressionContext) GetIndex() IExpressionContext { return s.index }

func (s *ExpressionContext) GetTernary() IExpressionContext { return s.ternary }

func (s *ExpressionContext) GetFrom() IExpressionContext { return s.from }

func (s *ExpressionContext) GetTo() IExpressionContext { return s.to }

func (s *ExpressionContext) GetPosition() IExpressionContext { return s.position }

func (s *ExpressionContext) SetSlice(v IExpressionContext) { s.slice = v }

func (s *ExpressionContext) SetIndex(v IExpressionContext) { s.index = v }

func (s *ExpressionContext) SetTernary(v IExpressionContext) { s.ternary = v }

func (s *ExpressionContext) SetFrom(v IExpressionContext) { s.from = v }

func (s *ExpressionContext) SetTo(v IExpressionContext) { s.to = v }

func (s *ExpressionContext) SetPosition(v IExpressionContext) { s.position = v }

func (s *ExpressionContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LushParserMINUS, 0)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LushParserNOT, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(LushParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(LushParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(LushParserMOD, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LushParserPLUS, 0)
}

func (s *ExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(LushParserEQ, 0)
}

func (s *ExpressionContext) NEQ() antlr.TerminalNode {
	return s.GetToken(LushParserNEQ, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(LushParserLT, 0)
}

func (s *ExpressionContext) LTE() antlr.TerminalNode {
	return s.GetToken(LushParserLTE, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(LushParserGT, 0)
}

func (s *ExpressionContext) GTE() antlr.TerminalNode {
	return s.GetToken(LushParserGTE, 0)
}

func (s *ExpressionContext) LAND() antlr.TerminalNode {
	return s.GetToken(LushParserLAND, 0)
}

func (s *ExpressionContext) LOR() antlr.TerminalNode {
	return s.GetToken(LushParserLOR, 0)
}

func (s *ExpressionContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(LushParserQUESTION, 0)
}

func (s *ExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(LushParserCOLON, 0)
}

func (s *ExpressionContext) LSQ() antlr.TerminalNode {
	return s.GetToken(LushParserLSQ, 0)
}

func (s *ExpressionContext) RSQ() antlr.TerminalNode {
	return s.GetToken(LushParserRSQ, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *LushParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, LushParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LushParserTRUE, LushParserFALSE, LushParserTRY, LushParserCMD, LushParserLPAREN, LushParserLSQ, LushParserENVVAR, LushParserID, LushParserSTRING, LushParserNUMBER:
		{
			p.SetState(336)
			p.Atom()
		}

	case LushParserNOT, LushParserMINUS:
		{
			p.SetState(337)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).unary_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == LushParserNOT || _la == LushParserMINUS) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).unary_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(338)
			p.expression(7)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(377)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(341)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(342)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).mul_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7516192768) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).mul_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(343)
					p.expression(7)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(345)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).add_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LushParserMINUS || _la == LushParserPLUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).add_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(346)
					p.expression(6)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(348)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).rel_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541165879296) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).rel_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(349)
					p.expression(5)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(351)
					p.Match(LushParserLAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(352)
					p.expression(4)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(353)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(354)
					p.Match(LushParserLOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(355)
					p.expression(3)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).ternary = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(356)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(357)
					p.Match(LushParserQUESTION)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(358)
					p.expression(0)
				}
				{
					p.SetState(359)
					p.Match(LushParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(360)
					p.expression(1)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).slice = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(362)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(363)
					p.Match(LushParserLSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(365)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1736147002620641376) != 0 {
					{
						p.SetState(364)

						var _x = p.expression(0)

						localctx.(*ExpressionContext).from = _x
					}

				}
				{
					p.SetState(367)
					p.Match(LushParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(369)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1736147002620641376) != 0 {
					{
						p.SetState(368)

						var _x = p.expression(0)

						localctx.(*ExpressionContext).to = _x
					}

				}
				{
					p.SetState(371)
					p.Match(LushParserRSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).index = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(372)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(373)
					p.Match(LushParserLSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(374)

					var _x = p.expression(0)

					localctx.(*ExpressionContext).position = _x
				}
				{
					p.SetState(375)
					p.Match(LushParserRSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNamespace returns the namespace token.
	GetNamespace() antlr.Token

	// SetNamespace sets the namespace token.
	SetNamespace(antlr.Token)

	// GetGroup returns the group rule contexts.
	GetGroup() IExpressionContext

	// SetGroup sets the group rule contexts.
	SetGroup(IExpressionContext)

	// Getter signatures
	String_() IStringContext
	Number() INumberContext
	Bool_() IBoolContext
	EnvVar() IEnvVarContext
	Var_() IVarContext
	Func_() IFuncContext
	TRY() antlr.TerminalNode
	DOT() antlr.TerminalNode
	ID() antlr.TerminalNode
	Cmd() ICmdContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	Array() IArrayContext

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	namespace antlr.Token
	group     IExpressionContext
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) GetNamespace() antlr.Token { return s.namespace }

func (s *AtomContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *AtomContext) GetGroup() IExpressionContext { return s.group }

func (s *AtomContext) SetGroup(v IExpressionContext) { s.group = v }

func (s *AtomContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *AtomContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *AtomContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *AtomContext) EnvVar() IEnvVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvVarContext)
}

func (s *AtomContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *AtomContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *AtomContext) TRY() antlr.TerminalNode {
	return s.GetToken(LushParserTRY, 0)
}

func (s *AtomContext) DOT() antlr.TerminalNode {
	return s.GetToken(LushParserDOT, 0)
}

func (s *AtomContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *AtomContext) Cmd() ICmdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserLPAREN, 0)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserRPAREN, 0)
}

func (s *AtomContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtomContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LushParserRULE_atom)
	var _la int

	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(382)
			p.String_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(383)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(384)
			p.Bool_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(385)
			p.EnvVar()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(386)
			p.Var_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserTRY {
			{
				p.SetState(387)
				p.Match(LushParserTRY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(390)

				var _m = p.Match(LushParserID)

				localctx.(*AtomContext).namespace = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(391)
				p.Match(LushParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(394)
			p.Func_()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(395)
			p.Cmd()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(396)
			p.Match(LushParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)

			var _x = p.expression(0)

			localctx.(*AtomContext).group = _x
		}
		{
			p.SetState(398)
			p.Match(LushParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(400)
			p.Array()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *FuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserLPAREN, 0)
}

func (s *FuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LushParserRPAREN, 0)
}

func (s *FuncContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FuncContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FuncContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *FuncContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LushParserRULE_func)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.Match(LushParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1736147002620641376) != 0 {
		{
			p.SetState(405)
			p.expression(0)
		}
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(406)
					p.Match(LushParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(407)
					p.expression(0)
				}

			}
			p.SetState(412)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCOMMA {
			{
				p.SetState(413)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(418)
		p.Match(LushParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNamespace returns the namespace token.
	GetNamespace() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetNamespace sets the namespace token.
	SetNamespace(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	namespace antlr.Token
	name      antlr.Token
}

func NewEmptyVarContext() *VarContext {
	var p = new(VarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_var
	return p
}

func InitEmptyVarContext(p *VarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_var
}

func (*VarContext) IsVarContext() {}

func NewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarContext {
	var p = new(VarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_var

	return p
}

func (s *VarContext) GetParser() antlr.Parser { return s.parser }

func (s *VarContext) GetNamespace() antlr.Token { return s.namespace }

func (s *VarContext) GetName() antlr.Token { return s.name }

func (s *VarContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *VarContext) SetName(v antlr.Token) { s.name = v }

func (s *VarContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LushParserID)
}

func (s *VarContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LushParserID, i)
}

func (s *VarContext) DOT() antlr.TerminalNode {
	return s.GetToken(LushParserDOT, 0)
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Var_() (localctx IVarContext) {
	localctx = NewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LushParserRULE_var)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(422)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(420)

			var _m = p.Match(LushParserID)

			localctx.(*VarContext).namespace = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(LushParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(424)

		var _m = p.Match(LushParserID)

		localctx.(*VarContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnvVarContext is an interface to support dynamic dispatch.
type IEnvVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENVVAR() antlr.TerminalNode

	// IsEnvVarContext differentiates from other interfaces.
	IsEnvVarContext()
}

type EnvVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvVarContext() *EnvVarContext {
	var p = new(EnvVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_envVar
	return p
}

func InitEmptyEnvVarContext(p *EnvVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_envVar
}

func (*EnvVarContext) IsEnvVarContext() {}

func NewEnvVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvVarContext {
	var p = new(EnvVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_envVar

	return p
}

func (s *EnvVarContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvVarContext) ENVVAR() antlr.TerminalNode {
	return s.GetToken(LushParserENVVAR, 0)
}

func (s *EnvVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitEnvVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) EnvVar() (localctx IEnvVarContext) {
	localctx = NewEnvVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LushParserRULE_envVar)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(LushParserENVVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(LushParserSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LushParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(428)
		p.Match(LushParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LushParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LushParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(LushParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_bool
	return p
}

func InitEmptyBoolContext(p *BoolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_bool
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(LushParserTRUE, 0)
}

func (s *BoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(LushParserFALSE, 0)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Bool_() (localctx IBoolContext) {
	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LushParserRULE_bool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LushParserTRUE || _la == LushParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LSQ() antlr.TerminalNode
	RSQ() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	PrimitiveType() IPrimitiveTypeContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LSQ() antlr.TerminalNode {
	return s.GetToken(LushParserLSQ, 0)
}

func (s *ArrayContext) RSQ() antlr.TerminalNode {
	return s.GetToken(LushParserRSQ, 0)
}

func (s *ArrayContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LushParserCOMMA)
}

func (s *ArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LushParserCOMMA, i)
}

func (s *ArrayContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LushParserRULE_array)
	var _la int

	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(434)
			p.Match(LushParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1736147002620641376) != 0 {
			{
				p.SetState(435)
				p.expression(0)
			}
			p.SetState(440)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == LushParserCOMMA {
				{
					p.SetState(436)
					p.Match(LushParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(437)
					p.expression(0)
				}

				p.SetState(442)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(445)
			p.Match(LushParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Match(LushParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Match(LushParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.PrimitiveType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArrayType returns the arrayType rule contexts.
	GetArrayType() IPrimitiveTypeContext

	// SetArrayType sets the arrayType rule contexts.
	SetArrayType(IPrimitiveTypeContext)

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	LSQ() antlr.TerminalNode
	RSQ() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	arrayType IPrimitiveTypeContext
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) GetArrayType() IPrimitiveTypeContext { return s.arrayType }

func (s *TypeContext) SetArrayType(v IPrimitiveTypeContext) { s.arrayType = v }

func (s *TypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeContext) LSQ() antlr.TerminalNode {
	return s.GetToken(LushParserLSQ, 0)
}

func (s *TypeContext) RSQ() antlr.TerminalNode {
	return s.GetToken(LushParserRSQ, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LushParserRULE_type)
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LushParserSTRING_TYPE, LushParserINT_TYPE, LushParserBOOL_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(451)
			p.PrimitiveType()
		}

	case LushParserLSQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(452)
			p.Match(LushParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.Match(LushParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)

			var _x = p.PrimitiveType()

			localctx.(*TypeContext).arrayType = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_TYPE() antlr.TerminalNode
	INT_TYPE() antlr.TerminalNode
	BOOL_TYPE() antlr.TerminalNode

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_primitiveType
	return p
}

func InitEmptyPrimitiveTypeContext(p *PrimitiveTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LushParserRULE_primitiveType
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LushParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(LushParserSTRING_TYPE, 0)
}

func (s *PrimitiveTypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(LushParserINT_TYPE, 0)
}

func (s *PrimitiveTypeContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(LushParserBOOL_TYPE, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LushVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LushParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LushParserRULE_primitiveType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *LushParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ReturnStatementContext = nil
		if localctx != nil {
			t = localctx.(*ReturnStatementContext)
		}
		return p.ReturnStatement_Sempred(t, predIndex)

	case 23:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LushParser) ReturnStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.GetTokenStream().LT(-1).GetLine() == p.GetTokenStream().LT(1).GetLine()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LushParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
