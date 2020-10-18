// Code generated from /Users/anqiansong/goland/go/go-zero/anqiansong/mock/grammar/MockParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MockParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 164,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 5, 2, 49, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	57, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 64, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 74, 10, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 5, 9, 94, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 5, 12, 104, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 116, 10, 14, 3, 14, 3, 14, 5, 14,
	120, 10, 14, 3, 14, 3, 14, 3, 15, 5, 15, 125, 10, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 131, 10, 15, 3, 16, 3, 16, 3, 16, 6, 16, 136, 10, 16,
	13, 16, 14, 16, 137, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 144, 10, 17, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 2, 2, 2, 160, 2, 48, 3, 2, 2, 2, 4, 56, 3, 2, 2, 2, 6, 63, 3, 2, 2,
	2, 8, 65, 3, 2, 2, 2, 10, 69, 3, 2, 2, 2, 12, 78, 3, 2, 2, 2, 14, 87, 3,
	2, 2, 2, 16, 93, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 98, 3, 2, 2, 2, 22,
	103, 3, 2, 2, 2, 24, 105, 3, 2, 2, 2, 26, 112, 3, 2, 2, 2, 28, 124, 3,
	2, 2, 2, 30, 132, 3, 2, 2, 2, 32, 141, 3, 2, 2, 2, 34, 145, 3, 2, 2, 2,
	36, 147, 3, 2, 2, 2, 38, 149, 3, 2, 2, 2, 40, 151, 3, 2, 2, 2, 42, 155,
	3, 2, 2, 2, 44, 161, 3, 2, 2, 2, 46, 49, 5, 4, 3, 2, 47, 49, 5, 28, 15,
	2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 3, 3, 2, 2, 2, 50, 57, 5,
	6, 4, 2, 51, 57, 5, 16, 9, 2, 52, 57, 5, 22, 12, 2, 53, 57, 5, 30, 16,
	2, 54, 57, 5, 34, 18, 2, 55, 57, 5, 44, 23, 2, 56, 50, 3, 2, 2, 2, 56,
	51, 3, 2, 2, 2, 56, 52, 3, 2, 2, 2, 56, 53, 3, 2, 2, 2, 56, 54, 3, 2, 2,
	2, 56, 55, 3, 2, 2, 2, 57, 5, 3, 2, 2, 2, 58, 64, 3, 2, 2, 2, 59, 64, 5,
	8, 5, 2, 60, 64, 5, 10, 6, 2, 61, 64, 5, 12, 7, 2, 62, 64, 5, 14, 8, 2,
	63, 58, 3, 2, 2, 2, 63, 59, 3, 2, 2, 2, 63, 60, 3, 2, 2, 2, 63, 61, 3,
	2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 7, 3, 2, 2, 2, 65, 66, 7, 4, 2, 2, 66,
	67, 7, 18, 2, 2, 67, 68, 7, 3, 2, 2, 68, 9, 3, 2, 2, 2, 69, 70, 7, 5, 2,
	2, 70, 73, 7, 9, 2, 2, 71, 72, 7, 27, 2, 2, 72, 74, 7, 17, 2, 2, 73, 71,
	3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 5, 2, 2, 2,
	76, 77, 7, 10, 2, 2, 77, 11, 3, 2, 2, 2, 78, 79, 7, 6, 2, 2, 79, 80, 7,
	9, 2, 2, 80, 81, 7, 27, 2, 2, 81, 82, 7, 17, 2, 2, 82, 83, 5, 2, 2, 2,
	83, 84, 7, 17, 2, 2, 84, 85, 5, 2, 2, 2, 85, 86, 7, 10, 2, 2, 86, 13, 3,
	2, 2, 2, 87, 88, 7, 7, 2, 2, 88, 89, 7, 18, 2, 2, 89, 90, 7, 29, 2, 2,
	90, 15, 3, 2, 2, 2, 91, 94, 5, 18, 10, 2, 92, 94, 5, 20, 11, 2, 93, 91,
	3, 2, 2, 2, 93, 92, 3, 2, 2, 2, 94, 17, 3, 2, 2, 2, 95, 96, 7, 26, 2, 2,
	96, 97, 5, 42, 22, 2, 97, 19, 3, 2, 2, 2, 98, 99, 7, 26, 2, 2, 99, 100,
	5, 40, 21, 2, 100, 21, 3, 2, 2, 2, 101, 104, 5, 24, 13, 2, 102, 104, 5,
	26, 14, 2, 103, 101, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 23, 3, 2, 2,
	2, 105, 106, 7, 26, 2, 2, 106, 107, 7, 9, 2, 2, 107, 108, 7, 8, 2, 2, 108,
	109, 7, 22, 2, 2, 109, 110, 5, 42, 22, 2, 110, 111, 7, 10, 2, 2, 111, 25,
	3, 2, 2, 2, 112, 113, 7, 26, 2, 2, 113, 115, 7, 9, 2, 2, 114, 116, 7, 8,
	2, 2, 115, 114, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2,
	117, 118, 7, 22, 2, 2, 118, 120, 7, 27, 2, 2, 119, 117, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 7, 10, 2, 2, 122, 27,
	3, 2, 2, 2, 123, 125, 5, 36, 19, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3,
	2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 7, 19, 2, 2, 127, 128, 7, 26,
	2, 2, 128, 130, 7, 20, 2, 2, 129, 131, 5, 38, 20, 2, 130, 129, 3, 2, 2,
	2, 130, 131, 3, 2, 2, 2, 131, 29, 3, 2, 2, 2, 132, 133, 7, 26, 2, 2, 133,
	135, 7, 9, 2, 2, 134, 136, 5, 32, 17, 2, 135, 134, 3, 2, 2, 2, 136, 137,
	3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2,
	2, 2, 139, 140, 7, 10, 2, 2, 140, 31, 3, 2, 2, 2, 141, 143, 7, 29, 2, 2,
	142, 144, 7, 17, 2, 2, 143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144,
	33, 3, 2, 2, 2, 145, 146, 7, 26, 2, 2, 146, 35, 3, 2, 2, 2, 147, 148, 7,
	29, 2, 2, 148, 37, 3, 2, 2, 2, 149, 150, 7, 29, 2, 2, 150, 39, 3, 2, 2,
	2, 151, 152, 7, 9, 2, 2, 152, 153, 7, 27, 2, 2, 153, 154, 7, 10, 2, 2,
	154, 41, 3, 2, 2, 2, 155, 156, 7, 13, 2, 2, 156, 157, 7, 27, 2, 2, 157,
	158, 7, 17, 2, 2, 158, 159, 7, 27, 2, 2, 159, 160, 7, 14, 2, 2, 160, 43,
	3, 2, 2, 2, 161, 162, 7, 24, 2, 2, 162, 45, 3, 2, 2, 2, 14, 48, 56, 63,
	73, 93, 103, 115, 119, 124, 130, 137, 143,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'bool'", "'slice'", "'map'", "'spec'", "", "'('", "')'", "'{'",
	"'}'", "'['", "']'", "'.'", "';'", "','", "'='", "'<<'", "'>>'", "'#'",
	"'->'", "'<-'", "'-'",
}
var symbolicNames = []string{
	"", "BOOL_VALUE", "BOOL", "SLICE", "MAP", "SPEC", "LANG", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "LBRACK", "RBRACK", "DOT", "SMICOLON", "COMMA", "ASSIGN",
	"QUOTE_LEFT", "QUOTE_RIGHT", "SHARP", "TO", "PUT", "BAR", "WS", "IDENT",
	"NUMBER", "PACKAGE", "VALUE", "ERRCHAR",
}

