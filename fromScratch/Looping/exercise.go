package looping

import "fmt"

func Exercise() {
	for i := 1; i <= 50; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
