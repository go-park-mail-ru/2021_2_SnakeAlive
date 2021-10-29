package placeUseCase

import (
	"snakealive/m/pkg/domain"
)

type PlaceUseCase interface {
	Get(key string) (domain.Places, bool)
}

func NewPlaceUseCase(placeStorage domain.PlaceStorage) PlaceUseCase {
	return placeUsecase{placeStorage: placeStorage}
}

type placeUsecase struct {
	placeStorage domain.PlaceStorage
}

func (u placeUsecase) Get(key string) (domain.Places, bool) {
	return u.placeStorage.Get(key)
}