var ruleNames = []string{
	"mockExpression", "normalExpression", "presetExpr", "presetBoolExpr", "presetSliceExpr",
	"presetMapExpr", "presetSpecExpr", "numberExpr", "numberRangeExpr", "numberLengthExpr",
	"stringExpr", "stringRange", "stringLength", "stringTemplate", "customArgSpec",
	"customArg", "constStatement", "prevConst", "afterConst", "lengthSpec",
	"rangeSpec", "ignoreStatement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MockParser struct {
	*antlr.BaseParser
}

func NewMockParser(input antlr.TokenStream) *MockParser {
	this := new(MockParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MockParser.g4"

	return this
}

// MockParser tokens.
const (
	MockParserEOF         = antlr.TokenEOF
	MockParserBOOL_VALUE  = 1
	MockParserBOOL        = 2
	MockParserSLICE       = 3
	MockParserMAP         = 4
	MockParserSPEC        = 5
	MockParserLANG        = 6
	MockParserLPAREN      = 7
	MockParserRPAREN      = 8
	MockParserLBRACE      = 9
	MockParserRBRACE      = 10
	MockParserLBRACK      = 11
	MockParserRBRACK      = 12
	MockParserDOT         = 13
	MockParserSMICOLON    = 14
	MockParserCOMMA       = 15
	MockParserASSIGN      = 16
	MockParserQUOTE_LEFT  = 17
	MockParserQUOTE_RIGHT = 18
	MockParserSHARP       = 19
	MockParserTO          = 20
	MockParserPUT         = 21
	MockParserBAR         = 22
	MockParserWS          = 23
	MockParserIDENT       = 24
	MockParserNUMBER      = 25
	MockParserPACKAGE     = 26
	MockParserVALUE       = 27
	MockParserERRCHAR     = 28
)

// MockParser rules.
const (
	MockParserRULE_mockExpression   = 0
	MockParserRULE_normalExpression = 1
	MockParserRULE_presetExpr       = 2
	MockParserRULE_presetBoolExpr   = 3
	MockParserRULE_presetSliceExpr  = 4
	MockParserRULE_presetMapExpr    = 5
	MockParserRULE_presetSpecExpr   = 6
	MockParserRULE_numberExpr       = 7
	MockParserRULE_numberRangeExpr  = 8
	MockParserRULE_numberLengthExpr = 9
	MockParserRULE_stringExpr       = 10
	MockParserRULE_stringRange      = 11
	MockParserRULE_stringLength     = 12
	MockParserRULE_stringTemplate   = 13
	MockParserRULE_customArgSpec    = 14
	MockParserRULE_customArg        = 15
	MockParserRULE_constStatement   = 16
	MockParserRULE_prevConst        = 17
	MockParserRULE_afterConst       = 18
	MockParserRULE_lengthSpec       = 19
	MockParserRULE_rangeSpec        = 20
	MockParserRULE_ignoreStatement  = 21
)

// IMockExpressionContext is an interface to support dynamic dispatch.
type IMockExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMockExpressionContext differentiates from other interfaces.
	IsMockExpressionContext()
}

type MockExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMockExpressionContext() *MockExpressionContext {
	var p = new(MockExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_mockExpression
	return p
}

func (*MockExpressionContext) IsMockExpressionContext() {}

func NewMockExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MockExpressionContext {
	var p = new(MockExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_mockExpression

	return p
}

func (s *MockExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MockExpressionContext) NormalExpression() INormalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INormalExpressionContext)
}

func (s *MockExpressionContext) StringTemplate() IStringTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringTemplateContext)
}

func (s *MockExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MockExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MockExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterMockExpression(s)
	}
}

func (s *MockExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitMockExpression(s)
	}
}

func (s *MockExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitMockExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) MockExpression() (localctx IMockExpressionContext) {
	localctx = NewMockExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MockParserRULE_mockExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MockParserBOOL, MockParserSLICE, MockParserMAP, MockParserSPEC, MockParserRPAREN, MockParserCOMMA, MockParserBAR, MockParserIDENT:
		{
			p.SetState(44)
			p.NormalExpression()
		}

	case MockParserQUOTE_LEFT, MockParserVALUE:
		{
			p.SetState(45)
			p.StringTemplate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INormalExpressionContext is an interface to support dynamic dispatch.
type INormalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNormalExpressionContext differentiates from other interfaces.
	IsNormalExpressionContext()
}

type NormalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormalExpressionContext() *NormalExpressionContext {
	var p = new(NormalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_normalExpression
	return p
}

func (*NormalExpressionContext) IsNormalExpressionContext() {}

func NewNormalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalExpressionContext {
	var p = new(NormalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_normalExpression

	return p
}

func (s *NormalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalExpressionContext) PresetExpr() IPresetExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetExprContext)
}

func (s *NormalExpressionContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NormalExpressionContext) StringExpr() IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *NormalExpressionContext) CustomArgSpec() ICustomArgSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICustomArgSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICustomArgSpecContext)
}

func (s *NormalExpressionContext) ConstStatement() IConstStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstStatementContext)
}

func (s *NormalExpressionContext) IgnoreStatement() IIgnoreStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIgnoreStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIgnoreStatementContext)
}

func (s *NormalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterNormalExpression(s)
	}
}

func (s *NormalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitNormalExpression(s)
	}
}

func (s *NormalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitNormalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) NormalExpression() (localctx INormalExpressionContext) {
	localctx = NewNormalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MockParserRULE_normalExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.PresetExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.NumberExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.StringExpr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.CustomArgSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(52)
			p.ConstStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.IgnoreStatement()
		}

	}

	return localctx
}

// IPresetExprContext is an interface to support dynamic dispatch.
type IPresetExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetExprContext differentiates from other interfaces.
	IsPresetExprContext()
}

type PresetExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetExprContext() *PresetExprContext {
	var p = new(PresetExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_presetExpr
	return p
}

func (*PresetExprContext) IsPresetExprContext() {}

func NewPresetExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetExprContext {
	var p = new(PresetExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_presetExpr

	return p
}

func (s *PresetExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetExprContext) PresetBoolExpr() IPresetBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetBoolExprContext)
}

func (s *PresetExprContext) PresetSliceExpr() IPresetSliceExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetSliceExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetSliceExprContext)
}

func (s *PresetExprContext) PresetMapExpr() IPresetMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetMapExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetMapExprContext)
}

func (s *PresetExprContext) PresetSpecExpr() IPresetSpecExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetSpecExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetSpecExprContext)
}

func (s *PresetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterPresetExpr(s)
	}
}

func (s *PresetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitPresetExpr(s)
	}
}

func (s *PresetExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitPresetExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) PresetExpr() (localctx IPresetExprContext) {
	localctx = NewPresetExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MockParserRULE_presetExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MockParserRPAREN, MockParserCOMMA:
		p.EnterOuterAlt(localctx, 1)

	case MockParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.PresetBoolExpr()
		}

	case MockParserSLICE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.PresetSliceExpr()
		}

	case MockParserMAP:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)
			p.PresetMapExpr()
		}

	case MockParserSPEC:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.PresetSpecExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPresetBoolExprContext is an interface to support dynamic dispatch.
type IPresetBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetBoolExprContext differentiates from other interfaces.
	IsPresetBoolExprContext()
}

type PresetBoolExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetBoolExprContext() *PresetBoolExprContext {
	var p = new(PresetBoolExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_presetBoolExpr
	return p
}

func (*PresetBoolExprContext) IsPresetBoolExprContext() {}

func NewPresetBoolExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetBoolExprContext {
	var p = new(PresetBoolExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_presetBoolExpr

	return p
}

func (s *PresetBoolExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetBoolExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MockParserBOOL, 0)
}

func (s *PresetBoolExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MockParserASSIGN, 0)
}

