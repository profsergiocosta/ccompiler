package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/profsergiocosta/ccompiler/symboltable"

	"github.com/profsergiocosta/ccompiler/lexer"
	"github.com/profsergiocosta/ccompiler/token"
	"github.com/profsergiocosta/ccompiler/ast"
)

const (
	XML = "XML"
	VM  = "VM"
)

type Parser struct {
	l             *lexer.Lexer
	curToken      token.Token
	peekToken     token.Token
	output        string
	errors        []string
	st            *symboltable.SymbolTable
	className     string
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

	p.st = symboltable.NewSymbolTable()



	p.nextToken()
	p.whileLabelNum = 0
	p.ifLabelNum = 0
	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Compile() {
	//p.CompileProgram()

}



func (p *Parser) ParseProgram() *ast.Program {

	

	p.expectPeek(token.INT)
	p.expectPeek(token.IDENT)
	
	fun := &ast.Function{Token: p.curToken}

	// parameters
	p.expectPeek(token.LPAREN)
	p.expectPeek(token.RPAREN)

	p.expectPeek(token.LBRACE)

	p.expectPeek(token.RETURN)

	stmt := p.parseReturnStatement()


	p.expectPeek(token.RBRACE)

	fun.Statement = stmt

	program := &ast.Program{Function: fun}

	return program



}



func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	p.nextToken()

	stmt.ReturnValue = p.parseIntegerLiteral()

	for p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}


func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
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

//// auxiliars

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (p *Parser) curTokenAsInt() int {
	i1, err := strconv.Atoi(p.curToken.Literal)
	check(err)
	return i1
}


