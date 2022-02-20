package main

import (
	"fmt"
	"strings"
)

func Pangram(s string) string {
	var hashmap = map[string]int{}
	for _, value := range strings.ToLower(s) {
		if value != ' ' {
			hashmap[string(value)]++
		}
	}
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
