package palindrome

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/*
	Problem:
	Mengecek Apakah kata atau kalimat input merupakan palindrome atau bukan
	- input:
		satu atau lebih kata bertipe string, (a-z,A-Z)
	- output:
		status apakah palindrome atau bukan (true or false)
	- example :
		1.	#input 	= "kodok"	#output = true
		2.	#input 	= "kakak"	#output = true
		3.	#input 	= "makan"	#output = false
*/

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

func TestIsPalindrome(t *testing.T) {
	t.Run("Test_IsPalindrome", func(t *testing.T) {
		assert.Equal(t, true, IsPalindrome("aaaaa"))
		assert.Equal(t, true, IsPalindrome("a"))
		assert.Equal(t, true, IsPalindrome(" "))
		assert.Equal(t, true, IsPalindrome("kodok"))
		assert.Equal(t, true, IsPalindrome("kasur ini rusak"))
		assert.Equal(t, true, IsPalindrome("Kasur Koh Ahok rusak"))
		assert.Equal(t, true, IsPalindrome("Ibu Ratna antar ubi"))
		assert.Equal(t, false, IsPalindrome("zoom"))
		assert.Equal(t, false, IsPalindrome("bubuk"))
	})
	// Result Pass
}

func BenchmarkFaktorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("Ibu Ratna antar ubi")
	}
	// Result 542.7 ns/op
}
