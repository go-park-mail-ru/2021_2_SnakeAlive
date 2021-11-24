package usecase

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	sight_service "snakealive/m/pkg/services/sight"
	trip_service "snakealive/m/pkg/services/trip"
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
	UpdateAlbum(ctx context.Context, in *trip_service.ModifyAlbumRequest, opts ...grpc.CallOption) (*trip_service.Album, error)
	DeleteAlbum(ctx context.Context, in *trip_service.AlbumRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UploadPhoto(ctx context.Context, in *trip_service.UploadRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SightsByTrip(ctx context.Context, in *trip_service.SightsRequest, opts ...grpc.CallOption) (*trip_service.Sights, error)
}