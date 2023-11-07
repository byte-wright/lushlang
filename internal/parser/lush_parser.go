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
		"program", "statement", "funcDef", "param", "if", "for", "returnStatement",
		"block", "assignment", "funcStatement", "expression", "atom", "func",
		"var", "envVar", "string", "number", "bool", "array", "type", "primitiveType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 223, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 5, 0, 45, 8, 0, 10, 0, 12, 0, 48, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 57, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5,
		2, 65, 8, 2, 10, 2, 12, 2, 68, 9, 2, 3, 2, 70, 8, 2, 1, 2, 1, 2, 3, 2,
		74, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 5,
		7, 98, 8, 7, 10, 7, 12, 7, 101, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 115, 8, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 141, 8, 10, 1, 10, 1, 10, 3, 10, 145, 8, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 153, 8, 10, 10, 10, 12, 10, 156,
		9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 3, 11, 169, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12,
		176, 8, 12, 10, 12, 12, 12, 179, 9, 12, 1, 12, 3, 12, 182, 8, 12, 3, 12,
		184, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 202, 8, 18,
		10, 18, 12, 18, 205, 9, 18, 3, 18, 207, 8, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 213, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 219, 8, 19, 1,
		20, 1, 20, 1, 20, 0, 1, 20, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 6, 1, 0, 13, 14, 1, 0, 16, 18,
		1, 0, 14, 15, 1, 0, 19, 24, 1, 0, 2, 3, 1, 0, 6, 8, 236, 0, 46, 1, 0, 0,
		0, 2, 56, 1, 0, 0, 0, 4, 58, 1, 0, 0, 0, 6, 77, 1, 0, 0, 0, 8, 80, 1, 0,
		0, 0, 10, 84, 1, 0, 0, 0, 12, 92, 1, 0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 104,
		1, 0, 0, 0, 18, 108, 1, 0, 0, 0, 20, 114, 1, 0, 0, 0, 22, 168, 1, 0, 0,
		0, 24, 170, 1, 0, 0, 0, 26, 187, 1, 0, 0, 0, 28, 189, 1, 0, 0, 0, 30, 191,
		1, 0, 0, 0, 32, 193, 1, 0, 0, 0, 34, 195, 1, 0, 0, 0, 36, 212, 1, 0, 0,
		0, 38, 218, 1, 0, 0, 0, 40, 220, 1, 0, 0, 0, 42, 45, 3, 2, 1, 0, 43, 45,
		3, 4, 2, 0, 44, 42, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0,
		46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 46, 1,
		0, 0, 0, 49, 50, 5, 0, 0, 1, 50, 1, 1, 0, 0, 0, 51, 57, 3, 16, 8, 0, 52,
		57, 3, 18, 9, 0, 53, 57, 3, 8, 4, 0, 54, 57, 3, 10, 5, 0, 55, 57, 3, 12,
		6, 0, 56, 51, 1, 0, 0, 0, 56, 52, 1, 0, 0, 0, 56, 53, 1, 0, 0, 0, 56, 54,
		1, 0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 3, 1, 0, 0, 0, 58, 59, 5, 9, 0, 0,
		59, 60, 5, 36, 0, 0, 60, 69, 5, 25, 0, 0, 61, 66, 3, 6, 3, 0, 62, 63, 5,
		31, 0, 0, 63, 65, 3, 6, 3, 0, 64, 62, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66,
		64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0,
		0, 69, 61, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73,
		5, 26, 0, 0, 72, 74, 3, 38, 19, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0,
		0, 74, 75, 1, 0, 0, 0, 75, 76, 3, 14, 7, 0, 76, 5, 1, 0, 0, 0, 77, 78,
		5, 36, 0, 0, 78, 79, 3, 38, 19, 0, 79, 7, 1, 0, 0, 0, 80, 81, 5, 4, 0,
		0, 81, 82, 3, 20, 10, 0, 82, 83, 3, 14, 7, 0, 83, 9, 1, 0, 0, 0, 84, 85,
		5, 5, 0, 0, 85, 86, 3, 16, 8, 0, 86, 87, 5, 34, 0, 0, 87, 88, 3, 20, 10,
		0, 88, 89, 5, 34, 0, 0, 89, 90, 3, 16, 8, 0, 90, 91, 3, 14, 7, 0, 91, 11,
		1, 0, 0, 0, 92, 93, 5, 10, 0, 0, 93, 94, 3, 20, 10, 0, 94, 13, 1, 0, 0,
		0, 95, 99, 5, 27, 0, 0, 96, 98, 3, 2, 1, 0, 97, 96, 1, 0, 0, 0, 98, 101,
		1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 102, 1, 0, 0,
		0, 101, 99, 1, 0, 0, 0, 102, 103, 5, 28, 0, 0, 103, 15, 1, 0, 0, 0, 104,
		105, 5, 36, 0, 0, 105, 106, 5, 37, 0, 0, 106, 107, 3, 20, 10, 0, 107, 17,
		1, 0, 0, 0, 108, 109, 3, 24, 12, 0, 109, 19, 1, 0, 0, 0, 110, 111, 6, 10,
		-1, 0, 111, 115, 3, 22, 11, 0, 112, 113, 7, 0, 0, 0, 113, 115, 3, 20, 10,
		9, 114, 110, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 154, 1, 0, 0, 0, 116,
		117, 10, 6, 0, 0, 117, 118, 7, 1, 0, 0, 118, 153, 3, 20, 10, 7, 119, 120,
		10, 5, 0, 0, 120, 121, 7, 2, 0, 0, 121, 153, 3, 20, 10, 6, 122, 123, 10,
		4, 0, 0, 123, 124, 7, 3, 0, 0, 124, 153, 3, 20, 10, 5, 125, 126, 10, 3,
		0, 0, 126, 127, 5, 11, 0, 0, 127, 153, 3, 20, 10, 4, 128, 129, 10, 2, 0,
		0, 129, 130, 5, 12, 0, 0, 130, 153, 3, 20, 10, 3, 131, 132, 10, 1, 0, 0,
		132, 133, 5, 32, 0, 0, 133, 134, 3, 20, 10, 0, 134, 135, 5, 33, 0, 0, 135,
		136, 3, 20, 10, 2, 136, 153, 1, 0, 0, 0, 137, 138, 10, 8, 0, 0, 138, 140,
		5, 29, 0, 0, 139, 141, 3, 20, 10, 0, 140, 139, 1, 0, 0, 0, 140, 141, 1,
		0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 5, 33, 0, 0, 143, 145, 3, 20,
		10, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0,
		146, 153, 5, 30, 0, 0, 147, 148, 10, 7, 0, 0, 148, 149, 5, 29, 0, 0, 149,
		150, 3, 20, 10, 0, 150, 151, 5, 30, 0, 0, 151, 153, 1, 0, 0, 0, 152, 116,
		1, 0, 0, 0, 152, 119, 1, 0, 0, 0, 152, 122, 1, 0, 0, 0, 152, 125, 1, 0,
		0, 0, 152, 128, 1, 0, 0, 0, 152, 131, 1, 0, 0, 0, 152, 137, 1, 0, 0, 0,
		152, 147, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 21, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 169, 3,
		30, 15, 0, 158, 169, 3, 32, 16, 0, 159, 169, 3, 34, 17, 0, 160, 169, 3,
		28, 14, 0, 161, 169, 3, 26, 13, 0, 162, 169, 3, 24, 12, 0, 163, 164, 5,
		25, 0, 0, 164, 165, 3, 20, 10, 0, 165, 166, 5, 26, 0, 0, 166, 169, 1, 0,
		0, 0, 167, 169, 3, 36, 18, 0, 168, 157, 1, 0, 0, 0, 168, 158, 1, 0, 0,
		0, 168, 159, 1, 0, 0, 0, 168, 160, 1, 0, 0, 0, 168, 161, 1, 0, 0, 0, 168,
		162, 1, 0, 0, 0, 168, 163, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 23, 1,
		0, 0, 0, 170, 171, 5, 36, 0, 0, 171, 183, 5, 25, 0, 0, 172, 177, 3, 20,
		10, 0, 173, 174, 5, 31, 0, 0, 174, 176, 3, 20, 10, 0, 175, 173, 1, 0, 0,
		0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178,
		181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 182, 5, 31, 0, 0, 181, 180,
		1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 172, 1, 0,
		0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 5, 26, 0, 0,
		186, 25, 1, 0, 0, 0, 187, 188, 5, 36, 0, 0, 188, 27, 1, 0, 0, 0, 189, 190,
		5, 35, 0, 0, 190, 29, 1, 0, 0, 0, 191, 192, 5, 38, 0, 0, 192, 31, 1, 0,
		0, 0, 193, 194, 5, 39, 0, 0, 194, 33, 1, 0, 0, 0, 195, 196, 7, 4, 0, 0,
		196, 35, 1, 0, 0, 0, 197, 206, 5, 29, 0, 0, 198, 203, 3, 20, 10, 0, 199,
		200, 5, 31, 0, 0, 200, 202, 3, 20, 10, 0, 201, 199, 1, 0, 0, 0, 202, 205,
		1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 207, 1, 0,
		0, 0, 205, 203, 1, 0, 0, 0, 206, 198, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0,
		207, 208, 1, 0, 0, 0, 208, 213, 5, 30, 0, 0, 209, 210, 5, 29, 0, 0, 210,
		211, 5, 30, 0, 0, 211, 213, 3, 40, 20, 0, 212, 197, 1, 0, 0, 0, 212, 209,
		1, 0, 0, 0, 213, 37, 1, 0, 0, 0, 214, 219, 3, 40, 20, 0, 215, 216, 5, 29,
		0, 0, 216, 217, 5, 30, 0, 0, 217, 219, 3, 38, 19, 0, 218, 214, 1, 0, 0,
		0, 218, 215, 1, 0, 0, 0, 219, 39, 1, 0, 0, 0, 220, 221, 7, 5, 0, 0, 221,
		41, 1, 0, 0, 0, 20, 44, 46, 56, 66, 69, 73, 99, 114, 140, 144, 152, 154,
		168, 177, 181, 183, 203, 206, 212, 218,
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
	LushParserEOF         = antlr.TokenEOF
	LushParserWHITESPACE  = 1
	LushParserTRUE        = 2
	LushParserFALSE       = 3
	LushParserIF          = 4
	LushParserFOR         = 5
	LushParserSTRING_TYPE = 6
	LushParserINT_TYPE    = 7
	LushParserBOOL_TYPE   = 8
	LushParserFUNC        = 9
	LushParserRETURN      = 10
	LushParserLAND        = 11
	LushParserLOR         = 12
	LushParserNOT         = 13
	LushParserMINUS       = 14
	LushParserPLUS        = 15
	LushParserMUL         = 16
	LushParserDIV         = 17
	LushParserMOD         = 18
	LushParserLT          = 19
	LushParserLTE         = 20
	LushParserGT          = 21
	LushParserGTE         = 22
	LushParserEQ          = 23
	LushParserNEQ         = 24
	LushParserLPAREN      = 25
	LushParserRPAREN      = 26
	LushParserLCUR        = 27
	LushParserRCUR        = 28
	LushParserLSQ         = 29
	LushParserRSQ         = 30
	LushParserCOMMA       = 31
	LushParserQUESTION    = 32
	LushParserCOLON       = 33
	LushParserSEMICOLON   = 34
	LushParserENVVAR      = 35
	LushParserID          = 36
	LushParserASSIGN      = 37
	LushParserSTRING      = 38
	LushParserNUMBER      = 39
)

