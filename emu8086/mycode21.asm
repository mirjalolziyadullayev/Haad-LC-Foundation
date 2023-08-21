
; You may customize this and other start-up templates; 
; The location of this template is c:\emu8086\inc\0_com_template.txt
                    
org 100h

include 'emu8086.inc'       ; library

;LOOP

lea bx, lst1 
lea bp, lst2

mov cx, 4
mov di, 0

takrorla:  
   
mov al, [bx+di] 
mov [bp+di], al

inc di 

loop takrorla


ret

lst1 db 11h, 22h, 33h, 44h    
lst2 db 4 dup(?)


