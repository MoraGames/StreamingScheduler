package models

type Qualities struct {
	Id         int    `json:"id"`
	Quality    string `json:"quality"`
	Resolution string `json:"resolution,omitempty"`
}
