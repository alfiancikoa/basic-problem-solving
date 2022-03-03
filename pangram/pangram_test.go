package pangram

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk mengecek kalimat Pangram
func Pangram(s string) string {
	// simpan seluruh karakter ke dalam map
	var hashmap = map[string]int{}
	// proses pengambilan karakter yang digunakan
	for _, value := range strings.ToLower(s) {
		if value != ' ' {
			hashmap[string(value)]++
		}
	}
	// cek apakah jumlah huruf sebanyak 26 berarti semua karakter huruf dipakai
	if len(hashmap) == 26 {
		return "pangram"
	}
	return "not pangram"
}

/*
	*NOTE:
	Run Unit Test = go test -v
	Run Benchmar = go test -bench=.
*/

// Unit Test untuk Function Pangram di Atas
func Test_Pangram(t *testing.T) {
	assert.Equal(t, "pangram", Pangram("We promptly judged antique ivory buckles for the next prize"))
	assert.Equal(t, "pangram", Pangram("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "not pangram", Pangram("We promptly judged antique ivory buckles for the prize"))
	// PASS ok
}

// Benchmark diunakan untuk menegcek seberapa cepat program dieksekusi
func Benchmark_PascalArray2D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pangram("We promptly judged antique ivory buckles for the next prize")
	}
	// PASS ok 10719 ns/op
}
