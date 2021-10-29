package placeRepository

import (
	ent "snakealive/m/internal/entities"
	"snakealive/m/pkg/domain"
	"sync"
)

type placeStorage struct {
	dataHolder map[string]domain.Places
	mu         *sync.RWMutex
}

func NewPlaceStorage() domain.PlaceStorage {
	return &placeStorage{
		dataHolder: ent.PlacesDB,
		mu:         &sync.RWMutex{},
	}
}

func (u *placeStorage) Get(key string) (value domain.Places, exist bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	value, exist = u.dataHolder[key]
	return
}
