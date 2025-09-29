package structures

import (
	"fmt"
)

type Rectangle struct {
	length int
	width  int
}

func area(rect Rectangle) int {
	return rect.length * rect.width
}

func perimeter(rect Rectangle) int {
	return (rect.length * 2) + (rect.width * 2)
}

func printInfo(rect Rectangle) {
	fmt.Println("Area is:", area(rect))
	fmt.Println("Perimeter is:", perimeter(rect))
}

func Exercise() {
	var r1 = Rectangle{
		length: 10,
		width:  20,
	}
	fmt.Println(r1)
	printInfo(r1)
	r1.length *= 2
	r1.width *= 2
	fmt.Println(r1)
	printInfo(r1)
}
