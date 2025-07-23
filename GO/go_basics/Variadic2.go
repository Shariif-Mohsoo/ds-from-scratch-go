package main

import "fmt"

// Variadic function that multiply numbers
func Mul(nums ...int) {
	fmt.Println("Numbers:", nums)
	if len(nums) == 0 {
		fmt.Println("Mul: 0")
		return
	}
	total := 1
	for _, num := range nums {
		total *= num
	}
	fmt.Println("Mul:", total)
}

func Variadic2() {
	Mul(1, 2)
	Mul(5, 10, 15, 20)
	Mul()
}
