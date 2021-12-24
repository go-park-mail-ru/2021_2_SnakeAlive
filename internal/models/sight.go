package models

type Sight struct {
	Description string `json:"description"`
	SightMetadata
}

type SightMetadata struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Photos      []string `json:"photos"`
	Country     string   `json:"country"`
	Rating      float32  `json:"rating"`
	Lat         float32  `json:"lat"`
	Lng         float32  `json:"lng"`
}

type SightSearch struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Lat    float32  `json:"lat"`
	Lng    float32  `json:"lng"`
	Photos []string `json:"photos"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SearchSights struct {
	Tags       []int64  `json:"tags"`
	Countries  []string `json:"countries"`
	Skip       int      `json:"skip"`
	Limit      int      `json:"limit"`
	Search     string   `json:"search"`
	MinRating  int      `json:"min_rating"`
	MinReviews int      `json:"min_amount_reviews"`
}
