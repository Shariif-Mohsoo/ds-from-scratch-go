package receiverfunctions

import "fmt"

type Coordinate struct {
	X, Y int
}

// regular function
func shiftBy(x, y int, coord *Coordinate) {
	coord.X += x
	coord.Y += y
}

// receiver function
func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}

func Receiverfunctions() {
	coord := Coordinate{5, 5}
	shiftBy(1, 1, &coord)
	fmt.Println(coord)
	coord.shiftBy(2, 2)
	fmt.Println(coord)
}
