package main

import (
	"fmt"
	"sort"
)

/*
	Problem:
	Hilangkan angka yang duplikat dan menampilkannya hanya sekali
	- input:
		deretan bilangan bulat yang dari 0 s/d n (0<n<=100)
	- output:
		deretan angka dalam array yang tidak duplikat dan terurut
	- example :
		#input 	= [1 1 2 3 4 4 5]
		#output = [1 2 3 4 5]
*/

func removeDup(arr []int) []int {
	var result []int
	hashmap := map[int]bool{}
	for _, digit := range arr {
		if !hashmap[digit] {
			result = append(result, digit)
		}
		hashmap[digit] = true
	}
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(removeDup([]int{1, 2, 2, 3, 4, 5, 5, 6, 7}))          // [1 2 3 4 5 6 7]
	fmt.Println(removeDup([]int{2, 2, 2, 3, 3, 3, 4, 5, 5, 5}))       // [2,3,4,5]
	fmt.Println(removeDup([]int{5, 4, 2, 1, 3}))                      // [1 2 3 4 5]
	fmt.Println(removeDup([]int{2, 3, 2, 4, 1, 2, 5, 4, 2, 1, 6, 6})) // [1 2 3 4 5 6]
}
