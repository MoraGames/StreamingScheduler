package models

import (
	"database/sql"
	"errors"
)

type Resources struct {
	Id int               `json:"id"`
	Url string           `json:"url"`
	Language *Languages  `json:"language,omitempty"`
	Subtitles *Languages `json:"subtitles,omitempty"`
	Quality *Qualities   `json:"quality,omitempty"`
	Episode *Episodes    `json:"episode"`
}

func (r *Resources) Exist(dbConn *sql.DB) (bool, error) {
	var exists bool
	err := dbConn.QueryRow("SELECT exists (SELECT * FROM Resources WHERE id = ?)", r.Id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.New("error checking if row exists " + err.Error())
	}
	return exists, nil
}

// GetUserById is a function that gets user information from id
func GetResourceById(dbConn *sql.DB, id int64) (*Resources, error) {
	var resource Resources

	//TODO: The whole method!!!!!

	return &resource, nil
}