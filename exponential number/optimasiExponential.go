package main

import "fmt"

/*
	Problem:
	Mencari hasil dari bilangan berpngkat
	- input:
		berupa input dengn dua parameter:
		parameter a = angka utama
		parameter b = angka pangkatnya
	- output:
		hasil pangkat dari bilangan a. seperti a^b = c
	- example :
		1.	#input 	= 2^0					#output = 1
		2.	#input 	= 2^1					#output = 2
		3.	#input 	= 2^8					#output = 256
		4.	#input 	= 9^-2					#output = 0.012345679012345678
*/

func PowOptimation(a int, b int) float64 {
	var result float64 = -1
	var pangkat int = b
	var temp = 1.0
	if b < 0 {
		pangkat = -b
	}
	if pangkat%2 != 0 {
		temp = float64(a)
	}
	for i := 0; i < pangkat/2; i++ {
		result = result * float64(a)
	}
	if b < 0 {
		return 1 / ((result * result) * temp)
	}
	return (result * result) * temp
}

func main() {
	fmt.Println(PowOptimation(2, 0))  // 1
	fmt.Println(PowOptimation(2, 1))  // 2
	fmt.Println(PowOptimation(2, 2))  // 4
	fmt.Println(PowOptimation(2, 4))  // 16
	fmt.Println(PowOptimation(2, 6))  // 64
	fmt.Println(PowOptimation(2, 8))  // 256
	fmt.Println(PowOptimation(2, 10)) // 1024
	fmt.Println(PowOptimation(9, -2)) // 0.012345679012345678
}
