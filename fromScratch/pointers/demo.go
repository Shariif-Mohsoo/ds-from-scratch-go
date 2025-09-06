package pointers

import "fmt"

type Counter struct {
	hits int
}

func increment1(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment1(counter)
}

func Demo() {
	counter := Counter{
		hits: 0,
	}
	hello := "Hello"
	world := "World!"

	fmt.Println(hello, world)
	fmt.Println(counter)
	increment1(&counter)
	fmt.Println(counter)

	replace(&hello, "Hi", &counter)
	fmt.Println(counter)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

}
