# Soal Dictionary

<p align=justify><b>Dictionary</b> merupakan konsep dasar kamus yang dapat ditambahkan dan datanya dapat diambil/dilihat.
</p>

### Problem:

Buatlah sebuah class Kamus yang akan menyimpan semua kata sudah yang ditambahkan ke dalamnya beserta dengan sinonim dari masing-masing kata yang sudah ditambahkan.<br>
<br>

### Input:

Adapun class tersebut harus memiliki fungsi-fungsi sebagai berikut :<br>
<b>Add(String kata, Array sinonim)</b> : tidak mengembalikan hasil (void)<br>
<b>Get(String kata)</b> : mengembalikan hasil array of strings
<br>

### Output:

Jika Perintah get/ambil data akan mengembalikan seluruh data yang berkaitan dengan input katanya
<br>

### example:

```
$kamus = new Kamus();
$kamus->Add(‘buah’, [‘pisang, ‘jeruk’]);
$kamus->Add(‘buah’, [‘naga’, ‘anggur’]);
$kamus->Add(‘hewan’, [‘kucing’, ‘kambing’]);

$kamus->Get(‘buah’);
// mengembalikan hasil [‘pisang’, ‘jeruk’, ‘naga’, ‘anggur’]

// Perhatikan baik-baik hasil pengujian di bawah ini
$kamus->Get(‘hewan’);
// mengembalikan hasil [‘kucing’, ‘kambing’]

$kamus->Get(‘kucing’);
// mengembalikan hasil [‘hewan’]

$kamus->Get(‘mawar’);
// mengembalikan hasil null
```

<br>

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```