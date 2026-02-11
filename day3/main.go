package main

import (
	"fmt"
)


func prompt_password() string {
	var password string
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)
	return password
}


func main() {
	secret_password := "secret123"
	var password string
	for {
		password = prompt_password()
		if password == secret_password {
			fmt.Println("Access Granted!")
			break
		} else {
			fmt.Println("Incorrect, try again.")
		}
	}
}