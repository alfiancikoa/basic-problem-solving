# Soal Pangram

<p align=justify><b>Pangram</b> atau kalimat holoalfabetis, adalah suatu kalimat yang menggunakan semua huruf dalam suatu aksara paling tidak satu kali. Pangram digunakan untuk menampilkan rupa huruf, menguji alat, serta mengembangkan keterampilan tulisan tangan, kaligrafi, dan mengetik. Contoh pangram yang terkenal dalam bahasa Inggris adalah kalimat "the quick brown fox jumps over the lazy dog" yang mencantumkan ke-26 huruf aksara Latin. <br>
Pangram sempurna adalah pangram yang mencantumkan setiap huruf hanya satu kali dan dapat dianggap sebagai anagram dari alfabet tersebut. Contoh pangram sempurna adalah kalimat "hana caraka; data sawala; padha jayanya; maga bathanga" yang mencantumkan lengkap ke-20 huruf aksara Jawa masing-masing satu kali.<br>
(Sumber: https://id.wikipedia.org/wiki/Pangram)
</p>

<b>Problem:</b><br>
Buatlah sebuah fungsi untuk mengetahui apakah kalimat yang dimasukkan sebagai input merupakan pangram atau bukan
<br>

### input:
Sebuah kalimat yang menggunakan huruf a-z
<br>

### output:
<b>pangram</b> jika kalimat tersebut merupakan pangram atau <b>not pangram</b> jika kalimat tersebut bukan pangram
<br>

### example:

1. Input = he quick brown fox jumps over the lazy dog; Output = pangram
2. Input = hari yang cerah untuk jiwa yang sepi; Output = not pangram



<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
