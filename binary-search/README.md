# Soal Binary Search Algorithm

<p align=justify><b>Binary Search</b> atau pencarian biner adalah sebuah teknik untuk menemukan nilai tertentu dalam sebuah larik (array) linear, dengan menghilangkan setengah data pada setiap langkah, dipakai secara luas tetapi tidak secara ekslusif dalam ilmu komputer. Sebuah pencarian biner mencari nilai te7ngah (median), melakukan sebuah pembandingan untuk menentukan apakah nilai yang dicari ada sebelum atau sesudahnya, kemudian mencari setengah sisanya dengan cara yang sama. Sebuah pencarian biner adalah salah satu contoh dari algoritme divide and conquer (atau lebih khusus algoritme decrease and conquer) dan sebuah pencarian dikotomi (lebih rinci di Algoritme pencarian).   <br>
(Sumber: https://id.wikipedia.org/wiki/Algoritma_pencarian_biner)
</p>

### Problem:

Buatlah sebuah fungsi untuk mencari bilangan target pada sebuah deretan array of integer 1 dimensi
dengan menggunakn algoritma binary search
<br>

### Input:

Berupa angka integer dengn dua parameter:<br>
Parameter <b>arr</b> = deretan angka array of integer yang telah terurut sortfirst<br>
Parameter <b>target</b> = angka int yang kan dicari pada deret arr
<br>

### Output:

Jika target terdapat pada list array of integer (arr) maka kembalikan indexnya<br>
jika tidak ditemukan, maka kembalikan nilai -1
<br>

### Example:

1. Input = arr: [1 2 3 4 5 6], target: 5 <br>Output = 4
2. Input = arr: [1 2 3 4 5 6], target: 2 <br>Output = 1
3. Input = arr: [1 2], target: 2 <br>Output = 1
4. Input = arr: [1 2 3 4 5 6], target: 9 <br>Output = -1

<br>

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
