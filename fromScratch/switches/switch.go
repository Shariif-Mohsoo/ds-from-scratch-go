package switches

import "fmt"

func Switch() {
	switchBasic(2)
}

func switchBasic(y int) {
	x := y
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Other", x)
	}
}
