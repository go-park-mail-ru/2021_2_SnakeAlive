package repository

import (
	"context"
	"snakealive/m/internal/services/trip/models"
)

type TripRepository interface {
	AddTrip(ctx context.Context, value *models.Trip, userID int) (int, error)
	GetTripById(ctx context.Context, id int) (value *models.Trip, err error)
	DeleteTrip(ctx context.Context, id int) error
	UpdateTrip(ctx context.Context, id int, value *models.Trip) error
	GetTripAuthor(ctx context.Context, id int) (int, error)

	AddAlbum(ctx context.Context, album *models.Album, userID int) (int, error)
	GetAlbumById(ctx context.Context, id int) (*models.Album, error)
	DeleteAlbum(ctx context.Context, id int) error
	UpdateAlbum(ctx context.Context, id int, album *models.Album) error
	GetAlbumAuthor(ctx context.Context, id int) (int, error)

	SightsByTrip(ctx context.Context, id int) (*[]int, error)
}
