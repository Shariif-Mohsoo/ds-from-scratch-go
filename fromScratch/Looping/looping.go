package looping

import "fmt"

func Looping() {
	forBasic(10)
	forWhile(5)
	infiniteLoop()
}

func forBasic(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}

func forWhile(count int) {
	i := 0
	for i < count {
		fmt.Println(i)
		i++
	}
}

func infiniteLoop() {
	x := -10
	for {
		if x == 10 {
			break
		}
		fmt.Println(x)
		x++
	}
}
