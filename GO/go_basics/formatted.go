package main

import "fmt"

func formattedData() {
	var person = "Moshin"
	var passion = "Coding"
	//default formatter is %v (placeholder)
	fmt.Printf("Person %v has the passion %v", person, passion)
}