// LushParser rules.
const (
	LushParserRULE_program         = 0
	LushParserRULE_statement       = 1
	LushParserRULE_funcDef         = 2
	LushParserRULE_param           = 3
	LushParserRULE_if              = 4
	LushParserRULE_for             = 5
	LushParserRULE_returnStatement = 6
	LushParserRULE_block           = 7
	LushParserRULE_assignment      = 8
	LushParserRULE_funcStatement   = 9
	LushParserRULE_expression      = 10
	LushParserRULE_atom            = 11
	LushParserRULE_func            = 12
	LushParserRULE_var             = 13
	LushParserRULE_envVar          = 14
	LushParserRULE_string          = 15
	LushParserRULE_number          = 16
	LushParserRULE_bool            = 17
	LushParserRULE_array           = 18
	LushParserRULE_type            = 19
	LushParserRULE_primitiveType   = 20
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719478320) != 0 {
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LushParserIF, LushParserFOR, LushParserRETURN, LushParserID:
			{
				p.SetState(42)
				p.Statement()
			}

		case LushParserFUNC:
			{
				p.SetState(43)
				p.FuncDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	FuncStatement() IFuncStatementContext
	If_() IIfContext
	For_() IForContext
	ReturnStatement() IReturnStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 2, LushParserRULE_statement)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.FuncStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)
			p.If_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.For_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)
			p.ReturnStatement()
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

	// GetReturn_ returns the return_ rule contexts.
	GetReturn_() ITypeContext

	// SetReturn_ sets the return_ rule contexts.
	SetReturn_(ITypeContext)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	AllParam() []IParamContext
	Param(i int) IParamContext
	Type_() ITypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	return_ ITypeContext
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

