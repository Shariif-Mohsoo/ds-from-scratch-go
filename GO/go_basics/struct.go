package main

import "fmt"

func structUsage() {
	//defining the struct.
	type personData struct {
		firstName     string
		lastName      string
		age           int
		martialStatus string
		employeed     bool
	}
	//declaring the slice for above struct
	var PersonsData = make([]personData, 0)
	//inserting data in sttruct
	// 1
	var dataAboutPerson = personData{
		firstName:     "Imran",
		lastName:      "Ali",
		age:           25,
		martialStatus: "Single",
		employeed:     false,
	}
	PersonsData = append(PersonsData, dataAboutPerson)
	// 2
	dataAboutPerson = personData{
		firstName:     "Ali",
		lastName:      "Hassan",
		age:           28,
		martialStatus: "Married",
		employeed:     true,
	}
	PersonsData = append(PersonsData, dataAboutPerson)
	//display the data
	fmt.Println("\n", dataAboutPerson)
	//display the data from slice
	fmt.Println(PersonsData)
}
