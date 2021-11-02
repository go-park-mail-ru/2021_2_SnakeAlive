package placeUseCase

import (
	"encoding/json"
	logs "snakealive/m/internal/logger"
	"snakealive/m/pkg/domain"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewPlaceUseCase(placeStorage domain.PlaceStorage) domain.PlaceUseCase {
	return placeUsecase{placeStorage: placeStorage}
}

type placeUsecase struct {
	placeStorage domain.PlaceStorage
}

func (u placeUsecase) GetById(id int) (value domain.Place, err error) {
	return u.placeStorage.GetById(id)
}

func (u placeUsecase) GetSight(sight domain.Place) (int, []byte) {
	logger := logs.GetLogger()

	response, err := json.Marshal(sight)
	if err != nil {
		logger.Error("error while marshalling JSON: ", zap.Error(err))
		return fasthttp.StatusBadRequest, []byte("{}")
	}
	return fasthttp.StatusOK, response
}

func (u placeUsecase) GetPlacesByCountry(value string) ([]byte, error) {
	logger := logs.GetLogger()

	places, err := u.placeStorage.GetPlacesByCountry(value)
	if err != nil {
		logger.Error("error while getting places by country")
		return []byte{}, err
	}

	bytes, err := json.Marshal(places)
	if err != nil {
		logger.Error("error while marshalling JSON: ", zap.Error(err))
	}
	return bytes, err
}
