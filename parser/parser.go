package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/profsergiocosta/ccompiler/ast"
	"github.com/profsergiocosta/ccompiler/lexer"
	"github.com/profsergiocosta/ccompiler/token"
)

const (
	XML = "XML"
	VM  = "VM"
)

type Parser struct {
	l             *lexer.Lexer
	curToken      token.Token
	peekToken     token.Token
	errors        []string
	fileName      string
	whileLabelNum int
	ifLabelNum    int
}

func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func New(pathName string) *Parser {

	input, err := ioutil.ReadFile(pathName)
	if err != nil {
		panic("erro")
	}
	l := lexer.New(string(input))

	p := &Parser{l: l}
	p.fileName = FilenameWithoutExtension(path.Base(pathName))

	p.nextToken()
	p.whileLabelNum = 0
	p.ifLabelNum = 0
	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {

	p.expectPeek(token.INT)
	p.expectPeek(token.IDENT)

	fun := &ast.Function{Token: p.curToken}

	p.expectPeek(token.LPAREN)
	p.expectPeek(token.RPAREN)

	p.expectPeek(token.LBRACE)

	fun.Statements = []ast.Statement{}

	for !p.peekTokenIs(token.RBRACE) {
		stmt := p.parseStatement()
		if stmt != nil {
			fun.Statements = append(fun.Statements, stmt)
		}
	}

	p.expectPeek(token.RBRACE)

	//fun.Statement = stmt

	program := &ast.Program{Function: fun}

	return program

}

func (p *Parser) parseStatement() ast.Statement {

	switch p.peekToken.Type {
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {

	fmt.Println("----")
	p.expectPeek(token.RETURN)

	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	stmt.ReturnValue = p.parseExpression()

	fmt.Println(p.curToken.Literal)

	p.expectPeek(token.SEMICOLON)

	fmt.Println(p.curToken.Literal)

	return stmt
}

func (p *Parser) parseExpression() ast.Expression {

	exp := p.parseTerm()

	for !p.peekTokenIs(token.EOF) && (p.peekTokenIs(token.PLUS) || p.peekTokenIs(token.MINUS)) {
		p.nextToken()
		operator := p.curToken
		left := exp
		exp = &ast.BinaryExpression{Operator: operator, Left: left, Right: p.parseTerm()}
	}
	return exp
}

func (p *Parser) parseTerm() ast.Expression {
	exp := p.parseFactor()

	for !p.peekTokenIs(token.EOF) && (p.peekTokenIs(token.SLASH) || p.peekTokenIs(token.ASTERISK)) {
		p.nextToken()
		operator := p.curToken
		left := exp
		exp = &ast.BinaryExpression{Operator: operator, Left: left, Right: p.parseTerm()}
	}
	return exp
}

func (p *Parser) parseFactor() ast.Expression {
	switch p.peekToken.Type {
	case token.INTCONST:
		return p.parseIntegerLiteral()
	case token.MINUS, token.NOT, token.BANG:
		return p.parseUnaryExpression()
	}
	return nil
}

func (p *Parser) parseUnaryExpression() ast.Expression {
	p.nextToken()
	operator := p.curToken
	exp := &ast.UnaryExpression{Operator: operator,
		Right: p.parseExpression()}
	fmt.Println(exp.Operator)
	return exp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	p.nextToken()

	lit := &ast.IntegerLiteral{Token: p.curToken}

	val, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = val
	return lit
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) {
	if p.peekTokenIs(t) {
		p.nextToken()
	} else {
		p.peekError(t, p.peekToken.Line)
		fmt.Println(p.errors)
		os.Exit(1)
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType, line int) {
	msg := fmt.Sprintf(" %v: expected next token to be %s, got %s instead",
		line, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
