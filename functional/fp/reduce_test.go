package fp_test

import (
	"testing"

	"github.com/jan-may/go-fp/fp"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {

	assert.Equal(t, 
		0,
		fp.Reduce([]int{}, func(sum int, value int) int {
			return sum + value
	}, 0))

	assert.Equal(t, 
		15,
		fp.Reduce([]int{1,2,3,4,5}, func(sum int, value int) int {
			return sum + value
	}, 0))

	assert.Equal(t, 
		120,
		fp.Reduce([]int{1,2,3,4,5}, func(product int, value int) int {
			return product * value
	}, 1))

}