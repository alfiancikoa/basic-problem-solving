package findindexsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk mencari index dari substring
func indexOf(a string, b string) int {
	for i := 0; i <= len(a)-len(b); i++ {
		str := a[i : len(b)+i]
		if b == str {
			return i
		}
	}
	return -1
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Testing dari Fungsi indexO di Atas
func Test_IndexOf(t *testing.T) {
	assert.Equal(t, -1, indexOf("abcdef", "df"))
	assert.Equal(t, 3, indexOf("abcdef", "de"))
	assert.Equal(t, 3, indexOf("abcdef", "def"))
	assert.Equal(t, -1, indexOf("abcdef", "cef"))
	assert.Equal(t, -1, indexOf("abcddef", "cde"))
	assert.Equal(t, 2, indexOf("abcdef", "cde"))
	assert.Equal(t, 5, indexOf("abccccdef", "cde"))
	assert.Equal(t, 5, indexOf("abccdcdecdf", "cde"))
	// PASS ok
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_IndexOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		indexOf("abccdcdecdf", "cde")
	}
	// Result 71.44 ns/op
}
