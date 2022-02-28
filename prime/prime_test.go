package prime

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Problem:
	Menentukan apakah bilangan yang di-inputkan merupakan bilangan prima atau bukan
	- input:
		Berupa angka positif
	- output:
		True Atau False
	- example :
		1.	#input 	= 0			#output = false
		2.	#input 	= 2			#output = true
		3.	#input 	= 11			#output = true
*/

func IsPrimeBasic(num int) bool {
	var faktor int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			faktor++
		}
	}
	return faktor == 2
}

func IsPrimeOptimize(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

/*
	*NOTE:
	Run Unit Test = go test -v
	Run Benchmar = go test -bench=.
*/

func TestPrime(t *testing.T) {
	t.Run("Test_IsPrime", func(t *testing.T) {
		assert.Equal(t, false, IsPrimeBasic(0))
		assert.Equal(t, true, IsPrimeBasic(2))
		assert.Equal(t, true, IsPrimeBasic(3))
		assert.Equal(t, false, IsPrimeBasic(4))
		assert.Equal(t, true, IsPrimeBasic(5))
		assert.Equal(t, true, IsPrimeBasic(11))
		assert.Equal(t, false, IsPrimeBasic(15))
		assert.Equal(t, true, IsPrimeBasic(1000000007))
	})
	// PASS ok
	t.Run("Test_IsPrimeOptimize", func(t *testing.T) {
		assert.Equal(t, false, IsPrimeOptimize(0))
		assert.Equal(t, true, IsPrimeOptimize(2))
		assert.Equal(t, true, IsPrimeOptimize(3))
		assert.Equal(t, false, IsPrimeOptimize(4))
		assert.Equal(t, true, IsPrimeOptimize(5))
		assert.Equal(t, true, IsPrimeOptimize(11))
		assert.Equal(t, false, IsPrimeOptimize(15))
		assert.Equal(t, true, IsPrimeOptimize(1000000007))
	})
	// PASS ok
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrimeBasic(1000000007)
	}
	//  Result = PASS ok  664420768 ns/op
}
func BenchmarkIsPrimeOptimize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrimeOptimize(1000000007)
	}
	//  Result = PASS ok  670954 ns/op
}
