package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"snakealive/m/internal/gateway/country/usecase"
	"snakealive/m/pkg/error_adapter"
)

type CountryDelivery interface {
	ListCountries(ctx *fasthttp.RequestCtx)
	GetCountryByID(ctx *fasthttp.RequestCtx)
	GetCountryByName(ctx *fasthttp.RequestCtx)
}

type countryDelivery struct {
	manager      usecase.CountryUsecase
	errorAdapter error_adapter.HttpAdapter
}

func (c *countryDelivery) ListCountries(ctx *fasthttp.RequestCtx) {
	countries, err := c.manager.GetCountriesList(ctx)
	if err != nil {
		httpError := c.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.Response.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(countries)
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(b)
}

func (c *countryDelivery) GetCountryByID(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	country, err := c.manager.GetById(ctx, id)
	if err != nil {
		httpError := c.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.Response.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(country)
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(b)
}

func (c *countryDelivery) GetCountryByName(ctx *fasthttp.RequestCtx) {
	name, _ := strconv.Atoi(ctx.UserValue("name").(string))
	country, err := c.manager.GetById(ctx, name)
	if err != nil {
		httpError := c.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.Response.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := easyjson.Marshal(country)
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(b)
}

func NewCountryDelivery(manager usecase.CountryUsecase, errorAdapter error_adapter.HttpAdapter) CountryDelivery {
	return &countryDelivery{
		manager:      manager,
		errorAdapter: errorAdapter,
	}
}
