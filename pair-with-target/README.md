# Soal Pair With Target

<p align=justify><b>Pair With Target</b> merupakan fungsi untuk mencari dua angka yang jika dijumlahkan akan menghasilkan angka yang sama persis dengan target.
</p>

<b>Problem:</b><br>
Buatlah sebuah fungsi yang dapat mengembalikan dua angka yang berasal dari input (nums) yang jika kedua angka tersebut dijumlahkan akan menghasilkan angka yang sama dengan target.
<br>

### input:

berupa angka integer dengn dua parameter:
parameter nums = deretan angka positif array of integer yang tidak duplikat terurut/acak
parameter target = angkka bilangan bulat positif
<br>

### output:

dua angka yang diambil dari nums, yang jika angka tersebut dijumlahkan akan menghasilkan angka yang sama dengan target.\
jika tidak terdapat angka yang dijumlahkan sama dengan target, maka kembalikan angka -1
<br>

### example:

1. input = nums[1 2 3 4 5 6], target = 3; Output = [1 2] // Karena 1 + 2 = 3 (target) 
2. input = nums[1 2 3 4 5 6], target = 5; Output = [2 3] // Karena 2 + 3 = 5 (target) 
3. input = nums[1, 4, 2, 5, 11, 12], target = 8; Output = [-1] // Karea tidak ada angka yang tepat


<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
