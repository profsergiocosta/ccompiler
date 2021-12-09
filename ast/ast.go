package ast

import (
	"bytes"

	"github.com/profsergiocosta/ccompiler/token"
)

/*
program = Program(function_declaration)
function_declaration = Function(string, statement) //string is the function name
statement = Return(exp)
exp = Constant(int)
*/

// Node represents an AST node.
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Function *Function
}

// TokenLiteral returns the first token literal of a program.
func (p *Program) TokenLiteral() string {
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	out.WriteString("(Program ")

	out.WriteString(p.Function.String())

	out.WriteString(" )")

	return out.String()
}

type Function struct {
	Token      token.Token
	Statements []Statement
}

func (fl *Function) expressionNode() {}

// TokenLiteral returns a token literal of function.
func (fl *Function) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *Function) String() string {
	var out bytes.Buffer

	out.WriteString("( " + fl.TokenLiteral() + " ")

	for _, st := range fl.Statements {
		out.WriteString(st.String())
	}

	out.WriteString(" )")

	return out.String()
}

// ReturnStatement represents a return statement.
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns a token literal of return statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + "  ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// IntegerLiteral represents an integer literal.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns a token literal of integer.
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type UnaryExpression struct {
	Operator token.Token
	Right    Expression
}

func (exp *UnaryExpression) expressionNode()      {}
func (exp *UnaryExpression) TokenLiteral() string { return exp.Operator.Literal }
func (exp *UnaryExpression) String() string {
	return exp.Operator.Literal
}

type BinaryExpression struct {
	Operator token.Token
	Left     Expression
	Right    Expression
}

func (exp *BinaryExpression) expressionNode()      {}
func (exp *BinaryExpression) TokenLiteral() string { return exp.Operator.Literal }

func (exp *BinaryExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(" + exp.Operator.Literal + " ")

	out.WriteString(exp.Left.String() + " ")
	out.WriteString(exp.Right.String() + " ")

	out.WriteString(" )")

	return out.String()
}
