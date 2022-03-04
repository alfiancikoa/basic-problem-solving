# Soal Find Index of Substring

<p align=justify><b>Substring</b> In formal language theory and computer science, a substring is a contiguous sequence of characters within a string. For instance, "the best of" is a substring of "It was the best of times". In contrast, "Itwastimes" is a subsequence of "It was the best of times", but not a substring.
Prefixes and suffixes are special cases of substrings. A prefix of a string {\displaystyle S}S is a substring of {\displaystyle S}S that occurs at the beginning of {S}S; likewise, a suffix of a string {S}S is a substring that occurs at the end of {S}S.
The list of all substrings of the string "apple" would be "apple", "appl", "pple", "app", "ppl", "ple", "ap", "pp", "pl", "le", "a", "p", "l", "e", "" (note the empty string at the end).
<br>(Sumber: https://en.wikipedia.org/wiki/Substring)
</p>

<b>Problem:</b><br>

<br>

### input:

berupa angka integer dengn dua parameter:
parameter a = angka utama (Bilangan Positif)
parameter b = angka pangkatnya (Bisa positif ataupun negatif)
<br>

### output:

hasil pangkat dari bilangan a. seperti a^b = c
<br>

### example:

1. input = 2^0      output = 1
2. input = 2^1      output = 2
3. input = 2^8      output = 258
4. input = 9^-2     output = 0.012345679012345678


<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
