package usecase

import (
	"context"
	"snakealive/m/internal/services/review/models"
	"snakealive/m/internal/services/review/repository"

	"github.com/microcosm-cc/bluemonday"
)

type ReviewUseCase interface {
	Add(ctx context.Context, review *models.Review, userID int) (int, error)
	Get(ctx context.Context, id int) (*models.Review, error)
	Delete(ctx context.Context, id int) error
	GetReviewsListByPlaceId(ctx context.Context, id int, userID int, limit int, skip int) (*[]models.Review, error)
	CheckAuthor(ctx context.Context, userID int, id int) bool
	SanitizeReview(ctx context.Context, review *models.Review) *models.Review
}

type reviewUseCase struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewUseCase(reviewRepository repository.ReviewRepository) ReviewUseCase {
	return &reviewUseCase{reviewRepository: reviewRepository}
}

func (u reviewUseCase) Add(ctx context.Context, review *models.Review, userID int) (int, error) {
	cleanReview := u.SanitizeReview(ctx, review)
	return u.reviewRepository.Add(ctx, cleanReview, userID)
}

func (u reviewUseCase) Get(ctx context.Context, id int) (*models.Review, error) {
	return u.reviewRepository.Get(ctx, id)
}

func (u reviewUseCase) Delete(ctx context.Context, id int) error {
	return u.reviewRepository.Delete(ctx, id)
}

func (u reviewUseCase) GetReviewsListByPlaceId(ctx context.Context, id int, userID int, limit int, skip int) (*[]models.Review, error) {
	return u.reviewRepository.GetListByPlace(ctx, id, limit, skip)
}

func (u reviewUseCase) CheckAuthor(ctx context.Context, userID int, id int) bool {
	author := u.reviewRepository.GetReviewAuthor(ctx, id)
	return author == id
}

func (u reviewUseCase) SanitizeReview(ctx context.Context, review *models.Review) *models.Review {
	sanitizer := bluemonday.UGCPolicy()

	review.Title = sanitizer.Sanitize(review.Title)
	review.Text = sanitizer.Sanitize(review.Text)
	return review
}
