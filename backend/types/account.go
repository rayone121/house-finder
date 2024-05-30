package types

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	UserName          string    `json:"userName"`
	EncryptedPassword string    `json:"_"`
	Balance           int64     `json:"balance"`
	Number            int64     `json:"number"`
	CreatedAt         time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName, userName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		UserName:          userName,
		EncryptedPassword: string(encpw),
		Number:            int64(rand.Intn(100000)),
		CreatedAt:         time.Now().UTC(),
		Balance:           0,
	}, nil
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}
