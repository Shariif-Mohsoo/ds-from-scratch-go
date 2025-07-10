package main

import "fmt"

func forLoop() {
	var count = 5
	for i := 0; i < count; i++ {

		fmt.Printf("\n%v th value is:%v", i, i+1)
	}
}
