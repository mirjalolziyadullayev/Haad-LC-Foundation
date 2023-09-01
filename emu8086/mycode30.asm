org 100h

include 'emu8086.inc'       ; library

;LOOP

lea bx, lst1 
lea bp, lst2

mov cx, 4  
mov si, 2
mov di, 0

takrorla:  

mov ax, [bx+di] 
div si  
   
mov [bp+di], al

inc di 



loop takrorla


ret

lst1 db 10, 40, 60, 30    
lst2 db 4 dup(?)
