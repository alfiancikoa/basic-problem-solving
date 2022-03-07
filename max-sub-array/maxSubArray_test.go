package maxsubarray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk mencari nilai max sub dari array
// Using Kadane's Algorithm
func maxSubSolution1(arr []int) int {
	var max, maxSoFar int
	for i := 0; i < len(arr); i++ {
		maxSoFar = maxSoFar + arr[i]
		if max < maxSoFar {
			max = maxSoFar
		}
		if maxSoFar < 0 {
			maxSoFar = 0
		}
	}
	return max
}

// Fungsi untuk mencari nilai max sub dari array
// Using MohitKumarAlgorithm
func maxSubSolution2(arr []int) int {
	var maxSoFar, currentMax int = arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		currentMax = int(math.Max(float64(arr[i]), float64(currentMax+arr[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(currentMax)))
	}
	return maxSoFar
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Test dari fungsi max Sub di atas
func Test_MaxSubArray(t *testing.T) {
	t.Run("Test_Kadane'sAlgorithm", func(t *testing.T) {
		assert.Equal(t, 0, maxSubSolution1([]int{0}))
		assert.Equal(t, 7, maxSubSolution1([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
		assert.Equal(t, 9, maxSubSolution1([]int{-2, 3, -1, 7, -1}))
		assert.Equal(t, 14, maxSubSolution1([]int{-1, -2, 8, -2, -1, 3, 6}))
	})
	t.Run("Test_MohitKumarAlgorithm", func(t *testing.T) {
		assert.Equal(t, 0, maxSubSolution2([]int{0}))
		assert.Equal(t, 7, maxSubSolution2([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
		assert.Equal(t, 9, maxSubSolution2([]int{-2, 3, -1, 7, -1}))
		assert.Equal(t, 14, maxSubSolution2([]int{-1, -2, 8, -2, -1, 3, 6}))
	})
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_Solution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSubSolution1([]int{-1, -2, 8, -2, -1, 3, 6})
	}
	// Result 14.54 ns/op
}
func Benchmark_Solution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSubSolution2([]int{-1, -2, 8, -2, -1, 3, 6})
	}
	// Result 159.7 ns/op
}
