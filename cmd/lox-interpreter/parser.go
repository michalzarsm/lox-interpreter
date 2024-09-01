package main

import "reflect"

type Parser struct {
	Lox     *Lox
	tokens  []Token
	current int
}

func newParser(tokens []Token, lox *Lox) *Parser {
	return &Parser{
		Lox:     lox,
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) parse() []Stmt {
	statements := make([]Stmt, 0)
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}

	return statements
}

func (p *Parser) expression() Expr {
	return p.assignment()
}

func (p *Parser) declaration() Stmt {
	if p.match(VAR) {
		return p.varDeclaration()
	}

	return p.statement()
}

func (p *Parser) statement() Stmt {
	if p.match(PRINT) {
		return p.printStatement()
	}

	if p.match(LEFT_BRACE) {
		return Block{p.block()}
	}

	return p.expressionStatement()
}

func (p *Parser) printStatement() Print {
	value := p.expression()
	p.consume(SEMICOLON, "Expect ';' after value.")
	return Print{Expression: value}
}

func (p *Parser) varDeclaration() Stmt {
	name := p.consume(IDENTIFIER, "Expect variable name.")

	var initializer Expr

	if p.match(EQUAL) {
		initializer = p.expression()
	}

	p.consume(SEMICOLON, "Expect ';' after variable declaration.")

	return VariableStmt{name, initializer}

}

func (p *Parser) expressionStatement() Expression {
	expr := p.expression()
	p.consume(SEMICOLON, "Expect ';' after expression.")
	return Expression{Expression: expr}
}

func (p *Parser) block() []Stmt {
	statements := make([]Stmt, 0)

	for !p.check(RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}

	p.consume(RIGHT_BRACE, "Expect '}' after block.")

	return statements
}

func (p *Parser) assignment() Expr {
	expr := p.equality()

	if p.match(EQUAL) {
		equals := p.previous()
		value := p.assignment()

		exprType := reflect.TypeOf(expr)

		if exprType == reflect.TypeFor[VariableExpr]() {
			name := expr.(VariableExpr).Name
			return Assign{name, value}
		}

		p.error(equals, "Invalid assignment target.", 65)
	}

	return expr
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match(SLASH, STAR) {
		operator := p.previous()
		right := p.unary()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary()
		return Unary{operator, right}
	}

	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(FALSE) {
		return Literal{false}
	}

	if p.match(TRUE) {
		return Literal{true}
	}

	if p.match(NIL) {
		return Literal{nil}
	}

	if p.match(NUMBER, STRING) {
		return Literal{p.previous().Literal}
	}

	if p.match(IDENTIFIER) {
		return VariableExpr{p.previous()}
	}

	if p.match(LEFT_PAREN) {
		expr := p.expression()
		p.consume(RIGHT_PAREN, "Expect ')' after expression.")
		return Grouping{expr}
	}

	p.error(p.peek(), "Expect expression.", 65)

	return nil
}

func (p *Parser) match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current += 1
	}

	return p.previous()
}

func (p *Parser) consume(tokenType TokenType, message string) Token {
	if p.check(tokenType) {
		return p.advance()
	}

	p.error(p.peek(), message, 65)

	return Token{}
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) error(token Token, message string, exitCode int) {
	p.Lox.errors = append(p.Lox.errors, Error{errorType: SyntaxError, token: token, message: message, exitCode: exitCode})
}

/*
func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == SEMICOLON {
			return
		}

		switch p.peek().Type {
		case CLASS:
		case FUN:
		case VAR:
		case FOR:
		case IF:
		case WHILE:
		case PRINT:
		case RETURN:
			return
		}

		p.advance()
	}
}
*/
