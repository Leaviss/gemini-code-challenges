package main

import (
	"fmt"
)

func promptUser() {
	var noun string
	var verb string
	var adj string

	fmt.Print("Provide a noun: ")
	fmt.Scan(&noun)

	fmt.Print("Provide a verb: ")
	fmt.Scan(&verb)

	fmt.Print("Provide a adj: ")
	fmt.Scan(&adj)

	fmt.Printf("The %s %s likes to %s the cloud server.\n", adj, noun, verb)
}

func main() {
	fmt.Println("I need your help telling a story.")
	promptUser()
}