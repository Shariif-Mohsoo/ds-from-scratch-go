package switches

import "fmt"

func Exercise() {
	giveAge(10)
	giveAge(18)
}

func giveAge(age int) {
	switch {
	case age == 0:
		fmt.Println("Newborn")
	case age >= 1 && age <= 3:
		fmt.Println("Toddler")
	case age >= 4 && age <= 12:
		fmt.Println("child")
	case age >= 13 && age <= 17:
		fmt.Println("teenager")
	default:
		fmt.Println("Adult")
	}
}
