package iota

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

// receiver function.
func (op Operation) calculate(lhs, rhs int) int {
	var res int
	switch op {
	case Add:
		res = lhs + rhs
	case Subtract:
		res = lhs - rhs
	case Multiply:
		res = lhs * rhs
	case Divide:
		{
			if rhs != 0 {

				res = lhs / rhs
			} else {
				res = 0
			}
		}
	default:
		panic("Unhandle able case")
	}
	return res
}

func Exercise() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2))
	sub := Operation(Subtract)
	fmt.Println(sub.calculate(4, 2))
	mul := Operation(Multiply)
	fmt.Println(mul.calculate(4, 2))
	div := Operation(Divide)
	fmt.Println(div.calculate(6, 2))
}
