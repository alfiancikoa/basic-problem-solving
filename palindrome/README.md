# Soal Palindrome

<p align=justify><b>Palindrom</b> adalah sebuah kata, frasa, angka maupun susunan lainnya yang dapat dibaca dengan sama baik dari depan maupun belakang (spasi antara huruf-huruf biasanya diperbolehkan). Kata "palindrom" berasal dari bahasa Yunani: palin (πάλιν; "lagi") and dromos (δρóμος; "arah").<br>(Sumber: https://id.wikipedia.org/wiki/Palindrom)</p>
Contoh Palindrome:
<br>

```
S A T O R
A R E P O
T E N E T
O P E R A
R O T A S
```

### Problem:

Buatlah Function untuk mengecek apakah kata atau kalimat yang dimasukkan sebagai input merupakan palindrome?
<br>

### Input:

Berupa kalimat atau kata bertipe string yang panjangnya sampai 65535 karakter
<br>

### Output:

Hasil berupa <b>True</b> atau <b>False</b>
<br>

### example:

1. Input = "katak" <br>         Output = true
2. Input = "saras" <br>         Output = true
3. Input = "makan" <br>         Output = false
4. Input = "nasigoreng"<br>     Output = false

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```

