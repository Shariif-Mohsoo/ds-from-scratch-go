package main

import "fmt"

func nestedMap() {
	myMap := make(map[int]map[string]int)
	// inserting values
	/*
		In below line issue is that we are tryping to access the myMap[0] before i has been initialized.
		In Go, make(map[int]map[string]int) only creates the outer map. Each inner map (map[string]int)
		must also be initialized individually before assigning values to them.
		If you try to assign to a key in an uninitialized map (nil), Go will panic at runtime
		myMap[0]["Mohsin"] = 1
	*/
	// Initializing and inserting values
	myMap[0] = map[string]int{"Mohsin": 1}
	myMap[1] = make(map[string]int)
	myMap[1]["Danish"] = 2
	myMap[2] = map[string]int{"Muneeb": 3}
	//displaying the nested map values
	fmt.Println("Nested Map:", myMap)
}
