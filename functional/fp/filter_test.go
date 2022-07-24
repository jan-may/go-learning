package fp_test

import (
	"testing"

	"github.com/jan-may/go-fp/fp"
	"github.com/stretchr/testify/assert"
)


func TestFilter(t *testing.T) {

	assert.Equal(t, 
		[]int{},
		fp.Filter([]int{}, func(value int) bool {
			return value < 10
	}))

	assert.Equal(t, 
		[]int{1,2},
		fp.Filter([]int{1,2,10,15,20}, func(value int) bool {
			return value < 10
	}))


}