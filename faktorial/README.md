# Soal Bilangan Prima

<p align=justify><b>Faktorial</b> dari bilangan bulat positif dari n yang dilambangkan dengan n!,
	adalah produk dari semua bilangan bulat positif yang kurang dari atau sama dengan n:
	n! = n * (n-1) * (n-2) * (n-3) * ... * 3 * 2 * 1. <br>
	contoh:<br>
	5! = 5 x 4 x 3 x 2 x 1. <br>
(Sumber: https://id.wikipedia.org/wiki/Faktorial)
</p>

### Problem:

Buatlah function untuk menentukan hasil dari faktorial bilangan yang di-inputkan
<br>

### Input:

Berupa angka integer positif > 0 sampai n
<br>

### Output:

Hasil faktorial bilangagn input n
<br>

### Example:

1. Input = 2 ;  Output = 4
2. Input = 3 ;  Output = 6
3. Input = 5 ;  Output = 120
4. Input = 8 ;  Output = 40.320

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
