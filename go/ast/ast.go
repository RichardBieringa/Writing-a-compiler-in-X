package ast

import "monkey/token"

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value *Expression
}

func (n *LetStatement) statementNode()       {}
func (n *LetStatement) TokenLiteral() string { return n.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// Identifier is an expression to allow `let a = anotherVar`
func (n *Identifier) expressionNode()      {}
func (n *Identifier) TokenLiteral() string { return n.Token.Literal }
