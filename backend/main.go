package main

import (
	"fmt"
	"log"

	"github.com/rayone121/house-finder/backend/api"
	"github.com/rayone121/house-finder/backend/storage"
)

func main() {
	store, err := storage.NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":3000", store)
	server.Start()
	fmt.Println("Hello, World!")
}
