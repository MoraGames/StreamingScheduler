package models

type Languages struct {
	Id int `json:"id"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Name string `json:"name"`
}