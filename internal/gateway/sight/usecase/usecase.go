package usecase

import (
	"context"

	"snakealive/m/internal/models"
	sight_service "snakealive/m/pkg/services/sight"
)

type SightUseCase interface {
	GetSightByID(ctx context.Context, id int) (models.SightMetadata, error)
	GetSightByCountry(ctx context.Context, country string) ([]models.Sight, error)
	SearchSights(ctx context.Context, search string, skip, limit int) ([]models.SightSearch, error)
	GetSightsByTag(ctx context.Context, tag string) ([]models.Sight, error)
}

type sightUseCase struct {
	sightGRPC sightGRPC
}

func (t *sightUseCase) SearchSights(ctx context.Context, search string, skip, limit int) ([]models.SightSearch, error) {
	response, err := t.sightGRPC.SearchSights(ctx, &sight_service.SearchSightRequest{
		Search: search,
		Skip:   int64(skip),
		Limit:  int64(limit),
	})
	if err != nil {
		return []models.SightSearch{}, err
	}

	adapted := make([]models.SightSearch, len(response.Sights))
	for i, sight := range response.Sights {
		adapted[i] = models.SightSearch{
			Id:   int(sight.Id),
			Name: sight.Name,
			Tags: sight.Tags,
			Lat:  sight.Lat,
			Lng:  sight.Lng,
		}
	}

	return adapted, nil
}

func (t *sightUseCase) GetSightByID(ctx context.Context, id int) (models.SightMetadata, error) {
	response, err := t.sightGRPC.GetSight(ctx, &sight_service.GetSightRequest{Id: int64(id)})
	if err != nil {
		return models.SightMetadata{}, nil
	}

	return t.adaptSightMeta(response.Sight), nil
}

func (t *sightUseCase) GetSightByCountry(ctx context.Context, country string) ([]models.Sight, error) {
	response, err := t.sightGRPC.GetSights(ctx, &sight_service.GetSightsRequest{CountryName: country})
	if err != nil {
		return []models.Sight{}, err
	}

	adapted := make([]models.Sight, len(response.Sights))
	for i := range adapted {
		adapted[i] = models.Sight{
			Description:   response.Sights[i].Description,
			SightMetadata: t.adaptSightMeta(response.Sights[i]),
		}
	}

	return adapted, nil
}

func (t *sightUseCase) GetSightsByTag(ctx context.Context, tag string) ([]models.Sight, error) {
	response, err := t.sightGRPC.GetSightsByTag(ctx, &sight_service.GetSightsByTagRequest{Tag: tag})
	if err != nil {
		return []models.Sight{}, err
	}

	adapted := make([]models.Sight, len(response.Sights))
	for i := range adapted {
		adapted[i] = models.Sight{
			Description:   response.Sights[i].Description,
			SightMetadata: t.adaptSightMeta(response.Sights[i]),
		}
	}

	return adapted, nil
}

func (t *sightUseCase) adaptSightMeta(sight *sight_service.Sight) models.SightMetadata {
	return models.SightMetadata{
		Id:      int(sight.Id),
		Name:    sight.Name,
		Tags:    sight.Tags,
		Photos:  sight.Photos,
		Country: sight.Country,
		Rating:  sight.Rating,
		Lat:     sight.Lat,
		Lng:     sight.Lng,
	}
}

func NewSightGatewayUseCase(grpc sightGRPC) SightUseCase {
	return &sightUseCase{sightGRPC: grpc}
}
