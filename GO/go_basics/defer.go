package main

import "fmt"

func printMess() {
	fmt.Println("I")
	defer fmt.Println("Kahuta")
	fmt.Println("Live")
	fmt.Println("In")
}

func deferKey() {
	//it will be invoked in the end
	defer fmt.Println("Dear")
	//begin from there.
	fmt.Println("\n\nHi")
	printMess()
}
