package placeUseCase

import (
	"snakealive/m/domain"
)

type PlaceUseCase interface {
	Get(key string) ([]domain.Place, bool)
}

func NewPlaceUseCase(placeStorage domain.PlaceStorage) PlaceUseCase {
	return placeUsecase{placeStorage: placeStorage}
}

type placeUsecase struct {
	placeStorage domain.PlaceStorage
}

func (u placeUsecase) Get(key string) ([]domain.Place, bool) {
	return u.placeStorage.Get(key)
}
