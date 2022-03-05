package findmostappearnumber

/*
	Problem:
	Mencari angka yang paling sering muncul
	jika porsi kemunculan angka sama, maka cari angka yang nilainya lebih tinggi
	- input:
		deretan angka berupa array 1 dimensi bilangan positif
	- output:
		angka yang paling sering muncul
	- example :
		1.	#input 	= [90 8 90 68 101 4 1 2 101 101]	#output = 101
		2.	#input 	= [1 2 3 4 5 6]						#output = 6
		3.	#input 	= [5 5 1 1 2 3 10 9]				#output = 5
*/

func mostAppear(num []int) int {
	hashmap := map[int]int{}
	mostSoFar, mostNumber := 0, 0
	for _, digit := range num {
		hashmap[digit]++
		if hashmap[digit] >= mostSoFar && digit >= mostNumber {
			mostNumber = digit
			mostSoFar = hashmap[digit]
		}
	}
	return mostNumber
}

// func main() {
// 	fmt.Println(mostAppear([]int{90, 8, 90, 68, 101, 4, 1, 2, 101, 101})) // 101
// 	fmt.Println(mostAppear([]int{92, 8, 92, 68, 100, 4, 1, 2, 100, 100})) // 100
// 	fmt.Println(mostAppear([]int{5, 5, 1, 1, 2, 3, 10, 9}))               // 5
// 	fmt.Println(mostAppear([]int{1, 2, 3, 4, 5, 6}))                      // 6
// 	fmt.Println(mostAppear([]int{4, 3, 6, 2, 1, 5}))                      // 6
// 	fmt.Println(mostAppear([]int{1, 1, 1, 1, 1, 1}))                      // 1
// }
