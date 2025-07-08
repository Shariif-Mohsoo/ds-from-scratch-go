package main

import "fmt"

func switchUsage() {
	city := ""
	fmt.Print("\nEnter city name: ")
	fmt.Scan(&city)

	switch city {
	case "Kahuta":
		{
			fmt.Println("You living in kahuta")
			break
		}
	case "RWP":
		{
			fmt.Println("You are living in rwp")
			break
		}
	case "Islamabad":
		{
			fmt.Println("You are living in Islamabad")
			break
		}
	default:
		{
			fmt.Println("Something is invalid dear")
		}
	}
}
