# Soal Most Appear Number

<p align=justify><b>Most Appear Number</b> adalah mencari bilangan yang paling sering muncul pada deret suatu bilangan. <br>
</p>

### Problem:

<p align=justify>Buatlah function untuk mengecek bilangan yang paling sering muncul. Apabila frekuensi kemunculan bilangan tersebut itu sama, maka pilih yang nilainya paling besar diantara yang lain. Untuk yang frekuensi kemunculan paling besar, nilainya tidak peduli lebih besar atau lebih kecil asalkan angka tersebut paling sering muncul.</p>

### Input:

Berupa deret bilangan array of integer > 0 sampai n, dengan panjang 1 < n < 100
<br>

### Output:

Bilangan integer yang frekuensi kemunculannya paling besar
<br>

### example:

1. Input = [1 2 3 4 1]<br>Output = 1
2. Input = [1 2 2 3 3 4 5 6]<br>Output = 3
3. Input = [1 1 1 1 1 2]<br>Output = 1
4. Input = [1 2 3 4 5 6]<br>Output = 6

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
