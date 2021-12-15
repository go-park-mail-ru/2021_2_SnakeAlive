package repository

import (
	"context"

	"snakealive/m/internal/services/trip/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TripRepository interface {
	AddTrip(ctx context.Context, value *models.Trip, userID int) (int, error)
	GetTripById(ctx context.Context, id int) (value *models.Trip, err error)
	DeleteTrip(ctx context.Context, id int) ([]int, error)
	UpdateTrip(ctx context.Context, id int, value *models.Trip) error
	GetTripAuthors(ctx context.Context, id int) ([]int, error)
	GetTripsByUser(ctx context.Context, id int) (*[]models.Trip, error)

	AddAlbum(ctx context.Context, album *models.Album, userID int) (int, error)
	GetAlbumById(ctx context.Context, id int) (*models.Album, error)
	DeleteAlbum(ctx context.Context, id int) error
	UpdateAlbum(ctx context.Context, id int, album *models.Album) error
	GetAlbumAuthor(ctx context.Context, id int) (int, error)

	SightsByTrip(ctx context.Context, id int) (*[]int, error)
	AlbumsByUser(ctx context.Context, id int) (*[]models.Album, error)

	AddTripUser(ctx context.Context, tripId int, userId int) error

	AddLinkToCache(ctx context.Context, uuid string, id int)
	CheckLink(ctx context.Context, uuid string, id int) bool
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

	var origin int
	if len(value.Sights) < 1 {
		origin = 0
	} else {
		origin = value.Sights[0].Id
	}

	err = conn.QueryRow(context.Background(),
		AddTripQuery,
		value.Title,
		value.Description,
		origin,
	).Scan(&tripId)
	if err != nil {
		return tripId, err
	}

	err = t.AddTripUser(ctx, tripId, userID)
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

	var place models.Place
	for rows.Next() {
		_ = rows.Scan(&place.Id, &place.Name, &place.Tags, &place.Description, &place.Rating, &place.Country, &place.Photos, &place.Day)
		trip.Sights = append(trip.Sights, place)
	}

	rows, err = conn.Query(context.Background(),
		GetAlbumsByTripQuery,
		id)
	if err != nil {
		return &trip, err
	}

	var album models.Album
	for rows.Next() {
		_ = rows.Scan(&album.Id, &album.Title, &album.Description, &album.Photos)
		trip.Albums = append(trip.Albums, album)
	}

	users, err := t.GetTripAuthors(ctx, trip.Id)
	if err != nil {
		return &trip, err
	}

	trip.Users = users

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

func (t *tripRepository) DeleteTrip(ctx context.Context, id int) ([]int, error) {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return []int{}, err
	}
	defer conn.Release()

	if err = t.deletePlaces(ctx, id); err != nil {
		return []int{}, err
	}

	ids, err := t.GetTripAuthors(ctx, id)
	if err != nil {
		return []int{}, err
	}

	_, err = conn.Exec(context.Background(),
		DeleteTripQuery,
		id,
	)
	return ids, err
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

func (t *tripRepository) addPlaces(ctx context.Context, value []models.Place, tripId int) error {
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

func (t *tripRepository) GetTripAuthors(ctx context.Context, id int) ([]int, error) {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return []int{}, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		GetTripUsersQuery,
		id)
	if err != nil {
		return []int{}, err
	}

	var ids []int
	var userId int
	for rows.Next() {
		_ = rows.Scan(&userId)
		ids = append(ids, userId)
	}
	return ids, err
}

func (t *tripRepository) GetTripsByUser(ctx context.Context, id int) (*[]models.Trip, error) {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		TripsByUserQuery,
		id)
	if err != nil {
		return nil, err
	}

	trips := make([]models.Trip, 0)
	var trip models.Trip
	for rows.Next() {
		_ = rows.Scan(&trip.Id, &trip.Title, &trip.Description)
		trips = append(trips, trip)
	}

	for i := range trips {
		rows, err := conn.Query(context.Background(),
			GetPlaceForTripQuery,
			trips[i].Id)
		if err != nil {
			return &trips, err
		}

		var place models.Place
		for rows.Next() {
			_ = rows.Scan(&place.Id, &place.Name, &place.Tags, &place.Description, &place.Rating, &place.Country, &place.Photos, &place.Day)
			trips[i].Sights = append(trips[i].Sights, place)
		}

		rows, err = conn.Query(context.Background(),
			GetAlbumsByTripQuery,
			trips[i].Id)
		if err != nil {
			return &trips, err
		}

		var album models.Album
		for rows.Next() {
			_ = rows.Scan(&album.Id, &album.Title, &album.Description, &album.Photos)
			trips[i].Albums = append(trips[i].Albums, album)
		}

		users, err := t.GetTripAuthors(ctx, trips[i].Id)
		if err != nil {
			return &trips, err
		}
		trips[i].Users = users
	}
	return &trips, nil
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
		album.Photos,
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
		album.Photos,
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

func (t *tripRepository) SightsByTrip(ctx context.Context, id int) (*[]int, error) {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		SightsByTripQuery,
		id)
	if err != nil {
		return nil, err
	}

	var ids []int
	var sightId int
	for rows.Next() {
		_ = rows.Scan(&sightId)
		ids = append(ids, sightId)
	}
	return &ids, nil
}

func (t *tripRepository) AlbumsByUser(ctx context.Context, id int) (*[]models.Album, error) {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		AlbumsByUserQuery,
		id)
	if err != nil {
		return nil, err
	}

	var albums []models.Album
	var album models.Album
	for rows.Next() {
		_ = rows.Scan(&album.Id, &album.Title, &album.Description, &album.TripId, &album.UserId, &album.Photos)
		albums = append(albums, album)
	}
	return &albums, nil
}

func (t *tripRepository) AddTripUser(ctx context.Context, tripId int, userId int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		AddTripUserQuery,
		tripId,
		userId,
	)

	if err != nil {
		return err
	}
	return err
}

func (t *tripRepository) AddLinkToCache(ctx context.Context, uuid string, id int) {
	links.mu.Lock()
	links.storage[uuid] = id
	links.mu.Unlock()
}

func (t *tripRepository) CheckLink(ctx context.Context, uuid string, id int) bool {
	links.mu.Lock()

	tripId, found := links.storage[uuid]
	if found && tripId == id {
		delete(links.storage, uuid)
		links.mu.Unlock()
		return true
	}

	links.mu.Unlock()
	return false
}