func (s *PresetBoolExprContext) BOOL_VALUE() antlr.TerminalNode {
	return s.GetToken(MockParserBOOL_VALUE, 0)
}

func (s *PresetBoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetBoolExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetBoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterPresetBoolExpr(s)
	}
}

func (s *PresetBoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitPresetBoolExpr(s)
	}
}

func (s *PresetBoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitPresetBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) PresetBoolExpr() (localctx IPresetBoolExprContext) {
	localctx = NewPresetBoolExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MockParserRULE_presetBoolExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(MockParserBOOL)
	}
	{
		p.SetState(64)
		p.Match(MockParserASSIGN)
	}
	{
		p.SetState(65)
		p.Match(MockParserBOOL_VALUE)
	}

	return localctx
}

// IPresetSliceExprContext is an interface to support dynamic dispatch.
type IPresetSliceExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetSliceExprContext differentiates from other interfaces.
	IsPresetSliceExprContext()
}

type PresetSliceExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetSliceExprContext() *PresetSliceExprContext {
	var p = new(PresetSliceExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_presetSliceExpr
	return p
}

func (*PresetSliceExprContext) IsPresetSliceExprContext() {}

func NewPresetSliceExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetSliceExprContext {
	var p = new(PresetSliceExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_presetSliceExpr

	return p
}

func (s *PresetSliceExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetSliceExprContext) SLICE() antlr.TerminalNode {
	return s.GetToken(MockParserSLICE, 0)
}

func (s *PresetSliceExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserLPAREN, 0)
}

func (s *PresetSliceExprContext) MockExpression() IMockExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMockExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMockExpressionContext)
}

func (s *PresetSliceExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserRPAREN, 0)
}

func (s *PresetSliceExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MockParserNUMBER, 0)
}

func (s *PresetSliceExprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MockParserCOMMA, 0)
}

func (s *PresetSliceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetSliceExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetSliceExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterPresetSliceExpr(s)
	}
}

func (s *PresetSliceExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitPresetSliceExpr(s)
	}
}

func (s *PresetSliceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitPresetSliceExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) PresetSliceExpr() (localctx IPresetSliceExprContext) {
	localctx = NewPresetSliceExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MockParserRULE_presetSliceExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(MockParserSLICE)
	}
	{
		p.SetState(68)
		p.Match(MockParserLPAREN)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MockParserNUMBER {
		{
			p.SetState(69)
			p.Match(MockParserNUMBER)
		}
		{
			p.SetState(70)
			p.Match(MockParserCOMMA)
		}

	}
	{
		p.SetState(73)
		p.MockExpression()
	}
	{
		p.SetState(74)
		p.Match(MockParserRPAREN)
	}

	return localctx
}

// IPresetMapExprContext is an interface to support dynamic dispatch.
type IPresetMapExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetMapExprContext differentiates from other interfaces.
	IsPresetMapExprContext()
}

type PresetMapExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetMapExprContext() *PresetMapExprContext {
	var p = new(PresetMapExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_presetMapExpr
	return p
}

func (*PresetMapExprContext) IsPresetMapExprContext() {}

func NewPresetMapExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetMapExprContext {
	var p = new(PresetMapExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_presetMapExpr

	return p
}

func (s *PresetMapExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetMapExprContext) MAP() antlr.TerminalNode {
	return s.GetToken(MockParserMAP, 0)
}

func (s *PresetMapExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserLPAREN, 0)
}

func (s *PresetMapExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MockParserNUMBER, 0)
}

func (s *PresetMapExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MockParserCOMMA)
}

func (s *PresetMapExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MockParserCOMMA, i)
}

func (s *PresetMapExprContext) AllMockExpression() []IMockExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMockExpressionContext)(nil)).Elem())
	var tst = make([]IMockExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMockExpressionContext)
		}
	}

	return tst
}

func (s *PresetMapExprContext) MockExpression(i int) IMockExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMockExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMockExpressionContext)
}

func (s *PresetMapExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserRPAREN, 0)
}

func (s *PresetMapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetMapExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetMapExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterPresetMapExpr(s)
	}
}

func (s *PresetMapExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitPresetMapExpr(s)
	}
}

func (s *PresetMapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitPresetMapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) PresetMapExpr() (localctx IPresetMapExprContext) {
	localctx = NewPresetMapExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MockParserRULE_presetMapExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(MockParserMAP)
	}
	{
		p.SetState(77)
		p.Match(MockParserLPAREN)
	}
	{
		p.SetState(78)
		p.Match(MockParserNUMBER)
	}
	{
		p.SetState(79)
		p.Match(MockParserCOMMA)
	}
	{
		p.SetState(80)
		p.MockExpression()
	}
	{
		p.SetState(81)
		p.Match(MockParserCOMMA)
	}
	{
		p.SetState(82)
		p.MockExpression()
	}
	{
		p.SetState(83)
		p.Match(MockParserRPAREN)
	}

	return localctx
}

