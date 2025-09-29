package interfaces

type MyInterface interface {
	Func1()
	Func2(x int) int
}

type MyType int

func (m MyType) Func1() {}
func (m MyType) Func2(x int) int {
	return x + x
}

// this func will receive any thing that implements MyInterface.
func execute(i MyInterface) {
	i.Func1()
}

func Interfaces() {
	m := MyType(1)
	execute(m)
	execute(&m)
}
