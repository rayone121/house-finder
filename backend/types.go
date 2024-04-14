package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Number    int64
}

func NewAccount(firstName, lastName, userName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Number:    int64(rand.Intn(1000000)),
	}
}
