package models

import "snakealive/m/internal/domain"

type Trip struct {
	Id          int
	Title       string
	Description string
	Sights      []domain.Place
}
