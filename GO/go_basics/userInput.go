package main

import "fmt"

func userInput() {
	var userName string
	var age int
	fmt.Print("Enter Your Name: ")
	//assigning user's input to userName
	fmt.Scan(&userName)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("\n Name:", userName, "Age:", age)
}
