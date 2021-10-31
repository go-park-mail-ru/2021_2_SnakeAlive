package domain

type Trip struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" valid:"required"`
	Description string    `json:"description"`
	Days        [][]Place `json:"days" valid:"required"`
}

type TripStorage interface {
	Add(value Trip, user User) (int, error)
	GetById(id int) (value Trip, err error)
	Delete(id int) error
	Update(id int, value Trip) error
	GetTripAuthor(id int) int
}

type TripUseCase interface {
	Add(value Trip, user User) (int, error)
	GetById(id int) (int, []byte)
	Delete(id int) error
	Update(id int, updatedTrip Trip) error
	CheckAuthor(user User, id int) bool
}
