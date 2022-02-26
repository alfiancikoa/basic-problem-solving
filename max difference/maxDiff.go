package main

import "fmt"

/*
	Problem:
		Mencari selisih tertinggi dari satu deret bilangan
	Input:
		Berupa deretan angka array bertipe satu dimensi
		Angka dapat berupa bilangan positif ataupun negatif
	Output:

	Example:

*/

func maximumDiff(arr []int) int {
	maxDiff, min := 0, arr[0]
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		} else if maxDiff < arr[i]-min {
			maxDiff = arr[i] - min
		}
	}
	return maxDiff
}

func main() {
	fmt.Println(maximumDiff([]int{1, 2, 6, 4}))           // 5
	fmt.Println(maximumDiff([]int{1, 3, 2, 4}))           // 3
	fmt.Println(maximumDiff([]int{-1, -3, 2, 4}))         // 7
	fmt.Println(maximumDiff([]int{3, 2, 4, 7}))           // 5
	fmt.Println(maximumDiff([]int{2, 7, 9, 5, 1, 3, 5}))  // 7
	fmt.Println(maximumDiff([]int{2, 7, 9, 5, 1, 5, 10})) // 9
}
