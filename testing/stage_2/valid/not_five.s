.globl main
main:
movl	$5,%eax

cmpl   $0, %eax    ;set ZF on if exp == 0, set it off otherwise
movl   $0, %eax    ;zero out EAX (doesn't change FLAGS)
sete   %al
ret
