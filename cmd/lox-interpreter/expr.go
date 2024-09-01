package main

type Expr interface {
	Accept(visitor ExprVisitor) any
}

type ExprVisitor interface {
	VisitAssignExpr(assign Assign) any
	VisitTernaryExpr(ternary Ternary) any
	VisitBinaryExpr(binary Binary) any
	VisitGroupingExpr(grouping Grouping) any
	VisitLiteralExpr(literal Literal) any
	VisitUnaryExpr(unary Unary) any
	VisitVariableExprExpr(variableexpr VariableExpr) any
}

type Assign struct {
	Name Token
	value Expr
}

func (thisAssign Assign) Accept(visitor ExprVisitor) any {
	return visitor.VisitAssignExpr(thisAssign)
}

type Ternary struct {
	Condition Expr
	TrueExpr Expr
	FalseExpr Expr
}

func (thisTernary Ternary) Accept(visitor ExprVisitor) any {
	return visitor.VisitTernaryExpr(thisTernary)
}

type Binary struct {
	Left Expr
	Operator Token
	Right Expr
}

func (thisBinary Binary) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinaryExpr(thisBinary)
}

type Grouping struct {
	Expression Expr
}

func (thisGrouping Grouping) Accept(visitor ExprVisitor) any {
	return visitor.VisitGroupingExpr(thisGrouping)
}

type Literal struct {
	Value any
}

func (thisLiteral Literal) Accept(visitor ExprVisitor) any {
	return visitor.VisitLiteralExpr(thisLiteral)
}

type Unary struct {
	Operator Token
	Right Expr
}

func (thisUnary Unary) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnaryExpr(thisUnary)
}

type VariableExpr struct {
	Name Token
}

func (thisVariableExpr VariableExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitVariableExprExpr(thisVariableExpr)
}

