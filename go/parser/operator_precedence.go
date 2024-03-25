package parser

import "monkey/token"

// Operator precedence levels for the Monkey programming language
// Ranges from 1 (lowest) - 7 (highest)
const (
	_ int = iota
	LOWEST
	EQUALS  // == LESSGREATER // > or <
	SUM     //+
	PRODUCT //*
	PREFIX  //-Xor!X
	CALL    // myFunction(X)
)

var precedenceMap = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       EQUALS,
	token.GT:       EQUALS,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}
