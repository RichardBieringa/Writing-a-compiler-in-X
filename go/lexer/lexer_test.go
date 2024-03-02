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
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
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

func TestMoreSourceCode(t *testing.T) {
	input := `let myIdentifier = 5;`

	lexer := New(input)

	tests := []struct {
		Type     token.TokenType
		Literal  string
		Position int
		PeekChar byte
	}{
		{
			Type:     token.LET,
			Literal:  "let",
			Position: 3,
			PeekChar: 'm',
		},
		{
			Type:     token.IDENT,
			Literal:  "myIdentifier",
			Position: 16,
			PeekChar: '=',
		},
		{
			Type:     token.ASSIGN,
			Literal:  "=",
			Position: 18,
			PeekChar: '5',
		},
		{
			Type:     token.INT,
			Literal:  "5",
			Position: 20,
			PeekChar: '\x00',
		},
		{
			Type:     token.SEMICOLON,
			Literal:  ";",
			Position: 20,
			PeekChar: '\x00',
		},
		{
			Type:     token.EOF,
			Literal:  "",
			Position: 20,
			PeekChar: '\x00',
		},
	}

	for i, testCase := range tests {

		current := lexer.NextToken()

		if current.Type != testCase.Type {
			t.Fatalf("Test[%d] - incorrect token type: expected: %q, got: %q", i, testCase.Type, current.Type)
		}

		if current.Literal != testCase.Literal {
			t.Fatalf("Test[%d] - incorrect token literal: expected: %q, got: %q", i, testCase.Literal, current.Literal)
		}

		if lexer.position != testCase.Position {
			t.Fatalf("Test[%d] - incorrect position: expected: %d, got: %d", i, testCase.Position, lexer.position)
		}

		if lexer.peekChar() != testCase.PeekChar {
			t.Fatalf("Test[%d] - incorrect peekchar: expected: %q, got: %q", i, testCase.PeekChar, lexer.peekChar())
		}
	}
}

func TestSourceCode(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expected := range tests {
		actual := lexer.NextToken()

		if actual.Type != expected.Type {
			t.Fatalf("tests[%d] - incorrect token type: expected=%q, got=%q (value=%x)", i, expected.Type, actual.Type, actual.Literal)
		}

		if actual.Literal != expected.Literal {
			t.Fatalf("tests[%d] - incorrect token literal: expected=%q, got=%q", i, expected.Literal, actual.Literal)
		}
	}
}
