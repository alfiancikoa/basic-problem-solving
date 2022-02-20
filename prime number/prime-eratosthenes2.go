package main

import (
	"fmt"
)

/*
	Problem:
	Menampilkan seluruh urutan bilangan prima sampai dengan angka yang diinputkan n
*/

func showPrimes(n int) {
	//  membuat tempat tampungan untuk bilanagn prima
	var prime = []int{}
	// set seluruh array menjadi true
	for i := 0; i < n; i++ {
		prime = append(prime, 1)
	}
	// set array ke 0 dan 1 menjadi false, karena bukan bilangan prima
	prime[0], prime[1] = 0, 0
	// coret bilangan yang bukan prima dengan mengeset menjadi false
	for i := 2; i < n; i++ {
		for j := 2; j*i < n; j++ {
			prime[j*i] = 0
		}
	}
	// tampilkan seluruh bilangan prima, nilai di dalam array adalah true
	for i := 0; i < n; i++ {
		if prime[i] != 0 {
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
