package helper

import "fmt"

func returnNum() func() int {
	num := 10
	return (func() int {
		num++
		return num
	})
}

func MyFunc() {
	next := returnNum()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
