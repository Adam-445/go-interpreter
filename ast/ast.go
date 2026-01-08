// Package ast implements the Abstract Syntax Tree
// AKA the representation of a program in memorry so a computer understands it
package ast

import "github.com/Adam-445/go-interpreter/token"

type Node interface {
	TokenLiteral() string // Every AST node must be able to tell what token it came from
}

type Statement interface {
	Node
	statementNode() // This type claims to be a statement
}

type Expression interface {
	Node
	expressionNode() // This type claims to be an Expression
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
