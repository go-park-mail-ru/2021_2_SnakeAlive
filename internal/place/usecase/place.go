package placeUseCase

import (
	"encoding/json"
	"log"
	"snakealive/m/pkg/domain"

	"github.com/valyala/fasthttp"
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
	response, err := json.Marshal(sight)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		return fasthttp.StatusBadRequest, []byte("{}")
	}
	return fasthttp.StatusOK, response
}
