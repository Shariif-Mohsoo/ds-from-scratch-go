package errors

import (
	"errors"
	"fmt"
)

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("can't divide by zero")
	} else {
		return rhs / lhs, nil
	}
}

func Errors() {
	fmt.Println(divide(2, 0))
	fmt.Println(
		divide(2, 3),
	)
}
