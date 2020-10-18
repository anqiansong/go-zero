// Code generated from /Users/anqiansong/goland/go/go-zero/anqiansong/mock/grammar/MockParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MockParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMockParserListener is a complete listener for a parse tree produced by MockParser.
type BaseMockParserListener struct{}

var _ MockParserListener = &BaseMockParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMockParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMockParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMockParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMockParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMockExpression is called when production mockExpression is entered.
func (s *BaseMockParserListener) EnterMockExpression(ctx *MockExpressionContext) {}

// ExitMockExpression is called when production mockExpression is exited.
func (s *BaseMockParserListener) ExitMockExpression(ctx *MockExpressionContext) {}

// EnterNormalExpression is called when production normalExpression is entered.
func (s *BaseMockParserListener) EnterNormalExpression(ctx *NormalExpressionContext) {}

// ExitNormalExpression is called when production normalExpression is exited.
func (s *BaseMockParserListener) ExitNormalExpression(ctx *NormalExpressionContext) {}

// EnterPresetExpr is called when production presetExpr is entered.
func (s *BaseMockParserListener) EnterPresetExpr(ctx *PresetExprContext) {}

// ExitPresetExpr is called when production presetExpr is exited.
func (s *BaseMockParserListener) ExitPresetExpr(ctx *PresetExprContext) {}

// EnterPresetBoolExpr is called when production presetBoolExpr is entered.
func (s *BaseMockParserListener) EnterPresetBoolExpr(ctx *PresetBoolExprContext) {}

// ExitPresetBoolExpr is called when production presetBoolExpr is exited.
func (s *BaseMockParserListener) ExitPresetBoolExpr(ctx *PresetBoolExprContext) {}

// EnterPresetSliceExpr is called when production presetSliceExpr is entered.
func (s *BaseMockParserListener) EnterPresetSliceExpr(ctx *PresetSliceExprContext) {}

// ExitPresetSliceExpr is called when production presetSliceExpr is exited.
func (s *BaseMockParserListener) ExitPresetSliceExpr(ctx *PresetSliceExprContext) {}

// EnterPresetMapExpr is called when production presetMapExpr is entered.
func (s *BaseMockParserListener) EnterPresetMapExpr(ctx *PresetMapExprContext) {}

// ExitPresetMapExpr is called when production presetMapExpr is exited.
func (s *BaseMockParserListener) ExitPresetMapExpr(ctx *PresetMapExprContext) {}

// EnterPresetSpecExpr is called when production presetSpecExpr is entered.
func (s *BaseMockParserListener) EnterPresetSpecExpr(ctx *PresetSpecExprContext) {}

// ExitPresetSpecExpr is called when production presetSpecExpr is exited.
func (s *BaseMockParserListener) ExitPresetSpecExpr(ctx *PresetSpecExprContext) {}

// EnterNumberExpr is called when production numberExpr is entered.
func (s *BaseMockParserListener) EnterNumberExpr(ctx *NumberExprContext) {}

// ExitNumberExpr is called when production numberExpr is exited.
func (s *BaseMockParserListener) ExitNumberExpr(ctx *NumberExprContext) {}

// EnterNumberRangeExpr is called when production numberRangeExpr is entered.
func (s *BaseMockParserListener) EnterNumberRangeExpr(ctx *NumberRangeExprContext) {}

// ExitNumberRangeExpr is called when production numberRangeExpr is exited.
func (s *BaseMockParserListener) ExitNumberRangeExpr(ctx *NumberRangeExprContext) {}

// EnterNumberLengthExpr is called when production numberLengthExpr is entered.
func (s *BaseMockParserListener) EnterNumberLengthExpr(ctx *NumberLengthExprContext) {}

// ExitNumberLengthExpr is called when production numberLengthExpr is exited.
func (s *BaseMockParserListener) ExitNumberLengthExpr(ctx *NumberLengthExprContext) {}

// EnterStringExpr is called when production stringExpr is entered.
func (s *BaseMockParserListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production stringExpr is exited.
func (s *BaseMockParserListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterStringRange is called when production stringRange is entered.
func (s *BaseMockParserListener) EnterStringRange(ctx *StringRangeContext) {}

// ExitStringRange is called when production stringRange is exited.
func (s *BaseMockParserListener) ExitStringRange(ctx *StringRangeContext) {}

// EnterStringLength is called when production stringLength is entered.
func (s *BaseMockParserListener) EnterStringLength(ctx *StringLengthContext) {}

// ExitStringLength is called when production stringLength is exited.
func (s *BaseMockParserListener) ExitStringLength(ctx *StringLengthContext) {}

// EnterStringTemplate is called when production stringTemplate is entered.
func (s *BaseMockParserListener) EnterStringTemplate(ctx *StringTemplateContext) {}

// ExitStringTemplate is called when production stringTemplate is exited.
func (s *BaseMockParserListener) ExitStringTemplate(ctx *StringTemplateContext) {}

// EnterCustomArgSpec is called when production customArgSpec is entered.
func (s *BaseMockParserListener) EnterCustomArgSpec(ctx *CustomArgSpecContext) {}

// ExitCustomArgSpec is called when production customArgSpec is exited.
func (s *BaseMockParserListener) ExitCustomArgSpec(ctx *CustomArgSpecContext) {}

// EnterCustomArg is called when production customArg is entered.
func (s *BaseMockParserListener) EnterCustomArg(ctx *CustomArgContext) {}

// ExitCustomArg is called when production customArg is exited.
func (s *BaseMockParserListener) ExitCustomArg(ctx *CustomArgContext) {}

// EnterConstStatement is called when production constStatement is entered.
func (s *BaseMockParserListener) EnterConstStatement(ctx *ConstStatementContext) {}

// ExitConstStatement is called when production constStatement is exited.
func (s *BaseMockParserListener) ExitConstStatement(ctx *ConstStatementContext) {}

// EnterPrevConst is called when production prevConst is entered.
func (s *BaseMockParserListener) EnterPrevConst(ctx *PrevConstContext) {}

// ExitPrevConst is called when production prevConst is exited.
func (s *BaseMockParserListener) ExitPrevConst(ctx *PrevConstContext) {}

// EnterAfterConst is called when production afterConst is entered.
func (s *BaseMockParserListener) EnterAfterConst(ctx *AfterConstContext) {}

// ExitAfterConst is called when production afterConst is exited.
func (s *BaseMockParserListener) ExitAfterConst(ctx *AfterConstContext) {}

// EnterLengthSpec is called when production lengthSpec is entered.
func (s *BaseMockParserListener) EnterLengthSpec(ctx *LengthSpecContext) {}

// ExitLengthSpec is called when production lengthSpec is exited.
func (s *BaseMockParserListener) ExitLengthSpec(ctx *LengthSpecContext) {}

// EnterRangeSpec is called when production rangeSpec is entered.
func (s *BaseMockParserListener) EnterRangeSpec(ctx *RangeSpecContext) {}

// ExitRangeSpec is called when production rangeSpec is exited.
func (s *BaseMockParserListener) ExitRangeSpec(ctx *RangeSpecContext) {}

// EnterIgnoreStatement is called when production ignoreStatement is entered.
func (s *BaseMockParserListener) EnterIgnoreStatement(ctx *IgnoreStatementContext) {}

// ExitIgnoreStatement is called when production ignoreStatement is exited.
func (s *BaseMockParserListener) ExitIgnoreStatement(ctx *IgnoreStatementContext) {}
