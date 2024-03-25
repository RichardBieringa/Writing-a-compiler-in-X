package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 8383838;
`
	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseProgram()
	checkParserErrors(t, parse)

	if program == nil {
		t.Fatal("ParseProgram returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Parse program did not return the correct amount of statements. Expected: %v, Got: %v", 3, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]

		testLetStatement(t, statement, tt.expectedIdentifier)
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parsing the program resulted in %d errors", len(errors))

	for i, err := range errors {
		t.Errorf("err[%d]: %q", i, err)
	}

	t.FailNow()
}

func TestReturnStatement(t *testing.T) {
	input := `
  return 5;
  return 10;
  return 993322;
  `

	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseProgram()

	if program == nil {
		t.Fatal("ParseProgram returned nil")
	}

	errors := parse.Errors()
	if len(errors) > 0 {
		for i, error := range errors {
			t.Logf("ERR[%d]: `%s`", i, error)
		}
		t.Fatalf("Parse program did not return the correct amount of errors. Excepted: %d, got %d", 3, len(errors))
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 return statements, got: %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	parse := New(lexer.New(input))

	program := parse.ParseProgram()

	if program == nil {
		t.Fatal("Parsing the program returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected one statement, got=`%d`", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected statement to be an ExpressionStatement, got=`%T`", statement)
	}

	ident, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Expected expression to be an ast.Identifier, got=`%T`", ident)
	}

	if ident.Identifier != "foobar" {
		t.Fatalf("Expected identifier's value to equal `foobar`, got=`%s`", ident.Identifier)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("Expected identifier's TokenLiteral to equal `foobar`, got=`%s`", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "7;"

	parse := New(lexer.New(input))

	program := parse.ParseProgram()

	if program == nil {
		t.Fatal("Parsing the program returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected one statement, got=`%d`", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected statement to be an ExpressionStatement, got=`%T`", statement)
	}

	literal, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Expected expression to be an ast.IntegerLiteral, got=`%T`", literal)
	}

	if literal.Value != "7" {
		t.Fatalf("Expected literal's value to equal `7`, got=`%s`", literal.Value)
	}

	if literal.TokenLiteral() != "foobar" {
		t.Fatalf("Expected literal's TokenLiteral to equal `7`, got=`%s`", literal.TokenLiteral())
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	t.Logf("testLetStatement: %+v, name=%q", statement, name)

	if statement.TokenLiteral() != "let" {
		t.Errorf("token.Literal was not 'let'. got=%+v", statement.TokenLiteral())
		return false
	}

	// Go type assertion, provides access to an interface's underlying type
	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		// Go print format for variable type '%T'
		t.Errorf("Statement was not a LetStatement. got=%T", statement)
		return false
	}

	if letStatement.Name.Identifier != name {
		t.Errorf("Statement (name) did not have expected name. Expected=%q, got=%q", name, letStatement.Name.Identifier)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("Statement (token literal) did not have expected name. Expected=%q, got=%q", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
