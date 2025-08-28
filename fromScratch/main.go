package main

import "fmt"

func main() {

	var ex = 3
	var exa int = 3
	var exam int
	exam = 5
	fmt.Println(ex, exa, exam)

	// compound declaration
	var a, b, c = 1, 2, "sample"
	fmt.Println("Compound Declaration", a, b, c)
	//Block creation
	var (
		a1 int = 1
		b1 int = 2
		c1     = "sample"
	)
	fmt.Println("Block Declaration", a1, b1, c1)
	// Constants
	const AppAuthor = "Bob"
	fmt.Println("Const ", AppAuthor)
}
