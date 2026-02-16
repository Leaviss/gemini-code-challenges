package main

import (
	"fmt"
)

func manageInventory() {
	servers := make(map[string]string)

	servers["darkstar"] = "192.168.1.1"

	fmt.Print("Enter server name: ")
	var s string
	fmt.Scan(&s)

	ip, exists := servers[s]

	if exists {
		fmt.Printf("Server IP for %s is %s\n", s, ip)
	} else {
		fmt.Println("Server not found.")
		var newServer string
		fmt.Printf("Enter IP for %s: ", s)
		fmt.Scan(&newServer)
		servers[s] = newServer
		fmt.Printf("Added: %s -> %s\n", s, newServer)
	}

	fmt.Println("\nCurrent Inventory:")
	for server, ip := range servers {
		fmt.Printf("%s: %s\n", server, ip)
	}
}



func main() {
	manageInventory()
}