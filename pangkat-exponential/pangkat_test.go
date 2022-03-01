package pangkatexponential

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function untukk Bilangan Berpangat Positif
func PangkatPositive(a int, b int) int {
	var result int = 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

// Function untukk Bilangan Berpangat Positif dan Negatif
func PangkatNegatif(a int, b int) float64 {
	var result float64 = 1
	var bPositif = int(math.Abs(float64(b)))
	for i := 0; i < bPositif; i++ {
		result *= float64(a)
	}
	if b < 0 {
		result = 1 / result
		return result
	}
	return result
}

// Function untukk Bilangan Berpangat Positif dan Negatif yang diOptimasi
func PangkatOptimasi(a int, b int) float64 {
	var result, temp float64 = 1, 1
	var bPositif = b
	if b < 0 {
		bPositif = -b
	}
	if bPositif%2 > 0 {
		temp = float64(a)
	}
	for i := 0; i < bPositif/2; i++ {
		result = result * float64(a)
	}
	if b < 0 {
		return 1 / ((result * result) * temp)
	}
	return result * result * temp
}

// Function untukk Bilangan Berpangat Positif dengan cara Rekursif
func PangkatRecursive(a int, b int) float64 {
	if b <= 0 {
		return 1
	}
	return float64(a) * PangkatRecursive(a, b-1)
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Testing dari Fungsi Eksponensil di Atas
func TestPangkat(t *testing.T) {
	t.Run("Test_PangkatPositive", func(t *testing.T) {
		assert.Equal(t, 1, PangkatPositive(2, 0))
		assert.Equal(t, 256, PangkatPositive(2, 8))
		assert.Equal(t, 1024, PangkatPositive(2, 10))
		// PASS ok
	})
	t.Run("Test_PangkatNegatif", func(t *testing.T) {
		assert.Equal(t, 1.0, PangkatNegatif(2, 0))
		assert.Equal(t, 256.0, PangkatNegatif(2, 8))
		assert.Equal(t, 1024.0, PangkatNegatif(2, 10))
		assert.Equal(t, 0.0009765625, PangkatNegatif(2, -10))
		assert.Equal(t, 0.012345679012345678, PangkatNegatif(9, -2))
		// PASS ok
	})
	t.Run("Test_PangkatRecursive", func(t *testing.T) {
		assert.Equal(t, 1.0, PangkatRecursive(2, 0))
		assert.Equal(t, 256.0, PangkatRecursive(2, 8))
		assert.Equal(t, 1024.0, PangkatRecursive(2, 10))
		// PASS ok
	})
	t.Run("Test_PangkatOptimasi", func(t *testing.T) {
		assert.Equal(t, 1.0, PangkatOptimasi(2, 0))
		assert.Equal(t, 256.0, PangkatOptimasi(2, 8))
		assert.Equal(t, 1024.0, PangkatOptimasi(2, 10))
		// PASS ok
	})
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func BenchmarkPangkat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PangkatPositive(2, 10)
	}
	// Result 16.13 ns/op
}
func BenchmarkPangkatRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PangkatRecursive(2, 10)
	}
	// Result 64.99 ns/op
}
func BenchmarkPangkatOptimasi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PangkatOptimasi(2, 10)
	}
	// Result 6.269 ns/op
}
