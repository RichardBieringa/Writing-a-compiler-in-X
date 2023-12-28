package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input string

	position     int  // current index position in the source code
	readPosition int  // position + 1
	currentChar  byte // current char
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()

	return lexer
}

func (l *Lexer) readChar() {
	// reached EOF
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
		return
	}

	// consume next character
	l.currentChar = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	t := token.Token{
		Literal: string(l.currentChar),
	}

	switch l.currentChar {
	case '=':
		t.Type = token.ASSIGN
	case '+':
		t.Type = token.PLUS
	case ',':
		t.Type = token.COMMA
	case ';':
		t.Type = token.SEMICOLON
	case '(':
		t.Type = token.LPAREN
	case ')':
		t.Type = token.RPAREN
	case '{':
		t.Type = token.LBRACE
	case '}':
		t.Type = token.RBRACE
	default:
		if isLetter(l.currentChar) {
			t.Literal = l.readIdentifier()
			t.Type = getIdentifier(t.Literal)
		} else {
			t.Type = token.ILLEGAL
		}
	}

	l.readChar()
	return t
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func getIdentifier(identifier string) token.TokenType {
	keywords := map[string]token.TokenType{
		"let": token.LET,
	}

	tokenType, ok := keywords[identifier]
	if ok {
		return tokenType
	}

	return token.IDENT
}

func (l *Lexer) readIdentifier() string {
	left := l.position

	for isLetter(l.currentChar) {
		l.readChar()
	}

	return l.input[left:l.position]
}
