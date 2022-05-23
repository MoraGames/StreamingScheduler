package models

type Series struct {
	Id int                `json:"id"`
	OriginalTitle *Titles `json:"originalTitle"`
	Plot string           `json:"plot,omitempty"`
	Format *Formats       `json:"format"`
	Favorites int         `json:"favorites"`
	AgeRestriction uint   `json:"ageRestriction,omitempty"`
}