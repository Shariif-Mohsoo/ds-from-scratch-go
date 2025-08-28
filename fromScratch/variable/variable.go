package variable

import "fmt"

func Variable() {
	fmt.Println("Let's explore variables")
	var ex = 30
	var exa int = 40
	var exam int
	exam = 50
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

	n1, n2 := 2, 3
	fmt.Println("Sum:", n1+n2)
	n3, n2 := 4, 4
	fmt.Println("n2 + n2:", n2+n2, "\nn3:", n3, "\n")
}
