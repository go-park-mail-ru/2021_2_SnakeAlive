package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"snakealive/m/internal/services/sight/models"
	"snakealive/m/pkg/errors"
)

type SightRepository interface {
	GetSightsByCountry(ctx context.Context, country string) ([]models.Sight, error)
	GetSightByID(ctx context.Context, id int) (models.Sight, error)
}

type sightRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (s *sightRepository) GetSightsByCountry(ctx context.Context, country string) ([]models.Sight, error) {
	request := s.queryFactory.CreateGetSightsByCountry(country)
	rows, err := s.conn.Query(ctx, request.Request, request.Params...)
	if err != nil {
		return []models.Sight{}, err
	}

	sights := make([]models.Sight, 0)
	for rows.Next() {
		var sight models.Sight

		if err = rows.Scan(
			&sight.Id, &sight.Name, &sight.Tags,
			&sight.Photos, &sight.Country, &sight.Rating,
		); err != nil {
			return []models.Sight{}, err
		}

		sights = append(sights, sight)
	}

	return sights, nil
}

func (s *sightRepository) GetSightByID(ctx context.Context, id int) (models.Sight, error) {
	var sight models.Sight
	request := s.queryFactory.CreateGetSightByID(id)
	if err := s.conn.QueryRow(ctx, request.Request, request.Params...).Scan(
		&sight.Id, &sight.Name, &sight.Country, &sight.Rating,
		&sight.Tags, &sight.Description, &sight.Photos,
	); err != nil {
		if err == pgx.ErrNoRows {
			return models.Sight{}, errors.SightDoesNotExist
		}

		return models.Sight{}, err
	}

	return sight, nil
}

func NewSightRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) SightRepository {
	return &sightRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}