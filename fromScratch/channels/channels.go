package channels

import (
	"fmt"
	"time"
)

func unBuffChannel() {
	channel := make(chan int)

	// send to channel
	go func() { channel <- 1 }()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	//receive from channel
	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}

func buffChannel() {
	channel := make(chan int, 2)

	channel <- 1
	channel <- 2

	go func() { channel <- 3 }()
	go func() { channel <- 4 }()

	//receive from channel
	first := <-channel
	second := <-channel
	third := <-channel
	fourth := <-channel

	fmt.Println(first, second, third, fourth)
}

func selectChannel() {
	one := make(chan int)
	two := make(chan int)

	for {
		select {
		case o := <-one:
			fmt.Println("one:", o)
		case t := <-two:
			fmt.Println("two:", t)
		default:
			fmt.Println("no data to receive")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Channels() {
	unBuffChannel()
	buffChannel()
	selectChannel()
}
