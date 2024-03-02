package ast

// Contains the common interfaces for the nodes in the AST
// Each Node in the AST should implement the Node interface
//
// Furthermore, we make the distinction between statement nodes and expression
// nodes:
// `Expressions` produce values whereas `Statements` do not

type Node interface {
	// Only used for debugging
	TokenLiteral() string
}

// An expression is a value producing node
// Examples: 5, "a-string", add(1, 2)
type Expression interface {
	Node
	// marker method to indicate type of node
	expressionNode()
}

// A statement does not produce a value
// Examples: return 5, let x = 3, switch expression vs statement in some langs
type Statement interface {
	Node
	// marker method to indicate type of node
	statementNode()
}
