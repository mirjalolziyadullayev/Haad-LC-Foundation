
; You may customize this and other start-up templates; 
; The location of this template is c:\emu8086\inc\0_com_template.txt

org 100h

mov al, 12h
mov ax, 1458h

;mov bh, 4567h

mov bx, 56h
mov cx, 0A1B4h

mov dh, 'h'
;mov dl, 'ab'
mov dx, 'ab'

mov ch, 0011b
mov cl, 00110011b
;mov dl, 111111111b 

mov dx, 1111111111111111b

ret

ozg1 db 11h
end