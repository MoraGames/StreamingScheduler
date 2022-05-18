package main

import (
	"errors"
	"strings"
)

type User struct {
	// TODO: Define user

	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilePicture,omitempty"`
}

//IsValid verifica che i dati utente inviati dal client in fase di registrazione siano corretti.
func (u *User) IsValid() error {

	//Check email
	if u.Email == "" || strings.Contains(u.Email, "@") == false {
		return errors.New("Email not valid")
	}
	if len(strings.Replace(u.Email, "@", "", -1)) < 3 {
		return errors.New("Lenght of email not valid")
	}

	//Check Username
	if u.Username == "" {
		return errors.New("Username not valid")
	}

	return nil
}

func (u *User) Exist() (bool, error) {

	row, err := dbConn.Query(`SELECT * FROM Users WHERE email = ?`, u.Email)
	if err != nil {
		return false, err
	}

	count := 0
	for row.Next() {
		count++
	}

	// If there isn't a user
	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (u *User) NewUser() (*User, error) {

	// TODO: Add new user to database

	return nil, nil
}
