package pairwithtarget

import "fmt"

/*
	Problem:
	Carilah dua angka apa saja yang terdapat di dalam arr[] yang jika dijumlahkan akan menghasilkan
	angka sama dengan target
	- input:
		Terdapat dua parameter input yaitu:
		# arr sebagai array kumpulan angka yang akan diproses untuk mencari angka pair dari target
		# target merupakan angka yang akan menjadi patokan untuk mencari angka berapa saja yang ada di arr
		yang jika dijumlahkan akan sama dengan target
		# kedua parameter input terdiri dari bilangan bulat positif
	- output:
		# array yang terdiri dari dua angka yang jika kedua angka tersebut dijumlahkan
		akan menghasilkan nilai yang sama dengan target.
		# output hanya terdiri dari 2 angka
	- example :
		1.	#input 	= [1, 2, 3, 4, 5] , 6		#output = [2 4]
		2.	#input 	= [2, 3, 4, 5, 6] , 8		#output = [3 5]
		3.	#input 	= [2 3 4 5 6] , 1			#output = [-1]
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
	if len(result) == 0 {
		result = append(result, -1)
	}
	return result
}

func main() {
	fmt.Println(Pair([]int{1, 2, 3, 4, 5}, 6)) // [2 4]
	fmt.Println(Pair([]int{2, 3, 4, 5, 6}, 8)) // [3 5]
	fmt.Println(Pair([]int{2, 2, 3, 4, 4}, 7)) // [3 4]
	fmt.Println(Pair([]int{2, 2, 3, 4, 4}, 1)) // [3 4]
}
