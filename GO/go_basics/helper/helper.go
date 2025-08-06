package helper

import "fmt"

func Hello() {
	fmt.Println("You are calling the helper function which is located in the helper folder")
	highOrderFunc()
}
func printNum(takeNum func() int) {
	fmt.Println("The number is:", takeNum())
}

func takeNum() int {
	x := 1
	x++
	return x
}

func highOrderFunc() {
	printNum(takeNum)
}
