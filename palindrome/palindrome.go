package main

import (
	"fmt"
	"strings"
)

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
