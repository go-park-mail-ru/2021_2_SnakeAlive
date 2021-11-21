package usecase

import (
	"context"
	"snakealive/m/internal/services/trip/models"
	"snakealive/m/internal/services/trip/repository"

	"github.com/microcosm-cc/bluemonday"
)

type TripUseCase interface {
	Add(ctx context.Context, value *models.Trip, userID int) (int, error)
	GetById(ctx context.Context, id int) (*models.Trip, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, updatedTrip *models.Trip) error

	CheckAuthor(ctx context.Context, userID int, id int) (bool, error)
	SanitizeTrip(ctx context.Context, trip *models.Trip) *models.Trip
}

type tripUseCase struct {
	tripRepository repository.TripRepository
}

func NewTripUseCase(tripRepository repository.TripRepository) TripUseCase {
	return &tripUseCase{tripRepository: tripRepository}
}

func (u tripUseCase) Add(ctx context.Context, value *models.Trip, userID int) (int, error) {
	cleanTrip := u.SanitizeTrip(ctx, value)

	return u.tripRepository.Add(ctx, cleanTrip, userID)
}

func (u tripUseCase) GetById(ctx context.Context, id int) (*models.Trip, error) {
	return u.tripRepository.GetById(ctx, id)
}

func (u tripUseCase) Update(ctx context.Context, id int, updatedTrip *models.Trip) error {
	cleanTrip := u.SanitizeTrip(ctx, updatedTrip)

	return u.tripRepository.Update(ctx, id, cleanTrip)
}

func (u tripUseCase) Delete(ctx context.Context, id int) error {
	return u.tripRepository.Delete(ctx, id)
}

func (u tripUseCase) CheckAuthor(ctx context.Context, userID int, id int) (bool, error) {
	author, err := u.tripRepository.GetTripAuthor(ctx, id)
	return author == userID, err
}

func (u tripUseCase) SanitizeTrip(ctx context.Context, trip *models.Trip) *models.Trip {
	sanitizer := bluemonday.UGCPolicy()

	trip.Title = sanitizer.Sanitize(trip.Title)
	trip.Description = sanitizer.Sanitize(trip.Description)
	return trip
}
