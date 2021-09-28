package storage

import (
	"sync"

	"snakealive/m/entities"
)

type UserStorage interface {
	Add(key string, value entities.User)
	Get(key string) (value entities.User, exist bool)
	Delete(key string)
}

type userStorage struct {
	dataHolder map[string]entities.User

	mu *sync.RWMutex
}

func NewUserStorage() UserStorage {
	return &userStorage{
		dataHolder: make(map[string]entities.User),
		mu:         &sync.RWMutex{},
	}
}

func (u *userStorage) Add(key string, value entities.User) {
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

func (u *userStorage) Get(key string) (value entities.User, exist bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	value, exist = u.dataHolder[key]
	return
}

type userSyncStorage struct {
	dataHolder *sync.Map
}

func NewUserSyncStorage() UserStorage {
	return &userSyncStorage{
		dataHolder: &sync.Map{},
	}
}

func (u *userSyncStorage) Add(key string, value entities.User) {
	u.dataHolder.Store(key, value)
}

func (u *userSyncStorage) Delete(key string) {
	u.dataHolder.Delete(key)
}

func (u *userSyncStorage) Get(key string) (value entities.User, exist bool) {
	valueCast, exist := u.dataHolder.Load(key)
	if !exist {
		return value, exist
	}
	value = valueCast.(entities.User)

	return
}
