package main

import (
	"database/sql"
	"errors"
	"github.com/MoraGames/StreamingScheduler/auth/internal/jwt"
	"strings"
)

type User struct {
	// TODO: Define user

	Id             int64          `json:"-"`
	Email          string         `json:"email"`
	Username       string         `json:"username"`
	Password       string         `json:"password"`
	ProfilePicture string         `json:"profilePicture,omitempty"`
	Enabled        bool           `json:"-"`
	Permissions    jwt.Permission `json:"permissions"`
}

//NewUser is a user method that creates the new user in the database
func (u *User) NewUser() (int64, error) {

	// TODO: Aggiungi offuscamento password

	// prepare the insert query
	stmt, err := dbConn.Prepare("INSERT INTO Users (email, username, password, permissions) VALUES (?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Email, u.Username, u.Password, u.Permissions)
	if err != nil {
		return -1, err
	}

	// get the id of the new user
	return res.LastInsertId()
}

//IsValid is a function that check the user info validity
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

// Exist is a function that checks if the user already exist
func (u *User) Exist() (bool, error) {

	var exists bool
	err := dbConn.QueryRow("SELECT exists (SELECT * FROM Users WHERE email = ?)", u.Email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.New("error checking if row exists " + err.Error())
	}

	return exists, nil
}

// Active is a function that sets the user to enable
func (u *User) Active() error {

	_, err := dbConn.Query(`UPDATE Users SET enabled=true WHERE id = ?`, u.Id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail is a function that gets user information from email
func GetUserByEmail(email string) (*User, error) {

	var user User
	var perm string
	var profileImg sql.NullString

	// Do the query
	err := dbConn.QueryRow(
		"SELECT id, username, email, password, profilePicture, enabled, permissions FROM Users WHERE email = ?",
		email,
	).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &profileImg, &user.Enabled, &perm)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	user.Permissions = jwt.Permission(perm)
	user.ProfilePicture = profileImg.String
	return &user, nil
}

// GetUserById is a function that gets user information from id
func GetUserById(id int64) (*User, error) {

	var user User
	var perm string
	var profileImg sql.NullString

	// Do the query
	err := dbConn.QueryRow(
		"SELECT id, username, email, password, profilePicture, enabled, permissions FROM Users WHERE id = ?",
		id,
	).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &profileImg, &user.Enabled, &perm)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	user.Permissions = jwt.Permission(perm)
	user.ProfilePicture = profileImg.String
	return &user, nil
}
