package main

import (
	"fmt"
	"strings"
)

/*
	Problem:
	mengecek apakah kalimat merupakan pangram atau bukan?
	pangram merupakan atau kalimat holoalfabetis, adalah suatu kalimat yang menggunakan semua huruf dalam suatu aksara paling tidak satu kali
	artinya dalam satu kalimat tersebut semua huruf alphabet digunakan mulai dari a - z
	- input:
		satu atau lebih kata yang menjadi kalimat bertipe string, (a-z,A-Z)
	- output:
		status apakah kalimat tersebut merupakan pangram atau bukan(pangram or not pangram)
	- example :
		1.	#input 	= "The quick brown fox jumps over the lazy dog"						#output = pangram
		2.	#input 	= "ini bukan pangram karena tidak menggunakan seluruh alphabet"		#output = not pangram
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
