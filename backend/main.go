package main

import "fmt"

func main() {
	server := newAPIServer(":3000")
	server.Start()
	fmt.Println("Hello, World!")
}
