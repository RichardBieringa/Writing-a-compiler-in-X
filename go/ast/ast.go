// Package ast contains the implementation of the Abstract Syntax Tree and the
// nodes it consists of.
//
// Each program is 'parsed' into an AST that is some internal representation
// of the source code of the program. This is in the form of a tree-like
// structure, with the root node being the program itself. It is 'abstract'
// since it can omits some less important details such as whitespace, certain
// symbols etc.
package ast

import "monkey/token"

// A let statement binds an identifier to some value produced by an expression
// Example: `let myIdentifier = 5;`
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
