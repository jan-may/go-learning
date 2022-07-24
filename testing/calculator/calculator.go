package calculator

import "errors"

func Add(left, rigth int) int {
	sum := left + rigth
	return sum
}

func Divide(left, right int) (int, int, error) {
	if right == 0 {
		return 0, 0, errors.New("division by zero")
	}

	return left / right, left % right, nil
}