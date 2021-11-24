package usecase

import (
	"context"

	"google.golang.org/grpc"
	sight_service "snakealive/m/pkg/services/sight"
)

type sightGRPC interface {
	GetSights(ctx context.Context, in *sight_service.GetSightsRequest, opts ...grpc.CallOption) (*sight_service.GetSightsReponse, error)
	GetSight(ctx context.Context, in *sight_service.GetSightRequest, opts ...grpc.CallOption) (*sight_service.GetSightResponse, error)
	SearchSights(ctx context.Context, in *sight_service.SearchSightRequest, opts ...grpc.CallOption) (*sight_service.SearchSightResponse, error)
}
