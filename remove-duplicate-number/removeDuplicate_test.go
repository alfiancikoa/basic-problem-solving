package removeduplicatenumber

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi Remove Duplicate with Complexity O(n^2)
func RmDupSolution1(nums []int) []int {
	var result []int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		isDuplicate := false
		for j := i; j < len(nums); j++ {
			if j == i {
				continue
			} else if nums[j] == nums[i] {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			result = append(result, nums[i])
		}
	}
	return result
}

// Fungsi Remove Duplicate with Complexity O(n)
func RmDupSolution2(nums []int) []int {
	var result []int
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			result = append(result, nums[i-1])
		}
	}
	result = append(result, nums[len(nums)-1])
	return result
}

// Fungsi Remove Duplicate Using Data Struct (Hashmap) with Complexity O(n)
func RmDupSolution3(nums []int) []int {
	var result []int
	var hash = map[int]bool{}
	sort.Ints(nums)
	for _, digit := range nums {
		if !hash[digit] {
			result = append(result, digit)
		}
		hash[digit] = true
	}
	return result
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

func Test_RemoveDuplicate(t *testing.T) {
	t.Run("Test_RemoveDupSolution1", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution1([]int{1, 1, 2, 3, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution1([]int{1, 3, 3, 2, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution1([]int{1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution1([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2}, RmDupSolution1([]int{1, 1, 1, 2, 2, 1, 1, 2, 1}))
		assert.Equal(t, []int{1, 2, 3}, RmDupSolution1([]int{3, 2, 1}))
		// PASS ok
	})
	t.Run("Test_RemoveDupSolution2", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution2([]int{1, 1, 2, 3, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution2([]int{1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution2([]int{1, 3, 3, 2, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution2([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2}, RmDupSolution2([]int{1, 1, 1, 2, 2, 1, 1, 2, 1}))
		assert.Equal(t, []int{1, 2, 3}, RmDupSolution2([]int{3, 2, 1}))
		// PASS ok
	})
	t.Run("Test_RemoveDupSolution3", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution3([]int{1, 1, 2, 3, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution3([]int{1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution3([]int{1, 3, 3, 2, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution3([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2}, RmDupSolution3([]int{1, 1, 1, 2, 2, 1, 1, 2, 1}))
		assert.Equal(t, []int{1, 2, 3}, RmDupSolution3([]int{3, 2, 1}))
		// PASS ok
	})
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_RemoveDuplicateSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RmDupSolution1(([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
	}
	// Result 1622 ns/op
}
func Benchmark_RemoveDuplicateSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RmDupSolution2(([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
	}
	// Result 1199 ns/op
}
func Benchmark_RemoveDuplicateSolution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RmDupSolution3(([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
	}
	// Result 1802 ns/op
}
