package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk mencari suatu bilangan dan mengembalikan index-nya
func LinearSearch(arr []int, target int) int {
	return -1
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Test dari fungsi Linear Search di atas
func Test_LinearSearch(t *testing.T) {
	assert.Equal(t, -1, LinearSearch([]int{}, 5))

}
