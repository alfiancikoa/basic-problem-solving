package faktorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi mencari nilai faktorial sederhana
func Faktorial(n int) int {
	if n <= 0 {
		return 1
	}
	var result int = 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

// Fungsi mencari nilai faktorial dengan cara recursive
func FaktorialRecursive(total int, n int) int {
	if n <= 0 {
		return total
	}
	return FaktorialRecursive(total*n, n-1)
}

/*
	*NOTE
	Running Unit Test -> go test -v
	Running Benchmark -> go test -bench=.
*/

// Unit Test untuk fungsi faktorial di atas
func TestFaktorial(t *testing.T) {
	t.Run("Test_FaktorialBiasa", func(t *testing.T) {
		// Faktorial biasa
		assert.Equal(t, 1, Faktorial(0))
		assert.Equal(t, 1, Faktorial(1))
		assert.Equal(t, 2, Faktorial(2))
		assert.Equal(t, 6, Faktorial(3))
		assert.Equal(t, 24, Faktorial(4))
		assert.Equal(t, 120, Faktorial(5))
		assert.Equal(t, 479001600, Faktorial(12))
		// Result PASS
	})
	t.Run("Test_FaktorialRecursive", func(t *testing.T) {
		// Faktorial Recursive
		assert.Equal(t, 1, FaktorialRecursive(1, 0))
		assert.Equal(t, 1, FaktorialRecursive(1, 1))
		assert.Equal(t, 2, FaktorialRecursive(1, 2))
		assert.Equal(t, 6, FaktorialRecursive(1, 3))
		assert.Equal(t, 24, FaktorialRecursive(1, 4))
		assert.Equal(t, 120, FaktorialRecursive(1, 5))
		assert.Equal(t, 479001600, FaktorialRecursive(1, 12))
		// Result PASS
	})

}

// Benchmark untuk mengetahui seberapa cepat program dieksekusi
func BenchmarkFaktorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Faktorial(12)
	}
	// Result 23.41 ns/op
}
func BenchmarkFaktorialRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FaktorialRecursive(1, 12)
	}
	// Result 46.95 ns/op
}
