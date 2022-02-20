package main

import "fmt"

func Pascal(num int) {
	for i := 1; i < num+1; i++ {
		pascal := 1
		for j := 1; j < i+1; j++ {
			fmt.Print(pascal, " ")
			pascal = int(pascal * (i - j) / j)
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
