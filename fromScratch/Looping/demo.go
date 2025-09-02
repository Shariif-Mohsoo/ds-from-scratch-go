package looping

import "fmt"

func Demo() {
	sum := 0
	fmt.Println("Sum is", sum)
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("INC Sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("DEC Sum is", sum)
	}

}
