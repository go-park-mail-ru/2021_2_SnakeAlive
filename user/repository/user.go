package userRepository

import (
	"snakealive/m/domain"
	ent "snakealive/m/entities"
	"sync"
)

type userStorage struct {
	dataHolder map[string]domain.User
	mu         *sync.RWMutex
}

func NewUserStorage() domain.UserStorage {
	return &userStorage{
		dataHolder: ent.AuthDB,
		mu:         &sync.RWMutex{},
	}
}

func (u *userStorage) Add(key string, value domain.User) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.dataHolder[key] = value
}

func (u *userStorage) Delete(key string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, exist := u.dataHolder[key]; exist {
		delete(u.dataHolder, key)
	}
}

func (u *userStorage) Get(key string) (value domain.User, exist bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	value, exist = u.dataHolder[key]
	return
}
