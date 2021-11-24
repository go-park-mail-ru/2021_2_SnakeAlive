package delivery

import (
	"context"

	"snakealive/m/internal/services/sight/models"
	"snakealive/m/internal/services/sight/usecase"
	"snakealive/m/pkg/error_adapter"
	sight_service "snakealive/m/pkg/services/sight"
)

type sightDelivery struct {
	usecase      usecase.SightUsecase
	errorAdapter error_adapter.ErrorAdapter
	sight_service.UnimplementedSightServiceServer
}

func (s *sightDelivery) GetSights(
	ctx context.Context, request *sight_service.GetSightsRequest,
) (*sight_service.GetSightsReponse, error) {
	sights, err := s.usecase.GetSightsByCountry(ctx, request.CountryName)
	if err != nil {
		return &sight_service.GetSightsReponse{}, s.errorAdapter.AdaptError(err)
	}

	adapted := &sight_service.GetSightsReponse{Sights: make([]*sight_service.Sight, len(sights))}
	for i := range sights {
		adapted.Sights[i] = s.adaptSight(sights[i])
	}

	return adapted, nil
}

func (s *sightDelivery) GetSight(
	ctx context.Context, request *sight_service.GetSightRequest,
) (*sight_service.GetSightResponse, error) {
	sight, err := s.usecase.GetSightByID(ctx, int(request.Id))
	if err != nil {
		return &sight_service.GetSightResponse{}, s.errorAdapter.AdaptError(err)
	}

	return &sight_service.GetSightResponse{
		Sight: s.adaptSight(sight),
	}, nil
}

func (s *sightDelivery) GetSightsByIDs(
	ctx context.Context, request *sight_service.GetSightsByIDsRequest,
) (*sight_service.GetSightsByIDsResponse, error) {
	sights, err := s.usecase.GetSightsByIDs(ctx, request.Ids)
	if err != nil {
		return &sight_service.GetSightsByIDsResponse{}, s.errorAdapter.AdaptError(err)
	}

	adapted := &sight_service.GetSightsByIDsResponse{Sights: make([]*sight_service.Sight, len(sights))}
	for i := range sights {
		adapted.Sights[i] = s.adaptSight(sights[i])
	}

	return adapted, nil
}

func (s *sightDelivery) adaptSight(sight models.Sight) *sight_service.Sight {
	return &sight_service.Sight{
		Id:          int64(sight.Id),
		Name:        sight.Name,
		Country:     sight.Country,
		Rating:      sight.Rating,
		Tags:        sight.Tags,
		Description: sight.Description,
		Photos:      sight.Photos,
		Lat:         sight.Lat,
		Lng:         sight.Lng,
	}
}

func NewSightDelivery(
	usecase usecase.SightUsecase, errorAdapter error_adapter.ErrorAdapter,
) sight_service.SightServiceServer {
	return &sightDelivery{
		usecase:      usecase,
		errorAdapter: errorAdapter,
	}
}
