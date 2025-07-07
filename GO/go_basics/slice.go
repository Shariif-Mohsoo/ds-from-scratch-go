package main

import "fmt"

func sliceUsage() {
	//creating a slice
	var animals []string
	//assigning or inserting values in the slice
	animals = append(animals, "Fox")
	animals = append(animals, "Tiger")
	animals = append(animals, "Loin")

	//displaying the info
	fmt.Println("\nThe slice elements:", animals)
	fmt.Println("Slice first element:", animals[0])
	fmt.Printf("Slice type %T", animals)
	fmt.Println("\nThe length of slice is:", len(animals))
}
