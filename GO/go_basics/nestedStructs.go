package main

import "fmt"

type messageToSendUser struct {
	message   string
	sender    User
	recipient User
}

type User struct {
	name   string
	number int
}

func nestedStructs() {
	//PLAY WITH IT TODAY
	//creating an instance
	//default in function these structs are passed by value.
	messageData := messageToSendUser{}
	messageData = getData(messageData)
	//displayng the data
	fmt.Printf("\n%v\n", messageData)
	messageData1 := messageToSendUser{
		message: "Hi Dear!",
		sender: User{
			name:   "Mohsin",
			number: 320,
		},
		recipient: User{
			name:   "Jawad",
			number: 321,
		},
	}
	fmt.Printf("\n%v\n", messageData1)
}

func getData(ms messageToSendUser) messageToSendUser {
	ms.message = "Hello! How Are You?"
	ms.recipient.name = "Imran"
	ms.recipient.number = 3210
	ms.sender.name = "Hashim"
	ms.sender.number = 3220
	return ms
}
