package main

import (
	"database/sql"
	"errors"
	"strings"
)

type User struct {
	// TODO: Define user

	Id             int64  `json:"-"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilePicture,omitempty"`
	Enabled        bool   `json:"-"`
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

func (u *User) Active() error {

	_, err := dbConn.Query(`UPDATE Users SET enabled=true WHERE id = ?`, u.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {

	var user User

	// prepare the insert query
	stmt, err := dbConn.Prepare("SELECT * FROM Users WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.ProfilePicture, &user.Enabled)
	}

	return &user, nil
}

func GetUserById(id int64) (*User, error) {

	var user User

	// prepare the insert query
	stmt, err := dbConn.Prepare("SELECT * FROM Users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.ProfilePicture, &user.Enabled)
	}

	return &user, nil
}
