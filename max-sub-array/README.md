# Soal Max Sub Array

<p align=justify><b>Max SubArray</b> merupakan pencarian angka maksimum dari penjumlahan berurut dari suatu deret angka.
</p>

### Problem:

Buatlah suatu fungsi yang dapat mencari nilai maksimum dari subArray satu dimensi
<br>

### Input:

Berupa deretan angka positif array of integer satu dimensi<br>
Angka merupakan bilangan bulat positif ataupun negatif
<br>

### Output:

Nilai penjumlahan array berupa bilangan positif integer
<br>

### example:

1. Input = [-2, 3, -1, 7, -1]<br>Output = 9 // hasil dari 3 + (-1) + 7 = 9
2. Input = [-2, -3, 4, -1, -2, 1, 5, -3]<br>Output = 7 // hasil dari 4 + (-1) + (-2) + 1 + 5 = 7


<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
