package placeRepository

import (
	"context"
	"fmt"
	"snakealive/m/pkg/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type placeStorage struct {
	dataHolder *pgxpool.Pool
}

func NewPlaceStorage(DB *pgxpool.Pool) domain.PlaceStorage {
	return &placeStorage{dataHolder: DB}
}

func (u *placeStorage) GetById(id int) (value domain.Place, err error) {

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error while getting sight")
		return domain.Place{}, err
	}
	defer conn.Release()

	var sight domain.Place
	err = conn.QueryRow(context.Background(),
		`SELECT id, name, country, rating, tags, description, photos
		FROM Places WHERE id = $1`,
		id,
	).Scan(&sight.Id, &sight.Name, &sight.Country, &sight.Rating, &sight.Tags, &sight.Description, &sight.Photos)

	return sight, err
}

func (u *placeStorage) GetPlacesByCountry(value string) (domain.TopPlaces, error) {
	places := make(domain.TopPlaces, 0)

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error while getting places")
		return places, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		`SELECT pl.id, pl.name, pl.tags, pl.photos, rw.user_id, rw.text
		FROM Places AS pl
		LEFT JOIN Reviews AS rw ON pl.id = rw.place_id
		WHERE pl.country = $1
		LIMIT 10`,
		value)
	if err != nil {
		fmt.Printf("Error while getting places")
		return places, err
	}

	var sight domain.TopPlace
	for rows.Next() {
		rows.Scan(&sight.Id, &sight.Name, &sight.Tags, &sight.Photos, &sight.UserId, &sight.Review)
		places = append(places, sight)
	}

	return places, err
}
