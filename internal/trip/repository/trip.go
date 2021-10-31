package tripRepository

import (
	"context"
	"fmt"
	"snakealive/m/pkg/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type tripStorage struct {
	dataHolder *pgxpool.Pool
}

func NewTripStorage(DB *pgxpool.Pool) domain.TripStorage {
	return &tripStorage{dataHolder: DB}
}

func (t *tripStorage) Add(value domain.Trip, user domain.User) (int, error) {
	var tripId int

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding trip ", err)
		return tripId, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`INSERT INTO Trips ("title", "description", "days", "user_id", "origin") VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		value.Title,
		value.Description,
		len(value.Days),
		user.Id,
		value.Days[0][0].Id,
	).Scan(&tripId)
	if err != nil {
		fmt.Printf("Connection error while adding trip ", err)
		return tripId, err
	}

	err = t.addPlaces(value.Days, tripId)
	return tripId, err
}

func (t *tripStorage) GetById(id int) (value domain.Trip, err error) {
	var trip domain.Trip

	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding trip ", err)
		return trip, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT id, title, description, days
		FROM Trips WHERE id = $1`,
		id,
	).Scan(&trip.Id, &trip.Title, &trip.Description, &trip.Days)

	rows, err := conn.Query(context.Background(),
		`SELECT pl.id, pl.name, pl.tags, pl.description, pl.rating, pl.country, pl.photos, tr.day
		FROM TripsPlaces AS tr
		JOIN Places AS pl ON tr.place_id = pl.id
		WHERE tr.trip_id = $1
		ORDER BY tr.day, tr.order`,
		id)
	if err != nil {
		fmt.Printf("Error while getting places")
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
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding trip ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`UPDATE Trips SET "title" = $1, "description" = $2, "days" = $3, "origin" = $4 WHERE id = $5`,
		value.Title,
		value.Description,
		len(value.Days),
		value.Days[0][0].Id,
		id,
	)
	if err != nil || t.deletePlaces(id) != nil {
		fmt.Printf("Connection error while adding trip ", err)
		return err
	}

	err = t.addPlaces(value.Days, id)
	return err
}

func (t *tripStorage) Delete(id int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while deleting places ", err)
		return err
	}
	defer conn.Release()

	if err = t.deletePlaces(id); err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(),
		`DELETE FROM Trips WHERE id = $1`,
		id,
	)
	return err
}

func (t *tripStorage) deletePlaces(tripId int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while deleting places ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`DELETE FROM TripsPlaces WHERE trip_id = $1`,
		tripId,
	)
	return err
}

func (t *tripStorage) addPlaces(value [][]domain.Place, tripId int) error {
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding places ", err)
		return err
	}
	defer conn.Release()

	for day, places := range value {
		for ind, place := range places {
			_, err = conn.Exec(context.Background(),
				`INSERT INTO TripsPlaces ("trip_id", "place_id", "day", "order") VALUES ($1, $2, $3, $4)`,
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
	conn, err := t.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding places ", err)
		return 0
	}
	defer conn.Release()

	var author int
	err = conn.QueryRow(context.Background(),
		`SELECT user_id
	FROM Trips WHERE id = $1`,
		id,
	).Scan(&author)

	if err != nil {
		return 0
	}
	return author
}
