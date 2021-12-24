package delivery

import (
	"context"
	"snakealive/m/internal/services/review/models"
	"snakealive/m/internal/services/review/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/errors"
	review_service "snakealive/m/pkg/services/review"

	"github.com/golang/protobuf/ptypes/empty"
)

type reviewDelivery struct {
	reviewUsecase usecase.ReviewUseCase
	errorAdapter  error_adapter.ErrorAdapter
	review_service.UnimplementedReviewServiceServer
}

func NewReviewDelivery(reviewUsecase usecase.ReviewUseCase, errorAdapter error_adapter.ErrorAdapter) review_service.ReviewServiceServer {
	return &reviewDelivery{
		reviewUsecase: reviewUsecase,
		errorAdapter:  errorAdapter,
	}
}

func (r *reviewDelivery) ReviewByPlace(ctx context.Context, request *review_service.ReviewRequest) (*review_service.Reviews, error) {
	reviews, err := r.reviewUsecase.GetReviewsListByPlaceId(ctx, int(request.PlaceId), int(request.Limit),
		int(request.Skip))

	if err != nil {
		return nil, err
	}

	var protoReviews []*review_service.Review
	var protoReview *review_service.Review
	for _, review := range *reviews {
		protoReview = &review_service.Review{
			Id:     int64(review.Id),
			Title:  review.Title,
			Text:   review.Text,
			Rating: int64(review.Rating),
			UserId: int64(review.UserId),
		}
		protoReviews = append(protoReviews, protoReview)
	}

	return &review_service.Reviews{
		Reviews: protoReviews,
	}, err
}

func (r *reviewDelivery) Add(ctx context.Context, request *review_service.AddReviewRequest) (*review_service.Review, error) {
	ind, err := r.reviewUsecase.Add(ctx, &models.Review{
		Title:   request.Review.Title,
		Text:    request.Review.Text,
		Rating:  int(request.Review.Rating),
		PlaceId: int(request.Review.PlaceId),
	}, int(request.UserId))
	if err != nil {
		return nil, err
	}

	review, err := r.reviewUsecase.Get(ctx, ind)
	if err != nil {
		return nil, err
	}

	return &review_service.Review{
		Id:     int64(review.Id),
		Title:  review.Title,
		Text:   review.Text,
		Rating: int64(review.Rating),
		UserId: int64(review.UserId),
	}, err

}
func (r *reviewDelivery) Delete(ctx context.Context, request *review_service.DeleteReviewRequest) (*empty.Empty, error) {
	authorized := r.reviewUsecase.CheckAuthor(ctx, int(request.UserId), int(request.ReviewId))
	if !authorized {
		return &empty.Empty{}, errors.DeniedAccess
	}

	err := r.reviewUsecase.Delete(ctx, int(request.ReviewId))
	if err != nil {
		return nil, r.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func (r *reviewDelivery) GetAmountOfReviewsBySight(ctx context.Context, request *review_service.AmountRequest) (*review_service.Amount, error) {
	amount, err := r.reviewUsecase.GetReviewsAmount(ctx, int(request.Id))
	if err != nil {
		return nil, r.errorAdapter.AdaptError(err)
	}

	return &review_service.Amount{
		Amount: int64(amount),
	}, nil
}
