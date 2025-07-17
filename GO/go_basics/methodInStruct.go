package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// creating a method for the struct
func (authI authenticationInfo) getBasicAuth() string {
	return "username:" + authI.username + "\n" + "password: " + authI.password
}

func display(authInfo authenticationInfo) {
	fmt.Println("====================")
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================")
}

func methodInStruct() {
	authIn := authenticationInfo{
		username: "username",
		password: "pass****d",
	}
	display(authIn)
}
