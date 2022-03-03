# Soal Pangram

<p align=justify><b>Pangram</b> atau kalimat holoalfabetis, adalah suatu kalimat yang menggunakan semua huruf dalam suatu aksara paling tidak satu kali. Pangram digunakan untuk menampilkan rupa huruf, menguji alat, serta mengembangkan keterampilan tulisan tangan, kaligrafi, dan mengetik. Contoh pangram yang terkenal dalam bahasa Inggris adalah kalimat "the quick brown fox jumps over the lazy dog" yang mencantumkan ke-26 huruf aksara Latin. <br>
Pangram sempurna adalah pangram yang mencantumkan setiap huruf hanya satu kali dan dapat dianggap sebagai anagram dari alfabet tersebut. Contoh pangram sempurna adalah kalimat "hana caraka; data sawala; padha jayanya; maga bathanga" yang mencantumkan lengkap ke-20 huruf aksara Jawa masing-masing satu kali.<br>
(Sumber: https://id.wikipedia.org/wiki/Pangram)
</p>

<b>Problem:</b><br>
Mencari hasil dari bilangan berpangkat <br>
2^5 = 2 * 2 * 2 * 2 * 2 = 32
<br>

### input:

berupa angka integer dengn dua parameter:
parameter a = angka utama (Bilangan Positif)
parameter b = angka pangkatnya (Bisa positif ataupun negatif)
<br>

### output:

hasil pangkat dari bilangan a. seperti a^b = c
<br>

### example:

1. input = 2^0      output = 1
2. input = 2^1      output = 2
3. input = 2^8      output = 258
4. input = 9^-2     output = 0.012345679012345678


<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
