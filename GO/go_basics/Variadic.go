package main

import "fmt"

func Variadic() {
	total := Sum(1, 2, 3)
	fmt.Println("Total:", total)
}
func Sum(num ...int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
