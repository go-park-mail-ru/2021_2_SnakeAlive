package usecase

import (
	"context"
	td "snakealive/m/internal/services/trip/delivery"
	"snakealive/m/internal/services/trip/models"
	"snakealive/m/pkg/errors"
	trip_service "snakealive/m/pkg/services/trip"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type TripGatewayUseCase interface {
	AddTrip(ctx context.Context, value *models.Trip, userID int) (*models.Trip, error)
	GetTripById(ctx context.Context, id int, userID int) (*models.Trip, error)
	DeleteTrip(ctx context.Context, id int, userID int) error
	UpdateTrip(ctx context.Context, id int, updatedTrip *models.Trip, userID int) (*models.Trip, error)

	AddAlbum(ctx context.Context, album *models.Album, userID int) (*models.Album, error)
	GetAlbumById(ctx context.Context, id int, userID int) (*models.Album, error)
	DeleteAlbum(ctx context.Context, id int, userID int) error
	UpdateAlbum(ctx context.Context, id int, updatedAlbum *models.Album, userID int) (*models.Album, error)
	UploadPhoto(ctx context.Context, filename string, userID int, id int) error
}

type tripGRPC interface {
	GetTrip(ctx context.Context, in *trip_service.TripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	AddTrip(ctx context.Context, in *trip_service.ModifyTripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	UpdateTrip(ctx context.Context, in *trip_service.ModifyTripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	DeleteTrip(ctx context.Context, in *trip_service.TripRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAlbum(ctx context.Context, in *trip_service.AlbumRequest, opts ...grpc.CallOption) (*trip_service.Album, error)
	AddAlbum(ctx context.Context, in *trip_service.ModifyAlbumRequest, opts ...grpc.CallOption) (*trip_service.Album, error)
	UpdateAlbum(ctx context.Context, in *trip_service.ModifyAlbumRequest, opts ...grpc.CallOption) (*trip_service.Album, error)
	DeleteAlbum(ctx context.Context, in *trip_service.AlbumRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UploadPhoto(ctx context.Context, in *trip_service.UploadRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tripGatewayUseCase struct {
	tripGRPC tripGRPC
}

func NewTripGatewayUseCase(grpc tripGRPC) TripGatewayUseCase {
	return &tripGatewayUseCase{tripGRPC: grpc}
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

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
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

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
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

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
	}, nil
}

func (u *tripGatewayUseCase) AddAlbum(ctx context.Context, album *models.Album, userID int) (*models.Album, error) {
	protoAlbum := &trip_service.Album{
		Id:          int64(album.Id),
		Title:       album.Title,
		TripId:      int64(album.TripId),
		Description: album.Description,
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

func (u *tripGatewayUseCase) UploadPhoto(ctx context.Context, filename string, userID int, id int) error {
	_, err := u.tripGRPC.UploadPhoto(ctx, &trip_service.UploadRequest{
		AlbumId:  int64(id),
		UserId:   int64(userID),
		Filename: filename,
	})
	return err
}
