package main

import (
	"fmt"
	"math"
)

/*
	Problem:
	Menentukan apakah bilangan yang di-inputkan merupakan bilangan prima atau bukan
	- input:

	- output:

	- example :
		1.	#input 	= 		#output =
		2.	#input 	= 		#output =
		3.	#input 	= 		#output =
*/

func isPrime(n int) bool {
	// jika nilai n lebih kecil dari dua, langsung direturn false karena sudah pasti bukan bilangan prima
	if n < 2 {
		return false
	}
	// proses pengecekan bilangan prima
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
