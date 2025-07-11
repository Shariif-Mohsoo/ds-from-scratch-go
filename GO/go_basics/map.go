package main

import "fmt"

func mapUsage() {
	//Creating a slice for map
	var userDataSlice []map[string]int
	//Creating an empty map list
	//var userDataSlice = make([]map[string]int, 0)

	//creating a map with string keys and value should be integer.
	var userData = make(map[string]int) //that will return an empty map
	userData["rollNo."] = 232202022
	userData["Semester"] = 5
	userData["Age"] = 22
	userDataSlice = append(userDataSlice, userData)

	userData["rollNo."] = 232202023
	userData["Semester"] = 5
	userData["Age"] = 22
	userDataSlice = append(userDataSlice, userData)

	//Displaying the data
	fmt.Println("\nRoll number:", userData["rollNo."], "\nSemester:", userData["Semester"], "\nAge:", userData["Age"])
	fmt.Println(userDataSlice)
}
