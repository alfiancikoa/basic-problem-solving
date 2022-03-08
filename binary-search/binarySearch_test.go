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

// Fungsi searching dengan menggunakan algoritma binary search dengan Recursive
func BinarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := ((right - left) / 2) + left
	if arr[mid] == target {
		return mid
	}
	if target < arr[mid] {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}
	return BinarySearchRecursive(arr, target, mid+1, right)
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
	t.Run("Test_BinarySearc", func(t *testing.T) {
		assert.Equal(t, -1, BinarySearch([]int{1, 2, 3, 4, 5}, 6))
		assert.Equal(t, 3, BinarySearch([]int{1, 2, 3, 4, 5}, 4))
		assert.Equal(t, 94, BinarySearch(arrNum, 95))
		// PASS ok
	})
	t.Run("Test_BinarySearchRecursive", func(t *testing.T) {
		assert.Equal(t, 3, BinarySearchRecursive([]int{1, 2, 3, 4, 5}, 4, 0, 5))
		assert.Equal(t, 3, BinarySearchRecursive([]int{1, 2, 3, 4, 5}, 4, 0, 5))
		assert.Equal(t, 94, BinarySearchRecursive(arrNum, 95, 0, len(arrNum)))
		// PASS ok
	})
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
	// Result 21.20 ns/op
}
