package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Kamus map[string][]string

func findElement(str []string, substring string) bool {
	for _, v := range str {
		if substring == v {
			return true
		}
	}
	return false
}

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

func Test_Dictionary(t *testing.T) {
	kamus := Kamus{}
	t.Run("Test_Add", func(t *testing.T) {
		assert.Equal(t, "success", kamus.Add("buah", []string{"pisang", "jeruk"}))
		assert.Equal(t, "success", kamus.Add("buah", []string{"anggur", "sirsak"}))
		assert.Equal(t, "success", kamus.Add("hewan", []string{"kucing", "biawak"}))
	})
	t.Run("Test_Get", func(t *testing.T) {
		assert.Equal(t, []string{"pisang", "jeruk", "anggur", "sirsak"}, kamus.Get("buah"))
		assert.Equal(t, []string{"buah"}, kamus.Get("sirsak"))
		assert.Equal(t, []string{"hewan"}, kamus.Get("kucing"))

	})
	t.Run("Test_GetIsEmpty", func(t *testing.T) {
		assert.Equal(t, []string([]string(nil)), kamus.Get("flora"))
	})
}
