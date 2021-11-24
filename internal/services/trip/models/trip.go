package models

import "snakealive/m/internal/domain"

type Trip struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Sights      []domain.Place `json:"sights"`
}

type Album struct {
	Id          int      `json:"id"`
	TripId      int      `json:"trip_id"`
	UserId      int      `json:"user_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}
