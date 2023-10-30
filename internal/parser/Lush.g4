grammar Lush
    ;

program: (statement)* EOF;

statement
    : assignment
    | funcStatement
    | if
    ;

if: IF expression block;

block: LCUR (statement)* RCUR;

assignment: ID ASSIGN expression;

funcStatement: func;

expression
    : atom
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
    | ternary = expression QUESTION expression COLON expression
    | slice = expression LSQ from=expression? COLON to=expression? RSQ
    ;

atom
    : string
    | number
    | bool
    | envVar
    | var
    | func
    | LPAREN group = expression RPAREN
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

WHITESPACE: [ \r\n\t]+ -> skip;

TRUE: 'true';
FALSE: 'false';
IF: 'if';

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

ENVVAR: '$' [a-zA-Z_0-0]+;

ID: [a-z_] [a-zA-Z_0-0]*;
ASSIGN: '=';

STRING
    : '"' (~["\\] | ESCAPED_VALUE)* '"'
    ;
NUMBER: [0-9]+;

fragment ESCAPED_VALUE: '\\' [nt\\"];