package models

type Sight struct {
	Id          int
	Name        string
	Country     string
	Rating      float32
	Tags        []string
	Description string
	Photos      []string
}
