package mutux

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func WaitGroupInMutux() {
	wg.Add(2)
	go func() {
		fmt.Println("Hi Dear!")
		wg.Done()
	}()

	go func() {
		fmt.Println("How are you?")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Done with waitgroup")
}
