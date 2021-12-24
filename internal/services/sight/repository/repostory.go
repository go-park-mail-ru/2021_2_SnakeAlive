package repository

import (
	"context"

	"snakealive/m/internal/services/sight/models"
	"snakealive/m/pkg/errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SightRepository interface {
	GetSightsByCountry(ctx context.Context, country string) ([]models.Sight, error)
	GetSightByID(ctx context.Context, id int) (models.Sight, error)
	GetSightByIDs(ctx context.Context, ids []int64) ([]models.Sight, error)
	GetSightByTag(ctx context.Context, tag int64) ([]models.Sight, error)
	SearchSights(ctx context.Context, req *models.SightsSearch) ([]models.Sight, error)
	GetTags(ctx context.Context) ([]models.Tag, error)
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
	defer rows.Close()

	sights := make([]models.Sight, 0)
	for rows.Next() {
		var sight models.Sight

		if err = rows.Scan(
			&sight.Id, &sight.Name, &sight.Description, &sight.Tags, &sight.Photos,
			&sight.Country, &sight.Rating,
			&sight.Lng, &sight.Lat,
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
		&sight.Tags, &sight.Description, &sight.Photos, &sight.Lng, &sight.Lat,
	); err != nil {
		if err == pgx.ErrNoRows {
			return models.Sight{}, errors.SightDoesNotExist
		}

		return models.Sight{}, err
	}

	return sight, nil
}

func (s *sightRepository) GetSightByIDs(ctx context.Context, ids []int64) ([]models.Sight, error) {
	request := s.queryFactory.CreateGetSightsByIDs(ids)
	rows, err := s.conn.Query(ctx, request.Request, request.Params...)
	if err != nil {
		return []models.Sight{}, err
	}
	defer rows.Close()

	sights := make([]models.Sight, 0)
	for rows.Next() {
		var sight models.Sight

		if err = rows.Scan(
			&sight.Id, &sight.Name, &sight.Country, &sight.Rating,
			&sight.Tags, &sight.Description, &sight.Photos, &sight.Lng, &sight.Lat,
		); err != nil {
			return []models.Sight{}, err
		}

		sights = append(sights, sight)
	}

	return sights, nil
}

func (s *sightRepository) GetSightByTag(ctx context.Context, tag int64) ([]models.Sight, error) {
	request := s.queryFactory.CreateGetSightsByTag(tag)
	rows, err := s.conn.Query(ctx, request.Request, request.Params...)
	if err != nil {
		return []models.Sight{}, err
	}
	defer rows.Close()

	sights := make([]models.Sight, 0)
	for rows.Next() {
		var sight models.Sight

		if err = rows.Scan(
			&sight.Id, &sight.Name, &sight.Country, &sight.Rating,
			&sight.Tags, &sight.Description, &sight.Photos, &sight.Lng, &sight.Lat,
		); err != nil {
			return []models.Sight{}, err
		}

		sights = append(sights, sight)
	}

	return sights, nil
}

func (s *sightRepository) SearchSights(ctx context.Context, req *models.SightsSearch) ([]models.Sight, error) {
	request := s.queryFactory.CreateSearchSights(req)
	rows, err := s.conn.Query(ctx, request.Request, request.Params...)
	if err != nil {
		return []models.Sight{}, err
	}
	defer rows.Close()

	sights := make([]models.Sight, 0)
	for rows.Next() {
		var sight models.Sight

		if err = rows.Scan(
			&sight.Id, &sight.Name, &sight.Country, &sight.Rating,
			&sight.Tags, &sight.Description, &sight.Photos, &sight.Lng, &sight.Lat,
		); err != nil {
			return []models.Sight{}, err
		}

		sights = append(sights, sight)
	}

	filteredSights := make([]models.Sight, 0)
	for _, sight := range sights {
		if req.MinRating != 0 && sight.Rating < float32(req.MinRating) {
			continue
		}
		filteredSights = append(filteredSights, sight)
	}

	return filteredSights, nil
}

func (s *sightRepository) GetTags(ctx context.Context) ([]models.Tag, error) {
	request := s.queryFactory.CreateGetTags()
	rows, err := s.conn.Query(ctx, request.Request, request.Params...)
	if err != nil {
		return []models.Tag{}, nil
	}
	defer rows.Close()

	tags := make([]models.Tag, 0)
	for rows.Next() {
		var tag models.Tag

		if err = rows.Scan(&tag.Id, &tag.Name); err != nil {
			return []models.Tag{}, nil
		}

		tags = append(tags, tag)
	}

	return tags, nil
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
