package main

import "fmt"

func print(name string) {
	fmt.Println("Name:", name)
}

func callMe(name string, action func(string)) {
	action(name)
}

func highOrderFunc1() {
	callMe("Mohsin", print)
}
