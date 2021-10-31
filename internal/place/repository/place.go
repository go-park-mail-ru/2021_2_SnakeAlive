package placeRepository

import (
	ent "snakealive/m/internal/entities"
	"snakealive/m/pkg/domain"
	"sync"
)

type PlaceStorage struct {
	dataHolder map[string]domain.Places
	mu         *sync.RWMutex
}

func NewPlaceStorage() domain.PlaceStorage {
	return &PlaceStorage{
		dataHolder: ent.PlacesDB,
		mu:         &sync.RWMutex{},
	}
}

func (u *PlaceStorage) Get(key string) (value domain.Places, exist bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	value, exist = u.dataHolder[key]
	return
}
