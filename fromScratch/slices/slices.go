package slices

import "fmt"

func Slices() {
	creatingSlice()
	creatingSliceFromArray()
	dynamicArray()
	preAllocation()
	multiDimensionalSlice()
}

func creatingSlice() {
	mySlice := []int{1, 2, 3, 4}
	fmt.Println(mySlice)
}

func creatingSliceFromArray() {
	numbers := [...]int{1, 2, 3, 4}
	slice1 := numbers[:] // 1 2 3 4
	fmt.Println(slice1)
	slice2 := numbers[1:] // 2 3 4
	fmt.Println(slice2)
	slice3 := slice2[:1] //2
	fmt.Println(slice3)
	slice4 := numbers[:2] // 1 2
	fmt.Println(slice4)
	slice5 := numbers[1:3] // 2 3
	fmt.Println(slice5)
}

func dynamicArray() {
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5, 6)
	fmt.Println(nums)
}

func preAllocation() {
	slice := make([]float32, 10)
	fmt.Println(slice)
}

func multiDimensionalSlice() {
	board := [][]string{
		{"o", "o", "x"},
		{"x", "o", "x"},
		{"x", "x", "x"},
	}
	fmt.Println(board)
}
