package fp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jan-may/go-fp/fp"
	"github.com/stretchr/testify/assert"
)

type agent struct {
	firstName string
	lastName  string
}

func TestMap(t *testing.T) {
	assert.Equal(t, 
		[]int{4, 9, 25, 49, 121},
		fp.Map([]int{2, 3, 5, 7, 11}, func(prime int) int {
			return prime * prime
	}))

	assert.Equal(t, 
		[]int{},
		fp.Map([]int{}, func(prime int) int {
			return prime * prime
	}))

	assert.Equal(t, 
		[]string{"HELLO", "WORLD"},
		fp.Map([]string{"hello", "world"}, func(word string) string {
			return strings.ToUpper(word)
	}))	

	assert.Equal(t, 
		[]string{"Dana Scully", "Fox Mulder"},
		fp.Map([]agent{
			{firstName: "Dana", lastName: "Scully"},
			{firstName: "Fox", lastName: "Mulder"},
		}, func(agent agent) string {
			return fmt.Sprintf("%s %s", agent.firstName, agent.lastName)
	}))	
}