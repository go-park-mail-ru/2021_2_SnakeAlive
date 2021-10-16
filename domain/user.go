package domain

import (
	"sync"
)

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStorage interface {
	Add(key string, value User)
	Get(key string) (value User, exist bool)
	Delete(key string)
}

type userStorage struct {
	dataHolder map[string]User

	mu *sync.RWMutex
}

func NewUserStorage() UserStorage {
	return &userStorage{
		dataHolder: AuthDB,
		mu:         &sync.RWMutex{},
	}
}

func (u *userStorage) Add(key string, value User) {
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

func (u *userStorage) Get(key string) (value User, exist bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	value, exist = u.dataHolder[key]
	return
}

type Usecase interface {
	Get(key string) (User, bool)
}

// type UserUsecase interface {
// 	Fetch(ctx *fasthttp.RequestCtx, cursor string, num int64) ([]User, string, error)
// 	Update(ctx *fasthttp.RequestCtx, u *User) error
// 	Delete(ctx *fasthttp.RequestCtx, id int64) error
// 	GetByEmail(ctx *fasthttp.RequestCtx, email string) (User, found bool)
// 	Validate() bool
// }

// // ArticleRepository represent the article's repository contract
// type UserRepository interface {
// 	Fetch(ctx *fasthttp.RequestCtx, cursor string, num int64) (res []User, nextCursor string, err error)
// 	GetByID(ctx *fasthttp.RequestCtx, id int64) (User, error)
// }
