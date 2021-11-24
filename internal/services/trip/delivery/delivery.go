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
	authorized, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.UserId), int(request.TripId))
	if !authorized || err != nil {
		return &trip_service.Trip{}, errors.DeniedAccess
	}

	trip, err := s.tripUsecase.GetTripById(ctx, int(request.TripId))
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

	id, err := s.tripUsecase.AddTrip(ctx, &models.Trip{
		Id:          int(request.Trip.Id),
		Title:       request.Trip.Title,
		Description: request.Trip.Description,
		Sights:      places,
	}, int(request.UserId))

	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	trip, err := s.tripUsecase.GetTripById(ctx, id)
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

func (s *tripDelivery) UpdateTrip(ctx context.Context, request *trip_service.ModifyTripRequest) (*trip_service.Trip, error) {
	authorized, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.UserId), int(request.Trip.Id))
	if !authorized || err != nil {
		return &trip_service.Trip{}, errors.DeniedAccess
	}

	places := PlacesFromProtoDays(request.Trip.Sights)

	err = s.tripUsecase.UpdateTrip(ctx, int(request.Trip.Id), &models.Trip{
		Id:          int(request.Trip.Id),
		Title:       request.Trip.Title,
		Description: request.Trip.Description,
		Sights:      places,
	})
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	trip, err := s.tripUsecase.GetTripById(ctx, int(request.Trip.Id))
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

func (s *tripDelivery) DeleteTrip(ctx context.Context, request *trip_service.TripRequest) (*empty.Empty, error) {
	authorized, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.UserId), int(request.TripId))
	if !authorized || err != nil {
		return &empty.Empty{}, errors.DeniedAccess
	}

	err = s.tripUsecase.DeleteTrip(ctx, int(request.TripId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func (s *tripDelivery) GetAlbum(ctx context.Context, request *trip_service.AlbumRequest) (*trip_service.Album, error) {
	authorized, err := s.tripUsecase.CheckAlbumAuthor(ctx, int(request.UserId), int(request.AlbumId))
	if !authorized || err != nil {
		return &trip_service.Album{}, errors.DeniedAccess
	}

	album, err := s.tripUsecase.GetAlbumById(ctx, int(request.AlbumId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	return &trip_service.Album{
		Id:          int64(album.Id),
		Title:       album.Title,
		Author:      int64(album.UserId),
		Description: album.Description,
		Photos:      album.Photos,
	}, nil
}

func (s *tripDelivery) AddAlbum(ctx context.Context, request *trip_service.ModifyAlbumRequest) (*trip_service.Album, error) {
	id, err := s.tripUsecase.AddAlbum(ctx, &models.Album{
		Id:          int(request.Album.Id),
		TripId:      int(request.Album.TripId),
		Title:       request.Album.Title,
		Description: request.Album.Description,
	}, int(request.UserId))

	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	album, err := s.tripUsecase.GetAlbumById(ctx, id)
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	return &trip_service.Album{
		Id:          int64(album.Id),
		Title:       album.Title,
		TripId:      int64(album.TripId),
		Author:      int64(album.UserId),
		Description: album.Description,
	}, nil
}

func (s *tripDelivery) UpdateAlbum(ctx context.Context, request *trip_service.ModifyAlbumRequest) (*trip_service.Album, error) {
	authorized, err := s.tripUsecase.CheckAlbumAuthor(ctx, int(request.UserId), int(request.Album.Id))
	if !authorized || err != nil {
		return &trip_service.Album{}, errors.DeniedAccess
	}

	err = s.tripUsecase.UpdateAlbum(ctx, int(request.Album.Id), &models.Album{
		Id:          int(request.Album.Id),
		Title:       request.Album.Title,
		TripId:      int(request.Album.TripId),
		Description: request.Album.Description,
	})
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	album, err := s.tripUsecase.GetAlbumById(ctx, int(request.Album.Id))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	return &trip_service.Album{
		Id:          int64(album.Id),
		Title:       album.Title,
		TripId:      int64(album.TripId),
		Author:      int64(album.UserId),
		Description: album.Description,
		Photos:      album.Photos,
	}, nil
}

func (s *tripDelivery) UploadPhoto(ctx context.Context, request *trip_service.UploadRequest) (*empty.Empty, error) {
	authorized, err := s.tripUsecase.CheckAlbumAuthor(ctx, int(request.UserId), int(request.AlbumId))
	if !authorized || err != nil {
		return &empty.Empty{}, errors.DeniedAccess
	}

	err = s.tripUsecase.UploadPhoto(ctx, request.Filename, int(request.AlbumId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func (s *tripDelivery) DeleteAlbum(ctx context.Context, request *trip_service.AlbumRequest) (*empty.Empty, error) {
	authorized, err := s.tripUsecase.CheckAlbumAuthor(ctx, int(request.UserId), int(request.AlbumId))
	if !authorized || err != nil {
		return &empty.Empty{}, errors.DeniedAccess
	}

	err = s.tripUsecase.DeleteAlbum(ctx, int(request.AlbumId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func (s *tripDelivery) SightsByTrip(ctx context.Context, request *trip_service.SightsRequest) (*trip_service.Sights, error) {
	sights, err := s.tripUsecase.SightsByTrip(ctx, int(request.TripId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	var ids []int64
	for _, id := range *sights {
		ids = append(ids, int64(id))
	}

	return &trip_service.Sights{
		Ids: ids,
	}, nil
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
