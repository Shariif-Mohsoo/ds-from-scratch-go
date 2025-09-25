package goroutines

import (
	"fmt"
	"time"
)

func count(amount int) {
	for i := 1; i <= amount; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func GoRoutines() {
	go count(5) //it will run in it's own world.
	fmt.Println("wait for goroutine")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("end program")

	// closure
	counter := 0
	wait := func(ms time.Duration) {
		time.Sleep(ms * time.Millisecond)
		counter += 1
	}

	fmt.Println("Launching go routines")
	go wait(100)
	go wait(900)
	go wait(1000)

	fmt.Println("Launched. Counter =", counter)
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("Waited 1100 ms. Counter =", counter)

}
