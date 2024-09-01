package main

type TokenType string

const (
	// Single-character tokens.
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
	COMMA       TokenType = ","
	DOT         TokenType = "."
	MINUS       TokenType = "-"
	PLUS        TokenType = "+"
	SEMICOLON   TokenType = ";"
	SLASH       TokenType = "/"
	STAR        TokenType = "*"

	// One or two character tokens.
	BANG          TokenType = "!"
	BANG_EQUAL    TokenType = "!="
	EQUAL         TokenType = "="
	EQUAL_EQUAL   TokenType = "=="
	GREATER       TokenType = ">"
	GREATER_EQUAL TokenType = ">="
	LESS          TokenType = "<"
	LESS_EQUAL    TokenType = "<="

	// Literals.
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

	// Keywords.
	AND    TokenType = "and"
	CLASS  TokenType = "class"
	ELSE   TokenType = "else"
	FALSE  TokenType = "false"
	FUN    TokenType = "fun"
	FOR    TokenType = "for"
	IF     TokenType = "if"
	NIL    TokenType = "nil"
	OR     TokenType = "or"
	PRINT  TokenType = "print"
	RETURN TokenType = "return"
	SUPER  TokenType = "super"
	THIS   TokenType = "this"
	TRUE   TokenType = "true"
	VAR    TokenType = "var"
	WHILE  TokenType = "while"

	EOF TokenType = "eof"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

var valueToTokenType = map[string]string{
	"(":          "LEFT_PAREN",
	")":          "RIGHT_PAREN",
	"{":          "LEFT_BRACE",
	"}":          "RIGHT_BRACE",
	",":          "COMMA",
	".":          "DOT",
	"-":          "MINUS",
	"+":          "PLUS",
	";":          "SEMICOLON",
	"/":          "SLASH",
	"*":          "STAR",
	"!":          "BANG",
	"!=":         "BANG_EQUAL",
	"=":          "EQUAL",
	"==":         "EQUAL_EQUAL",
	">":          "GREATER",
	">=":         "GREATER_EQUAL",
	"<":          "LESS",
	"<=":         "LESS_EQUAL",
	"IDENTIFIER": "IDENTIFIER",
	"STRING":     "STRING",
	"NUMBER":     "NUMBER",
	"and":        "AND",
	"class":      "CLASS",
	"else":       "ELSE",
	"false":      "FALSE",
	"fun":        "FUN",
	"for":        "FOR",
	"if":         "IF",
	"nil":        "NIL",
	"or":         "OR",
	"print":      "PRINT",
	"return":     "RETURN",
	"super":      "SUPER",
	"this":       "THIS",
	"true":       "TRUE",
	"var":        "VAR",
	"while":      "WHILE",
	"eof":        "EOF",
}

func getTokenTypeName(value string) string {
	if name, ok := valueToTokenType[value]; ok {
		return name
	}
	return "IDENTIFIER"
}
