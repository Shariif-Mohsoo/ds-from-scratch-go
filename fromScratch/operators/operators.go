package operators

import "fmt"

func Operators() {
	airthmaticOps()
	inc_decOps()
	relationOps()
	logicalOps()
}

func airthmaticOps() {
	// Airthmatic Operators
	var total = 3 + 3
	fmt.Println("Total:", total)
	a := 1
	a += 3
	fmt.Println("a=", a)
}

func inc_decOps() {
	//Increment/Decrement
	a := 5
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
}

func relationOps() {
	//Relational Ops
	a := 4
	b := 3
	res := a > b
	fmt.Println(res)
	res = a == b
	fmt.Println(res)
}

func logicalOps() {
	//Logical ops
	a := false
	b := true

	if a && b {
		fmt.Println(a)
	} else if a || b {
		fmt.Println(b)
	} else if a != b {
		fmt.Println(a, b, "not equal")
	} else {
		fmt.Println(a, b)
	}

}
