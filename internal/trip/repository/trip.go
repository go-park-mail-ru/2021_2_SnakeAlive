package tripRepository

import (
	"context"
	"snakealive/m/internal/domain"
	logs "snakealive/m/pkg/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

type tripStorage struct {
	dataHolder *pgxpool.Pool
}

func NewTripStorage(DB *pgxpool.Pool) domain.TripStorage {
	return &tripStorage{dataHolder: DB}
}

const AddTripQuery = `INSERT INTO Trips ("title", "description", "days", "user_id", "origin") 
						VALUES ($1, $2, $3, $4, $5) RETURNING id`

const GetTripQuery = `SELECT id, title, description FROM Trips WHERE id = $1`

const GetPlaceForTripQuery = `SELECT pl.id, pl.name, pl.tags, pl.description, pl.rating, pl.country, pl.photos, tr.day
								FROM TripsPlaces AS tr JOIN Places AS pl ON tr.place_id = pl.id WHERE tr.trip_id = $1
								ORDER BY tr.day, tr.order`

const UpdateTripQuery = `UPDATE Trips SET "title" = $1, "description" = $2, "days" = $3, "origin" = $4 WHERE id = $5`

const DeleteTripQuery = `DELETE FROM Trips WHERE id = $1`

const DeletePlacesForTripQuery = `DELETE FROM TripsPlaces WHERE trip_id = $1`

const AddPlacesForTripQuery = `INSERT INTO TripsPlaces ("trip_id", "place_id", "day", "order") VALUES ($1, $2, $3, $4)`

const GetTripAuthorQuery = `SELECT user_id FROM Trips WHERE id = $1`

func (t *tripStorage) Add(value domain.Trip, user domain.User) (int, error) {
	logger := logs.GetLogger()
	var tripId int

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return tripId, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		AddTripQuery,
		value.Title,
		value.Description,
		len(value.Days),
		user.Id,
		value.Days[0][0].Id,
	).Scan(&tripId)
	if err != nil {
		logger.Error("error while scanning trip info")
		return tripId, err
	}

	err = t.addPlaces(value.Days, tripId)
	return tripId, err
}

func (t *tripStorage) GetById(id int) (value domain.Trip, err error) {
	logger := logs.GetLogger()
	var trip domain.Trip

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return trip, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		GetTripQuery,
		id,
	).Scan(&trip.Id, &trip.Title, &trip.Description)
	if err != nil {
		logger.Error("error while scanning trip info")
		return trip, err
	}

	rows, err := conn.Query(context.Background(),
		GetPlaceForTripQuery,
		id)
	if err != nil {
		logger.Error("error while getting trip places")
		return trip, err
	}

	var place domain.Place
	var places []domain.Place
	var day int
	currentDay := 0
	for rows.Next() {
		rows.Scan(&place.Id, &place.Name, &place.Tags, &place.Description, &place.Rating, &place.Country, &place.Photos, &day)
		if day != currentDay && len(places) > 0 {
			currentDay = day
			trip.Days = append(trip.Days, places)
			places = []domain.Place{}
		}
		places = append(places, place)
	}
	trip.Days = append(trip.Days, places)

	return trip, err
}

func (t *tripStorage) Update(id int, value domain.Trip) error {
	logger := logs.GetLogger()

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		UpdateTripQuery,
		value.Title,
		value.Description,
		len(value.Days),
		value.Days[0][0].Id,
		id,
	)
	if err != nil || t.deletePlaces(id) != nil {
		logger.Error("error while removing previous places")
		return err
	}

	err = t.addPlaces(value.Days, id)
	return err
}

func (t *tripStorage) Delete(id int) error {
	logger := logs.GetLogger()

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	if err = t.deletePlaces(id); err != nil {
		logger.Error("error while deleting places")
		return err
	}
	_, err = conn.Exec(context.Background(),
		DeleteTripQuery,
		id,
	)
	return err
}

func (t *tripStorage) deletePlaces(tripId int) error {
	logger := logs.GetLogger()

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		DeletePlacesForTripQuery,
		tripId,
	)
	return err
}

func (t *tripStorage) addPlaces(value [][]domain.Place, tripId int) error {
	logger := logs.GetLogger()

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	for day, places := range value {
		for ind, place := range places {
			_, err = conn.Exec(context.Background(),
				AddPlacesForTripQuery,
				tripId,
				place.Id,
				day,
				ind,
			)
		}
	}
	return err
}

func (t *tripStorage) GetTripAuthor(id int) int {
	logger := logs.GetLogger()

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return 0
	}
	defer conn.Release()

	var author int
	err = conn.QueryRow(context.Background(),
		GetTripAuthorQuery,
		id,
	).Scan(&author)

	if err != nil {
		logger.Error("error while getting trip author from database")
		return 0
	}
	return author
}
