package main

import "fmt"

func counter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func closure() {
	next := counter()
	fmt.Print("\n\n\n")
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
