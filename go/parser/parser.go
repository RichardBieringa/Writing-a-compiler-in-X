package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{l: l}

	// reads the first two tokens such that
	// currentToken and peekToken are set
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
