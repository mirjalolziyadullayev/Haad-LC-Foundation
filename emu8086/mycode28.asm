org 100h

mov al, 50           ;   50 * 10 = 500 (MUL= (AX = AL*Operand))   2 ta 8 bitni kopaytirganda 16 bit boladi va 16 bitli registerga yozadi
mul numByte          ;

mov ax, sonWord      
mul numWord

mov al, 5
mul bl

mov ax, 500
mul bx

mov ax, 50           ;   50 / 10 = 5 (DIV= (AL = AX/Operand))   2 ta 16 bitni bolganda 8 bit boladi va 8 bitli registerlarga yoziladi
div numByte

mov ax, 50000
div numWord

ret 
numBayt db 10
numWord dw 10
end

;when operand is a byte:
;AX = AL * operand.         

;when operand is a word:
;(DX AX) = AX * operand.     
          
          
          
;when operand is a byte:
;AL = AX / operand
;AH = remainder (modulus)   

;when operand is a word:
;AX = (DX AX) / operand
;DX = remainder (modulus)