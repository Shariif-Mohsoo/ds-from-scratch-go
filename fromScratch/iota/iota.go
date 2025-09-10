package iota

import "fmt"

func first() {
	const (
		L0 = iota
		L1 = iota
		L2 = iota
		L3 = iota
		L4 = iota
	)
	fmt.Println(L0, L1, L2, L3, L4)
}

func second() {
	const (
		L0 = iota
		L1
		L2
		L3
		L4
	)
	fmt.Println(L0, L1, L2, L3, L4)
}

func third() {
	const (
		L0 = iota + 3
		L1
		L2
		L3
		L4
	)
	fmt.Println(L0, L1, L2, L3, L4)
}

func Fourth() {
	const (
		L0 = iota + 1
		_
		_
		L3
		L4
	)
	fmt.Println(L0, L3, L4)

}

func IOTA() {
	first()
	second()
	third()
	Fourth()
}
