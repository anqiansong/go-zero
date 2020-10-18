lexer grammar MockLexer;

// tokens

BOOL_VALUE:
    BOOL_TRUE
    |BOOL_FALSE
    ;
BOOL:
    'bool';
SLICE:
    'slice';
MAP:
    'map'
    ;
SPEC:
    'spec';

LANG:
    LANG_ZH
    |LANG_EN
    |LANG_NUMBER
    ;

// MAKERS
LPAREN:     '(';
RPAREN:     ')';
LBRACE:     '{';
RBRACE:     '}';
LBRACK:     '[';
RBRACK:     ']';
DOT:        '.';
SMICOLON:   ';';
COMMA:      ',';
ASSIGN:     '=';
QUOTE_LEFT:  '<<';
QUOTE_RIGHT:  '>>';
SHARP:        '#';
TO:          '->';
PUT:          '<-';
BAR:           '-';

WS:         [ \t\r\n]+ -> channel(HIDDEN) ;

// IDTIFIER
IDENT:       LOWER_CASE(ALPHA|DIGIT|UNDERSCORE)*;

// NUMBER
NUMBER:      '-'?DIGIT+;

PACKAGE:(LOWER_CASE+ DOT?)+;

VALUE: '{' ~[\r\n{}"]+ '}';

fragment LANG_ZH:
    'zh';
fragment LANG_EN:
    'en';
fragment LANG_NUMBER:
    'number';
fragment BOOL_TRUE:
    'true';
fragment BOOL_FALSE:
    'false';
fragment DIGIT: [0-9];

fragment LOWER_CASE:[a-z];

fragment UNDERSCORE: '_';

fragment ALPHA: [a-zA-Z];

ERRCHAR
	:	.	-> channel(HIDDEN)
	;