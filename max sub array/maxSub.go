package main

import "fmt"

/*
	Problem:
	Carilah hasil penjumlahan maksimum dari deretan angka
	- input:
		Berupa angka positif ataupun negatif:
	- output:
		hasil penjumlahan maksimum dari deretan angka input
	- example :
		1.	#input 	= [-2, 3, -1, 7, -1]					#output = 9
		2.	#input 	= [-2, -3, 4, -1, -2, 1, 5, -3]			#output = 7
*/

func MaxSubArray(arr []int) int {
	maxSoFar, max := 0, 0
	for i := 0; i < len(arr); i++ {
		maxSoFar += arr[i]
		if max < maxSoFar {
			max = maxSoFar
		}
		if maxSoFar < 0 {
			maxSoFar = 0
		}
	}
	return max
}

func main() {
	fmt.Println(MaxSubArray([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7
	fmt.Println(MaxSubArray([]int{-2, 3, -1, 7, -1}))            // 9
	fmt.Println(MaxSubArray([]int{-1, -2, 8, -2, -1, 3, 6}))     // 14
}
