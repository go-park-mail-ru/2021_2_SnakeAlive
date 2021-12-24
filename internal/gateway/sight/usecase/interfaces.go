package usecase

import (
	"context"

	review_service "snakealive/m/pkg/services/review"
	sight_service "snakealive/m/pkg/services/sight"

	"google.golang.org/grpc"
)

type reviewGRPC interface {
	GetAmountOfReviewsBySight(ctx context.Context, in *review_service.AmountRequest, opts ...grpc.CallOption) (*review_service.Amount, error)
}

type sightGRPC interface {
	GetSights(ctx context.Context, in *sight_service.GetSightsRequest, opts ...grpc.CallOption) (*sight_service.GetSightsReponse, error)
	GetSight(ctx context.Context, in *sight_service.GetSightRequest, opts ...grpc.CallOption) (*sight_service.GetSightResponse, error)
	SearchSights(ctx context.Context, in *sight_service.SearchSightRequest, opts ...grpc.CallOption) (*sight_service.SearchSightResponse, error)
	GetSightsByTag(ctx context.Context, in *sight_service.GetSightsByTagRequest, opts ...grpc.CallOption) (*sight_service.GetSightsByTagResponse, error)
	GetTags(ctx context.Context, in *sight_service.GetTagsRequest, opts ...grpc.CallOption) (*sight_service.GetTagsResponse, error)
	GetRandomTags(ctx context.Context, in *sight_service.GetTagsRequest, opts ...grpc.CallOption) (*sight_service.GetTagsResponse, error)
}
