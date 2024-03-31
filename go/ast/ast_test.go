package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token:      token.Token{Type: token.IDENT, Literal: "myVariableIdentifier"},
					Identifier: "myVariableIdentifier",
				},
				Value: &Identifier{
					Token:      token.Token{Type: token.IDENT, Literal: "someOtherIdent"},
					Identifier: "someOtherIdent",
				},
			},
		},
	}

	if program.String() != "let myVariableIdentifier = someOtherIdent;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
