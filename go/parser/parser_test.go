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

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	t.Logf("testLetStatement: %+v, name=%q", statement, name)

	if statement.TokenLiteral() != "let" {
		t.Errorf("token.Literal was not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	// Go type assertion, provides access to an interface's underlying type
	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		// Go print format for variable type '%T'
		t.Errorf("Statement was not a LetStatement. got=%T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("Statement (name) did not have expected name. Expected=%q, got=%q", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("Statement (token literal) did not have expected name. Expected=%q, got=%q", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
