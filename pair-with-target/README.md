# Soal Pair With Target

<p align=justify><b>Pair With Target</b> merupakan fungsi untuk mencari dua angka yang jika dijumlahkan akan menghasilkan angka yang sama persis dengan target.
</p>

### Problem:
Buatlah sebuah fungsi yang dapat mengembalikan dua angka yang berasal dari input (nums) yang jika kedua angka tersebut dijumlahkan akan menghasilkan angka yang sama dengan target.
<br>

### Input:

Berupa angka integer dengan dua parameter:<br>
Kedua parameter input terdiri dari bilangan bulat positif<br>
Parameter <b>nums</b> = deretan angka positif array of integer yang tidak duplikat terurut/acak<br>
Parameter <b>target</b> = target merupakan angka yang akan menjadi patokan untuk mencari angka berapa saja yang ada di <b>nums</b> yang jika dijumlahkan akan sama dengan <b>target</b>
<br>

### Output:

Array of integer yang berisi dua angka yang diambil dari nums, yang jika angka tersebut dijumlahkan akan menghasilkan angka yang sama dengan target.<br>
Jika tidak terdapat angka yang dijumlahkan sama dengan target, maka kembalikan angka -1
<br>

### example:

1. <b>Input</b> = nums[1 2 3 4 5 6], target = 3<br> <b>Output</b> = [1 2]<br> // Karena 1 + 2 = 3 (target) 
2. <b>Input</b> = nums[1 2 3 4 5 6], target = 5<br> <b>Output</b> = [2 3]<br> // Karena 2 + 3 = 5 (target) 
3. <b>Input</b> = nums[1, 4, 2, 5, 11, 12], target = 8<br> <b>Output</b> = [-1]<br> // Karea tidak ada angka yang tepat


<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
