package usecase

import (
	"context"

	td "snakealive/m/internal/services/trip/delivery"
	"snakealive/m/internal/services/trip/models"
	"snakealive/m/pkg/errors"
	sight_service "snakealive/m/pkg/services/sight"
	trip_service "snakealive/m/pkg/services/trip"
)

type TripGatewayUseCase interface {
	AddTrip(ctx context.Context, value *models.Trip, userID int) (*models.Trip, error)
	GetTripById(ctx context.Context, id int, userID int) (*models.Trip, error)
	DeleteTrip(ctx context.Context, id int, userID int) error
	UpdateTrip(ctx context.Context, id int, updatedTrip *models.Trip, userID int) (*models.Trip, error)
	GetTripsByUser(ctx context.Context, id int) (*[]models.Trip, error)

	AddAlbum(ctx context.Context, album *models.Album, userID int) (*models.Album, error)
	GetAlbumById(ctx context.Context, id int, userID int) (*models.Album, error)
	DeleteAlbum(ctx context.Context, id int, userID int) error
	UpdateAlbum(ctx context.Context, id int, updatedAlbum *models.Album, userID int) (*models.Album, error)
	SightsByTrip(ctx context.Context, id int) (*[]models.TripSight, error)
	GetAlbumsByUser(ctx context.Context, id int) (*[]models.Album, error)
}

type tripGatewayUseCase struct {
	tripGRPC  tripGRPC
	sightGRPC sightGRPC
}

func NewTripGatewayUseCase(grpc tripGRPC, sightGRPC sightGRPC) TripGatewayUseCase {
	return &tripGatewayUseCase{
		tripGRPC:  grpc,
		sightGRPC: sightGRPC,
	}
}

func (u *tripGatewayUseCase) AddTrip(ctx context.Context, value *models.Trip, userID int) (*models.Trip, error) {
	days := td.ProtoDaysFromPlaces(value.Sights)
	trip := &trip_service.Trip{
		Id:          int64(value.Id),
		Title:       value.Title,
		Description: value.Description,
		Sights:      days,
	}

	responce, err := u.tripGRPC.AddTrip(ctx,
		&trip_service.ModifyTripRequest{
			Trip:   trip,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	places := td.PlacesFromProtoDays(responce.Sights)
	albums := td.AlbumsFromProtoAlbums(responce.Albums)

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
		Albums:      albums,
	}, nil
}

func (u *tripGatewayUseCase) GetTripById(ctx context.Context, tripId int, userID int) (*models.Trip, error) {
	responce, err := u.tripGRPC.GetTrip(ctx, &trip_service.TripRequest{
		TripId: int64(tripId),
		UserId: int64(userID),
	})
	if err != nil {
		return nil, err
	}

	if responce.Id == 0 {
		return nil, errors.TripNotFound
	}

	places := td.PlacesFromProtoDays(responce.Sights)
	albums := td.AlbumsFromProtoAlbums(responce.Albums)

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
		Albums:      albums,
	}, nil
}

func (u *tripGatewayUseCase) DeleteTrip(ctx context.Context, id int, userID int) error {
	_, err := u.tripGRPC.DeleteTrip(ctx, &trip_service.TripRequest{
		TripId: int64(id),
		UserId: int64(userID),
	})
	return err
}

