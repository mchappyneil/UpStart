package main

import "errors"

type user struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	PasswordSalt []byte `json:"password_salt"`
	PasswordHash []byte `json:"password_hash"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

func registerNewUser(email, password string) (*user, error) {
	return nil, errors.New("placeholder")
}
