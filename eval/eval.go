package eval

import (
	"fmt"

	"github.com/profsergiocosta/ccompiler/ast"
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
		
		case *ast.IntegerLiteral:
			return evalIntegerLiteral(node)
			
	}

	return ""
}

func evalProgram(program *ast.Program) string {

	return Eval(program.Function)



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
	s := fmt.Sprintf("movl \t$%v, %%eax\n", val.Value )
	return s

}