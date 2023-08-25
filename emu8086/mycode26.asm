
data segment
    highBit db 2 dup(?)
    lowBit db 2 dup(?)
ends

stack segment
    dw   128  dup(0)
ends

code segment
start:
; set segment registers:
    mov ax, data
    mov ds, ax
    mov es, ax

    ; add your code here
    
    push 2244h
    push 3355h
    
    pop ax          ; 3355 -> AX == AH=33 AL=55
    
    mov highBit[0], ah
    mov lowBit[0], al
    
    pop ax          ; 2244 -> AX == AH=22 AL=44
    
    mov highBit[1], ah
    mov lowBit[1], al

    
    ;
    mov ax, 4c00h ; exit to operating system.
    int 21h    
ends

end start ; set entry point and stop the assembler.
