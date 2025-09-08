package ranges

import "fmt"

func Ranges() {
	slice := []string{"Hello", "World", "!"}
	for i, element := range slice {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Printf(" %q\n", ch)
		}
	}

}
