package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk mencari suatu bilangan dan mengembalikan index-nya
func LinearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if target == arr[i] {
			return i
		}
	}
	return -1
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Test dari fungsi Linear Search di atas
func Test_LinearSearch(t *testing.T) {
	assert.Equal(t, -1, LinearSearch([]int{}, 5))
	assert.Equal(t, 3, LinearSearch([]int{1, 2, 3, 4, 5}, 4))
	assert.Equal(t, 6, LinearSearch([]int{2, 3, 4, 1, 6, 5, 7, 8, 9}, 7))

}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_BinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LinearSearch([]int{2, 3, 4, 1, 6, 5, 7, 8, 9}, 7)
	}
	// Result 16.82 ns/op
}
