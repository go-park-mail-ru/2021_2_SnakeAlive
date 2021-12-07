package usecase

import (
	"context"

	"snakealive/m/internal/services/trip/models"
	"snakealive/m/internal/services/trip/repository"

	"github.com/microcosm-cc/bluemonday"
)

type TripUseCase interface {
	AddTrip(ctx context.Context, value *models.Trip, userID int) (int, error)
	GetTripById(ctx context.Context, id int) (*models.Trip, error)
	DeleteTrip(ctx context.Context, id int) error
	UpdateTrip(ctx context.Context, id int, updatedTrip *models.Trip) error

	CheckTripAuthor(ctx context.Context, userID int, id int) (bool, error)
	SanitizeTrip(ctx context.Context, trip *models.Trip) *models.Trip

	AddAlbum(ctx context.Context, album *models.Album, userID int) (int, error)
	GetAlbumById(ctx context.Context, id int) (*models.Album, error)
	DeleteAlbum(ctx context.Context, id int) error
	UpdateAlbum(ctx context.Context, id int, updatedAlbum *models.Album) error

	CheckAlbumAuthor(ctx context.Context, userID int, id int) (bool, error)
	SanitizeAlbum(ctx context.Context, album *models.Album) *models.Album

	SightsByTrip(ctx context.Context, id int) (*[]int, error)
	TripsByUser(ctx context.Context, id int) (*[]models.Trip, error)
	AlbumsByUser(ctx context.Context, id int) (*[]models.Album, error)

	AddTripUser(ctx context.Context, tripId int, userId int) error
}

type tripUseCase struct {
	tripRepository repository.TripRepository
}

func NewTripUseCase(tripRepository repository.TripRepository) TripUseCase {
	return &tripUseCase{tripRepository: tripRepository}
}

func (u tripUseCase) AddTrip(ctx context.Context, value *models.Trip, userID int) (int, error) {
	cleanTrip := u.SanitizeTrip(ctx, value)

	return u.tripRepository.AddTrip(ctx, cleanTrip, userID)
}

func (u tripUseCase) GetTripById(ctx context.Context, id int) (*models.Trip, error) {
	return u.tripRepository.GetTripById(ctx, id)
}

func (u tripUseCase) UpdateTrip(ctx context.Context, id int, updatedTrip *models.Trip) error {
	cleanTrip := u.SanitizeTrip(ctx, updatedTrip)

	return u.tripRepository.UpdateTrip(ctx, id, cleanTrip)
}

func (u tripUseCase) DeleteTrip(ctx context.Context, id int) error {
	return u.tripRepository.DeleteTrip(ctx, id)
}

func (u tripUseCase) CheckTripAuthor(ctx context.Context, userID int, id int) (bool, error) {
	authors, err := u.tripRepository.GetTripAuthors(ctx, id)
	for _, author := range authors {
		if author == userID {
			return true, err
		}
	}
	return false, err
}

func (u tripUseCase) SanitizeTrip(ctx context.Context, trip *models.Trip) *models.Trip {
	sanitizer := bluemonday.UGCPolicy()

	trip.Title = sanitizer.Sanitize(trip.Title)
	trip.Description = sanitizer.Sanitize(trip.Description)
	return trip
}

func (u tripUseCase) AddAlbum(ctx context.Context, album *models.Album, userID int) (int, error) {
	cleanAlbum := u.SanitizeAlbum(ctx, album)

	return u.tripRepository.AddAlbum(ctx, cleanAlbum, userID)
}

func (u tripUseCase) GetAlbumById(ctx context.Context, id int) (*models.Album, error) {
	return u.tripRepository.GetAlbumById(ctx, id)
}

func (u tripUseCase) DeleteAlbum(ctx context.Context, id int) error {
	return u.tripRepository.DeleteAlbum(ctx, id)
}

func (u tripUseCase) UpdateAlbum(ctx context.Context, id int, updatedAlbum *models.Album) error {
	cleanAlbum := u.SanitizeAlbum(ctx, updatedAlbum)
	return u.tripRepository.UpdateAlbum(ctx, id, cleanAlbum)
}

func (u tripUseCase) CheckAlbumAuthor(ctx context.Context, userID int, id int) (bool, error) {
	author, err := u.tripRepository.GetAlbumAuthor(ctx, id)
	return author == userID, err
}

func (u tripUseCase) SightsByTrip(ctx context.Context, id int) (*[]int, error) {
	return u.tripRepository.SightsByTrip(ctx, id)
}

func (u tripUseCase) TripsByUser(ctx context.Context, id int) (*[]models.Trip, error) {
	return u.tripRepository.GetTripsByUser(ctx, id)
}

func (u tripUseCase) SanitizeAlbum(ctx context.Context, album *models.Album) *models.Album {
	sanitizer := bluemonday.UGCPolicy()

	album.Title = sanitizer.Sanitize(album.Title)
	album.Description = sanitizer.Sanitize(album.Description)
	return album
}

func (u tripUseCase) AlbumsByUser(ctx context.Context, id int) (*[]models.Album, error) {
	return u.tripRepository.AlbumsByUser(ctx, id)
}

func (u tripUseCase) AddTripUser(ctx context.Context, tripId int, userId int) error {
	return u.tripRepository.AddTripUser(ctx, tripId, userId)
}
