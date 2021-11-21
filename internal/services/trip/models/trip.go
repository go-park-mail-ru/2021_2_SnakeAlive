package models

import "snakealive/m/internal/domain"

type Trip struct {
	Id          int
	Title       string
	Description string
	Days        [][]domain.Place
}
