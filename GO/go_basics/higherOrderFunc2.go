package main

import "fmt"

func mul1(x, y int) int {
	return (x * y)
}

func add1(x, y int) int {
	return (x + y)
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return (func(x int) int {
		return mathFunc(x, x)
	})
}

func highOrderFunc2() {
	squareFunc := selfMath(mul1)
	doubleFunc := selfMath(add1)
	// Anonymous function.
	anonFunc := selfMath(func(x, y int) int {
		return x * y
	})

	fmt.Println(squareFunc(5))
	fmt.Println(doubleFunc(5))
	fmt.Println(anonFunc(5))
}
