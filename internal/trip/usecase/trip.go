package TripUseCase

import (
	"encoding/json"
	"snakealive/m/internal/domain"
	logs "snakealive/m/pkg/logger"

	"github.com/asaskevich/govalidator"
	"github.com/microcosm-cc/bluemonday"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewTripUseCase(tripStorage domain.TripStorage, l *logs.Logger) domain.TripUseCase {
	return tripUsecase{
		tripStorage: tripStorage,
		l:           l,
	}
}

type tripUsecase struct {
	tripStorage domain.TripStorage
	l           *logs.Logger
}

func (u tripUsecase) Add(value domain.Trip, user domain.User) (int, error) {
	_, err := govalidator.ValidateStruct(value)
	if err != nil {
		u.l.Logger.Error("error while validating trip struct")
		return 0, err
	}
	cleanTrip := u.SanitizeTrip(value)

	return u.tripStorage.Add(cleanTrip, user)
}

func (u tripUsecase) GetById(id int) (int, []byte) {
	value, err := u.tripStorage.GetById(id)
	if err != nil {
		u.l.Logger.Error("error while getting trip: ", zap.Error(err))
		return fasthttp.StatusBadRequest, nil
	}

	if value.Id == 0 {
		u.l.Logger.Error("trip not found")
		return fasthttp.StatusNotFound, nil
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		return fasthttp.StatusBadRequest, nil
	}

	return fasthttp.StatusOK, bytes
}

func (u tripUsecase) Update(id int, updatedTrip domain.Trip) error {
	_, err := govalidator.ValidateStruct(updatedTrip)
	if err != nil {
		u.l.Logger.Error("error while validating trip struct")
		return err
	}
	cleanTrip := u.SanitizeTrip(updatedTrip)

	return u.tripStorage.Update(id, cleanTrip)
}

func (u tripUsecase) Delete(id int) error {
	return u.tripStorage.Delete(id)
}

func (u tripUsecase) CheckAuthor(user domain.User, id int) bool {
	author := u.tripStorage.GetTripAuthor(id)
	return author == user.Id
}

func (u tripUsecase) SanitizeTrip(trip domain.Trip) domain.Trip {
	sanitizer := bluemonday.UGCPolicy()

	trip.Title = sanitizer.Sanitize(trip.Title)
	trip.Description = sanitizer.Sanitize(trip.Description)
	return trip
}