func (s *FuncDefContext) GetReturn_() ITypeContext { return s.return_ }

func (s *FuncDefContext) SetReturn_(v ITypeContext) { s.return_ = v }

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

func (s *FuncDefContext) Type_() ITypeContext {
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
	p.EnterRule(localctx, 4, LushParserRULE_funcDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(LushParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(LushParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LushParserID {
		{
			p.SetState(61)
			p.Param()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LushParserCOMMA {
			{
				p.SetState(62)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(63)
				p.Param()
			}

			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(71)
		p.Match(LushParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536871360) != 0 {
		{
			p.SetState(72)

			var _x = p.Type_()

			localctx.(*FuncDefContext).return_ = _x
		}

	}
	{
		p.SetState(75)
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
	p.EnterRule(localctx, 6, LushParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
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

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *IfContext) Block() IBlockContext {
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
	p.EnterRule(localctx, 8, LushParserRULE_if)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(LushParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.expression(0)
	}
	{
		p.SetState(82)
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

// IForContext is an interface to support dynamic dispatch.
type IForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsForContext differentiates from other interfaces.
	IsForContext()
}

type ForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(LushParserFOR, 0)
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

func (s *ForContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(LushParserSEMICOLON)
}

func (s *ForContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(LushParserSEMICOLON, i)
}

func (s *ForContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 10, LushParserRULE_for)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(LushParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Assignment()
	}
	{
		p.SetState(86)
		p.Match(LushParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.expression(0)
	}
	{
		p.SetState(88)
		p.Match(LushParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Assignment()
	}
	{
		p.SetState(90)
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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

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

func (s *ReturnStatementContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 12, LushParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(LushParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.expression(0)
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
	p.EnterRule(localctx, 14, LushParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(LushParserLCUR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719477808) != 0 {
		{
			p.SetState(96)
			p.Statement()
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(102)
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

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(LushParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 16, LushParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(LushParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.expression(0)
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

	// Getter signatures
	Func_() IFuncContext

	// IsFuncStatementContext differentiates from other interfaces.
	IsFuncStatementContext()
}

type FuncStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 18, LushParserRULE_funcStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, LushParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LushParserTRUE, LushParserFALSE, LushParserLPAREN, LushParserLSQ, LushParserENVVAR, LushParserID, LushParserSTRING, LushParserNUMBER:
		{
			p.SetState(111)
			p.Atom()
		}

	case LushParserNOT, LushParserMINUS:
		{
			p.SetState(112)

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
			p.SetState(113)
			p.expression(9)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(117)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).mul_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&458752) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).mul_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(118)
					p.expression(7)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(120)

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
					p.SetState(121)
					p.expression(6)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(123)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).rel_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33030144) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).rel_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(124)
					p.expression(5)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(126)
					p.Match(LushParserLAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(127)
					p.expression(4)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(129)
					p.Match(LushParserLOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(130)
					p.expression(3)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).ternary = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(132)
					p.Match(LushParserQUESTION)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(133)
					p.expression(0)
				}
				{
					p.SetState(134)
					p.Match(LushParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(135)
					p.expression(2)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).slice = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LushParserRULE_expression)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(138)
					p.Match(LushParserLSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(140)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&928283385868) != 0 {
					{
						p.SetState(139)

						var _x = p.expression(0)

						localctx.(*ExpressionContext).from = _x
					}

				}
				{
					p.SetState(142)
					p.Match(LushParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(144)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&928283385868) != 0 {
					{
						p.SetState(143)

						var _x = p.expression(0)

						localctx.(*ExpressionContext).to = _x
					}

				}
				{
					p.SetState(146)
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
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(148)
					p.Match(LushParserLSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(149)

					var _x = p.expression(0)

					localctx.(*ExpressionContext).position = _x
				}
				{
					p.SetState(150)
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
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	Array() IArrayContext

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	group  IExpressionContext
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
	p.EnterRule(localctx, 22, LushParserRULE_atom)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.String_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.Bool_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
			p.EnvVar()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(161)
			p.Var_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(162)
			p.Func_()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(163)
			p.Match(LushParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

			var _x = p.expression(0)

			localctx.(*AtomContext).group = _x
		}
		{
			p.SetState(165)
			p.Match(LushParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(167)
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
	p.EnterRule(localctx, 24, LushParserRULE_func)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(LushParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(LushParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&928283385868) != 0 {
		{
			p.SetState(172)
			p.expression(0)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(173)
					p.Match(LushParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(174)
					p.expression(0)
				}

			}
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LushParserCOMMA {
			{
				p.SetState(180)
				p.Match(LushParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(185)
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

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *VarContext) ID() antlr.TerminalNode {
	return s.GetToken(LushParserID, 0)
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
	p.EnterRule(localctx, 26, LushParserRULE_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(LushParserID)
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
	p.EnterRule(localctx, 28, LushParserRULE_envVar)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
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
	p.EnterRule(localctx, 30, LushParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
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
	p.EnterRule(localctx, 32, LushParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
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
	p.EnterRule(localctx, 34, LushParserRULE_bool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
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
	p.EnterRule(localctx, 36, LushParserRULE_array)
	var _la int

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(LushParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&928283385868) != 0 {
			{
				p.SetState(198)
				p.expression(0)
			}
			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == LushParserCOMMA {
				{
					p.SetState(199)
					p.Match(LushParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(200)
					p.expression(0)
				}

				p.SetState(205)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(208)
			p.Match(LushParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.Match(LushParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(LushParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
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
	GetArrayType() ITypeContext

	// SetArrayType sets the arrayType rule contexts.
	SetArrayType(ITypeContext)

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	LSQ() antlr.TerminalNode
	RSQ() antlr.TerminalNode
	Type_() ITypeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	arrayType ITypeContext
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

func (s *TypeContext) GetArrayType() ITypeContext { return s.arrayType }

func (s *TypeContext) SetArrayType(v ITypeContext) { s.arrayType = v }

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

func (s *TypeContext) Type_() ITypeContext {
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
	p.EnterRule(localctx, 38, LushParserRULE_type)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LushParserSTRING_TYPE, LushParserINT_TYPE, LushParserBOOL_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.PrimitiveType()
		}

	case LushParserLSQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(LushParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(LushParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)

			var _x = p.Type_()

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
	p.EnterRule(localctx, 40, LushParserRULE_primitiveType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
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
	case 10:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LushParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
