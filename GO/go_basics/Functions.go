package main

import "fmt"

func airthmaticOperations() {
	x, y := getInput()
	fmt.Println("Sum:", sum(x, y))
	fmt.Println("Sub:", sub(x, y))
	fmt.Println("Mul:", mul(x, y))
	fmt.Println("Div:", div(x, y))
}

func getInput() (x, y int) {
	fmt.Println("Enter first number(integer):")
	fmt.Scan(&x)
	fmt.Println("Enter second number(integer):")
	fmt.Scan(&y)
	//This will implicitly return the x and y.
	return
}

func sum(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) float32 {
	if y > 0 {
		var res float32 = float32(x / y)
		return res
	} else {
		return 0
	}
}
