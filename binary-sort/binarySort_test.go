package binarysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi mengurutkan angka dengan menggunakan algoritma binary sort
func BinarySort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	tengah := arr[0]
	kiri, kanan := []int{}, []int{}
	for index, angka := range arr {
		if index == 0 {
			continue
		} else if angka < tengah {
			kiri = append(kiri, angka)
		} else {
			kanan = append(kanan, angka)
		}
	}
	kiri = BinarySort(kiri)
	kanan = BinarySort(kanan)
	result := append(kiri, tengah)
	result = append(result, kanan...)
	return result
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Test dari fungsi Binary Sort di atas
func Test_BinarySort(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4}, BinarySort([]int{3, 4, 2}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, BinarySort([]int{5, 4, 2, 3, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, BinarySort([]int{5, 1, 3, 2, 4}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, BinarySort([]int{2, 3, 1, 5, 10, 4, 7, 8, 6, 9}))
	// PASS ok
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_BinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySort([]int{2, 3, 1, 5, 10, 4, 7, 8, 6, 9})
	}
	// Result 4599 ns/op
}
