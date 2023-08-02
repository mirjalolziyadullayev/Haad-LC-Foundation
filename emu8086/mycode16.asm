org 100h

mov bx, offset son1; offset o'zgaruvchini 1 elementini ramdagi adresini oladi
mov al, [bx] ;qavs ichida yozilgan bx son 1 ni 1 chi elementini adresidagi qiymatini oladi

mov al, [bx+1] ;qavs ichidagi bx+1 bx ga son1 dagi 1 chi element dan keyingi 2 chi element ni oldi
mov al, [bx+2] ;qavs ichidagi bx+1 bx ga son1 dagi 1 chi element dan keyingi 3 chi element ni oldi

mov di, 3h    ;di ga 3 qiymati tenglandi
mov al, [bx+di]    ; qavs ichida bx ga di dagi 3 qiymatini qoshildi yani bx+3 boldi

mov [bx+1], 13h     ; bx+1, 13h son1 dagi 2 chi elementni 13ga ozgartiradi

ret 

son1 db 8h, 11h, 10h, 12h
end