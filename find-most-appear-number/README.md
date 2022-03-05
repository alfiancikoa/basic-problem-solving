# Soal Most Appear Number

<p align=justify><b>Most Appear Number</b> adalah mencari bilangan yang paling sering muncul pada deret suatu bilangan. <br>
</p>

<b>Problem:</b><br>
Buatlah function untuk mengecek bilangan yang paling sering muncul. Apabila frekuensi kemunculan bilangan tersebut itu sama, maka pilih yang nilainya paling besar diantara yang lain. Untuk yang frekuensi kemunculan paling besar, nilainya tidak peduli lebih besar atau lebih kecil asalkan angka tersebut paling sering muncul<br>
<br>

### input:

berupa deret bilangan array of integer > 0 sampai n, dengan panjang 1 < n < 100
<br>

### output:

bilangan integer yang frekuensi kemunculannya paling besar
<br>

### example:

1. input = [1 2 3 4 1]; Output = 1
2. input = [1 2 2 3 3 4 5 6]; Output = 3
3. input = [1 1 1 1 1 2]; Output = 1

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
