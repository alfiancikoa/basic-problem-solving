# Soal Bilangan Prima

<p align=justify><b>Faktorial</b> dari bilangan bulat positif dari n yang dilambangkan dengan n!,
	adalah produk dari semua bilangan bulat positif yang kurang dari atau sama dengan n:
	n! = n * (n-1) * (n-2) * (n-3) * ... * 3 * 2 * 1. <br>
	contoh:<br>
	5! = 5 x 4 x 3 x 2 x 1. <br>
(Sumber: https://id.wikipedia.org/wiki/Faktorial)
</p>

### Problem:

Buatlah function untuk menentukan hasil dari faktorial bilangan yang di-inputkan <br>
<br>

### input:

berupa angka integer positif > 0 sampai n
<br>

### output:

Hasil faktorial bilangagn input n
<br>

### example:

1. input = 2 ;  output = 4
2. input = 3 ;  output = 6
3. input = 5 ;  output = 120
4. input = 8 ;  output = 40.320

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
