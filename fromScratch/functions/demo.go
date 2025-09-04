package functions

import "fmt"

func Demo() {
	fmt.Println("Double 5:", double(5))
	greet()
	fmt.Println("Mul 5*4:", mul(5, 4))
	fmt.Println("double(5)*double(4):", mul(double(5), double(4)))

}

func greet() {
	fmt.Println("Hello from greet function")
}

func double(x int) int {
	return x + x
}

func mul(n1, n2 int) int {
	return n1 * n2
}
