package operators

import "fmt"

func Demo() {
	quiz()
}

func average(a, b, c int) float32 {
	return float32(a+b+c) / 3
}

func quiz() {
	quiz1, quiz2, quiz3 := 8, 9, 7

	if quiz1 > quiz2 {
		fmt.Println("quiz1 scored higher than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 scored higher than quiz1")
	} else {
		fmt.Println("quiz1  & quiz2 have same score")
	}
	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable grades")
	} else {
		fmt.Println("Not Acceptable grades")
	}
}
