# Soal Bilangan Prima

<p align=justify><b>Tapis Eratosthenes</b>  adalah suatu cara untuk menemukan semua bilangan prima di antara 1 dan suatu angka n. Tapis ini ditemukan oleh Eratosthenes, seorang ilmuwan Yunani kuno. Cara ini merupakan cara paling sederhana dan paling cepat untuk menemukan bilangan prima, sebelum Tapis Atkin ditemukan pada tahun 2004. Tapis Atkin merupakan cara yang lebih cepat namun lebih rumit dibandingkan dengan Tapis Eratosthenes. <br>
(Sumber: https://id.wikipedia.org/wiki/Tapis_Eratosthenes)
</p>

<b>Problem:</b><br>
Buatlah function untuk menampilkan deret bilangan prima hingga batas n yang diinputkan<br>
<br>

### input:

berupa angka integer positif > 0 sampai n
<br>

### output:

Array dari hasil kumpulan bilangan prima sampai n (input)
<br>

### example:

1. input = 5; Output = [2, 3, 5]
1. input = 15; Output = [2, 3, 5, 7, 11, 13]

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
