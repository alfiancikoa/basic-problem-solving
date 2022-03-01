package palindrome

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk Pengecekan Palindrome
func IsPalindrome(str string) bool {
	// remove blank space and make it lower
	input := strings.ToLower(strings.ReplaceAll(str, " ", ""))
	leftIndex, rightIndex := 0, len(input)-1
	for leftIndex < rightIndex {
		// jika karakter yang ada pada index kiri tidak sama dengan karakter index kanan maka langsung return false
		if input[leftIndex] != input[rightIndex-leftIndex] {
			return false
		}
		leftIndex++
	}
	return true
}

// Fungsi untuk Pengecekan Palindrome dengan menggunakan Recursive
func IsPalindromeRecursive(str string, i int) bool {
	if i > len(str)/2 {
		return true
	}
	leftIndex, rightIndex := i, len(str)-i-1
	if str[leftIndex] != str[rightIndex] {
		return false
	}
	return IsPalindromeRecursive(str, i+1)
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Testing dari Fungsi Palindrome di Atas
func TestIsPalindrome(t *testing.T) {
	t.Run("Test_IsPalindrome", func(t *testing.T) {
		assert.Equal(t, true, IsPalindrome("aaaaa"))
		assert.Equal(t, true, IsPalindrome("a"))
		assert.Equal(t, true, IsPalindrome(" "))
		assert.Equal(t, true, IsPalindrome("kodok"))
		assert.Equal(t, true, IsPalindrome("kasur ini rusak"))
		assert.Equal(t, false, IsPalindrome("zoom"))
		assert.Equal(t, false, IsPalindrome("bubuk"))
		// PASS ok
	})
	t.Run("Test_IsPalindromeRecursive", func(t *testing.T) {
		assert.Equal(t, true, IsPalindromeRecursive("aaaaa", 0))
		assert.Equal(t, true, IsPalindromeRecursive("a", 0))
		assert.Equal(t, true, IsPalindromeRecursive(" ", 0))
		assert.Equal(t, true, IsPalindromeRecursive("kodok", 0))
		assert.Equal(t, true, IsPalindromeRecursive("kasurinirusak", 0))
		assert.Equal(t, false, IsPalindromeRecursive("zoom", 0))
		assert.Equal(t, false, IsPalindromeRecursive("bubuk", 0))
		// PASS ok
	})
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("kasurinirusak")
	}
	// Result 117.7 ns/op
}
func BenchmarkIsPalindromeRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeRecursive("kasurinirusak", 0)
	}
	// Result 47.86 ns/op
}
