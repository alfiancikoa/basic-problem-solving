# Soal Binary Sort Algorithm

<p align=justify><b>Binary Sort</b> hampir sama konsepnya dengan binary search yaitu membagi dua suatu deretan angka dan dicek angka yang di sebelah kiri harus selalu lebih kecil dari angka yang sebelah kanan
</p>

<b>Problem:</b><br>
Buatlah fungsi untuk mengurutkan suatu deret bilangan array of integer satu dimensi dengan menggunakan algoritma binary sort
<br>

### input:

Berupa angka list deret array satu dimensi dengan kumpulan angkan yang acak<br>
Bilangan dapat berupa bilangan bulat positif ataupun negatif
<br>

### output:

Berupa array of integer satu dimensi yang telah terurut dari yang terkecil ke yang terbesar
<br>

### example:

1. Input = [3 4 2 1 5]  <br>Output = [1 2 3 4 5] 
2. Input = [3 2 5]  <br>Output = [2 3 5] 
3. Input = [8 5 7 4 6 9 3 1 2 10]  <br>Output = [1 2 3 4 5 6 7 8 9 10] 


<br>

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
