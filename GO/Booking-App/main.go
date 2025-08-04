package main

import (
	"fmt"
	"time"
)

// PACKAGE LEVEL VARIABLES.
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// creating an array to store persons who booked tickets
// var booking= [50]string{}
// var booking [50]string
// slices are must more convinent then the arrays so slice is
// var bookings []string
// map type slice
// var bookings = make([]map[string]string, 0)

// Creating a custom type (structure)
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Create a user data slice type(empty list)
var bookings = make([]UserData, 0)

func main() {
	//note space will be added by default when the variables value go to print.
	greetUsers()
	//using the infinite for loop for repeated input
	for {
		//start body
		//fetching the input from the user
		firstName, lastName, email, userTickets := getUserInput()
		//user input validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		// if userTickets > remainingTickets {
		// 	fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		// 	continue
		// }

		if isValidName && isValidEmail && isValidTicketNumber {
			//book the ticket
			bookTicket(userTickets, firstName, lastName, email)
			//send ticket info
			//it doesn' t give concurreny.
			// sendTicket(userTickets, firstName, lastName, email)

			//using goroutines (lightweight thread) to acheive concurrency
			go sendTicket(userTickets, firstName, lastName, email) //generating and sending ticket task runs now in the background.

			//call the function to print first name only
			firstNames := getFirstNames()
			//printing the first names
			fmt.Printf("These are all our bookings: %v\n", firstNames)
			//defining the condition
			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firstName or lastName you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
		//end body
	}
}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// function to print first name of people who book tickets
func getFirstNames() []string {
	//creating a new slice that will contains first name only
	firstNames := []string{}

	for _, booking := range bookings {
		//this will separate the string(fullName) on first name and last name then return the slice
		// var names = strings.Fields(booking)
		//stroring the firstName
		//var firstName = names[0]
		//inserting the value in the slice
		// firstNames = append(firstNames, firstName)

		//accessing data from the map
		// firstNames = append(firstNames, booking["firstName"])

		//accessing data from the structure now
		firstNames = append(firstNames, booking.firstName)
	}
	//returning the firstNames slice
	return firstNames
}

// getting the input from the user
func getUserInput() (string, string, string, uint) {
	//user variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// function for booking the ticket
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	//Create a map for a user.(empty map)
	//var userData = make(map[string]string)
	//inserting data into the map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//creating a struct of UserData
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//booking[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	//display the list of bookings
	//fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaning for %v\n", remainingTickets, conferenceName)
}

// Send Tickets
func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	//sleep for  10 seconds.
	time.Sleep(10 * time.Second)
	//after 10 second the below written code will be excecuted.
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("######################")

}
