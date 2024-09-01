package main

type Stmt interface {
	Accept(visitor StmtVisitor) any
}

type StmtVisitor interface {
	VisitBlockStmt(block Block) any
	VisitExpressionStmt(expression Expression) any
	VisitPrintStmt(print Print) any
	VisitVariableStmtStmt(variablestmt VariableStmt) any
}

type Block struct {
	Statements []Stmt
}

func (thisBlock Block) Accept(visitor StmtVisitor) any {
	return visitor.VisitBlockStmt(thisBlock)
}

type Expression struct {
	Expression Expr
}

func (thisExpression Expression) Accept(visitor StmtVisitor) any {
	return visitor.VisitExpressionStmt(thisExpression)
}

type Print struct {
	Expression Expr
}

func (thisPrint Print) Accept(visitor StmtVisitor) any {
	return visitor.VisitPrintStmt(thisPrint)
}

type VariableStmt struct {
	Name Token
	Initializer Expr
}

func (thisVariableStmt VariableStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitVariableStmtStmt(thisVariableStmt)
}

