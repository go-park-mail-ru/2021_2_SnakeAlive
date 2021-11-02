package reviewUseCase

import (
	"encoding/json"
	logs "snakealive/m/internal/logger"
	"snakealive/m/pkg/domain"

	"github.com/microcosm-cc/bluemonday"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewReviewUseCase(reviewStorage domain.ReviewStorage) domain.ReviewUseCase {
	return reviewUseCase{reviewStorage: reviewStorage}
}

type reviewUseCase struct {
	reviewStorage domain.ReviewStorage
}

func (u reviewUseCase) Add(review domain.Review, user domain.User) (int, error) {
	cleanReview := u.SanitizeReview(review)
	err := u.reviewStorage.Add(cleanReview, user.Id)
	if err != nil {
		return fasthttp.StatusBadRequest, err
	}
	return fasthttp.StatusOK, err
}

func (u reviewUseCase) Get(id int) (domain.Review, error) {
	return u.reviewStorage.Get(id)
}

func (u reviewUseCase) GetReviewsListByPlaceId(id int) (int, []byte) {
	logger := logs.GetLogger()

	response, err := u.reviewStorage.GetListByPlace(id)
	if err != nil {
		logger.Error("reviews not found: ", zap.Error(err))
		return fasthttp.StatusNotFound, []byte("{}")
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("error while marshalling JSON: ", zap.Error(err))
		return fasthttp.StatusOK, []byte("{}")
	}
	return fasthttp.StatusOK, bytes
}

func (u reviewUseCase) Delete(id int) error {
	return u.reviewStorage.Delete(id)
}

func (u reviewUseCase) CheckAuthor(user domain.User, id int) bool {
	author := u.reviewStorage.GetReviewAuthor(id)
	return author == user.Id
}

func (u reviewUseCase) SanitizeReview(review domain.Review) domain.Review {
	sanitizer := bluemonday.UGCPolicy()

	review.Title = sanitizer.Sanitize(review.Title)
	review.Text = sanitizer.Sanitize(review.Text)
	return review
}
