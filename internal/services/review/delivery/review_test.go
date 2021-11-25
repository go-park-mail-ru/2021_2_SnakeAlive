package delivery

import (
	"context"
	service_mocks "snakealive/m/internal/mocks"
	"snakealive/m/internal/services/review/models"
	review_usecase "snakealive/m/internal/services/review/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	review_service "snakealive/m/pkg/services/review"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestHandler_GetReviews(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	limit := 10
	skip := 0
	tripId := 1
	reviews := &[]models.Review{
		{Id: 1},
	}
	expectedReviews := &review_service.Reviews{
		Reviews: []*review_service.Review{
			{Id: 1},
		},
	}
	request := &review_service.ReviewRequest{
		PlaceId: 1,
		UserId:  1,
		Limit:   10,
		Skip:    0,
	}

	mockGetReviewsByTrip := func(r *service_mocks.MockReviewRepository, ctx context.Context, id int, limit int, skip int,
		reviews *[]models.Review) {
		r.EXPECT().GetListByPlace(ctx, id, limit, skip).Return(reviews, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	reviewRepo := service_mocks.NewMockReviewRepository(c)
	mockGetReviewsByTrip(reviewRepo, ctx, tripId, limit, skip, reviews)

	reviewUsecase := review_usecase.NewReviewUseCase(reviewRepo)
	reviewDelivery := NewReviewDelivery(reviewUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredReviews, err := reviewDelivery.ReviewByPlace(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedReviews, aquiredReviews)
}

func TestHandler_AddReview(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userId := 1
	tripId := 1
	review := &models.Review{
		Id:    1,
		Title: "Here",
	}
	expectedReview := &review_service.Review{
		Id:    1,
		Title: "Here",
	}
	request := &review_service.AddReviewRequest{
		UserId: 1,
		Review: expectedReview,
	}

	mockAddReview := func(r *service_mocks.MockReviewRepository, ctx context.Context, userId int, id int, review *models.Review) {
		r.EXPECT().Add(ctx, gomock.Any(), userId).Return(id, nil).AnyTimes()
	}
	mockGetReview := func(r *service_mocks.MockReviewRepository, ctx context.Context, id int, review *models.Review) {
		r.EXPECT().Get(ctx, id).Return(review, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	reviewRepo := service_mocks.NewMockReviewRepository(c)
	mockAddReview(reviewRepo, ctx, userId, tripId, review)
	mockGetReview(reviewRepo, ctx, tripId, review)

	reviewUsecase := review_usecase.NewReviewUseCase(reviewRepo)
	reviewDelivery := NewReviewDelivery(reviewUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredReview, err := reviewDelivery.Add(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedReview, aquiredReview)
}

func TestHandler_DeleteReview(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userId := 1
	reviewId := 1
	request := &review_service.DeleteReviewRequest{
		ReviewId: 1,
		UserId:   1,
	}

	mockDeleteReview := func(r *service_mocks.MockReviewRepository, ctx context.Context, id int) {
		r.EXPECT().Delete(ctx, id).Return(nil).AnyTimes()
	}
	mockGetReviewAuthor := func(r *service_mocks.MockReviewRepository, ctx context.Context, id int, userId int) {
		r.EXPECT().GetReviewAuthor(ctx, id).Return(userId).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	reviewRepo := service_mocks.NewMockReviewRepository(c)
	mockDeleteReview(reviewRepo, ctx, reviewId)
	mockGetReviewAuthor(reviewRepo, ctx, reviewId, userId)

	reviewUsecase := review_usecase.NewReviewUseCase(reviewRepo)
	reviewDelivery := NewReviewDelivery(reviewUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	_, err := reviewDelivery.Delete(ctx, request)

	assert.Nil(t, err)
}
