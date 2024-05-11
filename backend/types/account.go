package types

import "math/rand"

type Account struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	UserName    string `json:"userName"`
	Balance     int    `json:"balance"`
	PhoneNumber int64  `json:"phoneNumber"`
}

func NewAccount(firstName, lastName, userName string) *Account {
	return &Account{
		ID:          rand.Intn(10000),
		FirstName:   firstName,
		LastName:    lastName,
		UserName:    userName,
		PhoneNumber: int64(rand.Intn(1000000)),
		Balance:     rand.Intn(100000),
	}
}
