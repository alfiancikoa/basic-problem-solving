package main

import (
	"fmt"
	"math"
)

/*
	Problem:

	- input:

	- output:

	- example :
		1.	#input 	= 		#output =
		2.	#input 	= 		#output =
		3.	#input 	= 		#output =
*/

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

func DeretPrima(n int) []int {
	var count int
	var result []int
	i := 1
	for count < n {
		if isPrime(i) {
			result = append(result, i)
			count++
		}
		i++
	}
	return result
}

func main() {
	fmt.Println(DeretPrima(5))
	/*
		[2 3 5 7 11]
	*/
	fmt.Println(DeretPrima(10))
	/*
		[2 3 5 7 11 13 17 19 23 29]
	*/
}
