org 100h

mov bx, offset ozg1  
mov al, [bx]

inc bx
mov al, [bx]

inc bx
mov al, [bx]
 
inc bx
mov al, [bx]               ; increment ===== x = x + 1

ret
ozg1 db 11h, 22h, 33h, 44h
end