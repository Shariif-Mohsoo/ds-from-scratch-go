package main

import (
	"fmt"
	"math"
)

// 1. Struct
type Circle struct {
	Radius float64
}

// 2. Struct method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 3. Interface
type Shape1 interface {
	Area() float64
}

func describeShape1(s Shape1) {
	// Call method via interface
	fmt.Println("Area from interface:", s.Area())

	// 4. Type assertion
	c, ok := s.(Circle)
	if ok {
		fmt.Println("Type assertion successful! Radius is:", c.Radius)
	} else {
		fmt.Println("Type assertion failed")
	}
}

func myMain1() {
	// Create Circle struct
	myCircle := Circle{Radius: 5}

	// Pass to function that accepts interface
	describeShape1(myCircle)
}
