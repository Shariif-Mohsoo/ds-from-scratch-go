package main

import "fmt"

// an interface just contains the functions signature.
type Shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2 * r.height * 2 * r.width
}

func interFace() {
	var shape Shape
	shape = rect{
		height: 20.22,
		width:  10.44,
	}
	fmt.Println("Area:", shape.area())
	fmt.Println("Perimeter:", shape.perimeter())
}
