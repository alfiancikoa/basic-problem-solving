package pairwithtarget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk Pair Target
func pairTarget(nums []int, target int) []int {
	var result []int
	var hash = map[int]int{}
	for i := 0; i < len(nums); i++ {
		keys := target - nums[i]
		if hash[keys] > 0 {
			result = append(result, nums[hash[keys]-1], nums[i])
			break
		}
		hash[nums[i]] = i + 1
	}
	if result == nil {
		result = append(result, -1)
	}
	return result
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

func Test_PairTarget(t *testing.T) {
	assert.Equal(t, []int{1, 2}, pairTarget([]int{1, 2, 3, 4, 5, 6}, 3))
	assert.Equal(t, []int{2, 3}, pairTarget([]int{1, 2, 3, 4, 5, 6}, 5))
	assert.Equal(t, []int{2, 4}, pairTarget([]int{1, 2, 3, 4, 5, 6}, 6))
	assert.Equal(t, []int{3, 4}, pairTarget([]int{1, 2, 3, 4, 5, 6}, 7))
	assert.Equal(t, []int{-1}, pairTarget([]int{1, 4, 2, 5, 11, 12}, 8))
	// PASS ok
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_PairTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pairTarget([]int{1, 2, 3, 4, 5, 6}, 7)
	}
	// Result 299.3 ns/op
}
