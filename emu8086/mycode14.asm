org 100h

mov ax, ozg1

mov bx, ax ; registerlar orasida malumot almashtira olar ekanmiz
mov cx, bx 
mov dx, cx

;mov al, bx
;mov ax, bh

;mov es, 12h
mov es. ax

;mov cs, 12h    code segment ga hech qanaqa o'zgartirish kirita olmas ekanmiz
;mov cs, es
;mov cs, ax

mov cx, cs ; code segment dan malumot olishimiz mumkin
;mov cs, es

mov ax, cs  
mov es. ax

mov si, 11h     ; pointerlarni malumotini o'zgartirishimiz mumkin ekan
mov di, 22h
mov si, di                 

mov ah, [011Ah]  ; agar brackets ichida yozilsa u RAM dagi address hisoblanadi
ret
                
ozg1 dw 11h               
end                        