package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

// In Go, if the only type specified paramaeter is the last one, then it will apply to the rest of the parameters
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(100000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		// The default value of Go is the zero value of the type, in int case is 0
	}
}
