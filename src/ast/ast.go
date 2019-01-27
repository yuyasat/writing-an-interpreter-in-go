package ast

import "token"

// Node is [TODO]
type Node interface {
	TokenLiteral() string
}

// Statement is [TODO]
type Statement interface {
	Node
	statementNode()
}

// Expression is [TODO]
type Expression interface {
	Node
	expressionNode()
}

// Program is [TODO]
type Program struct {
	Statements []Statement
}

// TokenLiteral is [TODO]
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is [TODO]
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral is [TODO]
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is [TODO]
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral is [TODO]
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement is [TODO]
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is [TODO]
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
