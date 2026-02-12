package main

import (
	"fmt"
)


func promptPassword() string {
	var password string
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)
	return password
}


func main() {
	secretPassword := "secret123"
	var password string
	for {
		password = promptPassword()
		if password == secretPassword {
			fmt.Println("Access Granted!")
			break
		} else {
			fmt.Println("Incorrect, try again.")
		}
	}
}