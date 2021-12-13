package models

type Sight struct {
	Id          int
	Name        string
	Country     string
	Rating      float32
	Lat         float32
	Lng         float32
	Tags        []string
	Description string
	Photos      []string
}

type Tag struct {
	Id   int
	Name string
}
