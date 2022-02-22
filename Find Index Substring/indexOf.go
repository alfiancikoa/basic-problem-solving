package main

import "fmt"

/*
	Problem:
	Carilah pada index ke berapa posisi substring tersebut?
	jika ada maka return indexnya
	jika tidak maka return -1
*/

func indexOf(a string, b string) int {
	for i := 0; i <= len(a)-len(b); i++ {
		str := a[i : i+len(b)]
		if str == b {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(indexOf("abcdef", "de"))       // 3
	fmt.Println(indexOf("abcdef", "df"))       // -1
	fmt.Println(indexOf("abcdef", "def"))      // 3
	fmt.Println(indexOf("abcdef", "cef"))      // -1
	fmt.Println(indexOf("abcddef", "cde"))     // -1
	fmt.Println(indexOf("abcdef", "cde"))      // 2
	fmt.Println(indexOf("abccccdef", "cde"))   // 5
	fmt.Println(indexOf("abccdcdecdf", "cde")) // 5
}
