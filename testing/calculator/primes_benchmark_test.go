package calculator_test

import (
	"testing"

	"github.com/jan-may/go-tests/calculator"
)

func benchmarkIsPrime(b *testing.B, value int){
	// Setup ...
	b.ResetTimer()

	for i:= 0 ; i < b.N; i++ {
		calculator.IsPrime(value)
	}
}

func BenchmarkIsPrime987987654(b *testing.B) {
	benchmarkIsPrime(b, 9879876154)
}

func BenchmarkIsPrime2(b *testing.B) {
	benchmarkIsPrime(b, 2)
}

