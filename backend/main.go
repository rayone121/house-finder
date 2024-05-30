package main

// reBank is a simple bank application that allows users to create accounts, transfer money between accounts, and view account balances.

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayone121/reBank/backend/api"
	"github.com/rayone121/reBank/backend/storage"
	"github.com/rayone121/reBank/backend/types"
)

func seedAccount(store storage.Storage, fname, lname, username, password string) *types.Account {
	acc, err := types.NewAccount(fname, lname, username, password)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	return acc
}

func seedAccounts(store storage.Storage) {
	seedAccount(store, "John", "Doe", "johndoe", "pa123123")
	seedAccount(store, "Jane", "Doe", "janedoe", "p12323rd")
	seedAccount(store, "Alice", "Smith ", "alices ", "pa23123rd")
}
func main() {
	seed := flag.Bool("seed", false, "seed the database")
	flag.Parse()
	store, err := storage.NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding database")
		seedAccounts(store)
	}

	server := api.NewAPIServer(":3000", store)
	server.Start()
	fmt.Println("Hello, World!")
}
