parser grammar MockParser;

options {
    tokenVocab = MockLexer;
}
mockExpression:(normalExpression|stringTemplate);

normalExpression:
    presetExpr
    |numberExpr
    |stringExpr
    |customArgSpec
    |constStatement
    |ignoreStatement
    ;

presetExpr:
    |presetBoolExpr
    |presetSliceExpr
    |presetMapExpr
    |presetSpecExpr
    ;

presetBoolExpr:
    'bool' '=' BOOL_VALUE;

presetSliceExpr:
    'slice' '(' (NUMBER ',')? mockExpression ')';

presetMapExpr:
    'map' '(' NUMBER ',' mockExpression ',' mockExpression ')';

presetSpecExpr:
    'spec' '=' VALUE;


numberExpr:numberRangeExpr|numberLengthExpr;

numberRangeExpr:
    IDENT rangeSpec
    ;
numberLengthExpr:
    IDENT lengthSpec;

stringExpr:
     stringRange|stringLength;

stringRange:
    IDENT '(' LANG '->' rangeSpec ')';

stringLength:
    IDENT '(' LANG?  ('->' NUMBER )? ')';

stringTemplate:
    prevConst? '<<'IDENT'>>' afterConst?;

customArgSpec:
    IDENT '('(customArg)+')';

customArg: VALUE ','?;

constStatement: IDENT;

prevConst:VALUE;

afterConst:VALUE;

lengthSpec:'(' NUMBER ')';

rangeSpec: '[' NUMBER ',' NUMBER ']';

ignoreStatement: '-';
