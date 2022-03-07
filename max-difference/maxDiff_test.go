package maxdifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Fungsi untuk mencari max difference
func maxDiff(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var maxDiff, min int = 0, nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		} else if maxDiff < nums[i]-min {
			maxDiff = nums[i] - min
		}
	}
	return maxDiff
}

func Test_MaxDiff(t *testing.T) {
	assert.Equal(t, 0, maxDiff([]int{0}))
	assert.Equal(t, 1, maxDiff([]int{1, 2}))
	assert.Equal(t, 5, maxDiff([]int{1, 2, 6, 4}))
	assert.Equal(t, 3, maxDiff([]int{1, 3, 2, 4}))
	assert.Equal(t, 7, maxDiff([]int{-1, -3, 2, 4}))
	assert.Equal(t, 5, maxDiff([]int{3, 2, 4, 7}))
	assert.Equal(t, 7, maxDiff([]int{2, 7, 9, 5, 1, 3, 5}))
	assert.Equal(t, 9, maxDiff([]int{2, 7, 9, 5, 1, 5, 10}))
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func BenchmarkPangkat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDiff([]int{2, 7, 9, 5, 1, 5, 10})
	}
	// Result 30.57 ns/op
}
