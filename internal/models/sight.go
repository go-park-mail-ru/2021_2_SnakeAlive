package models

type Sight struct {
	Description string `json:"description"`
	SightMetadata
}

type SightMetadata struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Tags    []string `json:"tags"`
	Photos  []string `json:"photos"`
	Country string   `json:"country"`
	Rating  float32  `json:"rating"`
}

type SightSearch struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Lat  float32 `json:"lat"`
	Lng  float32 `json:"lng"`
}
