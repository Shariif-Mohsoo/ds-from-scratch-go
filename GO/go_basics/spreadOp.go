package main

import "fmt"

func speardOp() {
	allNames := []string{"Mohsin", "Jawad", "Saad", "Zain", "Gohar"}
	printNames(allNames...)
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}
