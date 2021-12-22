package usecase

import (
	"context"

	"snakealive/m/internal/models"
	review_service "snakealive/m/pkg/services/review"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type ReviewGatewayUseCase interface {
	Add(ctx context.Context, review *models.Review, userID int) (*models.Review, error)
	Delete(ctx context.Context, id int, userID int) error
	GetReviewsListByPlaceId(ctx context.Context, id int, limit int, skip int) (*[]models.Review, error)
}

type reviewGRPC interface {
	ReviewByPlace(ctx context.Context, in *review_service.ReviewRequest, opts ...grpc.CallOption) (*review_service.Reviews, error)
	Add(ctx context.Context, in *review_service.AddReviewRequest, opts ...grpc.CallOption) (*review_service.Review, error)
	Delete(ctx context.Context, in *review_service.DeleteReviewRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type reviewGatewayUseCase struct {
	reviewGRPC reviewGRPC
}

func NewReviewGatewayUseCase(grpc reviewGRPC) ReviewGatewayUseCase {
	return &reviewGatewayUseCase{reviewGRPC: grpc}
}

func (u *reviewGatewayUseCase) Add(ctx context.Context, review *models.Review, userID int) (*models.Review, error) {
	protoReview := &review_service.Review{
		Title:   review.Title,
		Text:    review.Text,
		Rating:  int64(review.Rating),
		UserId:  int64(userID),
		PlaceId: int64(review.PlaceId),
	}

	addedReview, err := u.reviewGRPC.Add(ctx, &review_service.AddReviewRequest{
		Review: protoReview,
		UserId: int64(userID),
	})
	if err != nil {
		return nil, err
	}

	return &models.Review{
		Id:     int(addedReview.Id),
		Title:  addedReview.Title,
		Text:   addedReview.Text,
		Rating: int(addedReview.Rating),
		UserId: int(addedReview.UserId),
	}, nil
}

func (u *reviewGatewayUseCase) Delete(ctx context.Context, id int, userID int) error {
	_, err := u.reviewGRPC.Delete(ctx, &review_service.DeleteReviewRequest{
		UserId:   int64(userID),
		ReviewId: int64(id),
	})
	return err
}

func (u *reviewGatewayUseCase) GetReviewsListByPlaceId(ctx context.Context, id int, limit int, skip int) (*[]models.Review, error) {
	reviews, err := u.reviewGRPC.ReviewByPlace(ctx, &review_service.ReviewRequest{
		PlaceId: int64(id),
		Limit:   int64(limit),
		Skip:    int64(skip),
	})
	if err != nil {
		return nil, err
	}

	var modelReview models.Review
	var modelReviews []models.Review

	for _, review := range reviews.Reviews {
		modelReview = models.Review{
			Id:     int(review.Id),
			Title:  review.Title,
			Text:   review.Text,
			Rating: int(review.Rating),
			UserId: int(review.UserId),
		}
		modelReviews = append(modelReviews, modelReview)
	}
	return &modelReviews, nil
}
