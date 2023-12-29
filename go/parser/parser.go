package parser

import (
	"log/slog"
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
	slog.Debug("Constructed a Parser")
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

// check if the current token matches the expected type
func (p *Parser) currentTokenIs(tokenType token.TokenType) bool {
	return p.currentToken.Type == tokenType
}

// check if the peek token matches the expected type
func (p *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}

// assert the peek token, and consume the current token
func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	}
	slog.Info(
		"Expect Peek",
		"expected", tokenType,
		"actual", p.peekToken.Type,
	)

	return false
}

// Parses the source code into one AST
// The result is a `Program`, that contains a series of `Statements`
// Each of these `Statements` are tree-like structures that contain
// other `Statements` and/or `Expressions`
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Look through all tokens of the lexer, try to produce statements
	for p.currentToken.Type != token.EOF {
		slog.Debug("ParseProgram", "token", p.currentToken)
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

// Tries parsing a statement based on the current token
func (p *Parser) parseStatement() ast.Statement {
	slog.Info(
		"Parser - parseStatement",
		"token", p.currentToken,
	)
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// let five = 5;
func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: assign expression value
	// p.expectPeek(token.INT)
	// statement.Value = &ast.Expression{}

	// Consume all tokens up to the semicolon
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	slog.Info("LET STATEMENT",
		"name", statement.Name,
		"value", statement.Value,
	)

	return statement
}
