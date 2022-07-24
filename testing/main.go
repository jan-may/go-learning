package main

import (
	"fmt"

	"github.com/jan-may/go-tests/calculator"
)

func main() {
	sum := calculator.Add(23, 42)
	fmt.Println(sum)
}