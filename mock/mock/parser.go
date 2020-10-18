package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tal-tech/go-zero/core/logx"
	parser "github.com/tal-tech/go-zero/mock/grammar"
	"github.com/tal-tech/go-zero/mock/render"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"reflect"
	"strconv"
	"strings"
)

const (
	UnKnown Type = iota
	TypeInt
	TypeUInt
	TypeFloat
	TypeBool
	TypeString
	TypeSlice
	TypeMap
	TypeInterface
)

var (
	lengthErr = errors.New("length can not be less than 0")
	rangeErr  = errors.New("range start can not be gather than the end")
)

type (
	Type        int
	SliceLooper func(count uint64, itemExpression string) error
	MapLooper   func(count uint64, keyExpression, valueExpression string) error
	Value       struct {
		reflect.Value
		Type        Type
		SliceLooper SliceLooper
		MapLooper   MapLooper
	}
	Parser struct {
		parser.BaseMockParserVisitor
		antlr.DefaultErrorListener
		invoker *InvokeMatcher
		value   *Value
	}
)

func NewParser(invoker *InvokeMatcher) *Parser {
	instance := &Parser{
		invoker: invoker,
	}
	return instance
}

func (p *Parser) Parse(expression string, value *Value) {
	defer func() {
		v := recover()
		if v != nil {
			logx.Error(v)
		}
	}()
	p.value = value
	inputStream := antlr.NewInputStream(expression)
	lexer := parser.NewMockLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	ps := parser.NewMockParser(tokens)
	ps.MockExpression().Accept(p)
}

func (p *Parser) visitRules(nodes ...antlr.RuleNode) interface{} {
	for _, node := range nodes {
		if node == nil {
			continue
		}
		v := node.Accept(p)
		if v != nil {
			return v
		}
	}
	return nil
}

func (p *Parser) VisitMockExpression(ctx *parser.MockExpressionContext) interface{} {
	return p.visitRules(ctx.NormalExpression(), ctx.StringTemplate())
}

// Visit a parse tree produced by MockParser#normalPression.
func (p *Parser) VisitNormalExpression(ctx *parser.NormalExpressionContext) interface{} {
	return p.visitRules(ctx.PresetExpr(), ctx.NumberExpr(), ctx.StringExpr(), ctx.ConstStatement(), ctx.CustomArgSpec())
}

// Visit a parse tree produced by MockParser#presetExpr.
func (p *Parser) VisitPresetExpr(ctx *parser.PresetExprContext) interface{} {
	return p.visitRules(ctx.PresetBoolExpr(), ctx.PresetSliceExpr(), ctx.PresetMapExpr(), ctx.PresetSpecExpr())
}

// Visit a parse tree produced by MockParser#presetBoolExpr.
func (p *Parser) VisitPresetBoolExpr(ctx *parser.PresetBoolExprContext) interface{} {
	if p.value.Type != TypeBool {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "bool"))
	}
	value := p.getText(ctx.BOOL_VALUE())
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	p.value.Value.SetBool(boolValue)
	return nil
}

// Visit a parse tree produced by MockParser#presetSliceExpr.
func (p *Parser) VisitPresetSliceExpr(ctx *parser.PresetSliceExprContext) interface{} {
	if p.value.Type != TypeSlice {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "slice"))
	}
	lengthText := p.getText(ctx.NUMBER())
	childExpression := ctx.MockExpression().GetText()
	length, _ := strconv.ParseUint(lengthText, 10, 64)
	if length < 0 {
		panic(lengthErr)
	}
	return p.value.SliceLooper(length, childExpression)
}

// Visit a parse tree produced by MockParser#presetMapExpr.
func (p *Parser) VisitPresetMapExpr(ctx *parser.PresetMapExprContext) interface{} {
	if p.value.Type != TypeMap {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "map"))
	}
	lengthText := p.getText(ctx.NUMBER())
	keyExpression := ctx.MockExpression(0).GetText()
	valueExpression := ctx.MockExpression(1).GetText()
	length, _ := strconv.ParseUint(lengthText, 10, 64)
	if length < 0 {
		panic(lengthErr)
	}
	return p.value.MapLooper(length, keyExpression, valueExpression)
}

