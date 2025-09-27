package testing

import "regexp"

func isValidEmail(addr string) bool {
	re, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("Failed to compile regex")
	} else {
		return re.Match([]byte(addr))
	}
}
