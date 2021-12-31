package gen

import (
	"fmt"

	"ccompiler/ast"
	"ccompiler/symboltable"
	"ccompiler/token"
)

var st = symboltable.NewSymbolTable() // global por agora

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

	case *ast.DeclareStatement:
		return genDeclareStatement(node)

	case *ast.AssignStatement:
		return genAssignStatement(node)

	}

	return ""
}

func genAssignStatement(dec *ast.AssignStatement) string {
	symbol, _ := st.Resolve(dec.VarName.Literal)
	index := symbol.Index
	s := genExpression(dec.Init)
	s = s + fmt.Sprintf("movl %%eax, %d(%%ebp)\n", index)
	return s

}

func genDeclareStatement(dec *ast.DeclareStatement) string {
	st.Define(dec.VarName.Literal)
	s := ""
	if dec.Init == nil {
		s = s + "movl  $0,%eax\n"
	} else {
		s = s + genExpression(dec.Init)
	}
	s = s + "pushl %eax\n"
	return s
}

func genProgram(program *ast.Program) string {

	s := ""
	for _, f := range program.Functions {
		s = s + genFunction(&f)
	}
	return s

}

func genExpression(expression ast.Expression) string {

	switch node := expression.(type) {
	case *ast.IntegerLiteral:
		return genIntegerLiteral(node)

	case *ast.UnaryExpression:
		return genUnaryExpression(node)

	case *ast.BinaryExpression:
		return genBinaryExpression(node)

	case *ast.VarExpression:
		return genVarExpression(node)

	}
	return ""
}

func genVarExpression(exp *ast.VarExpression) string {
	symbol, _ := st.Resolve(exp.Name.Literal)
	index := symbol.Index
	s := fmt.Sprintf("movl %d(%%ebp), %%eax\n", index)
	return s
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

func genFunctionEpilogue() string {
	s :=
		`;inicio epilogue
movl %ebp, %esp
pop %ebp
ret
;fim epilogue
`
	return s
}

func genFunction(function *ast.Function) string {

	s := fmt.Sprintf(".globl %s\n", function.Token.Literal)

	s = s + fmt.Sprintf("%s:\n", function.Token.Literal)
	s = s + ";inicio prologue\n"
	s = s + "push %ebp\n"
	s = s + "movl %esp, %ebp\n"
	s = s + ";fim prologue\n"

	for _, st := range function.Statements {
		s = s + Generate(st)
	}

	return s

}

func genReturnStatement(ret *ast.ReturnStatement) string {
	s := Generate(ret.ReturnValue)
	s = s + genFunctionEpilogue()
	return s
}

func genIntegerLiteral(val *ast.IntegerLiteral) string {
	s := fmt.Sprintf("movl\t$%v,%%eax\n", val.Value)
	return s

}
