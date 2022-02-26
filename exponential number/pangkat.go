package main

import "fmt"

func Pow(a int, b int) float64 {
	var result float64 = 1
	var pangkat int = b
	if b < 0 {
		pangkat = -b
	}
	for i := 0; i < pangkat; i++ {
		result = result * float64(a)
	}
	if b < 0 {
		return 1 / result
	}
	return result
}

func main() {
	fmt.Println(Pow(2, 0))  // 1
	fmt.Println(Pow(2, 1))  // 8
	fmt.Println(Pow(2, 2))  // 4
	fmt.Println(Pow(2, 4))  // 16
	fmt.Println(Pow(2, 8))  // 256
	fmt.Println(Pow(2, 10)) // 1024
	fmt.Println(Pow(9, -2)) // 0.012345679012345678
}
