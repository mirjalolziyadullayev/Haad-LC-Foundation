data segment
    
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
     
     push 1155h 
     push 2266h
                 
     pop bx 
     pop ax
     
     mov cl, al   
     mov al, bh   
     mov bh, cl       
 
    ;
    mov ax, 4c00h ; exit to operating system.
    int 21h    
ends

end start ; set entry point and stop the assembler.
