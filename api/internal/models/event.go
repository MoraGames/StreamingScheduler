package models

import (
	"time"
)

type Event struct {
	Id          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Resource    *Resource `json:"resource"`
}

// NewEvent is a series method that adds a new event into database
func (e *Event) NewEvent() (int64, error) {

	// Check original title exist
	exist, err := e.Resource.Exist()
	if err != nil {
		return -1, err
	}

	if !exist {
		e.Resource.Id = 0
		resource, err := e.Resource.NewResource()
		if err != nil {
			return -1, err
		}
		e.Resource.Id = resource
	}

	qp, err := DbConn.Prepare(`INSERT INTO Events(title, description, startTime, endTime, resource) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(e.Title, e.Description, e.StartTime, e.EndTime, e.Resource.Id)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetEventById is a function that gets the event from the database by id
func GetEventById(id int64) (*Event, error) {

	var e Event
	var resourceId int64

	qp, err := DbConn.Prepare(`SELECT * FROM Events WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&e.Id, &e.Title, &e.Description, &e.StartTime, &e.EndTime, &resourceId)
	}

	// populate resource
	e.Resource, err = GetResourceById(resourceId)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// Exist Check if the event exist
func (e *Event) Exist() (bool, error) {
	event, err := GetEventById(e.Id)
	if err != nil {
		return false, err
	}

	if event.Id == 0 {
		return false, nil
	}

	return true, nil
}

// GetEvents function that gets all events
func GetEvents() (events []*Event, err error) {

	rows, err := DbConn.Query(`SELECT * FROM Events`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event Event
		var resourceId int64

		rows.Scan(&event.Id, &event.Title, &event.Description, &event.StartTime, &resourceId, &event.EndTime)

		// populate resource
		event.Resource, err = GetResourceById(resourceId)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}

// DeleteEvent is a function that deletes the event from the database by id
func DeleteEvent(id int64) (error) {

	qp, err := DbConn.Prepare(`DELETE FROM Events WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = qp.Query(id)
	if err != nil {
		return err
	}

	return nil
}
