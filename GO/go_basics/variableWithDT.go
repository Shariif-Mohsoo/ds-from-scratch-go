package main

import "fmt"

func dataTypes() {
	var userName string
	var userAge int
	var ticketBooked bool

	userName = "Imran"
	userAge = 20
	ticketBooked = false

	fmt.Println("\nuser ", userName, " with age", userAge, " booked the ticket:", ticketBooked)

	//checking the types with the help of Printf and Placeholder (%T)
	fmt.Printf("userName is %T, userAge is %T and ticketBooked is %T\n", userName, userAge, ticketBooked)

}
