package generics

import "fmt"

func isEqual[T comparable](a, b T) bool {
	return a == b
}

type Integers32 interface {
	int32 | uint32
}

func SumNumbers[T Integers32](arr []T) T {
	var sum T
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Generics() {

	fmt.Println(isEqual('a', 'b'))
	fmt.Println(isEqual('a', 4))
	fmt.Println(isEqual(4, 4))

	fmt.Println([]int32{1, 2, 3}, "sum:", SumNumbers([]int32{1, 2, 3}))
	fmt.Println([]uint32{1, 2, 3}, "sum:", SumNumbers([]int32{1, 2, 3}))
}
