package main

import "fmt"

func twoXSlice() {
	var row int = 2
	var col int = 2

	matrix := make([][]int32, 0)
	for i := 1; i <= row; i++ {
		row := make([]int32, 0)
		for j := 1; j <= col; j++ {
			row = append(row, (int32)(i*j))
		}
		matrix = append(matrix, row)
	}
	fmt.Println(matrix)
}
