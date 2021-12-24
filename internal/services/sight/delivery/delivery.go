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

func (s *sightDelivery) SearchSights(
	ctx context.Context, request *sight_service.SearchSightRequest,
) (response *sight_service.SearchSightResponse, err error) {
	sights, err := s.usecase.SearchSights(ctx, &models.SightsSearch{
		Skip:       int(request.Skip),
		Limit:      int(request.Limit),
		Search:     request.Search,
		Tags:       request.Tags,
		Countries:  request.Countries,
		MinReviews: int(request.MinReviews),
		MinRating:  int(request.MinRating),
	})
	if err != nil {
		return &sight_service.SearchSightResponse{}, err
	}

	adapted := &sight_service.SearchSightResponse{Sights: make([]*sight_service.Sight, len(sights))}
	for i := range sights {
		adapted.Sights[i] = s.adaptSight(sights[i])
	}
	return adapted, nil
}

func (s *sightDelivery) GetSightsByTag(
	ctx context.Context, request *sight_service.GetSightsByTagRequest,
) (*sight_service.GetSightsByTagResponse, error) {
	sights, err := s.usecase.GetSightByTag(ctx, request.Tag)
	if err != nil {
		return &sight_service.GetSightsByTagResponse{}, s.errorAdapter.AdaptError(err)
	}

	adapted := &sight_service.GetSightsByTagResponse{Sights: make([]*sight_service.Sight, len(sights))}
	for i := range sights {
		adapted.Sights[i] = s.adaptSight(sights[i])
	}

	return adapted, nil
}

func (s *sightDelivery) GetTags(
	ctx context.Context, request *sight_service.GetTagsRequest,
) (response *sight_service.GetTagsResponse, err error) {
	tags, err := s.usecase.GetTags(ctx)
	if err != nil {
		return &sight_service.GetTagsResponse{}, err
	}

	response = &sight_service.GetTagsResponse{Tags: make([]*sight_service.Tag, len(tags))}
	for i, tag := range tags {
		response.Tags[i] = &sight_service.Tag{
			Id:   int64(tag.Id),
			Name: tag.Name,
		}
	}

	return response, nil
}

func (s *sightDelivery) GetRandomTags(
	ctx context.Context, request *sight_service.GetTagsRequest,
) (response *sight_service.GetTagsResponse, err error) {
	tags, err := s.usecase.GetRandomTags(ctx)
	if err != nil {
		return &sight_service.GetTagsResponse{}, err
	}

	response = &sight_service.GetTagsResponse{Tags: make([]*sight_service.Tag, len(tags))}
	for i, tag := range tags {
		response.Tags[i] = &sight_service.Tag{
			Id:   int64(tag.Id),
			Name: tag.Name,
		}
	}

	return response, nil
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
