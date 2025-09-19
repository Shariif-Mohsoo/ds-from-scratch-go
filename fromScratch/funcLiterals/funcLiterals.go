package funcliterals

import (
	"fmt"
	"strings"
)

func customMsg(fn func(m string), msg string) {
	msg = strings.ToUpper(msg)
	fn(msg)
}

func surround() func(msg string) {
	return func(msg string) {
		fmt.Printf("%.*s\n", len(msg), "---------------")
		fmt.Println(msg)
		fmt.Printf("%.*s\n", len(msg), "---------------")
	}
}

func FuncLiterals() {
	customMsg(surround(), "hello")
}
