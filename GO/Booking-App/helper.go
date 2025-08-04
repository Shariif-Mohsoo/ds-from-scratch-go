package main

import "strings"

// function to validate the users input
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	//firstName and lastName validation(length 2 or greater)
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	//email validation
	isValidEmail := strings.Contains(email, "@gmail.com")
	//Tickets Validation
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	//returning the desire values
	return isValidName, isValidEmail, isValidTicketNumber
}