// Visit a parse tree produced by MockParser#presetSpecExpr.
func (p *Parser) VisitPresetSpecExpr(ctx *parser.PresetSpecExprContext) interface{} {
	text := p.getText(ctx.VALUE())
	length := len(text)
	valueString := text[1 : length-1]
	switch p.value.Type {
	case TypeInt:
		vInt, err := strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			panic(err)
		}
		p.value.SetInt(vInt)
	case TypeUInt:
		vInt, err := strconv.ParseUint(valueString, 10, 64)
		if err != nil {
			panic(err)
		}
		p.value.SetUint(vInt)
	case TypeFloat:
		vFloat, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			panic(err)
		}
		p.value.SetFloat(vFloat)
	case TypeBool:
		vBool, err := strconv.ParseBool(valueString)
		if err != nil {
			panic(err)
		}
		p.value.SetBool(vBool)
	case TypeString, TypeInterface:
		p.value.SetString(valueString)
	}
	return nil
}

// Visit a parse tree produced by MockParser#numberExpr.
func (p *Parser) VisitNumberExpr(ctx *parser.NumberExprContext) interface{} {
	p.visitRules(ctx.NumberRangeExpr(), ctx.NumberLengthExpr())
	return nil
}

// Visit a parse tree produced by MockParser#numberRangeExpr.
func (p *Parser) VisitNumberRangeExpr(ctx *parser.NumberRangeExprContext) interface{} {
	if p.value.Type != TypeInt && p.value.Type != TypeUInt && p.value.Type != TypeFloat {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "number"))
	}
	expressionName := p.getText(ctx.IDENT())
	start, end := p.getRange(ctx.RangeSpec())
	if start >= end {
		panic(rangeErr)
	}
	invoke, ok := p.invoker.Match(expressionName)
	if !ok {
		panic(fmt.Errorf("invoke function %s not found", stringx.From(expressionName).UnTitle()))
	}

	vString := invoke(render.NewMockContext(expressionName, render.DefaultLang, 0, [2]int64{start, end}))
	switch vString.(type) {
	case uint8, uint16, uint32, uint64, uintptr:
		vInt, _ := strconv.ParseUint(fmt.Sprintf("%v", vString), 10, 64)
		p.value.SetUint(vInt)
	case int8, int16, int32, int64, int:
		vInt, _ := strconv.ParseInt(fmt.Sprintf("%v", vString), 10, 64)
		p.value.SetInt(vInt)
	case float32, float64:
		vFloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", vString), 64)
		p.value.SetFloat(vFloat)
	default:
		panic(fmt.Errorf("expected number value,but found %v", vString))
	}
	return nil
}

// Visit a parse tree produced by MockParser#numberLengthExpr.
func (p *Parser) VisitNumberLengthExpr(ctx *parser.NumberLengthExprContext) interface{} {
	if p.value.Type != TypeInt && p.value.Type != TypeUInt && p.value.Type != TypeFloat {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "number"))
	}
	expressionName := p.getText(ctx.IDENT())
	length := p.getLength(ctx.LengthSpec())
	if length < 0 {
		panic(lengthErr)
	}
	invoke, ok := p.invoker.Match(expressionName)
	if !ok {
		panic(fmt.Errorf("invoke function %s not found", stringx.From(expressionName).UnTitle()))
	}

	vString := invoke(render.NewMockContext(expressionName, render.DefaultLang, length, [2]int64{}))
	switch vString.(type) {
	case uint8, uint16, uint32, uint64, uintptr:
		vInt, _ := strconv.ParseUint(fmt.Sprintf("%v", vString), 10, 64)
		p.value.SetUint(vInt)
	case int8, int16, int32, int64, int:
		vInt, _ := strconv.ParseInt(fmt.Sprintf("%v", vString), 10, 64)
		p.value.SetInt(vInt)
	case float32, float64:
		vFloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", vString), 64)
		p.value.SetFloat(vFloat)
	default:
		panic(fmt.Errorf("expected number value,but found %v", reflect.TypeOf(vString).Name()))
	}
	return nil
}

