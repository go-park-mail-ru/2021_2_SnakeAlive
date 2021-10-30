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

func (u placeUsecase) Get(key string) (domain.Places, bool) {
	return u.placeStorage.Get(key)
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
