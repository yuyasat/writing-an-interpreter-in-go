package ast

import (
	"bytes"
	"token"
)

// Node is [TODO]
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier is [TODO]
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral is [TODO]
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// ReturnStatement is [TODO]
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is [TODO]
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is [TODO]
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral is [TODO]
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// IntegerLiteral is [TODO]
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral is [TODO]
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// PrefixExpression is [TODO]
type PrefixExpression struct {
	Token    token.Token // 前置トークン、例えば「!」
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral is [TODO]
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is [TODO]
type InfixExpression struct {
	Token    token.Token // 前置トークン、例えば「+」
	Left     Expression
	Operator string
	Right    Expression
}

func (pe *InfixExpression) expressionNode() {}

// TokenLiteral is [TODO]
func (pe *InfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(" " + pe.Operator + " ")
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
