package main

import "fmt"

func anonymousStruct() {
	//creating an anonymous struct
	myCar := struct {
		Make  string
		Model uint
	}{
		Make:  "Tesla",
		Model: 2025,
	}

	fmt.Print(myCar)
	//OVERWRITTEN TO THE PREVIOUS.
	myCar.Make = "Tesla"
	myCar.Model = 2022
	fmt.Print(myCar)
}
