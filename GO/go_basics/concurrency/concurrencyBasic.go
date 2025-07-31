package concurrency

import (
	"fmt"
	"time"
)

func printMsg(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func BasicsOfConcurrency() {
	go printMsg("Hello from Goroutine") // runs concurrently
	printMsg("Hello from Main")         // runs normally

	time.Sleep(2 * time.Second) // let goroutine finish before program ends
}
