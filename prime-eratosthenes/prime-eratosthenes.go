package main

import (
	"fmt"
	"math"
)

/*
	Problem:

	- input:
	Menampilkan seluruh urutan bilangan prima sampai dengan angka yang diinputkan n
	- output:

	- example :
		1.	#input 	= 		#output =
		2.	#input 	= 		#output =
		3.	#input 	= 		#output =
*/

func showPrimes(n int) {
	//  membuat tempat tampungan untuk bilanagn prima
	prime := make([]bool, n+1)
	// set seluruh array menjadi true
	for i := 0; i <= n; i++ {
		prime[i] = true
	}
	// set array ke 0 dan 1 menjadi false, karena bukan bilangan prima
	prime[0], prime[1] = false, false
	// coret bilangan yang bukan prima dengan mengeset menjadi false
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if prime[i] {
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}
	// tampilkan seluruh bilangan prima, nilai di dalam array adalah true
	for i := 0; i <= n; i++ {
		if prime[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func main() {
	showPrimes(100)
	/*
		2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
	*/
}