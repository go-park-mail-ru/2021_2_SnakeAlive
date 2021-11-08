package domain

import "github.com/valyala/fasthttp"

type Country struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}

type Countries []Country

type CountryHandler interface {
	GetCountriesList(ctx *fasthttp.RequestCtx)
	GetById(ctx *fasthttp.RequestCtx)
	GetByName(ctx *fasthttp.RequestCtx)
}

type CountryStorage interface {
	GetById(id int) (value Country, err error)
	GetByName(name string) (country Country, err error)
	GetCountriesList() (countries Countries, err error)
}

type CountryUseCase interface {
	GetById(id int) (code int, country []byte)
	GetByName(name string) (code int, country []byte)
	GetCountriesList() (code int, countries []byte)
}
