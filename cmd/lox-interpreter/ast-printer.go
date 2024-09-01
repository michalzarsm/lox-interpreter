package main

/*

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

func (t AstPrinter) VisitTernaryExpr(ternary Ternary) any {
	return t.Parenthesize("?", ternary.Condition, ternary.TrueExpr, ternary.FalseExpr)
}

func (t AstPrinter) VisitBinaryExpr(binary Binary) any {
	return t.Parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (t AstPrinter) VisitGroupingExpr(grouping Grouping) any {
	return t.Parenthesize("group", grouping.Expression)
}

func (t AstPrinter) VisitLiteralExpr(literal Literal) any {
	if literal.Value == nil {
		return "nil"
	}
	switch l := literal.Value.(type) {
	case float64:
		return FormatNumber(l)
	default:
		return fmt.Sprint(l)
	}
}

func (t AstPrinter) VisitUnaryExpr(unary Unary) any {
	return t.Parenthesize(unary.Operator.Lexeme, unary.Right)
}

func (t AstPrinter) Parenthesize(name string, exprs ...Expr) string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)
	for _, expr := range exprs {
		sb.WriteString(" ")
		sb.WriteString(expr.Accept(t).(string))
	}
	sb.WriteString(")")
	return sb.String()
}

func (t AstPrinter) Print(expr Expr) string {
	return expr.Accept(t).(string)
}

func NewAstPrinter() AstPrinter {
	return AstPrinter{}
}

func PrintAst(expr Expr) string {
	return NewAstPrinter().Print(expr)
}

*/
