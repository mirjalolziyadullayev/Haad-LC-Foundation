org 100h

include 'emu8086.inc'       ; library

;LOOP

lea bx, lst1 
lea bp, lst2

mov cx, 4  
mov si, 2
mov di, 0

takrorla:  

mov ah, [bx+di] 
div si  
   
mov [bp+di], ah

inc di 



loop takrorla

ret

lst1 db 2h, 4h, 6h, 8h   
lst2 db 4 dup(?)
