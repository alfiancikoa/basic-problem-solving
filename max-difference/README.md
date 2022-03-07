# Soal Max Difference

<p align=justify><b>Max Difference</b> merupakan selisih tertinggi antara tiap bilangan dengan bilangan yang paling minimum yang dihasilkan oleh suatu deret bilangan. <br>
</p>

<b>Problem:</b><br>
Mencari selisih maksimum dari suatu deret input
<br>

### input:

Berupa deretan angka array of integer yang terdiri dari bilangan positif ataupun negatif dengan panjang data 0 < n <= 100 
<br>

### output:

Hasil merupakan selisih maksimum yang ada pada deretan array input tersebut, output berupa bilangan integer positif
<br>

### example:

1. input = [1 2 6 4]<br>Output = 5
2. input = [-1 -3 2 4]<br>Output = 7
3. input = [2 7 9 5 1 5 10]<br>Output = 9

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
