package models

type Review struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Rating    int    `json:"rating"`
	UserId    int    `json:"user_id"`
	PlaceId   int    `json:"place_id"`
	CreatedAt string `json:"created_at"`
}