func (u *tripGatewayUseCase) UpdateTrip(ctx context.Context, id int, updatedTrip *models.Trip, userID int) (*models.Trip, error) {
	days := td.ProtoDaysFromPlaces(updatedTrip.Sights)
	trip := &trip_service.Trip{
		Id:          int64(id),
		Title:       updatedTrip.Title,
		Description: updatedTrip.Description,
		Sights:      days,
	}

	responce, err := u.tripGRPC.UpdateTrip(ctx,
		&trip_service.ModifyTripRequest{
			Trip:   trip,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	places := td.PlacesFromProtoDays(responce.Sights)
	albums := td.AlbumsFromProtoAlbums(responce.Albums)

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
		Albums:      albums,
	}, nil
}

func (u *tripGatewayUseCase) AddAlbum(ctx context.Context, album *models.Album, userID int) (*models.Album, error) {
	protoAlbum := &trip_service.Album{
		Id:          int64(album.Id),
		Title:       album.Title,
		TripId:      int64(album.TripId),
		Description: album.Description,
		Photos:      album.Photos,
	}

	responce, err := u.tripGRPC.AddAlbum(ctx,
		&trip_service.ModifyAlbumRequest{
			Album:  protoAlbum,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	return &models.Album{
		Id:          int(responce.Id),
		Title:       responce.Title,
		TripId:      int(responce.TripId),
		UserId:      int(responce.Author),
		Description: responce.Description,
	}, nil
}

func (u *tripGatewayUseCase) GetAlbumById(ctx context.Context, id int, userID int) (*models.Album, error) {
	responce, err := u.tripGRPC.GetAlbum(ctx, &trip_service.AlbumRequest{
		AlbumId: int64(id),
		UserId:  int64(userID),
	})
	if err != nil {
		return nil, err
	}

	if responce.Id == 0 {
		return nil, errors.TripNotFound
	}

	return &models.Album{
		Id:          int(responce.Id),
		Title:       responce.Title,
		TripId:      int(responce.TripId),
		UserId:      int(responce.Author),
		Description: responce.Description,
		Photos:      responce.Photos,
	}, nil
}

func (u *tripGatewayUseCase) DeleteAlbum(ctx context.Context, id int, userID int) error {
	_, err := u.tripGRPC.DeleteAlbum(ctx, &trip_service.AlbumRequest{
		AlbumId: int64(id),
		UserId:  int64(userID),
	})
	return err
}

func (u *tripGatewayUseCase) UpdateAlbum(ctx context.Context, id int, updatedAlbum *models.Album, userID int) (*models.Album, error) {
	album := &trip_service.Album{
		Id:          int64(id),
		Title:       updatedAlbum.Title,
		Description: updatedAlbum.Description,
		Photos:      updatedAlbum.Photos,
	}

	responce, err := u.tripGRPC.UpdateAlbum(ctx,
		&trip_service.ModifyAlbumRequest{
			Album:  album,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	return &models.Album{
		Id:          int(responce.Id),
		Title:       responce.Title,
		TripId:      int(responce.TripId),
		UserId:      int(responce.Author),
		Description: responce.Description,
		Photos:      responce.Photos,
	}, nil
}

func (u *tripGatewayUseCase) SightsByTrip(ctx context.Context, id int) (*[]models.TripSight, error) {
	ids, err := u.tripGRPC.SightsByTrip(ctx, &trip_service.SightsRequest{
		TripId: int64(id),
	})
	if err != nil {
		return nil, err
	}

	sights, err := u.sightGRPC.GetSightsByIDs(ctx, &sight_service.GetSightsByIDsRequest{Ids: ids.Ids})
	if err != nil {
		return nil, err
	}

	adapted := make([]models.TripSight, len(sights.Sights))
	for i, sight := range sights.Sights {
		adapted[i] = models.TripSight{
			Id:  int(sight.Id),
			Lng: sight.Lng,
			Lat: sight.Lat,
		}
	}

	return &adapted, nil
}

func (u *tripGatewayUseCase) GetTripsByUser(ctx context.Context, id int) (*[]models.Trip, error) {
	protoTrips, err := u.tripGRPC.GetTripsByUser(ctx, &trip_service.ByUserRequest{UserId: int64(id)})
	if err != nil {
		return nil, err
	}

	var trips []models.Trip
	for _, trip := range protoTrips.Trips {
		places := td.PlacesFromProtoDays(trip.Sights)
		albums := td.AlbumsFromProtoAlbums(trip.Albums)
		trips = append(trips, models.Trip{
			Id:          int(trip.Id),
			Title:       trip.Title,
			Description: trip.Description,
			Albums:      albums,
			Sights:      places,
		})
	}
	return &trips, nil
}

func (u *tripGatewayUseCase) GetAlbumsByUser(ctx context.Context, id int) (*[]models.Album, error) {
	protoAlbums, err := u.tripGRPC.GetAlbumsByUser(ctx, &trip_service.ByUserRequest{UserId: int64(id)})
	if err != nil {
		return nil, err
	}

	var albums []models.Album
	for _, album := range protoAlbums.Albums {
		albums = append(albums, models.Album{
			Id:          int(album.Id),
			Title:       album.Title,
			TripId:      int(album.TripId),
			UserId:      int(album.Author),
			Description: album.Description,
			Photos:      album.Photos,
		})
	}
	return &albums, nil
}
