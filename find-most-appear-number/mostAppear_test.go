package findmostappearnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk mencari angka yang paling sering muncul
func mostAppearNumber(nums []int) int {
	var hash = map[int]int{}
	var maxSoFar, mostNumber int
	for _, digit := range nums {
		hash[digit]++
		if hash[digit] >= maxSoFar && digit >= mostNumber {
			maxSoFar = hash[digit]
			mostNumber = digit
		}
	}
	return mostNumber
}

/*
	*NOTE:
	Run Unit Test = go test -v
	Run Benchmar = go test -bench=.
*/

// Unit Test dari Fungsi MostAppearNumber di Atas
func TestMostAppearNumber(t *testing.T) {
	assert.Equal(t, 1, mostAppearNumber([]int{1, 1, 1, 1, 1}))
	assert.Equal(t, 101, mostAppearNumber([]int{90, 8, 90, 68, 101, 4, 1, 2, 101, 101}))
	assert.Equal(t, 100, mostAppearNumber([]int{92, 8, 92, 68, 100, 4, 1, 2, 100, 100}))
	assert.Equal(t, 5, mostAppearNumber([]int{5, 5, 1, 1, 2, 3, 10, 9}))
	assert.Equal(t, 6, mostAppearNumber([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 6, mostAppearNumber([]int{4, 3, 6, 2, 1, 5}))
}

// Benchmark untuk mengetahui seberapa cepat program kita dieksekusi
func BenchmarkMostAppearNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mostAppearNumber([]int{92, 8, 92, 68, 100, 4, 1, 2, 100, 100})
	}
	//  Result = PASS ok  1758 ns/op
}
