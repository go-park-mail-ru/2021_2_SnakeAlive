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

func (u placeUsecase) GetPlacesByCountry(value string) ([]byte, error) {
	places, err := u.placeStorage.GetPlacesByCountry(value)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := json.Marshal(places)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
	}
	return bytes, err
}

func (u placeUsecase) GetPlaceListByName(param string) (int, []byte) {
	if _, found := u.Get(param); !found {
		log.Printf("country not found")
		return fasthttp.StatusNotFound, []byte("{}")
	}

	response, _ := u.Get(param)
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		return fasthttp.StatusOK, []byte("{}")
	}
	return fasthttp.StatusOK, bytes
}
