package placeUseCase

import (
	"encoding/json"
	"snakealive/m/internal/domain"
	logs "snakealive/m/pkg/logger"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewPlaceUseCase(placeStorage domain.PlaceStorage, l *logs.Logger) domain.PlaceUseCase {
	return placeUsecase{
		placeStorage: placeStorage,
		l:            l,
	}
}

type placeUsecase struct {
	placeStorage domain.PlaceStorage
	l            *logs.Logger
}

func (u placeUsecase) GetById(id int) (value domain.Place, err error) {
	return u.placeStorage.GetById(id)
}

func (u placeUsecase) GetSight(sight domain.Place) (int, []byte) {
	response, err := json.Marshal(sight)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		return fasthttp.StatusBadRequest, []byte("{}")
	}
	return fasthttp.StatusOK, response
}

func (u placeUsecase) GetPlacesByCountry(value string) ([]byte, error) {
	places, err := u.placeStorage.GetPlacesByCountry(value)
	if err != nil {
		u.l.Logger.Error("error while getting places by country")
		return []byte{}, err
	}

	bytes, err := json.Marshal(places)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
	}
	return bytes, err
}
