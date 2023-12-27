package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LBRACE, ")"},
		{token.RBRACE, ")"},
		{token.LPAREN, "{"},
		{token.RPAREN, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, expected := range tests {
		actual := lexer.NextToken()

		if actual.Type != expected.Type {
			t.Fatalf("tests[%d] - incorrect token type: expected=%q, got=%q", i, expected.Type, actual.Type)
		}

		if actual.Literal != expected.Literal {
			t.Fatalf("tests[%d] - incorrect token literal: expected=%q, got=%q", i, expected.Type, actual.Type)
		}
	}
}
