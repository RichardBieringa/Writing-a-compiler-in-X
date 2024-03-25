// Pagage parser contains a recursive descent parser implementation
// for the monkey programming language
//
// The parser constructs an Abstract Syntax Tree (AST) from a series of tokens
// produced by the lexer. It aims to accurately represent the program's source
// code through some of the language's rules.
//
// Example of the parsing process
// -- done by lexer
// source code `let x = 5;`
// tokens = `[LET, IDENTIFIER , EQ, INT, SEMICOLON]`;
// -- done by parser
// AST:
// - PROGRAM_NODE
//   - LET_STATEMENT_NODE
//   - NAME = IDENTIFIER_NODE
//   - VALUE = EXPRESSION_NODE
package parser

import (
	"fmt"
	"log/slog"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type (
	// parses prefix expressions, e.g. `++var`
	prefixParseFn func() ast.Expression

	// parses infix expressions, e.g. `5 + 5`
	infixParseFn func(ast.Expression) ast.Expression
)

// The Parses uses the tokens produces by the lexer
// this implementation only looks at the current, and next token (1 lookahead)
// which means that we can only the type of node by looking at these 2 tokens
type Parser struct {
	l *lexer.Lexer

	errors []string // Holds any errors that occured during parsing

	currentToken token.Token // The current token that the parser is consuming
	peekToken    token.Token // The next token, used for 1 node lookahead

	prefixParseMap map[token.TokenType]prefixParseFn
	infixParseMap  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	slog.Debug("Constructed a Parser")
	parser := &Parser{
		l:      l,
		errors: []string{},

		prefixParseMap: make(map[token.TokenType]prefixParseFn),
		infixParseMap:  make(map[token.TokenType]infixParseFn),
	}

	parser.registerPrefixParseFn(token.IDENT, parser.parseIdentifier)

	// reads the first two tokens such that
	// currentToken and peekToken are set
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (p *Parser) registerPrefixParseFn(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseMap[tokenType] = fn
}

func (p *Parser) registerInfixParseFn(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseMap[tokenType] = fn
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

// asserts that the peekToken matches the expected token type
func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	}

	p.peekError(tokenType)
	return false
}

func (p *Parser) peekError(tokenType token.TokenType) {
	err := fmt.Sprintf("Expected next token to be %q, received: %q",
		tokenType, p.peekToken)

	p.errors = append(p.errors, err)
}

// Parses the source code into one AST
// The result is a `Program`, that contains a series of `Statements`
// Each of these `Statements` are tree-like structures that contain
// other `Statements` and/or `Expressions`
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

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

// Returns a list of errors the parser encoutered
func (p *Parser) Errors() []string {
	return p.errors
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
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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
		Token:      p.currentToken,
		Identifier: p.currentToken.Literal,
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

	slog.Debug("LET STATEMENT",
		"name", statement.Name,
		"value", statement.Value,
	)

	return statement
}

// return 5; return myFunctionCall(2, 4);
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{
		Token: p.currentToken,
	}

	p.nextToken()

	// TODO: parse expressions
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	slog.Debug("RETURN STATEMENT",
		"value", statement.Expression,
	)

	return statement
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{
		Token: p.currentToken,
	}

	statement.Expression = p.parseExpression(LOWEST)

	// Allow expressiosn to be without semis
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedenceLevel int) ast.Expression {
	prefix := p.prefixParseMap[p.currentToken.Type]
	if prefix == nil {
		return nil
	}

	leftExpression := prefix()

	return leftExpression
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token:      p.currentToken,
		Identifier: p.currentToken.Literal,
	}
}
