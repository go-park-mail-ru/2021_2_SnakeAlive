package reviewUseCase

import (
	"encoding/json"
	"fmt"
	"snakealive/m/internal/domain"
	logs "snakealive/m/pkg/logger"

	"github.com/microcosm-cc/bluemonday"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewReviewUseCase(reviewStorage domain.ReviewStorage, userStorage domain.UserStorage, l *logs.Logger) domain.ReviewUseCase {
	return reviewUseCase{
		reviewStorage: reviewStorage,
		userStorage:   userStorage,
		l:             l}
}

type reviewUseCase struct {
	reviewStorage domain.ReviewStorage
	userStorage   domain.UserStorage
	l             *logs.Logger
}

type ReviewUser struct {
	Review domain.ReviewNoPlace `json:"review"`
	User   domain.PublicUser    `json:"user"`
	Owner  bool                 `json:"owner"`
}

func (u reviewUseCase) Add(review domain.Review, user domain.User) (int, []byte, error) {
	cleanReview := u.SanitizeReview(review)
	insertedId, err := u.reviewStorage.Add(cleanReview, user.Id)
	if err != nil {
		return fasthttp.StatusBadRequest, []byte("{}"), err
	}

	bytes, err := json.Marshal(insertedId)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		return fasthttp.StatusOK, []byte("{}"), err
	}
	return fasthttp.StatusOK, bytes, err
}

func (u reviewUseCase) Get(id int) (domain.Review, error) {
	return u.reviewStorage.Get(id)
}

func (u reviewUseCase) GetReviewsListByPlaceId(id int, user domain.User, limit int, skip int) (int, []byte) {
	user = u.SanitizeUser(user)
	reviews, err := u.reviewStorage.GetListByPlace(id, limit, skip)
	if err != nil {
		u.l.Logger.Error("reviews not found: ", zap.Error(err))
		return fasthttp.StatusNotFound, []byte("{}")
	}

	reviewData := make([]ReviewUser, 0)
	var ru ReviewUser

	for i := 0; i < len(reviews); i++ {
		ru.Review = reviews[i]
		ru.User, err = u.userStorage.GetPublicById(reviews[i].UserId)
		if err != nil {
			u.l.Logger.Error("error while get user in reviews list: ", zap.Error(err))
			return fasthttp.StatusOK, []byte("{}")
		}
		ru.Owner = u.CheckAuthor(user, reviews[i].Id)

		reviewData = append(reviewData, ru)
	}

	bytes, err := json.Marshal(reviewData)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		return fasthttp.StatusOK, []byte("{}")
	}
	fmt.Println("bytes= ", bytes)
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

func (u reviewUseCase) SanitizeUser(user domain.User) domain.User {
	sanitizer := bluemonday.UGCPolicy()

	user.Name = sanitizer.Sanitize(user.Name)
	user.Surname = sanitizer.Sanitize(user.Surname)
	user.Email = sanitizer.Sanitize(user.Email)
	user.Password = sanitizer.Sanitize(user.Password)
	return user
}
