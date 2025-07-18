package main

import "fmt"

// 1. Define interface
type Speaker interface {
	Speak()
}

// 2. Define a struct that will implement the interface
type Person struct {
	Name string
}

// 3. Implement the Speak method for Person
func (p Person) Speak() {
	fmt.Println("Hi, my name is", p.Name)
}

// 4. Another type implementing the same interface
type Robot struct {
	ID int
}

func (r Robot) Speak() {
	fmt.Println("Beep! I am robot", r.ID)
}

// Main function
func interfaceBasic() {
	var s Speaker

	s = Person{Name: "Ali"}
	s.Speak()

	s = Robot{ID: 101}
	s.Speak()
}
