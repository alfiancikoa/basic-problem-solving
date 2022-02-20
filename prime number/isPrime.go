package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(5))           // true
	fmt.Println(isPrime(7))           // true
	fmt.Println(isPrime(8))           // false
	fmt.Println(isPrime(1000000007))  // true
	fmt.Println(isPrime(1500450271))  // true
	fmt.Println(isPrime(1000000000))  // false
	fmt.Println(isPrime(10000000019)) // true
	fmt.Println(isPrime(10000000033)) // true
}
