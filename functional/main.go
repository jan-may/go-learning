package main

import (
	"fmt"
	"strings"

	"github.com/jan-may/go-fp/fp"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}

	squares := fp.Map(primes, func(prime int) int {
		return prime * prime
	})

	words := []string{"hello", "world"}
	capitalizedWords := fp.Map(words, func(word string) string {
		return strings.ToUpper(word)
	})
	fmt.Println(squares)
	fmt.Println(capitalizedWords)
}

