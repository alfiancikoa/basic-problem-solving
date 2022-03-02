# Soal Segitiga Pascal

<p align=justify><b>Segitiga Pascal</b> adalah suatu aturan geometri pada koefisien binomial dalam sebuah segitiga. Segitiga tersebut dinamai berdasarkan nama matematikawan Blaise Pascal, meskipun ahli matematika lain telah mengkajinya berabad-abad sebelum dia di India, Persia, Tiongkok, dan Italia. Barisan segitiga Pascal umumnya dihitung dimulai dengan baris kosong, dan nomor-nomor dalam barisan ganjil biasanya diatur agar terkait dengan nomor-nomor dalam baris genap. Konstruksi sederhana pada segitiga dilakukan dengan cara berikut. Di barisan nol, hanya tulis nomor 1. Kemudian, untuk membangun unsur-unsur barisan berikutnya, tambahkan nomor di atas dan di kiri dengan nomor secara langsung di atas dan di kanan untuk menemukan nilai baru. Jika nomor di kanan atau kiri tidak ada, gantikan suatu kosong pada tempatnya. Misalnya, nomor satu di barisan pertama adalah 0 + 1 = 1, di mana nomor 1 dan 3 dalam barisan ketiga ditambahkan untuk menghasilkan nomor 4 dalam barisan keempat. <br>
(Sumber: https://id.wikipedia.org/wiki/Segitiga_Pascal)
</p>

<b>Problem:</b><br>
Buatlah function untuk mengecek apakah bilangan yang di-inputkan merupakan bilangan prima atau bukan. Jika bilangan tersebut merupakan bilangan prima, maka tampilkan true jika tidak maka tampilkan false <br>
<br>

### input:

berupa angka integer positif > 0 sampai n
<br>

### output:

Tampilan segitiga pascal sesuai dengan inputan
<br>

### example:

1. input = 3 
   Output = 
   1
   1 1
   1 2 1

2. input = 5 
   Output = 
   1
   1 1
   1 2 1
   1 3 3 1
   1 4 6 4 1
   
<i>*Solusi telah dilengkapi dengan unit test dan benchmark test</i>
command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
