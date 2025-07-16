package main

import "fmt"

// creating a struct

type messageToSend struct {
	phoneNo int
	message string
}

func structTest() {
	test(messageToSend{
		phoneNo: 32011,
		message: "Hi Mohsoo!",
	})
	//creating a new object
	newData := messageToSend{
		phoneNo: 31121,
		message: "Hi Dear!",
	}

	test(newData)

	newData = messageToSend{}
	newData.phoneNo = 3222
	newData.message = "Go to hell!"
	test(newData)

}
func test(m messageToSend) {
	fmt.Printf("\nMessage:%v,from this number:0%v\n", m.message, m.phoneNo)
}
