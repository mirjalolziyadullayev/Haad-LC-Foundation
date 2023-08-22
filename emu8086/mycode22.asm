org 100h

lea bx, lst1 
lea bp, lst2

mov cx, 4
mov di, 3
mov si, 0

takrorla:  
   
mov al, [bx+di] 
mov [bp+si], al   

dec di
inc si

loop takrorla


ret

lst1 db 11h, 22h, 33h, 44h    
lst2 db 4 dup(?)
end


