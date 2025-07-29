package main

import "fmt"

func fizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			continue
			// fmt.Println(i)
		}
	}
}
