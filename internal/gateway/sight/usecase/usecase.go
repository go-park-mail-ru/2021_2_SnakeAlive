package usecase

import (
	"context"

	"snakealive/m/internal/models"
	sight_service "snakealive/m/pkg/services/sight"
)

type SightGatewayUseCase interface {
	GetSightByID(ctx context.Context, id int) (models.SightMetadata, error)
	GetSightByCountry(ctx context.Context, country string) ([]models.Sight, error)
}

type sightGatewayUseCase struct {
	sightGRPC sightGRPC
}

func (t *sightGatewayUseCase) GetSightByID(ctx context.Context, id int) (models.SightMetadata, error) {
	response, err := t.sightGRPC.GetSight(ctx, &sight_service.GetSightRequest{Id: int64(id)})
	if err != nil {
		return models.SightMetadata{}, nil
	}

	return t.adaptSightMeta(response.Sight), nil
}

func (t *sightGatewayUseCase) GetSightByCountry(ctx context.Context, country string) ([]models.Sight, error) {
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

func (t *sightGatewayUseCase) adaptSightMeta(sight *sight_service.Sight) models.SightMetadata {
	return models.SightMetadata{
		Id:      int(sight.Id),
		Name:    sight.Name,
		Tags:    sight.Tags,
		Photos:  sight.Photos,
		Country: sight.Country,
		Rating:  sight.Rating,
	}
}

func NewSightGatewayUseCase(grpc sightGRPC) SightGatewayUseCase {
	return &sightGatewayUseCase{sightGRPC: grpc}
}