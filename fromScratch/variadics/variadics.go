package variadics

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Variadic() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	all := append(a, b...)
	ans := sum(all...) // = sum(1,2,3,4,5,6)
	fmt.Println(ans)
}
