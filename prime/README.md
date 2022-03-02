# Soal Bilangan Prima

<p align=justify><b>Bilangan prima</b>adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. Bilangan 2 dan 3 adalah bilangan prima, sedangkan 4 bukan bilangan prima karena 4 memiliki faktor selain 1 dan 4, yakni 2. <br>
(Sumber: https://id.wikipedia.org/wiki/Bilangan_prima)
</p>

<b>Problem:</b><br>
Buatlah function untuk mengecek apakah bilangan yang di-inputkan merupakan bilangan prima atau bukan. Jika bilangan tersebut merupakan bilangan prima, maka tampilkan true jika tidak maka tampilkan false <br>
<br>

### input:

berupa angka integer positif > 0 sampai n
<br>

### output:

True jika bilangan prima, False jika bukan bilangan prima
<br>

### example:

1. input = 1 ;  output = false
2. input = 2 ;  output = true
3. input = 3 ;  output = true
4. input = 6 ;  output = false

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
