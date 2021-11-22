package domain

import "github.com/valyala/fasthttp"

type Trip struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" valid:"required"`
	Description string    `json:"description"`
	Days        [][]Place `json:"days" valid:"required"`
}

type TripHandler interface {
	Trip(ctx *fasthttp.RequestCtx)
	GetPlaceForTripQuery(ctx *fasthttp.RequestCtx)
	AddTrip(ctx *fasthttp.RequestCtx)
	Update(ctx *fasthttp.RequestCtx)
	Delete(ctx *fasthttp.RequestCtx)
}

type TripStorage interface {
	Add(value Trip, user User) (int, error)
	GetById(id int) (value Trip, err error)
	GetPlaceForTripQuery(id int) (value PlaceCoords, err error)
	Delete(id int) error
	Update(id int, value Trip) error
	GetTripAuthor(id int) int
}

type TripUseCase interface {
	Add(value Trip, user User) (int, error)
	GetById(id int) (int, []byte)
	GetPlaceForTripQuery(id int) (int, []byte)
	Delete(id int) error
	Update(id int, updatedTrip Trip) error
	CheckAuthor(user User, id int) bool
	SanitizeTrip(trip Trip) Trip
}
