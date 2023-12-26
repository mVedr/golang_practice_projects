package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type TransferRequest struct {
	ToAccount int `json:"to_account"`
	Amount    int `json:"amount"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Number            int64     `json:"number"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
	EncryptedPassword string    `json:"-"`
}

func NewAccount(first_name string, last_name string, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:                rand.Intn(10000),
		FirstName:         first_name,
		LastName:          last_name,
		Number:            int64(rand.Intn(1000000)),
		CreatedAt:         time.Now().UTC(),
		EncryptedPassword: string(encpw),
	}, nil
}

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}