// Visit a parse tree produced by MockParser#stringExpr.
func (p *Parser) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	p.visitRules(ctx.StringRange(), ctx.StringLength())
	return nil
}

// Visit a parse tree produced by MockParser#stringRange.
func (p *Parser) VisitStringRange(ctx *parser.StringRangeContext) interface{} {
	if p.value.Type != TypeString {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "string"))
	}
	expressionName := p.getText(ctx.IDENT())
	lang := p.getText(ctx.LANG())
	start, end := p.getRange(ctx.RangeSpec())
	if start >= end {
		panic(rangeErr)
	}
	l, err := render.InitLangByString(lang)
	if err != nil {
		panic(err)
	}
	invoke, ok := p.invoker.Match(expressionName)
	if !ok {
		panic(fmt.Errorf("invoke function %s not found", stringx.From(expressionName).UnTitle()))
	}

	v := invoke(render.NewMockContext(expressionName, l, render.DefaultLength, [2]int64{start, end}))
	vString, ok := v.(string)
	if !ok {
		panic(fmt.Errorf("expected string value,but found %v", reflect.TypeOf(v).Name()))
	}
	p.value.SetString(vString)
	return nil
}

// Visit a parse tree produced by MockParser#stringLength.
func (p *Parser) VisitStringLength(ctx *parser.StringLengthContext) interface{} {
	if p.value.Type != TypeString {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "string"))
	}
	lang := p.getText(ctx.LANG())
	expressionName := ctx.IDENT().GetText()
	length, _ := strconv.ParseUint(p.getText(ctx.NUMBER()), 10, 64)
	if length < 0 {
		panic(lengthErr)
	}
	l, err := render.InitLangByString(lang)
	if err != nil {
		panic(err)
	}
	invoke, ok := p.invoker.Match(expressionName)
	if !ok {
		panic(fmt.Errorf("invoke function %s not found", stringx.From(expressionName).UnTitle()))
	}

	v := invoke(render.NewMockContext(expressionName, l, length, render.DefaultRange))
	p.value.SetString(fmt.Sprintf("%v", v))
	return nil
}

// Visit a parse tree produced by MockParser#stringTemplate.
func (p *Parser) VisitStringTemplate(ctx *parser.StringTemplateContext) interface{} {
	if p.value.Type != TypeString {
		panic(p.expressionCheck(ctx.GetText(), p.value.Type, "string"))
	}
	expressionName := p.getText(ctx.IDENT())
	var prevConstText, afterConstText string
	if ctx.PrevConst() != nil {
		prevConstText = ctx.PrevConst().GetText()
	}
	if ctx.AfterConst() != nil {
		afterConstText = ctx.AfterConst().GetText()
	}
	prevConstText = p.trimValueFlag(prevConstText)
	afterConstText = p.trimValueFlag(afterConstText)

	t := render.NewDefaultTemplate(prevConstText, afterConstText, expressionName)
	v, err := t.Render(func(invokeFuncName string) (interface{}, error) {
		invoke, ok := p.invoker.Match(invokeFuncName)
		if !ok {
			return "", fmt.Errorf("invoke function %s not found", stringx.From(expressionName).UnTitle())
		}
		return invoke(render.EmptyMockContext), nil
	})
	if err != nil {
		panic(err)
	}
	p.value.SetString(fmt.Sprintf("%v", v))
	return nil
}

// Visit a parse tree produced by MockParser#customArgSpec.
func (p *Parser) VisitCustomArgSpec(ctx *parser.CustomArgSpecContext) interface{} {
	expressionName := p.getText(ctx.IDENT())
	ctxList := ctx.AllCustomArg()
	args := make([]interface{}, 0)
	for _, ctx := range ctxList {
		if ctx == nil {
			continue
		}
		text := p.getCustomArg(ctx)
		text = p.trimValueFlag(text)
		args = append(args, text)
	}
	invoke, ok := p.invoker.Match(expressionName)
	if !ok {
		panic(fmt.Errorf("invoke function %s not found", stringx.From(expressionName).UnTitle()))
	}

	v := invoke(render.NewMockContext(expressionName, render.DefaultLang, render.DefaultLength, render.DefaultRange, args...))
	p.setValue(v)
	return nil
}

