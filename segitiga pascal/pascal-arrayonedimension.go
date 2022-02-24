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
	// membuat array untuk menampung nilai yang sudah dihitung
	array := make([]int, num+1)
	prev := 1
	fmt.Println(1)
	for i := 1; i < num; i++ {
		array[i] = 1
		fmt.Print(1, " ")
		for j := 1; j < i; j++ {
			number := prev + array[j]
			prev = array[j]
			array[j] = number
			fmt.Print(number, " ")
		}
		fmt.Println(1)
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
