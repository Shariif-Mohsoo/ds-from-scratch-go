package main

import (
	"fmt"
)

// creating a structure
type UserData struct {
	firstName     string
	lastName      string
	age           uint
	martialStatus string
	employeed     bool
	disable       bool
}

func pracitce() {
	//get random input from the user
	getUserInput()
}

func getUserInput() {
	var n int
	fmt.Println("Enter any number")
	fmt.Scan(&n)
	fmt.Println("you entered:", n)
	//print the pattern
	/// Instead of limiting the user to enter specific limit number i should prefer to use threads (light weight thread okay)
	//a global one wg(waitGroup object from other file)
	wg.Add(1)

	go printPattern(n)
	//
	greetUser()
	getUserData()
	wg.Wait()
}
func printPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	wg.Done()
}

func greetUser() {
	fmt.Println("\n HI Dear! What going on dear?")
}

func getUserData() {
	//creating a slice for users
	var UsersDataSlice = make([]UserData, 0)

	//using infinite loop to get user info and terminates when user will be done
	getDataFromUser(UsersDataSlice)
}

func getDataFromUser(UsersDataSlice []UserData) {
	//creating a variables (declaring)
	var firstName, lastName, martialStatus string
	var employeed, disable bool
	var age uint
	for {
		fmt.Println("Enter firstName:")
		fmt.Scan(&firstName)
		fmt.Println("Enter lastName:")
		fmt.Scan(&lastName)
		fmt.Println("Enter age:")
		fmt.Scan(&age)
		fmt.Println("Enter martial status:")
		fmt.Scan(&martialStatus)
		fmt.Println("Are you employeed (true/false):")
		fmt.Scan(&employeed)
		fmt.Println("Are you disable (true/false):")
		fmt.Scan(&disable)
		//inserting data into struct
		var userData = UserData{
			firstName:     firstName,
			lastName:      lastName,
			age:           age,
			martialStatus: martialStatus,
			employeed:     employeed,
			disable:       disable,
		}
		//appending structre into slice
		UsersDataSlice = append(UsersDataSlice, userData)

		var ch string
		fmt.Print("\nTo stop press y: ")
		fmt.Scan(&ch)
		if ch == "y" || ch == "Y" {
			displayData(UsersDataSlice)
			break
		} else {
			continue
		}
	}
}

func displayData(UserDataSlice []UserData) {
	fmt.Print("\n\n Users:", UserDataSlice)
}
