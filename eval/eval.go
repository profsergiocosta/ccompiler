package eval

import (

	"fmt"
	"os"

	"github.com/profsergiocosta/ccompiler/ast"
)

func Eval(node ast.Node, asm *os.File) {

	switch node := node.(type) {
		// Statements
	
		case *ast.Program:
			evalProgram(node, asm)

		case *ast.Function:
			evalFunction(node, asm)
			
	}


}

func evalProgram(program *ast.Program, asm *os.File) {

	Eval(program.Function, asm)

}

func evalFunction(function *ast.Function, asm *os.File) {

	s := fmt.Sprintf(".globl _%s\n",function.Token.Literal)
	asm.WriteString(s)

	s = fmt.Sprintf("_%s\n",function.Token.Literal)
	asm.WriteString(s)

	Eval(function.Statement, asm)

}