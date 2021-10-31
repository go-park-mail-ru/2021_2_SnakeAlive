package TripUseCase

import (
	"encoding/json"
	"log"
	"snakealive/m/pkg/domain"

	"github.com/asaskevich/govalidator"
	"github.com/valyala/fasthttp"
)

func NewTripUseCase(tripStorage domain.TripStorage) domain.TripUseCase {
	return tripUsecase{tripStorage: tripStorage}
}

type tripUsecase struct {
	tripStorage domain.TripStorage
}

func (u tripUsecase) Add(value domain.Trip, user domain.User) (int, error) {
	_, err := govalidator.ValidateStruct(value)
	if err != nil {
		log.Printf("error while validating trip")
		return 0, err
	}

	return u.tripStorage.Add(value, user)
}

func (u tripUsecase) GetById(id int) (int, []byte) {
	value, err := u.tripStorage.GetById(id)
	if err != nil {
		log.Printf("error while getting trip: %s", err)
		return fasthttp.StatusBadRequest, nil
	}

	if value.Id == 0 {
		return fasthttp.StatusNotFound, nil
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		return fasthttp.StatusBadRequest, nil
	}

	return fasthttp.StatusOK, bytes
}

func (u tripUsecase) Update(id int, updatedTrip domain.Trip) error {
	_, err := govalidator.ValidateStruct(updatedTrip)
	if err != nil {
		log.Printf("error while validating trip")
		return err
	}

	return u.tripStorage.Update(id, updatedTrip)
}

func (u tripUsecase) Delete(id int) error {
	return u.tripStorage.Delete(id)
}

func (u tripUsecase) CheckAuthor(user domain.User, id int) bool {
	author := u.tripStorage.GetTripAuthor(id)
	return author == user.Id
}
