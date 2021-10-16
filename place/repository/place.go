package placeRepository

import (
	"snakealive/m/domain"
	ent "snakealive/m/entities"
	"sync"
)

type placeStorage struct {
	dataHolder map[string][]domain.Place
	mu         *sync.RWMutex
}

func NewPlaceStorage() domain.PlaceStorage {
	return &placeStorage{
		dataHolder: ent.PlacesDB,
		mu:         &sync.RWMutex{},
	}
}

func (u *placeStorage) Get(key string) (value []domain.Place, exist bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	value, exist = u.dataHolder[key]
	return
}
