package repository

import (
	"context"
	"snakealive/m/internal/domain"
	"snakealive/m/internal/services/trip/models"

	"github.com/jackc/pgx/v4/pgxpool"
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
}

type tripRepository struct {
	dataHolder *pgxpool.Pool
}

func NewTripRepository(DB *pgxpool.Pool) TripRepository {
	return &tripRepository{dataHolder: DB}
}

func (t *tripRepository) AddTrip(ctx context.Context, value *models.Trip, userID int) (int, error) {
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

func (t *tripRepository) GetTripById(ctx context.Context, id int) (*models.Trip, error) {
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

func (t *tripRepository) UpdateTrip(ctx context.Context, id int, value *models.Trip) error {
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

func (t *tripRepository) DeleteTrip(ctx context.Context, id int) error {
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

func (t *tripRepository) AddAlbum(ctx context.Context, album *models.Album, userID int) (int, error) {
	var albumId int

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return albumId, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		AddAlbumQuery,
		album.Title,
		album.Description,
		album.TripId,
		userID,
	).Scan(&albumId)

	return albumId, err
}

func (t *tripRepository) GetAlbumById(ctx context.Context, id int) (*models.Album, error) {
	var album models.Album

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return &album, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		GetAlbumQuery,
		id,
	).Scan(&album.Id, &album.Title, &album.Description, &album.TripId, &album.UserId, &album.Photos)

	return &album, err
}

func (t *tripRepository) DeleteAlbum(ctx context.Context, id int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		DeleteAlbumQuery,
		id,
	)
	return err
}

func (t *tripRepository) UpdateAlbum(ctx context.Context, id int, album *models.Album) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		UpdateAlbumQuery,
		album.Title,
		album.Description,
		id,
	)

	return err
}

func (t *tripRepository) GetAlbumAuthor(ctx context.Context, id int) (int, error) {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var author int
	err = conn.QueryRow(context.Background(),
		GetAlbumAuthorQuery,
		id,
	).Scan(&author)

	if err != nil {
		return 0, err
	}
	return author, err
}
