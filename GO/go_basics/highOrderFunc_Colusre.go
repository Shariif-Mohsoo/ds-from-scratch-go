package main

import "fmt"

// Higher-order function: takes a function and returns a closure
func logExecutes(f func()) func() {
	return func() {
		defer fmt.Println("Finished Execution")
		fmt.Println("\n\nStarting Execution")
		f() //calling the orignal function.
	}
}

func sayHello() {
	fmt.Println("Hello From Inside Function")
}

func practiceHighOrder_Closure() {
	loggedFunc := logExecutes(sayHello)
	loggedFunc()
}
