package models

type Titles struct {
	Id       int        `json:"id"`
	Title    string     `json:"title"`
	Language *Languages `json:"language"`
}
