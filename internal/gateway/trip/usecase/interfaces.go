package usecase

import (
	"context"

	sight_service "snakealive/m/pkg/services/sight"
	trip_service "snakealive/m/pkg/services/trip"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type sightGRPC interface {
	GetSightsByIDs(ctx context.Context, in *sight_service.GetSightsByIDsRequest, opts ...grpc.CallOption) (*sight_service.GetSightsByIDsResponse, error)
}

type tripGRPC interface {
	GetTrip(ctx context.Context, in *trip_service.TripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	AddTrip(ctx context.Context, in *trip_service.ModifyTripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	UpdateTrip(ctx context.Context, in *trip_service.ModifyTripRequest, opts ...grpc.CallOption) (*trip_service.Trip, error)
	DeleteTrip(ctx context.Context, in *trip_service.TripRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAlbum(ctx context.Context, in *trip_service.AlbumRequest, opts ...grpc.CallOption) (*trip_service.Album, error)
	AddAlbum(ctx context.Context, in *trip_service.ModifyAlbumRequest, opts ...grpc.CallOption) (*trip_service.Album, error)
	GetTripsByUser(ctx context.Context, in *trip_service.TripByUserRequest, opts ...grpc.CallOption) (*trip_service.Trips, error)
	UpdateAlbum(ctx context.Context, in *trip_service.ModifyAlbumRequest, opts ...grpc.CallOption) (*trip_service.Album, error)
	DeleteAlbum(ctx context.Context, in *trip_service.AlbumRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SightsByTrip(ctx context.Context, in *trip_service.SightsRequest, opts ...grpc.CallOption) (*trip_service.Sights, error)
}
