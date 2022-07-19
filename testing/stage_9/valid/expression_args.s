.globl add
add:
#inicio prologue
push %ebp
movl %esp, %ebp
#fim prologue
movl	$5,%eax
pushl %eax
movl 8(%ebp), %eax
push %eax  
movl 12(%ebp), %eax
pop %ecx
addl %ecx, %eax
push %eax  
movl -4(%ebp), %eax
pop %ecx
addl %ecx, %eax
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
movl	$15,%eax
pushl %eax
movl	$30,%eax
pushl %eax
call add
addl $0x8, %esp
push %eax  
movl	$2,%eax
pop %ecx
imul %ecx, %eax
pushl %eax
movl -16(%ebp), %eax
#inicio epilogue
movl %ebp, %esp
pop %ebp
ret
#fim epilogue
