package main

import (
	"fmt"
	"strings"
)

/*
	Problem:
	mengecek apakah kalimat merupakan pangram atau bukan?
	pangram merupakan atau kalimat holoalfabetis, adalah suatu kalimat yang menggunakan semua huruf dalam suatu aksara paling tidak satu kal
*/

func Pangram(s string) string {
	// simpan seluruh karakter ke dalam map
	var hashmap = map[string]int{}
	// proses pengambilan karakter yang digunakan
	for _, value := range strings.ToLower(s) {
		if value != ' ' {
			hashmap[string(value)]++
		}
	}
	// cek apakah jumlah huruf sebanyak 26 berarti semua karakter huruf dipakai
	if len(hashmap) == 26 {
		return "pangram"
	}
	return "not pangram"
}

func main() {
	fmt.Println(Pangram("We promptly judged antique ivory buckles for the next prize")) // panggram
	fmt.Println(Pangram("The quick brown fox jumps over the lazy dog"))                 //pangram
	fmt.Println(Pangram("We promptly judged antique ivory buckles for the prize"))      // not pangram

}
