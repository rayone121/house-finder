package main

import (
	"flag"
	"fmt"
)

func main() {
	listenAddr := flag.String("listenaddr", "3000", "the server address")

	server := newAPIServer(":3000")
	server.Start()
	fmt.Println("Hello, World!")
}