// Visit a parse tree produced by MockParser#constStatement.
func (p *Parser) VisitConstStatement(ctx *parser.ConstStatementContext) interface{} {
	expressionName := p.getText(ctx.IDENT())
	invoke, ok := p.invoker.Match(expressionName)
	if !ok {
		panic(fmt.Errorf("%v: invoke function %s not found", ctx.GetText(), stringx.From(expressionName).UnTitle()))
	}

	v := invoke(render.EmptyMockContext)
	p.setValue(v)
	return nil
}

func (p *Parser) setValue(v interface{}) {
	vString := fmt.Sprintf("%v", v)
	number := json.Number(vString)
	switch p.value.Type {
	case TypeInt, TypeUInt:
		vInt, err := number.Int64()
		if err != nil {
			panic(err)
		}
		p.value.SetInt(vInt)
	case TypeFloat:
		vFloat, err := number.Float64()
		if err != nil {
			panic(err)
		}
		p.value.SetFloat(vFloat)
	case TypeBool:
		vBool, err := strconv.ParseBool(vString)
		if err != nil {
			panic(err)
		}
		p.value.SetBool(vBool)
	case TypeSlice, TypeMap:
		// ignore
	case TypeString, TypeInterface:
		p.value.SetString(vString)
	}
}

//func (p *Parser) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, msg string, _ antlr.RecognitionException) {
//	if len(msg) > 0 {
//		logx.Errorf("%+v", msg)
//	}
//}

func (p *Parser) getText(node antlr.TerminalNode) string {
	if node == nil {
		return ""
	}
	return node.GetText()
}
func (p *Parser) getLength(ctx parser.ILengthSpecContext) uint64 {
	switch v := ctx.(type) {
	case *parser.LengthSpecContext:
		text := v.NUMBER().GetText()
		length, _ := strconv.ParseUint(text, 10, 64)
		return length
	default:
		return 0
	}
}

func (p *Parser) getCustomArg(ctx parser.ICustomArgContext) string {
	switch v := ctx.(type) {
	case *parser.CustomArgContext:
		return p.getText(v.VALUE())
	default:
		return ""
	}
}

func (p *Parser) getRange(ctx parser.IRangeSpecContext) (int64, int64) {
	switch v := ctx.(type) {
	case *parser.RangeSpecContext:
		startText := v.NUMBER(0).GetText()
		endText := v.NUMBER(1).GetText()
		start, _ := strconv.ParseInt(startText, 10, 64)
		end, _ := strconv.ParseInt(endText, 10, 64)
		return start, end
	default:
		return 0, 0
	}
}

func (p *Parser) expressionCheck(expression string, tp Type, expressionType string) error {
	switch tp {
	case TypeInt, TypeUInt, TypeFloat:
		return fmt.Errorf("%v: expected number expression,but found %v expression", expression, expressionType)
	case TypeBool:
		return fmt.Errorf("%v: expected bool	 expression,but found %v expression", expression, expressionType)
	case TypeString:
		return fmt.Errorf("%v: expected string expression,but found %v expression", expression, expressionType)
	case TypeSlice:
		return fmt.Errorf("%v: expected slice expression,but found %v expression", expression, expressionType)
	case TypeMap:
		return fmt.Errorf("%v: expected map expression,but found %v expression", expression, expressionType)
	case TypeInterface:
		return fmt.Errorf("%v: expected interface expression,but found %v expression", expression, expressionType)
	}
	return nil
}

func (p *Parser) trimValueFlag(text string) string {
	if strings.Contains(text, "{") {
		text = strings.ReplaceAll(text, "{", "")
	}
	if strings.Contains(text, "}") {
		text = strings.ReplaceAll(text, "}", "")
	}
	return text
}
