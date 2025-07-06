package main

import "fmt"

func arrayUsage() {
	//creating an array
	var users [5]string //1st way to create an array
	//var users = [5]string //2nd way to create an array

	//storing values in an array
	users[0] = "Mohsin"
	users[1] = "Jawad"
	users[2] = "Saad"
	users[3] = "Zain"
	users[4] = "Gohar"

	//displaying info

	fmt.Printf("The whole array: %v\n", users)
	fmt.Printf("The first value: %v\n", users[0])
	fmt.Printf("The Array Type: %T\n", users)
	fmt.Printf("The whole array length: %v\n", len(users))

}
