package models

type Episodes struct {
	Id int `json:"id"`
	Series *Series `json:"series"`
	Number int `json:"number,omitempty"`
	OriginalTitle *Titles `json:"originalTitle,omitempty"`
	Plot string `json:"plot,omitempty"`
	restrictedAge int `json:"restrictedAge,omitempty"`
}