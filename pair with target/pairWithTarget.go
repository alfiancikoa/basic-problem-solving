package main

import "fmt"

/*
	Problem:

	- input:

	- output:

	- example :
		1.	#input 	= 		#output =
		2.	#input 	= 		#output =
		3.	#input 	= 		#output =
*/

func Pair(arr []int, target int) []int {
	hashmap := map[int]int{}
	var result []int
	for i := 0; i < len(arr); i++ {
		pair := target - arr[i]
		if hashmap[pair] > 0 {
			result = append(result, arr[hashmap[pair]], arr[i])
			break
		}
		hashmap[arr[i]] = i
	}
	return result
}

func main() {
	fmt.Println(Pair([]int{1, 2, 3, 4, 5}, 6)) // 2 4
	fmt.Println(Pair([]int{2, 3, 4, 5, 6}, 8)) // 3 5
	fmt.Println(Pair([]int{2, 2, 3, 4, 4}, 7)) // 3 4
}
