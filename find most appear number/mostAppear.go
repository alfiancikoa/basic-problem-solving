package main

import "fmt"

/*
	Problem:
	Carilah angka yang sering muncul pada deret array input
	Jika porsi kemuunculan angka sama semua, maka carilah angka yang nilainya lebih tinggi
*/

func mostAppear(num []int) int {
	hashmap := map[int]int{}
	mostSoFar, mostNumber := 0, 0
	for _, digit := range num {
		hashmap[digit]++
		if hashmap[digit] >= mostSoFar && digit >= mostNumber {
			mostNumber = digit
			mostSoFar = hashmap[digit]
		}
	}
	return mostNumber
}

func main() {
	fmt.Println(mostAppear([]int{90, 8, 90, 68, 101, 4, 1, 2, 101, 101})) // 101
	fmt.Println(mostAppear([]int{92, 8, 92, 68, 100, 4, 1, 2, 100, 100})) // 100
	fmt.Println(mostAppear([]int{5, 5, 1, 1, 2, 3, 10, 9}))               // 5
	fmt.Println(mostAppear([]int{1, 2, 3, 4, 5, 6}))                      // 6
	fmt.Println(mostAppear([]int{4, 3, 6, 2, 1, 5}))                      // 6
	fmt.Println(mostAppear([]int{1, 1, 1, 1, 1, 1}))                      // 1
}
