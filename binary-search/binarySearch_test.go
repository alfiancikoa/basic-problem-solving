package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi searching dengan menggunakan algoritma binary search
func BinarySearch(arr []int, target int) int {
	var left, right int = 0, len(arr) - 1
	for left <= right {
		mid := ((right - left) / 2) + left
		if target > arr[mid] {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Test dari fungsi Binary Search di atas
func Test_BinarySearch(t *testing.T) {
	arrNum := make([]int, 100)
	for i := 0; i < 100; i++ {
		arrNum[i] = i + 1
	}
	assert.Equal(t, -1, BinarySearch([]int{1, 2, 3, 4, 5}, 6))
	assert.Equal(t, 3, BinarySearch([]int{1, 2, 3, 4, 5}, 4))
	assert.Equal(t, 94, BinarySearch(arrNum, 95))
	// PASS ok
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_BinarySearch(b *testing.B) {
	arrNum := make([]int, 100)
	for i := 0; i < 100; i++ {
		arrNum[i] = i + 1
	}
	for i := 0; i < b.N; i++ {
		BinarySearch(arrNum, 99)
	}
	// Result 18.06 ns/op
}
