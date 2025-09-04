package functions

import "fmt"

func Functions() {
	fmt.Println("Get started with functions")
	sum()
	fmt.Println("Sum:", add(5, 6))
	a, b, _ := multiReturn()
	fmt.Println("Multi Return:", a, b)
}
func sum() {
	fmt.Println("Sum: ", 1+2)
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func multiReturn() (int, int, int) {
	return 1, 2, 3
}
