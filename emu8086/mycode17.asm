org 100h
mov bx, offset lst1    ;var ni 1chi elementini addresini oldi
mov al, [bx]   ; al ga bx dagi addresni kochirdi

inc bx   ;bx = bx + 1

mov al, [bx]

inc bx

mov al, [bx]

mov dl, [bx]
ret
lst1 db 11h, 22h, 33h, 44h
end


