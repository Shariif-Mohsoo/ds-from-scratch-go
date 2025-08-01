package channels

import (
	"fmt"
	"time"
)

func BasicOfChannels() {
	ch := make(chan string) //creating a channel
	go func() {
		time.Sleep(time.Second * 3)
		ch <- "Hello from goroutines!" // sending happens after 3 seconds
	}()

	msg := <-ch // waits (blocks) here until value is sent into channel
	fmt.Println("Main Received:", msg)

}
