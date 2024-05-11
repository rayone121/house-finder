package main

// reBank is a simple bank application that allows users to create accounts, transfer money between accounts, and view account balances.

import (
	"fmt"
	"log"

	"github.com/rayone121/reBank/backend/api"
	"github.com/rayone121/reBank/backend/storage"
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
