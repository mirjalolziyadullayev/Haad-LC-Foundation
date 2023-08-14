org 100h

lea bx, son1 
mov ah, [bx]
mov [bx+1], 15h

lea bx, son2
mov dx, [bx+6]
         
mov ah, 11h
mov al, 22h

xchg al, ah        ;    2 ta register o'rnini almashtiradi

mov dl, ah         ;    shu uch satrda xchg bajaradigan operatsiyani mov commandi bilan bajardik
mov ah, al         ;
mov al, dl         ;
                  
ret
son1 db 11h, 12h
son2 dw 12h, 13h, 14h, 15h