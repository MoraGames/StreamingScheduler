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

func (u *User) Exist() (bool, error){

	//TODO: Add exist check in the database

	return false, nil
}

func (u *User) NewUser() (*User, error) {

	// TODO: Add new user to database

	return nil, nil
}