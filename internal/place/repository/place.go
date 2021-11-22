package placeRepository

import (
	"context"
	"snakealive/m/internal/domain"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

type placeStorage struct {
	dataHolder *pgxpool.Pool
}

func NewPlaceStorage(DB *pgxpool.Pool) domain.PlaceStorage {
	return &placeStorage{dataHolder: DB}
}

func (u *placeStorage) GetById(id int) (value domain.Place, err error) {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return domain.Place{}, err
	}
	defer conn.Release()

	var sight domain.Place
	err = conn.QueryRow(context.Background(),
		cnst.GetPlaceByIdQuery,
		id,
	).Scan(&sight.Id, &sight.Name, &sight.Country, &sight.Lat, &sight.Lng, &sight.Rating, &sight.Tags, &sight.Description, &sight.Photos)

	return sight, err
}

func (u *placeStorage) GetPlacesByCountry(value string) (domain.TopPlaces, error) {
	logger := logs.GetLogger()
	places := make(domain.TopPlaces, 0)

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return places, err
	}
	defer conn.Release()
	const GetPlacesByCountryQuery = `select id, name, tags, photos, country, avg(rating) over (partition by id) as total_rating from Places where country = $1 order by total_rating desc limit 10`
	rows, err := conn.Query(context.Background(),
		GetPlacesByCountryQuery,
		value)
	if err != nil {
		logger.Error("error while getting list of places from database")

		return places, err
	}

	var sight domain.TopPlace
	for rows.Next() {
		rows.Scan(&sight.Id, &sight.Name, &sight.Tags, &sight.Photos, &sight.Country, &sight.Rating)
		places = append(places, sight)
	}

	return places, err
}
