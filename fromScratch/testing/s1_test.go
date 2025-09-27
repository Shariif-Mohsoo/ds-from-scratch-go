package testing

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	data := "email@example.com"
	// fmt.Println(data)
	if !isValidEmail(data) {
		t.Errorf("isValidEmail(%v)=false,want true", data)
	}
	data = "mr@gmail.com"
	if !isValidEmail(data) {
		t.Errorf("isValidEmail(%v)=false,want true", data)
	}
}
