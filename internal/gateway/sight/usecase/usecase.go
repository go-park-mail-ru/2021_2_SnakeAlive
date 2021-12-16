package usecase

import (
	"context"

	"snakealive/m/internal/models"
	sight_service "snakealive/m/pkg/services/sight"
)

type SightUseCase interface {
	GetSightByID(ctx context.Context, id int) (models.SightMetadata, error)
	GetSightByCountry(ctx context.Context, country string) ([]models.Sight, error)
	SearchSights(ctx context.Context, req *models.SearchSights) ([]models.SightSearch, error)
	GetSightsByTag(ctx context.Context, tag int) ([]models.Sight, error)

	GetTags(ctx context.Context) ([]models.Tag, error)
	GetRandomTags(ctx context.Context) ([]models.Tag, error)
}

type sightUseCase struct {
	sightGRPC sightGRPC
}

func (t *sightUseCase) SearchSights(ctx context.Context, req *models.SearchSights) ([]models.SightSearch, error) {
	response, err := t.sightGRPC.SearchSights(ctx, &sight_service.SearchSightRequest{
		Search:    req.Search,
		Skip:      int64(req.Skip),
		Limit:     int64(req.Limit),
		Countries: req.Countries,
		Tags:      req.Tags,
	})
	if err != nil {
		return []models.SightSearch{}, err
	}

	adapted := make([]models.SightSearch, len(response.Sights))
	for i, sight := range response.Sights {
		adapted[i] = models.SightSearch{
			Id:     int(sight.Id),
			Name:   sight.Name,
			Tags:   sight.Tags,
			Photos: sight.Photos,
			Lat:    sight.Lat,
			Lng:    sight.Lng,
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

func (t *sightUseCase) GetTags(ctx context.Context) ([]models.Tag, error) {
	response, err := t.sightGRPC.GetTags(ctx, &sight_service.GetTagsRequest{})
	if err != nil {
		return []models.Tag{}, err
	}

	tags := make([]models.Tag, len(response.Tags))
	for i, tag := range response.Tags {
		tags[i] = models.Tag{
			Id:   int(tag.Id),
			Name: tag.Name,
		}
	}

	return tags, nil
}

func (t *sightUseCase) GetRandomTags(ctx context.Context) ([]models.Tag, error) {
	response, err := t.sightGRPC.GetRandomTags(ctx, &sight_service.GetTagsRequest{})
	if err != nil {
		return []models.Tag{}, err
	}

	tags := make([]models.Tag, len(response.Tags))
	for i, tag := range response.Tags {
		tags[i] = models.Tag{
			Id:   int(tag.Id),
			Name: tag.Name,
		}
	}

	return tags, nil
}

func (t *sightUseCase) GetSightsByTag(ctx context.Context, tag int) ([]models.Sight, error) {
	response, err := t.sightGRPC.GetSightsByTag(ctx, &sight_service.GetSightsByTagRequest{Tag: int64(tag)})
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
		Id:          int(sight.Id),
		Name:        sight.Name,
		Tags:        sight.Tags,
		Photos:      sight.Photos,
		Description: sight.Description,
		Country:     sight.Country,
		Rating:      sight.Rating,
		Lat:         sight.Lat,
		Lng:         sight.Lng,
	}
}

func NewSightGatewayUseCase(grpc sightGRPC) SightUseCase {
	return &sightUseCase{sightGRPC: grpc}
}
