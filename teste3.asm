.globl dobro
dobro:
;inicio prologue
push %ebp
movl %esp, %ebp
;fim prologue
movl	$2,%eax
;inicio epilogue
movl %ebp, %esp
pop %ebp
ret
;fim epilogue
.globl main
main:
;inicio prologue
push %ebp
movl %esp, %ebp
;fim prologue
movl	$2,%eax
;inicio epilogue
movl %ebp, %esp
pop %ebp
ret
;fim epilogue
