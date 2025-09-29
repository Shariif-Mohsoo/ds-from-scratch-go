package structures

import (
	"fmt"
)

// Creating a struct
type Sample struct {
	field string
	a, b  int
}

func Structures() {
	//Initializing a struct
	d1 := Sample{"word", 1, 2}
	fmt.Println(d1)
	d2 := Sample{
		field: "word",
		a:     1,
		b:     2,
	}
	fmt.Println(d2)
	//Default struct
	d3 := Sample{}
	fmt.Println(d3)

	//Accessing fields
	word, a, b := d1.field, d1.a, d1.b
	fmt.Println(word, a, b)
	d1.a = 10
	d1.b = 20
	fmt.Println(d1.a, d1.b)

	anonymousStruct()
}

func anonymousStruct() {
	var sample struct {
		field string
		a, b  int
	}
	sample.field = "word1"
	sample.a = 22
	sample.b = 25
	fmt.Println("Anonymous struct:", sample)
}
