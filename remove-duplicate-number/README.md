# Soal Remove Duplicate

<p align=justify><b>Remove Duplicate</b> adalah sebuah fungsi untuk menghilangkan atau mengabaikan angka yang sama atau duplikat, sehingga angka yang akan dimunculkan merupakan deretan angka yang tidak mempunyai kembaran atau kesamaan.</p>

<br>

<b>Problem:</b><br>
Buatlah Function menghilangkan atau mengabaikan angka yang duplikat atau sama
<br>

### input:

Berupa deretan bilangan bulat yang dari 0 s/d n (0<n<=100)
<br>

### output:

Hasil deretan angka dalam array yang tidak duplikat sesuai dengan angka pada input-an dan terurut
<br>

### example:

1. input = [1 1 2 3 4 4 5];  output = [1 2 3 4 5]
2. input = [3 4 2 5 4 1];    output = [1 2 3 4 5]
3. input = [3 2 1]           output = [1 2 3]

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```

