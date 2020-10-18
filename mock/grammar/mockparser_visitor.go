// Code generated from /Users/anqiansong/goland/go/go-zero/anqiansong/mock/grammar/MockParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MockParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MockParser.
type MockParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MockParser#mockExpression.
	VisitMockExpression(ctx *MockExpressionContext) interface{}

	// Visit a parse tree produced by MockParser#normalExpression.
	VisitNormalExpression(ctx *NormalExpressionContext) interface{}

	// Visit a parse tree produced by MockParser#presetExpr.
	VisitPresetExpr(ctx *PresetExprContext) interface{}

	// Visit a parse tree produced by MockParser#presetBoolExpr.
	VisitPresetBoolExpr(ctx *PresetBoolExprContext) interface{}

	// Visit a parse tree produced by MockParser#presetSliceExpr.
	VisitPresetSliceExpr(ctx *PresetSliceExprContext) interface{}

	// Visit a parse tree produced by MockParser#presetMapExpr.
	VisitPresetMapExpr(ctx *PresetMapExprContext) interface{}

	// Visit a parse tree produced by MockParser#presetSpecExpr.
	VisitPresetSpecExpr(ctx *PresetSpecExprContext) interface{}

	// Visit a parse tree produced by MockParser#numberExpr.
	VisitNumberExpr(ctx *NumberExprContext) interface{}

	// Visit a parse tree produced by MockParser#numberRangeExpr.
	VisitNumberRangeExpr(ctx *NumberRangeExprContext) interface{}

	// Visit a parse tree produced by MockParser#numberLengthExpr.
	VisitNumberLengthExpr(ctx *NumberLengthExprContext) interface{}

	// Visit a parse tree produced by MockParser#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by MockParser#stringRange.
	VisitStringRange(ctx *StringRangeContext) interface{}

	// Visit a parse tree produced by MockParser#stringLength.
	VisitStringLength(ctx *StringLengthContext) interface{}

	// Visit a parse tree produced by MockParser#stringTemplate.
	VisitStringTemplate(ctx *StringTemplateContext) interface{}

	// Visit a parse tree produced by MockParser#customArgSpec.
	VisitCustomArgSpec(ctx *CustomArgSpecContext) interface{}

	// Visit a parse tree produced by MockParser#customArg.
	VisitCustomArg(ctx *CustomArgContext) interface{}

	// Visit a parse tree produced by MockParser#constStatement.
	VisitConstStatement(ctx *ConstStatementContext) interface{}

	// Visit a parse tree produced by MockParser#prevConst.
	VisitPrevConst(ctx *PrevConstContext) interface{}

	// Visit a parse tree produced by MockParser#afterConst.
	VisitAfterConst(ctx *AfterConstContext) interface{}

	// Visit a parse tree produced by MockParser#lengthSpec.
	VisitLengthSpec(ctx *LengthSpecContext) interface{}

	// Visit a parse tree produced by MockParser#rangeSpec.
	VisitRangeSpec(ctx *RangeSpecContext) interface{}

	// Visit a parse tree produced by MockParser#ignoreStatement.
	VisitIgnoreStatement(ctx *IgnoreStatementContext) interface{}
}
