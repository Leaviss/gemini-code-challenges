package main

import (
	"fmt"
	"strings"
)

func check_access() {
	var role string
	fmt.Print("Enter your role: ")
	fmt.Scan(&role)
	
	switch strings.ToLower(role) {
	case "admin":
		fmt.Print("Full Access Granted\n")
	case "manager":
		fmt.Print("Limited Access Granted\n")
	default:
		fmt.Print("Access Denied\n")
	}
}

func main() {
	check_access()
}