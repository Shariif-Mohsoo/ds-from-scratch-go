package datetest

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
)

// USING THE GLOBAL PACKAGE.
func DateTest() {
	tt := tinytime.New(1585750374)
	tt = tt.Add(time.Hour * 40)
	fmt.Println(tt)
}
