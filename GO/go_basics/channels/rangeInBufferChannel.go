package channels

import (
	"fmt"
	"time"
)

func RangeWithBufferChannel() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for item := range ch {
			fmt.Println(item)
		}
	}()

	time.Sleep(time.Second * 3)
}
