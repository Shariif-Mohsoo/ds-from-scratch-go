package pointers

import "fmt"

func increment(x *int) {
	*x += 1
}

func Pointers() {
	val := 10
	var valPtr *int
	valPtr = &val
	fmt.Println(valPtr, "=", val)
	increment(&val)
	fmt.Println(valPtr, "=", val)

}
