package segitigapascal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk menampilkan segitiga pascal dengan menggunakan bantuan Array 1 Dimensi
func Pascal_Array1D(n int) string {
	var result string = "1 \n"
	array := make([]int, n+1)
	prev := 1
	for i := 1; i < n; i++ {
		array[i] = 1
		result += fmt.Sprint(1, " ")
		for j := 1; j < i; j++ {
			number := prev + array[j]
			prev = array[j]
			array[j] = number
			result += fmt.Sprint(number, " ")
		}
		result += fmt.Sprintln(1)
	}
	return result
}

// Fungsi untuk menampilkan segitiga pascal dengan menggunakan bantuan Array 2 Dimensi
func Pascal_Array2D(n int) string {
	var result string
	array := make([][]int, n)
	for i := 0; i < n; i++ {
		array[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				array[i][j] = 1
			} else {
				array[i][j] = array[i-1][j-1] + array[i-1][j]
			}
			result += fmt.Sprintf("%d ", array[i][j])
		}
		result += "\n"
	}
	return result
}

// Fungsi untuk menampilkan segitiga pascal dengan menggunakan bantuan rumus formula
func Pascal_Formula(n int) string {
	var result string
	for i := 1; i < n+1; i++ {
		var Pascal int = 1
		for j := 1; j < i+1; j++ {
			result += fmt.Sprint(Pascal, " ")
			Pascal = int(Pascal * (i - j) / j)
		}
		result += "\n"
	}
	return result
}

// Unit Test untuk Function Segitiga Pascal di Atas
func TestPascalFunction(t *testing.T) {
	t.Run("Test_PascalArray2D", func(t *testing.T) {
		assert.Equal(t, "1 \n1 1 \n1 2 1 \n", Pascal_Array2D(3))
		assert.Equal(t, "1 \n1 1 \n1 2 1 \n1 3 3 1 \n1 4 6 4 1 \n", Pascal_Array2D(5))
		fmt.Println(Pascal_Array2D(5))
		/*
			1
			1 1
			1 2 1
			1 3 3 1
			1 4 6 4 1
		*/
		// PASS ok
	})
	t.Run("Test_PascalFormula", func(t *testing.T) {
		assert.Equal(t, "1 \n1 1 \n1 2 1 \n", Pascal_Formula(3))
		assert.Equal(t, "1 \n1 1 \n1 2 1 \n1 3 3 1 \n1 4 6 4 1 \n", Pascal_Formula(5))
		fmt.Println(Pascal_Formula(5))
		/*
			1
			1 1
			1 2 1
			1 3 3 1
			1 4 6 4 1
		*/
		// PASS ok
	})
	t.Run("Test_PascalArray1D", func(t *testing.T) {
		assert.Equal(t, "1 \n1 1\n1 2 1\n", Pascal_Array1D(3))
		assert.Equal(t, "1 \n1 1\n1 2 1\n1 3 3 1\n1 4 6 4 1\n", Pascal_Array1D(5))
		fmt.Println(Pascal_Array1D(5))
		/*
			1
			1 1
			1 2 1
			1 3 3 1
			1 4 6 4 1
		*/
		// PASS ok
	})
}

// Benchmark diunakan untuk menegcek seberapa cepat program dieksekusi
func Benchmark_PascalArray2D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pascal_Array2D(6)
	}
	// PASS ok 11244 ns/op
}
func Benchmark_PascalFormula(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pascal_Formula(6)
	}
	// PASS ok 12327 ns/op
}
func Benchmark_PascalArray1D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pascal_Array1D(6)
	}
	// PASS ok 12324 ns/op
}
