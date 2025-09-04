package functions

import "fmt"

func Exercise() {
	greetPerson("Alice")
	fmt.Println(hiThere())
	a, b := twoTwos()

	answer := addThree(five(), a, b)
	fmt.Println(answer)
}

// Write a function that returns any two values
func twoTwos() (int, int) {
	return 2, 2
}

// write a function that returns any number
func five() int {
	return 5
}

// write a function to add 3 numbers together, supplied as arguments, and return the answer
func addThree(a, b, c int) int {
	return a + b + c
}

// write a function that accepts a person's name as a function parameter and displays a greetings to that person
func greetPerson(name string) {
	fmt.Println("Hi! ", name)
}

// write a function that returns any message, and call it from within fmt.Println()
func hiThere() string {
	return "Hi There"
}
