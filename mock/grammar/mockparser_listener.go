// Code generated from /Users/anqiansong/goland/go/go-zero/anqiansong/mock/grammar/MockParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MockParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MockParserListener is a complete listener for a parse tree produced by MockParser.
type MockParserListener interface {
	antlr.ParseTreeListener

	// EnterMockExpression is called when entering the mockExpression production.
	EnterMockExpression(c *MockExpressionContext)

	// EnterNormalExpression is called when entering the normalExpression production.
	EnterNormalExpression(c *NormalExpressionContext)

	// EnterPresetExpr is called when entering the presetExpr production.
	EnterPresetExpr(c *PresetExprContext)

	// EnterPresetBoolExpr is called when entering the presetBoolExpr production.
	EnterPresetBoolExpr(c *PresetBoolExprContext)

	// EnterPresetSliceExpr is called when entering the presetSliceExpr production.
	EnterPresetSliceExpr(c *PresetSliceExprContext)

	// EnterPresetMapExpr is called when entering the presetMapExpr production.
	EnterPresetMapExpr(c *PresetMapExprContext)

	// EnterPresetSpecExpr is called when entering the presetSpecExpr production.
	EnterPresetSpecExpr(c *PresetSpecExprContext)

	// EnterNumberExpr is called when entering the numberExpr production.
	EnterNumberExpr(c *NumberExprContext)

	// EnterNumberRangeExpr is called when entering the numberRangeExpr production.
	EnterNumberRangeExpr(c *NumberRangeExprContext)

	// EnterNumberLengthExpr is called when entering the numberLengthExpr production.
	EnterNumberLengthExpr(c *NumberLengthExprContext)

	// EnterStringExpr is called when entering the stringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterStringRange is called when entering the stringRange production.
	EnterStringRange(c *StringRangeContext)

	// EnterStringLength is called when entering the stringLength production.
	EnterStringLength(c *StringLengthContext)

	// EnterStringTemplate is called when entering the stringTemplate production.
	EnterStringTemplate(c *StringTemplateContext)

	// EnterCustomArgSpec is called when entering the customArgSpec production.
	EnterCustomArgSpec(c *CustomArgSpecContext)

	// EnterCustomArg is called when entering the customArg production.
	EnterCustomArg(c *CustomArgContext)

	// EnterConstStatement is called when entering the constStatement production.
	EnterConstStatement(c *ConstStatementContext)

	// EnterPrevConst is called when entering the prevConst production.
	EnterPrevConst(c *PrevConstContext)

	// EnterAfterConst is called when entering the afterConst production.
	EnterAfterConst(c *AfterConstContext)

	// EnterLengthSpec is called when entering the lengthSpec production.
	EnterLengthSpec(c *LengthSpecContext)

	// EnterRangeSpec is called when entering the rangeSpec production.
	EnterRangeSpec(c *RangeSpecContext)

	// EnterIgnoreStatement is called when entering the ignoreStatement production.
	EnterIgnoreStatement(c *IgnoreStatementContext)

	// ExitMockExpression is called when exiting the mockExpression production.
	ExitMockExpression(c *MockExpressionContext)

	// ExitNormalExpression is called when exiting the normalExpression production.
	ExitNormalExpression(c *NormalExpressionContext)

	// ExitPresetExpr is called when exiting the presetExpr production.
	ExitPresetExpr(c *PresetExprContext)

	// ExitPresetBoolExpr is called when exiting the presetBoolExpr production.
	ExitPresetBoolExpr(c *PresetBoolExprContext)

	// ExitPresetSliceExpr is called when exiting the presetSliceExpr production.
	ExitPresetSliceExpr(c *PresetSliceExprContext)

	// ExitPresetMapExpr is called when exiting the presetMapExpr production.
	ExitPresetMapExpr(c *PresetMapExprContext)

	// ExitPresetSpecExpr is called when exiting the presetSpecExpr production.
	ExitPresetSpecExpr(c *PresetSpecExprContext)

	// ExitNumberExpr is called when exiting the numberExpr production.
	ExitNumberExpr(c *NumberExprContext)

	// ExitNumberRangeExpr is called when exiting the numberRangeExpr production.
	ExitNumberRangeExpr(c *NumberRangeExprContext)

	// ExitNumberLengthExpr is called when exiting the numberLengthExpr production.
	ExitNumberLengthExpr(c *NumberLengthExprContext)

	// ExitStringExpr is called when exiting the stringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitStringRange is called when exiting the stringRange production.
	ExitStringRange(c *StringRangeContext)

	// ExitStringLength is called when exiting the stringLength production.
	ExitStringLength(c *StringLengthContext)

	// ExitStringTemplate is called when exiting the stringTemplate production.
	ExitStringTemplate(c *StringTemplateContext)

	// ExitCustomArgSpec is called when exiting the customArgSpec production.
	ExitCustomArgSpec(c *CustomArgSpecContext)

	// ExitCustomArg is called when exiting the customArg production.
	ExitCustomArg(c *CustomArgContext)

	// ExitConstStatement is called when exiting the constStatement production.
	ExitConstStatement(c *ConstStatementContext)

	// ExitPrevConst is called when exiting the prevConst production.
	ExitPrevConst(c *PrevConstContext)

	// ExitAfterConst is called when exiting the afterConst production.
	ExitAfterConst(c *AfterConstContext)

	// ExitLengthSpec is called when exiting the lengthSpec production.
	ExitLengthSpec(c *LengthSpecContext)

	// ExitRangeSpec is called when exiting the rangeSpec production.
	ExitRangeSpec(c *RangeSpecContext)

	// ExitIgnoreStatement is called when exiting the ignoreStatement production.
	ExitIgnoreStatement(c *IgnoreStatementContext)
}
