# Soal Linear Search Algorithm

<p align=justify><b>Linear Search</b> atau pencarian linear adalah sebuah algoritme pencarian, juga dikenal sebagai pencarian sekuensial, yang cocok untuk mencari sebuah nilai tertentu pada sebuah himpunan data.<br>

Algoritme ini beroperasi dengan memeriksa setiap elemen dari sebuah list sampai sebuah kecocokan ditemukan. Pencarian linear bekerja dalam O(n). Jika data terdistribusi secara acak, rata-rata ada n/2 pembandingan akan dilakukan. Kasus terbaik adalah ketika nilai yang dicari adalah elemen pertama dari list, kasus ini hanya memerlukan 1 pembandingan. Kasus terburuk adalah ketika nilai yang dicari tidak ada dalam list, yang memerlukan n pembadingan.<br>
(Sumber: https://id.wikipedia.org/wiki/Pencarian_linear)
</p>

### Problem:
Buatlah sebuah fungsi untuk mencari bilangan target pada sebuah deretan array of integer 1 dimensi
dengan menggunakn algoritma Linear search
<br>

### Input:

Berupa angka integer dengn dua parameter:<br>
Parameter <b>arr</b> = deretan angka array of integer, angka tersebut dapat terurut ataupun acak<br>
Parameter <b>target</b> = angka int yang kan dicari pada deret arr
<br>

### Output:

Jika target terdapat pada list array of integer (arr) maka kembalikan indexnya<br>
Jika tidak ditemukan, maka kembalikan nilai -1
<br>

### example:

1. Input = arr: [1 2 3 4 5 6], target: 5 <br>Output = 4
2. Input = arr: [2 4 5 1 3], target: 2 <br>Output = 0
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
