
co n bien co

a0 la xac suat bien co 0 xay ra
a1 la xac suat bien co 1 xay ra
b0 la..................0 ko xay ra
b1 la..................1 ko xay ra

a: xay ra
b = 1 - a: ko xay ra

n = 2
xay ra: [a0, a1]
ko xay ra: [b0, b1]

0           1                       2
b0 * b1     a0 * b1 + a1 * b0       a0 * a1

n = 3
xay ra: [a0, a1, a2] = p_list
ko xay ra: [b0, b1, b2]

000 -> 0

010
100 -> 1
001

011
110 -> 2
101

111 -> 3

[
0: b0 * b1 * b2    
1: a0 * b1 * b2 + b0 * a1 * b2 + b0 * b1 * a2
2: a0 * a1 * b2 + a0 * b1 * a2 + b0 * a1 * a2
3: a0 * a1 * a2
]

FFT -> Fast Fourier Transform

(a + bx)(c + dx)

a * c + (a * d + b * c)x + b * d * x^2