package main

import (
	"fmt"
	"os"
)

type Lox struct {
	errors []Error
}

type ErrorType = string

const (
	SyntaxError       ErrorType = "SyntaxError"
	RuntimeError      ErrorType = "RuntimeError"
	ValueConvertError ErrorType = "ValueConvertError"
)

type Error struct {
	errorType ErrorType
	token     Token
	message   string
	exitCode  int
}

func newLox() *Lox {
	return &Lox{
		errors: make([]Error, 0),
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./lox-interpreter.sh <command> <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" && command != "parse" && command != "run" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		lox := newLox()
		switch command {
		case "tokenize":
			{
				scanner := newScanner(string(fileContents), lox)
				scanner.scanTokens()

				tokens := scanner.Tokens

				for _, token := range tokens {
					if token.Type == STRING {
						fmt.Printf("STRING \"%s\" %s\n", token.Lexeme, token.Literal)
					} else if token.Type == NUMBER {
						num, ok := token.Literal.(float64)
						if !ok {
							scanner.Lox.errors = append(scanner.Lox.errors, Error{errorType: ValueConvertError, token: token, message: "Not a float", exitCode: 65})
							return
						}

						if num == float64(int(num)) {
							fmt.Printf("NUMBER %s %.1f\n", token.Lexeme, num)
						} else {
							fmt.Printf("NUMBER %s %g\n", token.Lexeme, num)
						}
					} else {
						fmt.Printf("%s %s null\n", getTokenTypeName(string(token.Type)), token.Lexeme)
					}
				}

				if len(scanner.Lox.errors) > 0 {
					lox.error()
				}
			}
		case "parse":
			{
				scanner := newScanner(string(fileContents), lox)
				scanner.scanTokens()
				if len(scanner.Lox.errors) > 0 {
					lox.error()
				}
				tokens := scanner.Tokens
				parser := newParser(tokens, lox)
				statements := parser.parse()
				if len(parser.Lox.errors) > 0 {
					lox.error()
				}
				fmt.Printf("%v\n", statements)
			}
		case "run":
			{
				scanner := newScanner(string(fileContents), lox)
				scanner.scanTokens()
				if len(scanner.Lox.errors) > 0 {
					lox.error()
				}
				tokens := scanner.Tokens
				parser := newParser(tokens, lox)
				statements := parser.parse()
				if len(parser.Lox.errors) > 0 {
					lox.error()
				}
				env := newEnvironment(nil)
				interpreter := newInterpreter(env)
				interpreter.interpret(statements)
			}
		}
	} else {
		fmt.Println("EOF  null")
	}
}

func (lox *Lox) error() {
	for _, err := range lox.errors {
		token := err.token
		message := err.message
		fmt.Fprintln(os.Stderr, fmt.Errorf("[line %d] Error: %s", token.Line, message))
	}
	os.Exit(lox.errors[0].exitCode)
}
