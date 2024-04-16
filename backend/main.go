package main

import (
	"fmt"

	"github.com/rayone121/house-finder/backend/api"
)

func main() {
	server := api.NewAPIServer(":3000")
	server.Start()
	fmt.Println("Hello, World!")
}
