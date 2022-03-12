package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Kamus map[string][]string

// Fungsi untuk mencari element dalam kamus
func findElement(str []string, substring string) bool {
	for _, v := range str {
		if substring == v {
			return true
		}
	}
	return false
}

// Fungsi untuk menamabah item ke dalam kamus
func (k Kamus) Add(key string, values []string) string {
	if k[key] == nil {
		k[key] = append(k[key], values...)
	} else {
		for _, value := range values {
			if isExist := findElement(k[key], value); !isExist {
				k[key] = append(k[key], value)
			}
		}
	}
	return "success"
}

// Fungsi untuk mendapatkan data dari kamus
func (k Kamus) Get(word string) []string {
	var result []string
	for key, value := range k {
		if isExist := findElement(value, word); isExist {
			result = append(result, key)
		}
	}
	values := k[word]
	result = append(result, values...)
	return result
}

/*
	*Note:
	Untuk menjalankan Unit Test = go test -v
	Untuk menjalankan Benchmark = go test -bench=.
*/

// Unit Test dari fungsi keseluruhan di atas
func Test_Dictionary(t *testing.T) {
	kamus := Kamus{}
	t.Run("Test_Add", func(t *testing.T) {
		assert.Equal(t, "success", kamus.Add("buah", []string{"pisang", "jeruk"}))
		assert.Equal(t, "success", kamus.Add("buah", []string{"anggur", "sirsak"}))
		assert.Equal(t, "success", kamus.Add("hewan", []string{"kucing", "biawak"}))
		// Pass OK
	})
	t.Run("Test_Get", func(t *testing.T) {
		assert.Equal(t, []string{"pisang", "jeruk", "anggur", "sirsak"}, kamus.Get("buah"))
		assert.Equal(t, []string{"buah"}, kamus.Get("sirsak"))
		assert.Equal(t, []string{"hewan"}, kamus.Get("kucing"))
		// Pass OK
	})
	t.Run("Test_GetIsEmpty", func(t *testing.T) {
		assert.Equal(t, []string([]string(nil)), kamus.Get("flora"))
		// Pass OK
	})
}

// Benchmark untuk menghitung seberapa cepat program kita dieksekusi
func Benchmark_Dictionary(b *testing.B) {
	kamus := Kamus{}
	for i := 0; i < b.N; i++ {
		kamus.Get("biawak")
	}
	// Result 28.62 ns/op
}
