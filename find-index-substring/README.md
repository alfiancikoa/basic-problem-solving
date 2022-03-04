# Soal Find Index of Substring

<p align=justify><b>Substring</b> In formal language theory and computer science, a substring is a contiguous sequence of characters within a string. For instance, "the best of" is a substring of "It was the best of times". In contrast, "Itwastimes" is a subsequence of "It was the best of times", but not a substring.
Prefixes and suffixes are special cases of substrings. A prefix of a string {\displaystyle S}S is a substring of {S}S that occurs at the beginning of {S}S; likewise, a suffix of a string {S}S is a substring that occurs at the end of {S}S.
The list of all substrings of the string "apple" would be "apple", "appl", "pple", "app", "ppl", "ple", "ap", "pp", "pl", "le", "a", "p", "l", "e", "" (note the empty string at the end).
<br>(Sumber: https://en.wikipedia.org/wiki/Substring)
</p>

<b>Problem:</b><br>
Buatlah sebuah fungsi untuk mencari keberadaan substring. kembalikan posisi indexnya
<br>

### input:

berupa string dengan dua parameter:
parameter a = string utama
parameter b = substring yang akan dicari pada string utam (a)
<br>

### output:

Index keberadaan dari substring (b) jika terdapat di string utama (a). Jika substring (b) tidak terdapat pada string utama (a) maka kembalikan nilai -1
<br>

### example:

1. input = "abcdef" , "de"      output = 3
2. input = "abcdef" , "df"      output =-1


<i><b>*Solusi telah dilengkapi dengan unit test dan benchmark test</b></i>

### command untuk menjalankan unit test dan benchmark:

```
# Unit Test
$ go test -v

# Benchmark
$ go test -bench=.
```
