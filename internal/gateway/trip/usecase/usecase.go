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
	Add(ctx context.Context, value *models.Trip, userID int) (*models.Trip, error)
	GetById(ctx context.Context, id int, userID int) (*models.Trip, error)
	Delete(ctx context.Context, id int, userID int) error
	Update(ctx context.Context, id int, updatedTrip *models.Trip, userID int) (*models.Trip, error)
}

type tripGRPC interface {
	GetTrip(ctx context.Context, in *trip_service.TripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	AddTrip(ctx context.Context, in *trip_service.ModifyTripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	Update(ctx context.Context, in *trip_service.ModifyTripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	Delete(ctx context.Context, in *trip_service.TripRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tripGatewayUseCase struct {
	tripGRPC tripGRPC
}

func NewTripGatewayUseCase(grpc tripGRPC) TripGatewayUseCase {
	return &tripGatewayUseCase{tripGRPC: grpc}
}

func (u *tripGatewayUseCase) Add(ctx context.Context, value *models.Trip, userID int) (*models.Trip, error) {
	days := td.ProtoDaysFromPlaces(value.Days)
	trip := &trip_service.Trip{
		Id:          int64(value.Id),
		Title:       value.Title,
		Description: value.Description,
		Days:        days,
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

	places := td.PlacesFromProtoDays(responce.Days)

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Days:        places,
	}, nil
}

func (u *tripGatewayUseCase) GetById(ctx context.Context, tripId int, userID int) (*models.Trip, error) {
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

	places := td.PlacesFromProtoDays(responce.Days)

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Days:        places,
	}, nil
}

func (u *tripGatewayUseCase) Delete(ctx context.Context, id int, userID int) error {
	_, err := u.tripGRPC.Delete(ctx, &trip_service.TripRequest{
		TripId: int64(id),
		UserId: int64(userID),
	})
	return err
}

func (u *tripGatewayUseCase) Update(ctx context.Context, id int, updatedTrip *models.Trip, userID int) (*models.Trip, error) {
	days := td.ProtoDaysFromPlaces(updatedTrip.Days)
	trip := &trip_service.Trip{
		Id:          int64(updatedTrip.Id),
		Title:       updatedTrip.Title,
		Description: updatedTrip.Description,
		Days:        days,
	}

	responce, err := u.tripGRPC.Update(ctx,
		&trip_service.ModifyTripRequest{
			Trip:   trip,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	places := td.PlacesFromProtoDays(responce.Days)

	return &models.Trip{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Days:        places,
	}, nil
}
