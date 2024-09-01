package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ast-generator <output directory>")
		os.Exit(64)
	}

	outputDir := os.Args[1]

	defineAst(outputDir, "Expr", []string{
		"Assign       : Name Token, value Expr",
		"Ternary      : Condition Expr, TrueExpr Expr, FalseExpr Expr",
		"Binary       : Left Expr, Operator Token, Right Expr",
		"Grouping     : Expression Expr",
		"Literal      : Value any",
		"Unary        : Operator Token, Right Expr",
		"VariableExpr : Name Token",
	})

	defineAst(outputDir, "Stmt", []string{
		"Block        : Statements []Stmt",
		"Expression   : Expression Expr",
		"Print        : Expression Expr",
		"VariableStmt : Name Token, Initializer Expr",
	})
}

func defineAst(outputDir string, baseName string, types []string) {
	path := outputDir + "/" + strings.ToLower(baseName) + ".go"

	file, err := os.Create(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}

	file.WriteString("package main\n\n")
	file.WriteString("type" + " " + baseName + " " + "interface {\n")
	file.WriteString("	Accept(visitor" + " " + baseName + "Visitor) any\n")
	file.WriteString("}\n\n")

	defineVisitor(file, baseName, types)

	for _, t := range types {
		splitType := strings.Split(t, ":")
		structName := strings.Trim(splitType[0], " ")
		fields := strings.Trim(splitType[1], " ")
		defineType(file, structName, fields)
		defineAccept(file, baseName, structName)
	}
}

func defineType(file *os.File, structName string, fieldList string) {
	fields := strings.Split(fieldList, ", ")
	file.WriteString("type" + " " + structName + " " + "struct {\n")
	for _, field := range fields {
		splitField := strings.Split(field, " ")
		fieldName := splitField[0]
		fieldType := splitField[1]
		file.WriteString("	" + fieldName + " " + fieldType + "\n")
	}
	file.WriteString("}\n\n")
}

func defineAccept(file *os.File, baseName, structName string) {
	file.WriteString("func (this" + structName + " " + structName + ") " + "Accept(visitor" + " " + baseName + "Visitor) any {\n")
	file.WriteString("	return visitor.Visit" + structName + baseName + "(this" + structName + ")\n")
	file.WriteString("}\n\n")
}

func defineVisitor(file *os.File, baseName string, types []string) {
	fmt.Fprintf(file, "type"+" "+baseName+"Visitor interface {\n")
	for _, t := range types {
		className := strings.Trim(strings.Split(t, ":")[0], " ")
		fmt.Fprintf(file, "\tVisit%s%s(%s %s) any\n", className, baseName, strings.ToLower(className), className)
	}
	fmt.Fprintf(file, "}\n\n")
}
