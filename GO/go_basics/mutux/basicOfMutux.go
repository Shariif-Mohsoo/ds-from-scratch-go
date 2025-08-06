package mutux

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var count int = 0

func increment(wg *sync.WaitGroup) {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func BasicOfMutux() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final Count:", count)
}