// IPresetSpecExprContext is an interface to support dynamic dispatch.
type IPresetSpecExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetSpecExprContext differentiates from other interfaces.
	IsPresetSpecExprContext()
}

type PresetSpecExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetSpecExprContext() *PresetSpecExprContext {
	var p = new(PresetSpecExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_presetSpecExpr
	return p
}

func (*PresetSpecExprContext) IsPresetSpecExprContext() {}

func NewPresetSpecExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetSpecExprContext {
	var p = new(PresetSpecExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_presetSpecExpr

	return p
}

func (s *PresetSpecExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetSpecExprContext) SPEC() antlr.TerminalNode {
	return s.GetToken(MockParserSPEC, 0)
}

func (s *PresetSpecExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MockParserASSIGN, 0)
}

func (s *PresetSpecExprContext) VALUE() antlr.TerminalNode {
	return s.GetToken(MockParserVALUE, 0)
}

func (s *PresetSpecExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetSpecExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetSpecExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterPresetSpecExpr(s)
	}
}

func (s *PresetSpecExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitPresetSpecExpr(s)
	}
}

func (s *PresetSpecExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitPresetSpecExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) PresetSpecExpr() (localctx IPresetSpecExprContext) {
	localctx = NewPresetSpecExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MockParserRULE_presetSpecExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(MockParserSPEC)
	}
	{
		p.SetState(86)
		p.Match(MockParserASSIGN)
	}
	{
		p.SetState(87)
		p.Match(MockParserVALUE)
	}

	return localctx
}

// INumberExprContext is an interface to support dynamic dispatch.
type INumberExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberExprContext differentiates from other interfaces.
	IsNumberExprContext()
}

type NumberExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberExprContext() *NumberExprContext {
	var p = new(NumberExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_numberExpr
	return p
}

func (*NumberExprContext) IsNumberExprContext() {}

func NewNumberExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberExprContext {
	var p = new(NumberExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_numberExpr

	return p
}

func (s *NumberExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberExprContext) NumberRangeExpr() INumberRangeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberRangeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberRangeExprContext)
}

func (s *NumberExprContext) NumberLengthExpr() INumberLengthExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLengthExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLengthExprContext)
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterNumberExpr(s)
	}
}

func (s *NumberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitNumberExpr(s)
	}
}

func (s *NumberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitNumberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) NumberExpr() (localctx INumberExprContext) {
	localctx = NewNumberExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MockParserRULE_numberExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.NumberRangeExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.NumberLengthExpr()
		}

	}

	return localctx
}

// INumberRangeExprContext is an interface to support dynamic dispatch.
type INumberRangeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberRangeExprContext differentiates from other interfaces.
	IsNumberRangeExprContext()
}

type NumberRangeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberRangeExprContext() *NumberRangeExprContext {
	var p = new(NumberRangeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_numberRangeExpr
	return p
}

func (*NumberRangeExprContext) IsNumberRangeExprContext() {}

func NewNumberRangeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberRangeExprContext {
	var p = new(NumberRangeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_numberRangeExpr

	return p
}

func (s *NumberRangeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberRangeExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MockParserIDENT, 0)
}

func (s *NumberRangeExprContext) RangeSpec() IRangeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeSpecContext)
}

func (s *NumberRangeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberRangeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberRangeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterNumberRangeExpr(s)
	}
}

func (s *NumberRangeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitNumberRangeExpr(s)
	}
}

func (s *NumberRangeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitNumberRangeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) NumberRangeExpr() (localctx INumberRangeExprContext) {
	localctx = NewNumberRangeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MockParserRULE_numberRangeExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(MockParserIDENT)
	}
	{
		p.SetState(94)
		p.RangeSpec()
	}

	return localctx
}

// INumberLengthExprContext is an interface to support dynamic dispatch.
type INumberLengthExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLengthExprContext differentiates from other interfaces.
	IsNumberLengthExprContext()
}

type NumberLengthExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLengthExprContext() *NumberLengthExprContext {
	var p = new(NumberLengthExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_numberLengthExpr
	return p
}

func (*NumberLengthExprContext) IsNumberLengthExprContext() {}

func NewNumberLengthExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLengthExprContext {
	var p = new(NumberLengthExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_numberLengthExpr

	return p
}

func (s *NumberLengthExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLengthExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MockParserIDENT, 0)
}

func (s *NumberLengthExprContext) LengthSpec() ILengthSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthSpecContext)
}

func (s *NumberLengthExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLengthExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLengthExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterNumberLengthExpr(s)
	}
}

func (s *NumberLengthExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitNumberLengthExpr(s)
	}
}

func (s *NumberLengthExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitNumberLengthExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) NumberLengthExpr() (localctx INumberLengthExprContext) {
	localctx = NewNumberLengthExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MockParserRULE_numberLengthExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(MockParserIDENT)
	}
	{
		p.SetState(97)
		p.LengthSpec()
	}

	return localctx
}

// IStringExprContext is an interface to support dynamic dispatch.
type IStringExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringExprContext differentiates from other interfaces.
	IsStringExprContext()
}

type StringExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringExprContext() *StringExprContext {
	var p = new(StringExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_stringExpr
	return p
}

func (*StringExprContext) IsStringExprContext() {}

func NewStringExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringExprContext {
	var p = new(StringExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_stringExpr

	return p
}

func (s *StringExprContext) GetParser() antlr.Parser { return s.parser }

func (s *StringExprContext) StringRange() IStringRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringRangeContext)
}

func (s *StringExprContext) StringLength() IStringLengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLengthContext)
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) StringExpr() (localctx IStringExprContext) {
	localctx = NewStringExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MockParserRULE_stringExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.StringRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.StringLength()
		}

	}

	return localctx
}

// IStringRangeContext is an interface to support dynamic dispatch.
type IStringRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringRangeContext differentiates from other interfaces.
	IsStringRangeContext()
}

type StringRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringRangeContext() *StringRangeContext {
	var p = new(StringRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_stringRange
	return p
}

func (*StringRangeContext) IsStringRangeContext() {}

func NewStringRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringRangeContext {
	var p = new(StringRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_stringRange

	return p
}

func (s *StringRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *StringRangeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MockParserIDENT, 0)
}

func (s *StringRangeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserLPAREN, 0)
}

func (s *StringRangeContext) LANG() antlr.TerminalNode {
	return s.GetToken(MockParserLANG, 0)
}

func (s *StringRangeContext) TO() antlr.TerminalNode {
	return s.GetToken(MockParserTO, 0)
}

func (s *StringRangeContext) RangeSpec() IRangeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeSpecContext)
}

func (s *StringRangeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserRPAREN, 0)
}

func (s *StringRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterStringRange(s)
	}
}

func (s *StringRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitStringRange(s)
	}
}

func (s *StringRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitStringRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) StringRange() (localctx IStringRangeContext) {
	localctx = NewStringRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MockParserRULE_stringRange)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(MockParserIDENT)
	}
	{
		p.SetState(104)
		p.Match(MockParserLPAREN)
	}
	{
		p.SetState(105)
		p.Match(MockParserLANG)
	}
	{
		p.SetState(106)
		p.Match(MockParserTO)
	}
	{
		p.SetState(107)
		p.RangeSpec()
	}
	{
		p.SetState(108)
		p.Match(MockParserRPAREN)
	}

	return localctx
}

// IStringLengthContext is an interface to support dynamic dispatch.
type IStringLengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLengthContext differentiates from other interfaces.
	IsStringLengthContext()
}

type StringLengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLengthContext() *StringLengthContext {
	var p = new(StringLengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_stringLength
	return p
}

func (*StringLengthContext) IsStringLengthContext() {}

func NewStringLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLengthContext {
	var p = new(StringLengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_stringLength

	return p
}

func (s *StringLengthContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLengthContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MockParserIDENT, 0)
}

func (s *StringLengthContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserLPAREN, 0)
}

func (s *StringLengthContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserRPAREN, 0)
}

func (s *StringLengthContext) LANG() antlr.TerminalNode {
	return s.GetToken(MockParserLANG, 0)
}

func (s *StringLengthContext) TO() antlr.TerminalNode {
	return s.GetToken(MockParserTO, 0)
}

func (s *StringLengthContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MockParserNUMBER, 0)
}

func (s *StringLengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterStringLength(s)
	}
}

func (s *StringLengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitStringLength(s)
	}
}

func (s *StringLengthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitStringLength(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) StringLength() (localctx IStringLengthContext) {
	localctx = NewStringLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MockParserRULE_stringLength)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(MockParserIDENT)
	}
	{
		p.SetState(111)
		p.Match(MockParserLPAREN)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MockParserLANG {
		{
			p.SetState(112)
			p.Match(MockParserLANG)
		}

	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MockParserTO {
		{
			p.SetState(115)
			p.Match(MockParserTO)
		}
		{
			p.SetState(116)
			p.Match(MockParserNUMBER)
		}

	}
	{
		p.SetState(119)
		p.Match(MockParserRPAREN)
	}

	return localctx
}

// IStringTemplateContext is an interface to support dynamic dispatch.
type IStringTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringTemplateContext differentiates from other interfaces.
	IsStringTemplateContext()
}

type StringTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringTemplateContext() *StringTemplateContext {
	var p = new(StringTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_stringTemplate
	return p
}

func (*StringTemplateContext) IsStringTemplateContext() {}

func NewStringTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTemplateContext {
	var p = new(StringTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_stringTemplate

	return p
}

func (s *StringTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTemplateContext) QUOTE_LEFT() antlr.TerminalNode {
	return s.GetToken(MockParserQUOTE_LEFT, 0)
}

func (s *StringTemplateContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MockParserIDENT, 0)
}

func (s *StringTemplateContext) QUOTE_RIGHT() antlr.TerminalNode {
	return s.GetToken(MockParserQUOTE_RIGHT, 0)
}

func (s *StringTemplateContext) PrevConst() IPrevConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrevConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrevConstContext)
}

func (s *StringTemplateContext) AfterConst() IAfterConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAfterConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAfterConstContext)
}

func (s *StringTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterStringTemplate(s)
	}
}

func (s *StringTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitStringTemplate(s)
	}
}

