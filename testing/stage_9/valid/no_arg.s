.globl three
three:
#inicio prologue
push %ebp
movl %esp, %ebp
#fim prologue
movl	$35,%eax
#inicio epilogue
movl %ebp, %esp
pop %ebp
ret
#fim epilogue
.globl main
main:
#inicio prologue
push %ebp
movl %esp, %ebp
#fim prologue
call three
#inicio epilogue
movl %ebp, %esp
pop %ebp
ret
#fim epilogue
