.globl twice
twice:
#inicio prologue
push %ebp
movl %esp, %ebp
#fim prologue
movl	$2,%eax
push %eax  
movl 8(%ebp), %eax
pop %ecx
imul %ecx, %eax
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
movl	$3,%eax
pushl %eax
call twice
addl $0x4, %esp
#inicio epilogue
movl %ebp, %esp
pop %ebp
ret
#fim epilogue
