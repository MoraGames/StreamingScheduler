package models

import "time"

type Event struct {
	Id          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Time        time.Time `json:"time"`
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

	qp, err := DbConn.Prepare(`INSERT INTO Events(title, description, time, resource) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(e.Title, e.Description, e.Time, e.Resource.Id)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetEventById is a function that gets the event from the database by id
func GetEventById(id int64) (*Event, error) {

	var event Event
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
		rows.Scan(&event.Id, &event.Title, &event.Description, &event.Time, &resourceId)
	}

	// populate resource
	event.Resource, err = GetResourceById(resourceId)
	if err != nil {
		return nil, err
	}

	return &event, nil
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

	qp, err := DbConn.Prepare(`SELECT * FROM Events`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event Event
		var resourceId int64

		rows.Scan(&event.Id, &event.Title, &event.Description, &event.Time, &resourceId)

		// populate resource
		event.Resource, err = GetResourceById(resourceId)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}
