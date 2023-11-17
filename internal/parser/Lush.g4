grammar Lush
    ;

program
    : importStatement* (
        statement
        | funcDef
    )* EOF
    ;

library
    : importStatement*
        funcDef
    * EOF
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

param: ID type;

if: IF expression block;

for
    : FOR assignment SEMICOLON expression SEMICOLON
        assignment block
    ;

returnStatement
    : RETURN (
        expression (COMMA expression)*
    )
    ;

block: LCUR (statement)* RCUR;

assignment
    : ID (COMMA ID)* ASSIGN expression (
        COMMA expression
    )*
    ;

funcStatement: func;

expression
    : atom
    | expression DOT func
    | unary_op = (MINUS | NOT) expression
    | slice = expression LSQ from = expression? COLON to =
        expression? RSQ
    | index = expression LSQ position = expression RSQ
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
    | ternary = expression QUESTION expression COLON
        expression
    ;

atom
    : string
    | number
    | bool
    | envVar
    | var
    | func
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
    | LSQ RSQ arrayType = type
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

ENVVAR: '$' [a-zA-Z_0-0]+;

ID: [a-z_] [a-zA-Z_0-0]*;
ASSIGN: '=';

STRING
    : '"' (~["\\] | ESCAPED_VALUE)* '"'
    ;
NUMBER: [0-9]+;

fragment ESCAPED_VALUE: '\\' [nt\\"];