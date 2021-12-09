package eval

import (
	"fmt"

	"github.com/profsergiocosta/ccompiler/ast"
	"github.com/profsergiocosta/ccompiler/token"
)

func Eval(node ast.Node) string {

	switch node := node.(type) {
		// Statements
	
		case *ast.Program:
			return evalProgram(node)

		case *ast.Function:
			return evalFunction(node)

		case *ast.ReturnStatement:
			return evalReturnStatement(node)
		
		case ast.Expression:
			return evalExpression(node)


			
	}

	return ""
}

func evalProgram(program *ast.Program) string {

	return Eval(program.Function)



}

func evalExpression (expression ast.Expression ) string {

	switch node := expression.(type) {
		case *ast.IntegerLiteral:
			return evalIntegerLiteral(node)

		case *ast.UnaryExpression:
			return evalUnaryExpression(node)
		
		case *ast.BinaryExpression:
			return evalBinaryExpression(node)

	}
	return  ""
}

func evalBinaryExpression(exp *ast.BinaryExpression ) string {
	s := evalExpression(exp.Left)
	s = s + "push %eax  \n";
	s = s + evalExpression(exp.Right)
	s = s + "pop %ecx\n"
	switch exp.Operator.Type {
			case token.PLUS:
				s = s + "addl %ecx, %eax\n"
			case token.ASTERISK:
				s = s + "imul %ecx, %eax\n"

	}

	return s
}

func evalUnaryExpression(exp *ast.UnaryExpression ) string {
	s := evalExpression(exp.Right)
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

func evalFunction(function *ast.Function) string {

	s := fmt.Sprintf(".globl %s\n",function.Token.Literal)

	s = s + fmt.Sprintf("%s:\n",function.Token.Literal)

	return s + Eval(function.Statement)


}

func evalReturnStatement (ret *ast.ReturnStatement) string {
	s := Eval(ret.ReturnValue)
	s = s + "ret\n"
	return s

}

func evalIntegerLiteral (val *ast.IntegerLiteral) string {
	s := fmt.Sprintf("movl\t$%v,%%eax\n", val.Value )
	return s

}