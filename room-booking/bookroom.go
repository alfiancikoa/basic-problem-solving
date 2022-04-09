package main

import (
	"fmt"
	"math"
)

func kamar(guest, room int) []int {
	var result []int
	for guest > 0 {
		c := math.Ceil(float64(guest) / float64(room))
		result = append(result, int(c))
		guest -= int(c)
		room--
	}
	return result
}

func main() {
	fmt.Println(kamar(4, 2))  // [2 2]
	fmt.Println(kamar(2, 1))  // [2]
	fmt.Println(kamar(5, 3))  // [2 2 1]
	fmt.Println(kamar(7, 2))  // [4 3]
	fmt.Println(kamar(4, 4))  // [1 1 1 1]
	fmt.Println(kamar(10, 4)) // [3 3 2 2]
}
