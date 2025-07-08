package main

import "fmt"

func ifElse() {
	var age int
	fmt.Println("Enter Your Age:")
	fmt.Scan(&age)

	if age > 0 {
		fmt.Println("You are", age, "years old")
	} else {
		fmt.Println("Something is wrong")
	}

}
