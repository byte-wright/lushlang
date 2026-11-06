grammar Lush
    ;

program
    : importStatement* (
        statement
        | funcDef
    )* EOF
    ;

package
    : importStatement* (
        funcDef
        | externalFuncDef
    )* EOF
    ;

importStatement: IMPORT STRING;

statement
    : assignment
    | funcStatement
    | if
    | for
    | returnStatement
    ;

funcDef
    : FUNC ID LPAREN (
        param (COMMA param)*
    )? RPAREN (type (COMMA type)*)? block
    ;

externalFuncDef
    : FUNC ID LPAREN (
        param (COMMA param)*
    )? RPAREN (type (COMMA type)*)? EXTERNAL_CODE
    ;

param: ID type;

if
    : IF expression then = block (
        ELSE (
            elseIf = if
            | elseBlock = block
        )
    )?
    ;

for
    : FOR assignment SEMICOLON expression SEMICOLON
        assignment block
    ;

// like in Go, a return at the end of a line has no values, so code after it isn't
// taken as its value
returnStatement
    : RETURN (
        {p.GetTokenStream().LT(-1).GetLine() == p.GetTokenStream().LT(1).GetLine()}? expression (
            COMMA expression
        )*
    )?
    ;

block: LCUR (statement)* RCUR;

assignment
    : ID (COMMA ID)* ASSIGN expression (
        COMMA expression
    )*
    ;

funcStatement: (namespace = ID DOT)? func;

expression
    : atom
    | slice = expression LSQ from = expression? COLON to =
        expression? RSQ
    | index = expression LSQ position = expression RSQ
    | unary_op = (MINUS | NOT) expression
    | expression mul_op = (
        MUL
        | DIV
        | MOD
    ) expression
    | expression add_op = (
        PLUS
        | MINUS
    ) expression
    | expression rel_op = (
        EQ
        | NEQ
        | LT
        | LTE
        | GT
        | GTE
    ) expression
    | expression LAND expression
    | expression LOR expression
    | <assoc = right> ternary = expression QUESTION expression COLON
        expression
    ;

atom
    : string
    | number
    | bool
    | envVar
    | var
    | (namespace = ID DOT)? func
    | LPAREN group = expression RPAREN
    | array
    ;

func
    : ID LPAREN (
        expression (COMMA expression)* COMMA?
    )? RPAREN
    ;

var: ID;

envVar: ENVVAR;

string: STRING;

number: NUMBER;

bool: TRUE | FALSE;

array
    : LSQ (
        expression (COMMA expression)*
    )? RSQ
    | LSQ RSQ primitiveType
    ;

type
    : primitiveType
    // bash arrays can't be nested
    | LSQ RSQ arrayType = primitiveType
    ;

primitiveType
    : STRING_TYPE
    | INT_TYPE
    | BOOL_TYPE
    ;

WHITESPACE: [ \r\n\t]+ -> skip;

IMPORT: 'import';
TRUE: 'true';
FALSE: 'false';
IF: 'if';
ELSE: 'else';
FOR: 'for';
STRING_TYPE: 'string';
INT_TYPE: 'int';
BOOL_TYPE: 'bool';
FUNC: 'func';
RETURN: 'return';

LAND: '&&';
LOR: '||';
NOT: '!';

MINUS: '-';
PLUS: '+';
MUL: '*';
DIV: '/';
MOD: '%';

LT: '<';
LTE: '<=';
GT: '>';
GTE: '>=';
EQ: '==';
NEQ: '!=';

LPAREN: '(';
RPAREN: ')';
LCUR: '{';
RCUR: '}';
LSQ: '[';
RSQ: ']';
COMMA: ',';
QUESTION: '?';
COLON: ':';
SEMICOLON: ';';
DOT: '.';

ENVVAR: '$' [a-zA-Z_0-9]+;

// names must start lowercase, uppercase ones are only lexed so that they get one error
// instead of one per letter
ID: [a-zA-Z_] [a-zA-Z_0-9]*;
ASSIGN: '=';

STRING
    : '"' (~["\\] | ESCAPED_VALUE)* '"'
    ;
NUMBER: [0-9]+;

EXTERNAL_CODE
    : '```' (. | '\r' | '\n')*? '```'
    ;

fragment ESCAPED_VALUE: '\\' [nt\\"];