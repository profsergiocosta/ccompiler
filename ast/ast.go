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


type Function struct {
	Token      token.Token
	Statement  Statement
}

func (fl *Function) expressionNode() {}

// TokenLiteral returns a token literal of function.
func (fl *Function) TokenLiteral() string {
	return fl.Token.Literal
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