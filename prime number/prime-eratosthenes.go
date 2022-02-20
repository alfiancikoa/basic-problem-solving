package main

import (
	"fmt"
	"math"
)

func showPrimes(n int) {
	prime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		prime[i] = true
	}
	prime[0], prime[1] = false, false
	fmt.Println(prime)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if prime[i] {
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}
	for i := 0; i <= n; i++ {
		if prime[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func main() {
	showPrimes(10)
}
