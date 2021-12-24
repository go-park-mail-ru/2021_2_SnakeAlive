package delivery

import (
	"context"
	"strconv"

	"snakealive/m/internal/services/trip/models"
	"snakealive/m/internal/services/trip/usecase"
	cnst "snakealive/m/pkg/constants"
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
	protoAlbums := ProtoAlbumsFromAlbums(trip.Albums)

	users := make([]int64, 0)
	for _, id := range trip.Users {
		users = append(users, int64(id))
	}

	return &trip_service.Trip{
		Id:          int64(trip.Id),
		Title:       trip.Title,
		Description: trip.Description,
		Sights:      protoDays,
		Albums:      protoAlbums,
		Users:       users,
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
	protoAlbums := ProtoAlbumsFromAlbums(trip.Albums)

	users := make([]int64, 0)
	for _, id := range trip.Users {
		users = append(users, int64(id))
	}

	return &trip_service.Trip{
		Id:          int64(trip.Id),
		Title:       trip.Title,
		Description: trip.Description,
		Sights:      days,
		Albums:      protoAlbums,
		Users:       users,
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
	protoAlbums := ProtoAlbumsFromAlbums(trip.Albums)

	users := make([]int64, 0)
	for _, id := range trip.Users {
		users = append(users, int64(id))
	}

	return &trip_service.Trip{
		Id:          int64(trip.Id),
		Title:       trip.Title,
		Description: trip.Description,
		Sights:      days,
		Albums:      protoAlbums,
		Users:       users,
	}, nil
}

func (s *tripDelivery) DeleteTrip(ctx context.Context, request *trip_service.TripRequest) (*trip_service.Users, error) {
	authorized, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.UserId), int(request.TripId))
	if !authorized || err != nil {
		return &trip_service.Users{}, errors.DeniedAccess
	}

	responce, err := s.tripUsecase.DeleteTrip(ctx, int(request.TripId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	var users trip_service.Users
	for _, user := range responce {
		users.Users = append(users.Users, int64(user))
	}

	return &users, nil
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
		TripId:      int64(album.TripId),
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
		Photos:      request.Album.Photos,
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
		Photos:      request.Album.Photos,
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

func (s *tripDelivery) GetTripsByUser(ctx context.Context, request *trip_service.ByUserRequest) (*trip_service.Trips, error) {
	trips, err := s.tripUsecase.TripsByUser(ctx, int(request.UserId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	var protoTrips trip_service.Trips
	for _, trip := range *trips {
		protoAlbums := ProtoAlbumsFromAlbums(trip.Albums)
		protoSights := ProtoDaysFromPlaces(trip.Sights)

		users := make([]int64, 0)
		for _, id := range trip.Users {
			users = append(users, int64(id))
		}

		protoTrips.Trips = append(protoTrips.Trips, &trip_service.Trip{
			Id:          int64(trip.Id),
			Title:       trip.Title,
			Description: trip.Description,
			Sights:      protoSights,
			Albums:      protoAlbums,
			Users:       users,
		})
	}
	return &protoTrips, nil
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

func (s *tripDelivery) GetAlbumsByUser(ctx context.Context, request *trip_service.ByUserRequest) (*trip_service.Albums, error) {
	responce, err := s.tripUsecase.AlbumsByUser(ctx, int(request.UserId))
	if err != nil {
		return nil, s.errorAdapter.AdaptError(err)
	}

	var protoAlbums trip_service.Albums
	for _, album := range *responce {
		protoAlbums.Albums = append(protoAlbums.Albums, &trip_service.Album{
			Id:          int64(album.Id),
			Title:       album.Title,
			TripId:      int64(album.TripId),
			Author:      int64(album.UserId),
			Description: album.Description,
			Photos:      album.Photos,
		})
	}
	return &protoAlbums, nil

}

func (s *tripDelivery) AddTripUser(ctx context.Context, request *trip_service.AddTripUserRequest) (*empty.Empty, error) {
	authorized, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.Author), int(request.TripId))
	if !authorized || err != nil {
		return &empty.Empty{}, errors.DeniedAccess
	}

	userIsAuthor, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.UserId), int(request.TripId))
	if userIsAuthor || err != nil {
		return &empty.Empty{}, errors.UserIsAlreadyAuthor
	}

	err = s.tripUsecase.AddTripUser(ctx, int(request.TripId), int(request.UserId))
	if err != nil {
		return &empty.Empty{}, s.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func (s *tripDelivery) ShareLink(ctx context.Context, request *trip_service.ShareRequest) (*trip_service.Link, error) {
	authorized, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.UserId), int(request.TripId))
	if !authorized || err != nil {
		return nil, errors.DeniedAccess
	}

	link := s.tripUsecase.ShareLink(ctx, int(request.TripId))

	return &trip_service.Link{Link: link}, nil
}

func (s *tripDelivery) AddUserByLink(ctx context.Context, request *trip_service.AddByShareRequest) (*trip_service.Link, error) {
	authorized, err := s.tripUsecase.CheckTripAuthor(ctx, int(request.UserId), int(request.TripId))
	id := strconv.Itoa(int(request.TripId))
	if err != nil {
		return nil, err
	}

	if authorized {
		return &trip_service.Link{
			Link: cnst.TripPostURL + "/" + id,
		}, nil
	}

	if !s.tripUsecase.CheckLink(ctx, request.Uuid, int(request.TripId)) {
		return nil, errors.DeniedAccess
	}

	err = s.tripUsecase.AddTripUser(ctx, int(request.TripId), int(request.UserId))
	if err != nil {
		return nil, err
	}

	return &trip_service.Link{
		Link: cnst.TripPostURL + "/" + id,
	}, nil
}

func ProtoDaysFromPlaces(places []models.Place) []*trip_service.Sight {
	var protoDays []*trip_service.Sight
	for _, sight := range places {
		var tags []*trip_service.Tag
		for _, tag := range sight.Tags {
			tags = append(tags, &trip_service.Tag{
				Id:   int64(tag.Id),
				Name: tag.Name,
			})
		}

		protoSight := trip_service.Sight{
			Id:          int64(sight.Id),
			Name:        sight.Name,
			Country:     sight.Country,
			Rating:      sight.Rating,
			Tags:        tags,
			Description: sight.Description,
			Photos:      sight.Photos,
			Day:         int64(sight.Day),
		}
		protoDays = append(protoDays, &protoSight)
	}
	return protoDays
}

func PlacesFromProtoDays(sights []*trip_service.Sight) []models.Place {
	var places []models.Place
	for _, sight := range sights {
		var tags []models.Tag
		for _, tag := range sight.Tags {
			tags = append(tags, models.Tag{
				Id:   int(tag.Id),
				Name: tag.Name,
			})
		}

		placesSight := models.Place{
			Id:          int(sight.Id),
			Name:        sight.Name,
			Country:     sight.Country,
			Rating:      sight.Rating,
			Tags:        tags,
			Description: sight.Description,
			Photos:      sight.Photos,
			Day:         int(sight.Day),
		}
		places = append(places, placesSight)
	}
	return places
}

func ProtoAlbumsFromAlbums(albums []models.Album) []*trip_service.Album {
	var protoAlbums []*trip_service.Album
	for _, album := range albums {
		protoAlbum := trip_service.Album{
			Id:          int64(album.Id),
			Title:       album.Title,
			Description: album.Description,
			Photos:      album.Photos,
		}
		protoAlbums = append(protoAlbums, &protoAlbum)
	}
	return protoAlbums
}

func AlbumsFromProtoAlbums(protoAlbums []*trip_service.Album) []models.Album {
	var albums []models.Album
	for _, album := range protoAlbums {
		modelAlbum := models.Album{
			Id:          int(album.Id),
			Title:       album.Title,
			Description: album.Description,
			Photos:      album.Photos,
		}
		albums = append(albums, modelAlbum)
	}
	return albums
}
