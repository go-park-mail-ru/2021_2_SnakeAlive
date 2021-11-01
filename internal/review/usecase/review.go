package reviewUseCase

import (
	"encoding/json"
	"log"
	"snakealive/m/pkg/domain"

	"github.com/valyala/fasthttp"
)

func NewReviewUseCase(reviewStorage domain.ReviewStorage) domain.ReviewUseCase {
	return reviewUseCase{reviewStorage: reviewStorage}
}

type reviewUseCase struct {
	reviewStorage domain.ReviewStorage
}

func (u reviewUseCase) Add(review domain.Review, user domain.User) (int, error) {
	err := u.reviewStorage.Add(review, user.Id)
	if err != nil {
		return fasthttp.StatusBadRequest, err
	}
	return fasthttp.StatusOK, err
}

func (u reviewUseCase) Get(id int) (domain.Review, error) {
	return u.reviewStorage.Get(id)
}

func (u reviewUseCase) GetReviewsListByPlaceId(id int) (int, []byte) {
	response, err := u.reviewStorage.GetListByPlace(id)
	if err != nil {
		log.Print("reviews not found", err)
		return fasthttp.StatusNotFound, []byte("{}")
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
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