package main

import (
	"database/sql"
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

	var exists bool
	err := dbConn.QueryRow("SELECT exists (SELECT * FROM Users WHERE email = ?)", u.Email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.New("error checking if row exists " + err.Error())
	}

	return exists, nil
}

func (u *User) NewUser() (int64, error) {

	// prepare the insert query
	stmt, err := dbConn.Prepare("INSERT INTO Users (email, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Email, u.Username, u.Password)
	if err != nil {
		return -1, err
	}

	// get the id of the new user
	return res.LastInsertId()
}
