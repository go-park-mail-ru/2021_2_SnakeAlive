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
	var sight domain.Place

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error while getting sight")
		return sight, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT id, name, country, rating, tags, description
		FROM Places WHERE id = $1`,
		id,
	).Scan(&sight.Id, &sight.Name, &sight.Country, &sight.Rating, &sight.Tags, &sight.Description)

	return sight, err
}
