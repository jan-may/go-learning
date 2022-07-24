package calculator_test

import (
	"fmt"
	"testing"

	"github.com/jan-may/go-tests/calculator"
	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T){
	tests:= []struct {
		value int
		isPrime bool
}{
	{-2, false},
	{-1, false},
	{0, false},
	{2, true},
	{3, true},
	{4, false},
	{5, true},
	{6, false},
	{7, true},
	{8, false},
	{9, false},
	{10, false},
	{11, true},
}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.value), func(t *testing.T){
			isPrime := calculator.IsPrime(test.value)
			assert.Equal(t, test.isPrime, isPrime)
		})
	}
}

		
		