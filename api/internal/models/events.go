package models

import "time"

type Events struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Time        time.Time  `json:"time"`
	Resource    *Resources `json:"resource"`
}