func (s *StringTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitStringTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) StringTemplate() (localctx IStringTemplateContext) {
	localctx = NewStringTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MockParserRULE_stringTemplate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MockParserVALUE {
		{
			p.SetState(121)
			p.PrevConst()
		}

	}
	{
		p.SetState(124)
		p.Match(MockParserQUOTE_LEFT)
	}
	{
		p.SetState(125)
		p.Match(MockParserIDENT)
	}
	{
		p.SetState(126)
		p.Match(MockParserQUOTE_RIGHT)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MockParserVALUE {
		{
			p.SetState(127)
			p.AfterConst()
		}

	}

	return localctx
}

// ICustomArgSpecContext is an interface to support dynamic dispatch.
type ICustomArgSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCustomArgSpecContext differentiates from other interfaces.
	IsCustomArgSpecContext()
}

type CustomArgSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustomArgSpecContext() *CustomArgSpecContext {
	var p = new(CustomArgSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_customArgSpec
	return p
}

func (*CustomArgSpecContext) IsCustomArgSpecContext() {}

func NewCustomArgSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CustomArgSpecContext {
	var p = new(CustomArgSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_customArgSpec

	return p
}

func (s *CustomArgSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CustomArgSpecContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MockParserIDENT, 0)
}

func (s *CustomArgSpecContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserLPAREN, 0)
}

func (s *CustomArgSpecContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserRPAREN, 0)
}

func (s *CustomArgSpecContext) AllCustomArg() []ICustomArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICustomArgContext)(nil)).Elem())
	var tst = make([]ICustomArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICustomArgContext)
		}
	}

	return tst
}

func (s *CustomArgSpecContext) CustomArg(i int) ICustomArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICustomArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICustomArgContext)
}

func (s *CustomArgSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomArgSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CustomArgSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterCustomArgSpec(s)
	}
}

func (s *CustomArgSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitCustomArgSpec(s)
	}
}

func (s *CustomArgSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitCustomArgSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) CustomArgSpec() (localctx ICustomArgSpecContext) {
	localctx = NewCustomArgSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MockParserRULE_customArgSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(MockParserIDENT)
	}
	{
		p.SetState(131)
		p.Match(MockParserLPAREN)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MockParserVALUE {
		{
			p.SetState(132)
			p.CustomArg()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(137)
		p.Match(MockParserRPAREN)
	}

	return localctx
}

// ICustomArgContext is an interface to support dynamic dispatch.
type ICustomArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCustomArgContext differentiates from other interfaces.
	IsCustomArgContext()
}

type CustomArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustomArgContext() *CustomArgContext {
	var p = new(CustomArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_customArg
	return p
}

func (*CustomArgContext) IsCustomArgContext() {}

func NewCustomArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CustomArgContext {
	var p = new(CustomArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_customArg

	return p
}

func (s *CustomArgContext) GetParser() antlr.Parser { return s.parser }

func (s *CustomArgContext) VALUE() antlr.TerminalNode {
	return s.GetToken(MockParserVALUE, 0)
}

func (s *CustomArgContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MockParserCOMMA, 0)
}

func (s *CustomArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CustomArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterCustomArg(s)
	}
}

func (s *CustomArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitCustomArg(s)
	}
}

func (s *CustomArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitCustomArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) CustomArg() (localctx ICustomArgContext) {
	localctx = NewCustomArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MockParserRULE_customArg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(MockParserVALUE)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MockParserCOMMA {
		{
			p.SetState(140)
			p.Match(MockParserCOMMA)
		}

	}

	return localctx
}

// IConstStatementContext is an interface to support dynamic dispatch.
type IConstStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstStatementContext differentiates from other interfaces.
	IsConstStatementContext()
}

type ConstStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstStatementContext() *ConstStatementContext {
	var p = new(ConstStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_constStatement
	return p
}

func (*ConstStatementContext) IsConstStatementContext() {}

func NewConstStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstStatementContext {
	var p = new(ConstStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_constStatement

	return p
}

func (s *ConstStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstStatementContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MockParserIDENT, 0)
}

func (s *ConstStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterConstStatement(s)
	}
}

func (s *ConstStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitConstStatement(s)
	}
}

func (s *ConstStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitConstStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) ConstStatement() (localctx IConstStatementContext) {
	localctx = NewConstStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MockParserRULE_constStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(MockParserIDENT)
	}

	return localctx
}

// IPrevConstContext is an interface to support dynamic dispatch.
type IPrevConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrevConstContext differentiates from other interfaces.
	IsPrevConstContext()
}

type PrevConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrevConstContext() *PrevConstContext {
	var p = new(PrevConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_prevConst
	return p
}

func (*PrevConstContext) IsPrevConstContext() {}

func NewPrevConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrevConstContext {
	var p = new(PrevConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_prevConst

	return p
}

func (s *PrevConstContext) GetParser() antlr.Parser { return s.parser }

func (s *PrevConstContext) VALUE() antlr.TerminalNode {
	return s.GetToken(MockParserVALUE, 0)
}

func (s *PrevConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrevConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrevConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterPrevConst(s)
	}
}

func (s *PrevConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitPrevConst(s)
	}
}

