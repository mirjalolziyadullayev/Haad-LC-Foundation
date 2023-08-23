
org 100h

lea bx, numbers 


mov di, 0
mov cx, 3 

takrorla:
  
mov al, [bx+di]   
mov []



inc di 

loop takrorla


ret

numbers db 1h, 2h, 3h, 4h
    
end




