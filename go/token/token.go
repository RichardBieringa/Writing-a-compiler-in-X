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
