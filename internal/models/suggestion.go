package models

type Suggestion struct {
	Text      string `json:"text"`
	Frequency int    `json:"frequency"`
}
