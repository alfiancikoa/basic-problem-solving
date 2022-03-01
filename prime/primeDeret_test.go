package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi menampilkan deretan angka prima sebanyak n yang diinputkan
func deretPrime(n int) []int {
	var result []int
	var i, count = 1, 0
	for count < n {
		if IsPrimeOptimize(i) {
			result = append(result, i)
			count++
		}
		i++
	}
	return result
}

/*
	*NOTE:
	Run Unit Test = go test -v
	Run Benchmar = go test -bench=.
*/

// Unit Testing dari Fungsi DeretPrima di Atas
func Test_DeretPrime(t *testing.T) {
	assert.Equal(t, []int([]int(nil)), deretPrime(0))
	assert.Equal(t, []int{2}, deretPrime(1))
	assert.Equal(t, []int{2, 3}, deretPrime(2))
	assert.Equal(t, []int{2, 3, 5, 7, 11}, deretPrime(5))
	assert.Equal(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}, deretPrime(10))
	// PASS ok
}

// Benchmark untuk mengetahui seberapa cepat program kita dieksekusi
func Benchmark_DeretPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deretPrime(10)
	}
	//  Result = PASS ok  2165 ns/op
}
