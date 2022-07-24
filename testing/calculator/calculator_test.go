package calculator_test

import (
	"testing"

	"github.com/jan-may/go-tests/calculator"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		left, right int
		sum int
	}{
		{"positive integers", 1, 2, 3},
		{"negative integers", -23, -42, -65},
		{"zeros", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T) {
			sum := calculator.Add(test.left, test.right)
			assert.Equal(t, test.sum, sum)
		})
	}
}


func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		left, right int
		quotient, remainder int
		expectError bool
	}{
		{"positive integers", 17, 3, 5, 2, false},
		{"zero", 17, 0, 0 , 0, true},
		{"negative integers", -17, 3, -5, -2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T) {
			quotient, remainder, err := calculator.Divide(test.left, test.right)

			if test.expectError {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, test.quotient, quotient)
			assert.Equal(t, test.remainder, remainder)
			assert.NoError(t, err)
		})
	}
}