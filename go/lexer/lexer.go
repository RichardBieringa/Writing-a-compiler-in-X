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
	return token.Token{
		Type:    "Test",
		Literal: "Test",
	}
}
