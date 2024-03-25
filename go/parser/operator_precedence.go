package parser

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
