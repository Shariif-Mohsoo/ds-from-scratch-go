package channels

import "fmt"

func SelectStatementWithChannels() {
	chInt := make(chan int)
	chStr := make(chan string)

	go func() { chInt <- 2 }()
	go func() { chStr <- "Select Channel" }()

	select {
	case myInt := <-chInt:
		fmt.Println(myInt)
	case myStr := <-chStr:
		fmt.Println(myStr)
	}
}
