package main

import (
	"fmt"
	"strings"
)

/*
	Problem:
	Mengecek Apakah kata atau kalimat input merupakan palindrome atau bukan
	- input:
		satu atau lebih kata bertipe string, (a-z,A-Z)
	- output:
		status apakah palindrome atau bukan (true or false)
	- example :
		1.	#input 	= "kodok"	#output = true
		2.	#input 	= "kakak"	#output = true
		3.	#input 	= "makan"	#output = false
*/

func palindrome(str string) bool {
	// remove blank space and make it lower
	input := strings.ToLower(strings.ReplaceAll(str, " ", ""))
	leftIndex, rightIndex := 0, len(input)-1
	for leftIndex <= rightIndex {
		// jika karakter yang ada pada index kiri tidak sama dengan karakter index kanan maka langsung return false
		if input[leftIndex] != input[rightIndex-leftIndex] {
			return false
		}
		leftIndex++
	}
	return true
}

func main() {
	fmt.Println(palindrome("kodok"))                // true
	fmt.Println(palindrome("kasur ini rusak"))      // true
	fmt.Println(palindrome("Kasur Koh Ahok rusak")) // true
	fmt.Println(palindrome("Ibu Ratna antar ubi"))  // true
	fmt.Println(palindrome("zoom"))                 // false
	fmt.Println(palindrome("bubuk"))                // false
	fmt.Println(palindrome("madam"))                // true
}
