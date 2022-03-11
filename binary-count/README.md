# Soal Binary Count

<p align=justify><b>Binary Count</b> atau hitung biner merupakan challenge atau tantangan untuk menghitung berapa banyak angka biner 1 atau 0 pada bilangan desimal. Bilangan desimal harus dikonversi ke dalam bilangan biner terlebih dahulu agar dapat dihitung berapa banyak angka 1 atau 0 pada bilangan biner tersebut</p>

### Problem

Buatlah sebuah fungsi dengan input-an dua parameter yang dapat mengembalikan jumlah angka 1 atau angka 0 pada bilangan desimal yang telah dikonversi ke dalam bilangan biner sesuai dengan inputan parameter 
<br>

### Input

Input memiliki dua parameter yaitu:<br>
Parameter <b>angka</b>: merupakan bilangan desimal positif integer dari 0 sampai 65536 
Parameter <b>nomorBit</b>: merupakan flag bilangan yang akan dihitung jumlahnya pada bilangan biner dari hasil konversi bilangan desimal
<br>

### Output

Hasil berupa bilangan bulat integer positif yang merepresentasikan jumlah bit 1 atau 0 pada bilangan biner yang dari hasil konversi bilangan desimal
<br>

### Example

<br>

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
