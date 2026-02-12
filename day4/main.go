package main

import (
	"fmt"
)

func requestServer() string {
	var s string
	fmt.Print("Enter a new server name: ")
	fmt.Scan(&s)
	return s
}

func manageServers(server string) (servers []string) {
	staticServers := []string{"darkstar", "ai-core", "omni-1"}
	for _, s := range staticServers {
		servers = append(servers, s)
	}
	servers = append(servers, server)
	return 
}

func main() {
	server := requestServer()
	servers := manageServers(server)
	fmt.Print("Server List:\n")

	for i, s := range servers {
		fmt.Printf("[%d] %s is online!\n", i, s)
	}
}