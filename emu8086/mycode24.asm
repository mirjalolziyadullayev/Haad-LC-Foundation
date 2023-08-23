data segment
    list1 db 11h, 22h, 33h, 44h
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
     
     push 12h       ;   LIFO (Last In First Out) standarti boyicha ishlaydi
     push 22h       ;   stack ga joylaydi
     push 33h       ;
     
     pop bx         ;   stack dagi oxrgi qiymatni yani 33h ni oladi va BX ga tenglaydi va bitta ortga qaytadi yani 22h
     
     
     
     
    ; 
    mov ax, 4c00h ; exit to operating system.
    int 21h    
ends

end start ; set entry point and stop the assembler.
