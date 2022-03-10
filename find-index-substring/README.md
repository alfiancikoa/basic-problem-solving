# Soal Find Index of Substring

<p align=justify><b>Substring</b> In formal language theory and computer science, a substring is a contiguous sequence of characters within a string. For instance, "the best of" is a substring of "It was the best of times". In contrast, "Itwastimes" is a subsequence of "It was the best of times", but not a substring.
Prefixes and suffixes are special cases of substrings. A prefix of a string {\displaystyle S}S is a substring of {S}S that occurs at the beginning of {S}S; likewise, a suffix of a string {S}S is a substring that occurs at the end of {S}S.
The list of all substrings of the string "apple" would be "apple", "appl", "pple", "app", "ppl", "ple", "ap", "pp", "pl", "le", "a", "p", "l", "e", "" (note the empty string at the end).
<br>(Sumber: https://en.wikipedia.org/wiki/Substring)
</p>

### Problem:
Buatlah sebuah fungsi untuk mencari keberadaan substring. kembalikan posisi indexnya
<br>

### input:

Berupa string dengan dua parameter:<br>
Parameter <b>a</b> = string utama<br>
Parameter <b>b</b> = substring yang akan dicari pada string utama (a)
<br>

### Output:

Index keberadaan dari substring (b) jika terdapat di string utama (a).<br>Jika substring (b) tidak terdapat pada string utama (a) maka kembalikan nilai -1
<br>

### example:

1. Input = a: "abcdef", b: "de"<br>Output = 3
2. Input = a: "abccccdef", b: "cde"<br>Output = 5
3. Input = a: "abcdef", b: "df"<br>Output = -1

<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
