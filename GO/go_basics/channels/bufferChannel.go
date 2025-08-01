package channels

import "fmt"

func BufferChannel() {
	ch := make(chan string, 3)
	ch <- "HI"
	ch <- "Dear"
	ch <- "!"
	fmt.Println(<-ch, <-ch, <-ch)
}
