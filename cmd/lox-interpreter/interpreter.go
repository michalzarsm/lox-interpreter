package main

import (
	"fmt"
	"os"
)

type Interpreter struct {
	env *Environment
}

func newInterpreter(env *Environment) *Interpreter {
	return &Interpreter{
		env,
	}
}

func (i *Interpreter) interpret(statements []Stmt) {
	for _, statement := range statements {
		i.execute(statement)
	}
}

func (i *Interpreter) execute(stmt Stmt) {
	stmt.Accept(i)
}

func (i *Interpreter) executeBlock(statements []Stmt, env *Environment) {
	previous := i.env
	i.env = env

	for _, statement := range statements {
		i.execute(statement)
	}

	i.env = previous
}

func (i *Interpreter) VisitBlockStmt(stmt Block) any {
	i.executeBlock(stmt.Statements, newEnvironment(i.env))
	return nil
}

func (i *Interpreter) VisitLiteralExpr(literal Literal) any {
	return literal.Value
}

func (i *Interpreter) VisitGroupingExpr(grouping Grouping) any {
	return i.evaluate(grouping.Expression)
}

func (i *Interpreter) evaluate(expr Expr) any {
	return expr.Accept(i)
}

func (i *Interpreter) VisitUnaryExpr(unary Unary) any {
	right := i.evaluate(unary.Right)

	switch unary.Operator.Type {
	case BANG:
		{
			return !isTruthy(right)
		}
	case MINUS:
		{
			ok := checkNumberOperand(right)
			if ok {
				return -right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operand must be a number.\n[line %d]\n", unary.Operator.Line)
			os.Exit(70)
		}
	}

	return nil
}

func (i *Interpreter) VisitVariableExprExpr(expr VariableExpr) any {
	getVar, err := i.env.get(expr.Name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[line %d] %v\n", expr.Name.Line, err.Error())
		os.Exit(70)
	}

	return getVar
}

func (i *Interpreter) VisitAssignExpr(expr Assign) any {
	value := i.evaluate(expr.value)
	i.env.assign(expr.Name, value)
	return value
}

func (i *Interpreter) VisitBinaryExpr(binary Binary) any {
	left := i.evaluate(binary.Left)
	right := i.evaluate(binary.Right)

	switch binary.Operator.Type {
	case GREATER:
		{
			ok := checkNumberOperands(left, right)
			if ok {
				return left.(float64) > right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	case GREATER_EQUAL:
		{
			ok := checkNumberOperands(left, right)
			if ok {
				return left.(float64) >= right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	case LESS:
		{
			ok := checkNumberOperands(left, right)
			if ok {
				return left.(float64) < right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	case LESS_EQUAL:
		{
			ok := checkNumberOperands(left, right)
			if ok {
				return left.(float64) <= right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	case BANG_EQUAL:
		{
			return !isEqual(left, right)
		}
	case EQUAL_EQUAL:
		{
			return isEqual(left, right)
		}
	case MINUS:
		{
			ok := checkNumberOperands(left, right)
			if ok {
				return left.(float64) - right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	case PLUS:
		{
			okNumber := checkNumberOperands(left, right)
			if okNumber {
				return left.(float64) + right.(float64)
			}

			okString := checkStringOperands(left, right)
			if okString {
				return left.(string) + right.(string)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	case SLASH:
		{
			ok := checkNumberOperands(left, right)
			if ok {
				return left.(float64) / right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	case STAR:
		{
			ok := checkNumberOperands(left, right)
			if ok {
				return left.(float64) * right.(float64)
			}

			fmt.Fprintf(os.Stderr, "Operands must be a numbers.\n[line %d]\n", binary.Operator.Line)
			os.Exit(70)
		}
	}

	return nil
}

func (i *Interpreter) VisitTernaryExpr(ternary Ternary) any { return nil }

func (i *Interpreter) VisitExpressionStmt(stmt Expression) any {
	i.evaluate(stmt.Expression)

	return nil
}

func (i *Interpreter) VisitPrintStmt(stmt Print) any {
	value := i.evaluate(stmt.Expression)
	if value == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(value)
	}

	return nil
}

func (i *Interpreter) VisitVariableStmtStmt(stmt VariableStmt) any {
	var value any
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}

	i.env.define(stmt.Name.Lexeme, value)

	return nil
}

func isTruthy(val any) bool {
	if val == nil {
		return false
	}

	valBool, ok := val.(bool)

	if ok {
		return valBool
	}

	return true
}

func isEqual(a any, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}

	return a == b
}

func checkNumberOperand(operand any) bool {
	_, ok := operand.(float64)

	return ok
}

func checkNumberOperands(operands ...any) bool {
	for _, operand := range operands {
		ok := checkNumberOperand(operand)
		if !ok {
			return false
		}
	}

	return true
}

func checkStringOperand(operand any) bool {
	_, ok := operand.(string)
	return ok
}

func checkStringOperands(operands ...any) bool {
	for _, operand := range operands {
		ok := checkStringOperand(operand)
		if !ok {
			return false
		}
	}

	return true
}
