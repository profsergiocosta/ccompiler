.globl main
main:
#inicio prologue
push %ebp
movl %esp, %ebp
#fim prologue
movl	$72,%eax
pushl %eax
movl	$72,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$101,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$108,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$108,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$111,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$44,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$32,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$87,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$111,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$114,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$108,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$100,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$33,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$10,%eax
pushl %eax
call putchar
addl $0x4, %esp
movl %eax, -4(%ebp)
movl	$0,%eax
#inicio epilogue
movl %ebp, %esp
pop %ebp
ret
#fim epilogue
