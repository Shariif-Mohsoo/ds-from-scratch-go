/*
🧪 Problem: Simple Shape Area Calculator
Create a program that:

Defines two shapes (Circle, Rectangle)

# Implements an Area() method for each

# Uses an interface Shape

# Uses type assertion to check the concrete type

Returns an error if invalid dimensions are passed
*/
package main

import (
	"errors"
	"fmt"
	"math"
)

type Circle2 struct {
	Radius float64
}

type Rectangle2 struct {
	Width, Height float64
}

type Shape2 interface {
	Area() (float64, error)
}

// Method for cirle
func (c Circle2) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0, errors.New("Radius must be positive")
	}
	return math.Pi * c.Radius * c.Radius, nil
}

// Method for rectangle
func (r Rectangle2) Area() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0, errors.New("Width and Height must be positive")
	}
	return r.Height * r.Width, nil
}

func Error_Handling() {
	var s Shape2
	s = Circle2{Radius: 5}
	printArea(s)

	s = Rectangle2{Width: 4, Height: 5}
	printArea(s)

	s = Circle2{Radius: -2}
	printArea(s)
}

// Helper function with type assertion
func printArea(s Shape2) {
	area, err := s.Area()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("Area: %.2f \n", area)
	// rect,ok:=s.(Rectangle2)
	// cir,ok:=s.(Circle2)
	// if(!ok)
	//  Type assertion
	switch shape := s.(type) {
	case Circle2:
		fmt.Println("Shape is Circle with radius", shape.Radius)
	case Rectangle2:
		fmt.Println("Shape is Rectangle with width and height:", shape.Width, shape.Height)
	default:
		fmt.Println("Unknown Shape")
	}
}
