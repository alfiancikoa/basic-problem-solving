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

func Pascal(num int) {
	array := make([][]int, num)
	for i := 0; i < num; i++ {
		array[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				array[i][j] = 1
			} else {
				array[i][j] = array[i-1][j-1] + array[i-1][j]
			}
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}
}

func main() {
	Pascal(6)
	/*
		1
		1 1
		1 2 1
		1 3 3 1
		1 4 6 4 1
		1 5 10 10 5 1
	*/
}
