// Package ast contains the implementation of the Abstract Syntax Tree and the
// nodes it consists of.
//
// Each program is 'parsed' into an AST that is some internal representation
// of the source code of the program. This is in the form of a tree-like
// structure, with the root node being the program itself. It is 'abstract'
// since it can omits some less important details such as whitespace, certain
// symbols etc.
package ast

import (
	"bytes"
	"monkey/token"
)

// A let statement binds an identifier to some value produced by an expression
// Example: `let myIdentifier = 5;`
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (n *LetStatement) statementNode()       {}
func (n *LetStatement) TokenLiteral() string { return n.Token.Literal }
func (n *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(n.TokenLiteral() + " ")
	out.WriteString(n.Name.String())
	out.WriteString(" = ")

	if n.Value != nil {
		out.WriteString(n.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token      token.Token // token.IDENT
	Identifier string
}

// Identifier is an expression to allow `let a = anotherVar`
func (n *Identifier) expressionNode()      {}
func (n *Identifier) TokenLiteral() string { return n.Token.Literal }
func (n *Identifier) String() string       { return n.Identifier }

// A return statement returns a value
// Example: `return add(5);`
type ReturnStatement struct {
	Token      token.Token // Token.RETURN
	Expression Expression
}

func (n *ReturnStatement) statementNode()       {}
func (n *ReturnStatement) TokenLiteral() string { return n.Token.Literal }
func (n *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(n.TokenLiteral() + " ")

	if n.Expression != nil {
		out.WriteString(n.Expression.String())
	}

	out.WriteString(";")
	return out.String()
}

// A line that only contains some expression
// Example: `x + 10;`
type ExpressionStatement struct {
	Token      token.Token // Token.IDENT
	Expression Expression
}

func (n *ExpressionStatement) statementNode()       {}
func (n *ExpressionStatement) TokenLiteral() string { return n.Token.Literal }
func (n *ExpressionStatement) String() string {
	if n.Expression != nil {
		return n.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token // Token.INT
	Value int64
}

func (n *IntegerLiteral) expressionNode()      {}
func (n *IntegerLiteral) TokenLiteral() string { return n.Token.Literal }
func (n *IntegerLiteral) String() string       { return n.Token.Literal }

type PrefixExpression struct {
	Token    token.Token // the operator token, e.g. `+`
	Operator string
	Value    Expression
}

func (n *PrefixExpression) expressionNode()      {}
func (n *PrefixExpression) TokenLiteral() string { return n.Token.Literal }
func (n *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString(n.Operator)

	if n.Value != nil {
		out.WriteString(n.Value.String())
	}

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // the operator token, e.g. `+`
	Left     Expression
	Operator string
	Right    Expression
}

func (n *InfixExpression) expressionNode()      {}
func (n *InfixExpression) TokenLiteral() string { return n.Token.Literal }
func (n *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")

	if n.Left != nil {
		out.WriteString(n.Left.String())
	}

	out.WriteString(" " + n.Operator + " ")

	if n.Right != nil {
		out.WriteString(n.Right.String())
	}

	out.WriteString(")")

	return out.String()
}
