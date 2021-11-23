package repository

import (
	"context"
	"snakealive/m/internal/domain"
	"snakealive/m/internal/services/trip/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TripRepository interface {
	Add(ctx context.Context, value *models.Trip, userID int) (int, error)
	GetById(ctx context.Context, id int) (value *models.Trip, err error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, value *models.Trip) error
	GetTripAuthor(ctx context.Context, id int) (int, error)
}

type tripRepository struct {
	dataHolder *pgxpool.Pool
}

func NewTripRepository(DB *pgxpool.Pool) TripRepository {
	return &tripRepository{dataHolder: DB}
}

func (t *tripRepository) Add(ctx context.Context, value *models.Trip, userID int) (int, error) {
	var tripId int

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return tripId, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		AddTripQuery,
		value.Title,
		value.Description,
		userID,
		value.Sights[0].Id,
	).Scan(&tripId)
	if err != nil {
		return tripId, err
	}

	err = t.addPlaces(ctx, value.Sights, tripId)
	return tripId, err
}

func (t *tripRepository) GetById(ctx context.Context, id int) (*models.Trip, error) {
	var trip models.Trip

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return &trip, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		GetTripQuery,
		id,
	).Scan(&trip.Id, &trip.Title, &trip.Description)
	if err != nil {
		return &trip, err
	}

	rows, err := conn.Query(context.Background(),
		GetPlaceForTripQuery,
		id)
	if err != nil {
		return &trip, err
	}

	var place domain.Place
	for rows.Next() {
		rows.Scan(&place.Id, &place.Name, &place.Tags, &place.Description, &place.Rating, &place.Country, &place.Photos, &place.Day)
		trip.Sights = append(trip.Sights, place)
	}

	return &trip, err
}

func (t *tripRepository) Update(ctx context.Context, id int, value *models.Trip) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		UpdateTripQuery,
		value.Title,
		value.Description,
		value.Sights[0].Id,
		id,
	)
	if err != nil || t.deletePlaces(ctx, id) != nil {
		return err
	}

	err = t.addPlaces(ctx, value.Sights, id)
	return err
}

func (t *tripRepository) Delete(ctx context.Context, id int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	if err = t.deletePlaces(ctx, id); err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(),
		DeleteTripQuery,
		id,
	)
	return err
}

func (t *tripRepository) deletePlaces(ctx context.Context, tripId int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		DeletePlacesForTripQuery,
		tripId,
	)
	return err
}

func (t *tripRepository) addPlaces(ctx context.Context, value []domain.Place, tripId int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	for ind, place := range value {
		_, err = conn.Exec(context.Background(),
			AddPlacesForTripQuery,
			tripId,
			place.Id,
			place.Day,
			ind,
		)
	}
	return err
}

func (t *tripRepository) GetTripAuthor(ctx context.Context, id int) (int, error) {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var author int
	err = conn.QueryRow(context.Background(),
		GetTripAuthorQuery,
		id,
	).Scan(&author)

	if err != nil {
		return 0, err
	}
	return author, err
}
