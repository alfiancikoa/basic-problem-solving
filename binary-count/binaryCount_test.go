package binarycount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk menghitung berapa banyak bit 1 atau 0 pada bilangan desimal yang diinputkan
func BinaryCount(angka int, nomorBit int) interface{} {
	if nomorBit > 1 || nomorBit < 0 {
		return nil
	}
	var bit int
	for angka > 0 {
		if angka%2 == nomorBit {
			bit++
		}
		angka /= 2
	}
	return bit
}

// Unit Test untuk fungsi Binary Count di atas
func TestBinaryConvertion(t *testing.T) {
	// 13 = 1 1 0 1
	assert.Equal(t, 3, BinaryCount(13, 1))
	assert.Equal(t, 1, BinaryCount(13, 0))
	assert.Equal(t, nil, BinaryCount(13, 2))
	// 10 = 1 0 1 0
	assert.Equal(t, 2, BinaryCount(10, 1))
	assert.Equal(t, 2, BinaryCount(10, 0))
	assert.Equal(t, nil, BinaryCount(10, 5))
	// 15 = 1 1 1 1
	assert.Equal(t, 4, BinaryCount(15, 1))
	assert.Equal(t, 0, BinaryCount(15, 0))
	assert.Equal(t, nil, BinaryCount(15, -1))
	// Pass
}

func Benchmark_BinaryCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinaryCount(15, 1)
	}
	// Result 6.759 ns/op
}
