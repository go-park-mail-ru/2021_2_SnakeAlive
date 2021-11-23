package delivery

import (
	"context"
	"snakealive/m/internal/domain"
	"snakealive/m/internal/services/trip/models"
	"snakealive/m/internal/services/trip/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/errors"
	trip_service "snakealive/m/pkg/services/trip"

	"github.com/golang/protobuf/ptypes/empty"
)

type tripDelivery struct {
	tripUsecase  usecase.TripUseCase
	errorAdapter error_adapter.ErrorAdapter
	trip_service.UnimplementedTripServiceServer
}

func NewTripDelivery(tripUsecase usecase.TripUseCase, errorAdapter error_adapter.ErrorAdapter) trip_service.TripServiceServer {
	return &tripDelivery{
		tripUsecase:  tripUsecase,
		errorAdapter: errorAdapter,
	}
}

func (s *tripDelivery) GetTrip(ctx context.Context, request *trip_service.TripRequest) (*trip_service.Trip, error) {
	authorized, err := s.tripUsecase.CheckAuthor(ctx, int(request.UserId), int(request.TripId))
	if !authorized || err != nil {
		return &trip_service.Trip{}, errors.DeniedAccess
	}

	trip, err := s.tripUsecase.GetById(ctx, int(request.TripId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	protoDays := ProtoDaysFromPlaces(trip.Sights)

	return &trip_service.Trip{
		Id:          int64(trip.Id),
		Title:       trip.Title,
		Description: trip.Description,
		Sights:      protoDays,
	}, nil
}

func (s *tripDelivery) AddTrip(ctx context.Context, request *trip_service.ModifyTripRequest) (*trip_service.Trip, error) {
	places := PlacesFromProtoDays(request.Trip.Sights)

	id, err := s.tripUsecase.Add(ctx, &models.Trip{
		Id:          int(request.Trip.Id),
		Title:       request.Trip.Title,
		Description: request.Trip.Description,
		Sights:      places,
	}, int(request.UserId))

	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	trip, err := s.tripUsecase.GetById(ctx, id)
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	days := ProtoDaysFromPlaces(trip.Sights)
	return &trip_service.Trip{
		Id:          int64(trip.Id),
		Title:       trip.Title,
		Description: trip.Description,
		Sights:      days,
	}, nil
}

func (s *tripDelivery) Update(ctx context.Context, request *trip_service.ModifyTripRequest) (*trip_service.Trip, error) {
	authorized, err := s.tripUsecase.CheckAuthor(ctx, int(request.UserId), int(request.Trip.Id))
	if !authorized || err != nil {
		return &trip_service.Trip{}, errors.DeniedAccess
	}

	places := PlacesFromProtoDays(request.Trip.Sights)

	err = s.tripUsecase.Update(ctx, int(request.Trip.Id), &models.Trip{
		Id:          int(request.Trip.Id),
		Title:       request.Trip.Title,
		Description: request.Trip.Description,
		Sights:      places,
	})
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	trip, err := s.tripUsecase.GetById(ctx, int(request.Trip.Id))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	days := ProtoDaysFromPlaces(trip.Sights)
	return &trip_service.Trip{
		Id:          int64(trip.Id),
		Title:       trip.Title,
		Description: trip.Description,
		Sights:      days,
	}, nil
}

func (s *tripDelivery) Delete(ctx context.Context, request *trip_service.TripRequest) (*empty.Empty, error) {
	authorized, err := s.tripUsecase.CheckAuthor(ctx, int(request.UserId), int(request.TripId))
	if !authorized || err != nil {
		return &empty.Empty{}, errors.DeniedAccess
	}

	err = s.tripUsecase.Delete(ctx, int(request.TripId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func ProtoDaysFromPlaces(places []domain.Place) []*trip_service.Sight {
	var protoDays []*trip_service.Sight
	for _, sight := range places {
		protoSight := trip_service.Sight{
			Id:          int64(sight.Id),
			Name:        sight.Name,
			Country:     sight.Country,
			Rating:      sight.Rating,
			Tags:        sight.Tags,
			Description: sight.Description,
			Photos:      sight.Photos,
			Day:         int64(sight.Day),
		}
		protoDays = append(protoDays, &protoSight)
	}
	return protoDays
}

func PlacesFromProtoDays(sights []*trip_service.Sight) []domain.Place {
	var places []domain.Place
	for _, sight := range sights {
		placesSight := domain.Place{
			Id:          int(sight.Id),
			Name:        sight.Name,
			Country:     sight.Country,
			Rating:      sight.Rating,
			Tags:        sight.Tags,
			Description: sight.Description,
			Photos:      sight.Photos,
			Day:         int(sight.Day),
		}
		places = append(places, placesSight)
	}
	return places
}
