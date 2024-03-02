// Package token contains the various types of tokens used in the monkey
// programming language
//
// These tokens are constructed by the Lexer when it consumes the source code.
// Which basically means, trying to map the text to some meaninful concept.
//
// There are various kinds of Tokens that we try to identify, for example
// language specific keywords such as 'return', symbols that carry a meaning
// as parenthesis, or identifiers for variables or functions.
package token

// Represents the type of the token, e.g. an INT or an IDENTIFIER
type TokenType string

type Token struct {
	Type    TokenType
	Literal string // The literal value of the token
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Comparison
	GT     = ">"
	LT     = "<"
	EQ     = "=="
	NOT_EQ = "!="

	// Bools
	FALSE = "FALSE"
	TRUE  = "TRUE"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)
