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
        | globalAssignment
    )* EOF
    ;

importStatement: IMPORT STRING (AS alias = ID)?;

statement
    : assignment catchClause?
    | globalAssignment catchClause?
    | funcStatement catchClause?
    | if
    | for
    | returnStatement
    | failStatement
    | BREAK
    | CONTINUE
    ;

funcDef
    : FUNC ID LPAREN (
        param (COMMA param)*
    )? RPAREN results? block
    ;

externalFuncDef
    : FUNC ID LPAREN (
        param (COMMA param)*
    )? RPAREN results? EXTERNAL_CODE
    ;

param: ID type;

// error is only valid as the last one, the visitor reports it instead of a syntax error
results: result (COMMA result)*;

result: type | ERROR_TYPE;

if
    : IF expression then = block (
        ELSE (
            elseIf = if
            | elseBlock = block
        )
    )?
    ;

for
    : FOR block
    | FOR cond = expression block
    | FOR init = assignment SEMICOLON cond = expression SEMICOLON
        step = assignment block
    | FOR value = ID (COMMA index = ID)? IN (
        from = expression UNTIL
    )? over = expression block
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

failStatement: FAIL LPAREN expression RPAREN;

catchClause: CATCH name = ID block;

block: LCUR (statement)* RCUR;

assignment
    : ID (COMMA ID)* op = (
        ASSIGN
        | PLUS_ASSIGN
        | MINUS_ASSIGN
        | MUL_ASSIGN
        | DIV_ASSIGN
        | MOD_ASSIGN
    ) expression (COMMA expression)*
    ;

globalAssignment: GLOBAL assignment;

funcStatement: TRY? (namespace = ID DOT)? func;

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
    | TRY? (namespace = ID DOT)? func
    | LPAREN group = expression RPAREN
    | array
    ;

func
    : ID LPAREN (
        expression (COMMA expression)* COMMA?
    )? RPAREN
    ;

var: (namespace = ID DOT)? name = ID;

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

LINE_COMMENT: '//' ~[\r\n]* -> skip;

IMPORT: 'import';
AS: 'as';
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
BREAK: 'break';
CONTINUE: 'continue';
IN: 'in';
UNTIL: 'until';
GLOBAL: 'global';
TRY: 'try';
CATCH: 'catch';
FAIL: 'fail';
ERROR_TYPE: 'error';

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
PLUS_ASSIGN: '+=';
MINUS_ASSIGN: '-=';
MUL_ASSIGN: '*=';
DIV_ASSIGN: '/=';
MOD_ASSIGN: '%=';

STRING
    : '"' (~["\\] | ESCAPED_VALUE)* '"'
    ;
NUMBER: [0-9]+;

EXTERNAL_CODE
    : '```' (. | '\r' | '\n')*? '```'
    ;

fragment ESCAPED_VALUE: '\\' [nt\\"];