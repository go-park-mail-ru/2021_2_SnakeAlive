package placeRepository

import (
	"context"
	"snakealive/m/internal/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type placeStorage struct {
	dataHolder *pgxpool.Pool
}

func NewPlaceStorage(DB *pgxpool.Pool) domain.PlaceStorage {
	return &placeStorage{dataHolder: DB}
}

const GetPlaceByIdQuery = `SELECT id, name, country, rating, tags, description, photos FROM Places WHERE id = $1`

func (u *placeStorage) GetById(id int) (value domain.Place, err error) {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return domain.Place{}, err
	}
	defer conn.Release()

	var sight domain.Place
	err = conn.QueryRow(context.Background(),
		GetPlaceByIdQuery,
		id,
	).Scan(&sight.Id, &sight.Name, &sight.Country, &sight.Rating, &sight.Tags, &sight.Description, &sight.Photos)

	return sight, err
}

func (u *placeStorage) GetPlacesByCountry(value string) (domain.TopPlaces, error) {
	places := make(domain.TopPlaces, 0)

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return places, err
	}
	defer conn.Release()
	const GetPlacesByCountryQuery = `select id, name, tags, photos, country, avg(rating) over (partition by id) as total_rating from Places where country = $1 order by total_rating desc limit 10`
	rows, err := conn.Query(context.Background(),
		GetPlacesByCountryQuery,
		value)
	if err != nil {
		return places, err
	}

	var sight domain.TopPlace
	for rows.Next() {
		rows.Scan(&sight.Id, &sight.Name, &sight.Tags, &sight.Photos, &sight.Country, &sight.Rating)
		places = append(places, sight)
	}

	return places, err
}
