package calculator

import "math"

func IsPrime(value int) bool {
	if value <= 1 {
		return false
	}
	if value == 2 {
		return true
	}

	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}