package types

import (
	"time"
)

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	UserName    string `json:"userName"`
	PhoneNumber int64  `json:"phoneNumber"`
}
type Account struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	UserName    string    `json:"userName"`
	Balance     int       `json:"balance"`
	PhoneNumber int64     `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName, userName string, phoneNumber int64) *Account {
	return &Account{
		FirstName:   firstName,
		LastName:    lastName,
		UserName:    userName,
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now().UTC(),
		Balance:     0,
	}
}