func (s *PrevConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitPrevConst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) PrevConst() (localctx IPrevConstContext) {
	localctx = NewPrevConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MockParserRULE_prevConst)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(MockParserVALUE)
	}

	return localctx
}

// IAfterConstContext is an interface to support dynamic dispatch.
type IAfterConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAfterConstContext differentiates from other interfaces.
	IsAfterConstContext()
}

type AfterConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAfterConstContext() *AfterConstContext {
	var p = new(AfterConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_afterConst
	return p
}

func (*AfterConstContext) IsAfterConstContext() {}

func NewAfterConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AfterConstContext {
	var p = new(AfterConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_afterConst

	return p
}

func (s *AfterConstContext) GetParser() antlr.Parser { return s.parser }

func (s *AfterConstContext) VALUE() antlr.TerminalNode {
	return s.GetToken(MockParserVALUE, 0)
}

func (s *AfterConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AfterConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterAfterConst(s)
	}
}

func (s *AfterConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitAfterConst(s)
	}
}

func (s *AfterConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitAfterConst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) AfterConst() (localctx IAfterConstContext) {
	localctx = NewAfterConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MockParserRULE_afterConst)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(MockParserVALUE)
	}

	return localctx
}

// ILengthSpecContext is an interface to support dynamic dispatch.
type ILengthSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthSpecContext differentiates from other interfaces.
	IsLengthSpecContext()
}

type LengthSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthSpecContext() *LengthSpecContext {
	var p = new(LengthSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_lengthSpec
	return p
}

func (*LengthSpecContext) IsLengthSpecContext() {}

func NewLengthSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthSpecContext {
	var p = new(LengthSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_lengthSpec

	return p
}

func (s *LengthSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthSpecContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserLPAREN, 0)
}

func (s *LengthSpecContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MockParserNUMBER, 0)
}

func (s *LengthSpecContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MockParserRPAREN, 0)
}

func (s *LengthSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterLengthSpec(s)
	}
}

func (s *LengthSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitLengthSpec(s)
	}
}

func (s *LengthSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitLengthSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) LengthSpec() (localctx ILengthSpecContext) {
	localctx = NewLengthSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MockParserRULE_lengthSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(MockParserLPAREN)
	}
	{
		p.SetState(150)
		p.Match(MockParserNUMBER)
	}
	{
		p.SetState(151)
		p.Match(MockParserRPAREN)
	}

	return localctx
}

// IRangeSpecContext is an interface to support dynamic dispatch.
type IRangeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeSpecContext differentiates from other interfaces.
	IsRangeSpecContext()
}

type RangeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeSpecContext() *RangeSpecContext {
	var p = new(RangeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_rangeSpec
	return p
}

func (*RangeSpecContext) IsRangeSpecContext() {}

func NewRangeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeSpecContext {
	var p = new(RangeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_rangeSpec

	return p
}

func (s *RangeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeSpecContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MockParserLBRACK, 0)
}

func (s *RangeSpecContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(MockParserNUMBER)
}

func (s *RangeSpecContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(MockParserNUMBER, i)
}

func (s *RangeSpecContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MockParserCOMMA, 0)
}

func (s *RangeSpecContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MockParserRBRACK, 0)
}

func (s *RangeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterRangeSpec(s)
	}
}

func (s *RangeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitRangeSpec(s)
	}
}

func (s *RangeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitRangeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) RangeSpec() (localctx IRangeSpecContext) {
	localctx = NewRangeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MockParserRULE_rangeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(MockParserLBRACK)
	}
	{
		p.SetState(154)
		p.Match(MockParserNUMBER)
	}
	{
		p.SetState(155)
		p.Match(MockParserCOMMA)
	}
	{
		p.SetState(156)
		p.Match(MockParserNUMBER)
	}
	{
		p.SetState(157)
		p.Match(MockParserRBRACK)
	}

	return localctx
}

// IIgnoreStatementContext is an interface to support dynamic dispatch.
type IIgnoreStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIgnoreStatementContext differentiates from other interfaces.
	IsIgnoreStatementContext()
}

type IgnoreStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIgnoreStatementContext() *IgnoreStatementContext {
	var p = new(IgnoreStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MockParserRULE_ignoreStatement
	return p
}

func (*IgnoreStatementContext) IsIgnoreStatementContext() {}

func NewIgnoreStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoreStatementContext {
	var p = new(IgnoreStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MockParserRULE_ignoreStatement

	return p
}

func (s *IgnoreStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IgnoreStatementContext) BAR() antlr.TerminalNode {
	return s.GetToken(MockParserBAR, 0)
}

func (s *IgnoreStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgnoreStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IgnoreStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.EnterIgnoreStatement(s)
	}
}

func (s *IgnoreStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MockParserListener); ok {
		listenerT.ExitIgnoreStatement(s)
	}
}

func (s *IgnoreStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MockParserVisitor:
		return t.VisitIgnoreStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MockParser) IgnoreStatement() (localctx IIgnoreStatementContext) {
	localctx = NewIgnoreStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MockParserRULE_ignoreStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(MockParserBAR)
	}

	return localctx
}
