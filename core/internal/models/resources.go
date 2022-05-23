package models

type Resources struct {
	Id int               `json:"id"`
	Url string           `json:"url"`
	Language *Languages  `json:"language,omitempty"`
	Subtitles *Languages `json:"subtitles,omitempty"`
	Quality *Qualities   `json:"quality,omitempty"`
	Episode *Episodes    `json:"episode"`
}