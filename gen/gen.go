package gen

import (
	"fmt"

	"github.com/profsergiocosta/ccompiler/ast"
	"github.com/profsergiocosta/ccompiler/token"
)

func Generate(node ast.Node) string {

	switch node := node.(type) {
	// Statements

	case *ast.Program:
		return genProgram(node)

	case *ast.Function:
		return genFunction(node)

	case *ast.ReturnStatement:
		return genReturnStatement(node)

	case ast.Expression:
		return genExpression(node)

	}

	return ""
}

func genProgram(program *ast.Program) string {

	return genFunction(program.Function)

}

func genExpression(expression ast.Expression) string {

	switch node := expression.(type) {
	case *ast.IntegerLiteral:
		return genIntegerLiteral(node)

	case *ast.UnaryExpression:
		return genUnaryExpression(node)

	case *ast.BinaryExpression:
		return genBinaryExpression(node)

	}
	return ""
}

func genBinaryExpression(exp *ast.BinaryExpression) string {
	s := genExpression(exp.Left)
	s = s + "push %eax  \n"
	s = s + genExpression(exp.Right)
	s = s + "pop %ecx\n"
	switch exp.Operator.Type {
	case token.PLUS:
		s = s + "addl %ecx, %eax\n"
	case token.ASTERISK:
		s = s + "imul %ecx, %eax\n"

	}

	return s
}

func genUnaryExpression(exp *ast.UnaryExpression) string {
	s := genExpression(exp.Right)
	fmt.Println(exp.Operator)
	switch exp.Operator.Type {
	case token.MINUS:
		s = s + "neg\t%eax\n"
	case token.NOT:
		s = s + "not\t%eax\n"
	case token.BANG:
		s = s + `
cmpl   $0, %eax    ;set ZF on if exp == 0, set it off otherwise
movl   $0, %eax    ;zero out EAX (doesn't change FLAGS)
sete   %al` + "\n"

	}

	return s
}

func genFunction(function *ast.Function) string {

	s := fmt.Sprintf(".globl %s\n", function.Token.Literal)

	s = s + fmt.Sprintf("%s:\n", function.Token.Literal)

	return s + Generate(function.Statement)

}

func genReturnStatement(ret *ast.ReturnStatement) string {
	s := Generate(ret.ReturnValue)
	s = s + "ret\n"
	return s

}

func genIntegerLiteral(val *ast.IntegerLiteral) string {
	s := fmt.Sprintf("movl\t$%v,%%eax\n", val.Value)
	return s

}
