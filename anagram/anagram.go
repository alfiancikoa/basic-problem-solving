package main

import "fmt"

/*
	Problem:
	Mencari jumlah karakter yang akan diubah agar bisa menjadi anagram
	Anagram adalah salah satu jenis permainan kata yang huruf-huruf di kata awal biasanya diacak untuk membentuk kata lain atau sebuah kalimat. Anagram sering dipakai sebagai kode
	- input:

	- output:

	- example :
		1.	#input 	= 		#output =
		2.	#input 	= 		#output =
		3.	#input 	= 		#output =
*/

func anagram(s string) int32 {
	var result int32
	var hashmap = map[string]int{}
	if len(s)%2 != 0 {
		return -1
	}
	strA := s[:len(s)/2]
	strB := s[len(s)/2:]
	for _, value := range strA {
		hashmap[string(value)]++
	}
	for _, value := range strB {
		if _, isExist := hashmap[string(value)]; isExist {
			hashmap[string(value)]--
			if hashmap[string(value)] < 0 {
				result++
			}
		} else {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(anagram("asdfjoieufoa"))     // 3
	fmt.Println(anagram("fdhlvosfpafhalll")) // 5
	fmt.Println(anagram("mvdalvkiopaufl"))   // 5
	fmt.Println(anagram("mnop"))             // 2
	fmt.Println(anagram("xyyx"))             // 0
	fmt.Println(anagram("xaxbbbxx"))         // 1
}
