
; You may customize this and other start-up templates; 
; The location of this template is c:\emu8086\inc\0_com_template.txt

org 100h

lea bx, son
mov al, 2h

xlatb    ;al ga bx + al ni qiymatini yozadi  (al = bx + al)  (bx da address bor)  (al da 2 bor) 

ret
son db 11h, 12h, 13h, 14h
end     
     



