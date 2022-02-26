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
	fmt.Println(Pow(2, 0))   // 1
	fmt.Println(Pow(2, 1))   // 2
	fmt.Println(Pow(2, 2))   // 4
	fmt.Println(Pow(2, 4))   // 16
	fmt.Println(Pow(2, 6))   // 64
	fmt.Println(Pow(2, 8))   // 256
	fmt.Println(Pow(2, 10))  // 1024
	fmt.Println(Pow(2, -10)) // 0.0009765625
	fmt.Println(Pow(9, -2))  // 0.012345679012345678
}
