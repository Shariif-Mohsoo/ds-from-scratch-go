package arrays

import "fmt"

func Arrays() {
	arr1()
	arr2()
	arr3()
	accessElements()
	iteration()
}

func arr1() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)
}

func accessElements() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
}

func arr2() {
	myArray := [...]int{7, 8, 9}
	fmt.Println(myArray)
}

func arr3() {
	myArray := [4]int{7, 8, 9}
	fmt.Println(myArray)
}

func iteration() {
	myArray := [...]int{7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}
