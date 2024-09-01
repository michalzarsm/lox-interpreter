package main

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	Lox     *Lox
	Source  string
	Tokens  []Token
	Start   int
	Current int
	Line    int
}

func newScanner(source string, lox *Lox) *Scanner {
	return &Scanner{
		Lox:     lox,
		Source:  source,
		Tokens:  make([]Token, 0),
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

func (s *Scanner) scanTokens() {
	for s.Current < len(s.Source) {
		s.scanToken()
	}
	s.Tokens = append(s.Tokens, Token{Type: EOF})
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch TokenType(c) {
	case LEFT_PAREN:
		s.addToken(Token{Type: LEFT_PAREN, Lexeme: string(LEFT_PAREN), Literal: nil, Line: s.Line})
	case RIGHT_PAREN:
		s.addToken(Token{Type: RIGHT_PAREN, Lexeme: string(RIGHT_PAREN), Literal: nil, Line: s.Line})
	case LEFT_BRACE:
		s.addToken(Token{Type: LEFT_BRACE, Lexeme: string(LEFT_BRACE), Literal: nil, Line: s.Line})
	case RIGHT_BRACE:
		s.addToken(Token{Type: RIGHT_BRACE, Lexeme: string(RIGHT_BRACE), Literal: nil, Line: s.Line})
	case COMMA:
		s.addToken(Token{Type: COMMA, Lexeme: string(COMMA), Literal: nil, Line: s.Line})
	case DOT:
		s.addToken(Token{Type: DOT, Lexeme: string(DOT), Literal: nil, Line: s.Line})
	case MINUS:
		s.addToken(Token{Type: MINUS, Lexeme: string(MINUS), Literal: nil, Line: s.Line})
	case PLUS:
		s.addToken(Token{Type: PLUS, Lexeme: string(PLUS), Literal: nil, Line: s.Line})
	case SEMICOLON:
		s.addToken(Token{Type: SEMICOLON, Lexeme: string(SEMICOLON), Literal: nil, Line: s.Line})
	case SLASH:
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(Token{Type: SLASH, Lexeme: string(SLASH), Literal: nil, Line: s.Line})
		}
	case STAR:
		s.addToken(Token{Type: STAR, Lexeme: string(STAR), Literal: nil, Line: s.Line})
	case "!":
		if s.match('=') {
			s.addToken(Token{Type: BANG_EQUAL, Lexeme: string(BANG_EQUAL), Literal: nil, Line: s.Line})

		} else {
			s.addToken(Token{Type: BANG, Lexeme: string(BANG), Literal: nil, Line: s.Line})
		}
	case "=":
		if s.match('=') {
			s.addToken(Token{Type: EQUAL_EQUAL, Lexeme: string(EQUAL_EQUAL), Literal: nil, Line: s.Line})
		} else {
			s.addToken(Token{Type: EQUAL, Lexeme: string(EQUAL), Literal: nil, Line: s.Line})
		}
	case "<":
		if s.match('=') {
			s.addToken(Token{Type: LESS_EQUAL, Lexeme: string(LESS_EQUAL), Literal: nil, Line: s.Line})
		} else {
			s.addToken(Token{Type: LESS, Lexeme: string(LESS), Literal: nil, Line: s.Line})
		}
	case ">":
		if s.match('=') {
			s.addToken(Token{Type: GREATER_EQUAL, Lexeme: string(GREATER_EQUAL), Literal: nil, Line: s.Line})
		} else {
			s.addToken(Token{Type: GREATER, Lexeme: string(GREATER), Literal: nil, Line: s.Line})
		}
	case " ":
		{
		}
	case "\r":
		{
		}
	case "\t":
		{
			break
		}
	case "\n":
		{
			s.Line += 1
			break
		}
	case "\"":
		{
			s.string()
		}
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			s.Lox.errors = append(s.Lox.errors, Error{errorType: SyntaxError, token: Token{Line: s.Line}, message: fmt.Sprintf("Unexpected character: %c", c), exitCode: 65})
		}
	}
}

func (s *Scanner) string() {
	stringStart := s.Current

	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.Line += 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.Lox.errors = append(s.Lox.errors, Error{errorType: SyntaxError, token: Token{Line: s.Line}, message: "Unterminated string.", exitCode: 65})
		return
	}

	s.advance()

	value := s.Source[stringStart : s.Current-1]

	s.addToken(Token{Type: STRING, Lexeme: value, Literal: value, Line: s.Line})
}

func (s *Scanner) number() {
	digitStart := s.Current - 1

	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	value := s.Source[digitStart:s.Current]

	parsedValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		s.Lox.errors = append(s.Lox.errors, Error{errorType: ValueConvertError, token: Token{Line: s.Line}, message: "Float Parse Error.", exitCode: 65})
		return
	}

	s.addToken(Token{Type: NUMBER, Lexeme: value, Literal: parsedValue, Line: s.Line})
}

func (s *Scanner) identifier() {
	identifierStart := s.Current
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	value := s.Source[identifierStart-1 : s.Current]

	switch TokenType(value) {
	case AND:
		{
			s.addToken(Token{Type: AND, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case CLASS:
		{
			s.addToken(Token{Type: CLASS, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case ELSE:
		{
			s.addToken(Token{Type: ELSE, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case FALSE:
		{
			s.addToken(Token{Type: FALSE, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case FOR:
		{
			s.addToken(Token{Type: FOR, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case FUN:
		{
			s.addToken(Token{Type: FUN, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case IF:
		{
			s.addToken(Token{Type: IF, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case NIL:
		{
			s.addToken(Token{Type: NIL, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case OR:
		{
			s.addToken(Token{Type: OR, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case PRINT:
		{
			s.addToken(Token{Type: PRINT, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case RETURN:
		{
			s.addToken(Token{Type: RETURN, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case SUPER:
		{
			s.addToken(Token{Type: SUPER, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case THIS:
		{
			s.addToken(Token{Type: THIS, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case TRUE:
		{
			s.addToken(Token{Type: TRUE, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case VAR:
		{
			s.addToken(Token{Type: VAR, Lexeme: value, Literal: nil, Line: s.Line})
		}
	case WHILE:
		{
			s.addToken(Token{Type: WHILE, Lexeme: value, Literal: nil, Line: s.Line})
		}
	default:
		{
			s.addToken(Token{Type: IDENTIFIER, Lexeme: value, Literal: nil, Line: s.Line})
		}
	}

}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.Source[s.Current] != expected {
		return false
	}

	s.Current += 1
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.Source[s.Current]
}

func (s *Scanner) peekNext() byte {
	if s.Current+1 >= len(s.Source) {
		return 0
	}

	return s.Source[s.Current+1]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) addToken(token Token) {
	s.Tokens = append(s.Tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) advance() byte {
	c := s.Source[s.Current]
	s.Current += 1
	return c
}
