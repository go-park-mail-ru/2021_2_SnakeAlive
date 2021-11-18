package countryUseCase

import (
	"encoding/json"
	"snakealive/m/internal/domain"
	logs "snakealive/m/pkg/logger"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewCountryUseCase(countryStorage domain.CountryStorage, l *logs.Logger) domain.CountryUseCase {
	return countryUsecase{
		countryStorage: countryStorage,
		l:              l,
	}
}

type countryUsecase struct {
	countryStorage domain.CountryStorage
	l              *logs.Logger
}

func (u countryUsecase) GetCountriesList() (int, []byte) {
	var code int
	countries, err := u.countryStorage.GetCountriesList()
	if err != nil {
		u.l.Logger.Error("error while GetCountriesList: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}

	bytes, err := json.Marshal(countries)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}
	code = fasthttp.StatusOK
	return code, bytes
}

func (u countryUsecase) GetById(id int) (int, []byte) {
	var code int
	code = fasthttp.StatusOK
	country, err := u.countryStorage.GetById(id)
	if err != nil {
		u.l.Logger.Error("error while GetById: ", zap.Error(err))
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	if (country == domain.Country{}) {
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	bytes, err := json.Marshal(country)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}

	return code, bytes
}

func (u countryUsecase) GetByName(name string) (int, []byte) {
	var code int
	code = fasthttp.StatusOK
	country, err := u.countryStorage.GetByName(name)
	if err != nil {
		u.l.Logger.Error("error while GetByName: ", zap.Error(err))
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	if (country == domain.Country{}) {
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	bytes, err := json.Marshal(country)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}
	return code, bytes
}
