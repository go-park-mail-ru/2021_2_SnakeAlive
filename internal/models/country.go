package models

type Country struct {
	Id          int    `json:"id"`
	Translated  string `json:"translated"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}

//easyjson:json
type Countries []Country
