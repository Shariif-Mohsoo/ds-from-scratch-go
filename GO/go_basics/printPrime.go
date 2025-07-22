package main

import "fmt"

func printPrime() {
	var num int
	fmt.Println("Enter number:")
	fmt.Scan(&num)
	for i := 2; i < num; i++ {
		if num == 2 {
			fmt.Println("number is prime", num)
			return
		}
		if num%i == 0 {
			fmt.Println("Number is not prime", num)
			return
		}
	}

	fmt.Println("Number is prime", num)
}
