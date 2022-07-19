.globl _putchar
.globl main
main:
#inicio prologue
push %ebp
movl %esp, %ebp
#fim prologue
movl	$72,%eax
pushl %eax
movl -4(%ebp), %eax
pushl %eax
call putchar@PLT
addl $0x4, %esp
movl %eax, -4(%ebp)
