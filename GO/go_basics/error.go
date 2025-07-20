package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func check(err error) bool {
	if err != nil {
		fmt.Println("Error:", err)
		return true
	}
	return false
}

func covertingError() {
	result, err := divide(10, 0)
	if !check(err) {
		fmt.Println("Result:", result)
	}
	result, err = divide(20, 5)
	if !check(err) {
		fmt.Println("Result:", result)
	}
}
