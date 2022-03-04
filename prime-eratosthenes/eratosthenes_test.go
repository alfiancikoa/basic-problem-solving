package primeeratosthenes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk menampilkan bilangan Prime dengan menggunakan algoritma Eratosthhenes yang pertama
func primeEratosthenes1(n int) []int {
	//  membuat tempat tampungan untuk bilanagn prima
	var prime, result []int
	// set seluruh array menjadi true
	for i := 0; i < n; i++ {
		prime = append(prime, 1)
	}
	// set array ke 0 dan 1 menjadi false, karena bukan bilangan prima
	prime[0], prime[1] = 0, 0
	// coret bilangan yang bukan prima dengan mengeset menjadi false
	for i := 2; i < n; i++ {
		for j := 2; j*i < n; j++ {
			prime[j*i] = 0
		}
	}
	// tampilkan seluruh bilangan prima, nilai di dalam array adalah true
	for i := 0; i < n; i++ {
		if prime[i] == 1 {
			result = append(result, i)
		}
	}
	return result
}

/*
	*NOTE:
	Run Unit Test = go test -v
	Run Benchmar = go test -bench=.
*/

// Unit Testing dari Fungsi Prime di Atas
func TestPrimeEratosthenes(t *testing.T) {
	t.Run("Test_Eratosthenes1", func(t *testing.T) {
		assert.Equal(t, []int{2, 3}, primeEratosthenes1(5))
		assert.Equal(t, []int{2, 3, 5, 7}, primeEratosthenes1(10))
		assert.Equal(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}, primeEratosthenes1(50))
	})
	// PASS ok
}

// Benchmark untuk mengetahui seberapa cepat program kita dieksekusi
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeEratosthenes1(50)
	}
	//  Result = PASS ok  3248 ns/op
}
