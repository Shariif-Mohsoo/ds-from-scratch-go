package main

import "fmt"

func syntaticSugar() {
	var conferenceName string = "Go Conference"

	fmt.Println("Conference Name:", conferenceName)

	//syntatic sugar (a term to describe the feature in language that let you do something more easily)
	//var conferenceNme string := "Go Conference" // will throw an error
	conference := "Go Confrerence"
	fmt.Println("After applying syntatic sugar")
	fmt.Println("Go Conference:", conference)
}
