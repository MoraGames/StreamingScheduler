package models

import (
	"database/sql"
	"log"
	"time"
)

type Events struct {
	Id int              `json:"id"`
	Title string        `json:"title"`
	Description string  `json:"description,omitempty"`
	Time time.Time      `json:"time"`
	Resource *Resources `json:"resource"`
}

func (e *Events) NewEvents(dbConn *sql.DB) (int64, error) {
	// prepare the insert query
	stmt, err := dbConn.Prepare("INSERT INTO Events (title, description, time, resource) VALUES (?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	if exist, err := e.Resource.Exist(dbConn); err == nil || exist == false {
		return -1, err
	}

	res, err := stmt.Exec(e.Title, e.Description, e.Time.Unix(), e.Resource.Id)
	if err != nil {
		return -1, err
	}

	// get the id of the new event
	return res.LastInsertId()
}

func GetEventById(dbConn *sql.DB, id int64) (*Events, error) {
	var event Events
	var timestamp int64
	var resourceId int64

	// Do the query
	err := dbConn.QueryRow(
		"SELECT id, title, description, time, resource FROM Events WHERE id = ?",
		id,
	).Scan(&event.Id, &event.Title, &event.Description, &timestamp, &resourceId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	event.Time = time.Unix(timestamp, 0)
	event.Resource, err = GetResourceById(dbConn, resourceId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &event, nil
}