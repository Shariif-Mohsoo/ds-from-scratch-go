package main

import "fmt"

func sameLineVariableDeclaraion() {
	firstName, lastName, age, gender, isMarried := "Ali", "Ahmad", 22, "Male", false
	fmt.Printf("firstName:%v,lastName:%v,Age:%v,gender:%v,isMarried:%v", firstName, lastName, age, gender, isMarried)
}
