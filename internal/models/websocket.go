package models

import "snakealive/m/internal/services/trip/models"

type TripRequest struct {
	Message string `json:"message"`
	TripId  int    `json:"trip_id"`
}

type TripResponce struct {
	Message string      `json:"message"`
	Trip    models.Trip `json:"trip"`
}
