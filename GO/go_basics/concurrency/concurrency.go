package concurrency

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("\nEmail Received '%s' \n", message)
	}()
	fmt.Printf("\nEmail sent: '%s' \n", message)

}

func TestConcurrency() {

	sendEmail("Hi dear!")
	sendEmail("Hey dear!")
	sendEmail("What'up dear?")
	sendEmail("How are you dear?")
	sendEmail("I am anonymous dear!")

	// Let all goroutines finish before main exits
	time.Sleep(1 * time.Second)

}
