// Code generated from /Users/anqiansong/goland/go/go-zero/anqiansong/mock/grammar/MockParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MockParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMockParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMockParserVisitor) VisitMockExpression(ctx *MockExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitNormalExpression(ctx *NormalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitPresetExpr(ctx *PresetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitPresetBoolExpr(ctx *PresetBoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitPresetSliceExpr(ctx *PresetSliceExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitPresetMapExpr(ctx *PresetMapExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitPresetSpecExpr(ctx *PresetSpecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitNumberExpr(ctx *NumberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitNumberRangeExpr(ctx *NumberRangeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitNumberLengthExpr(ctx *NumberLengthExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitStringRange(ctx *StringRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitStringLength(ctx *StringLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitStringTemplate(ctx *StringTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitCustomArgSpec(ctx *CustomArgSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitCustomArg(ctx *CustomArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitConstStatement(ctx *ConstStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitPrevConst(ctx *PrevConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitAfterConst(ctx *AfterConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitLengthSpec(ctx *LengthSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitRangeSpec(ctx *RangeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMockParserVisitor) VisitIgnoreStatement(ctx *IgnoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}
