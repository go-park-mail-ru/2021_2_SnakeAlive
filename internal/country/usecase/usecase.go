package countryUseCase

import (
	"encoding/json"
	logs "snakealive/m/internal/logger"
	"snakealive/m/pkg/domain"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewCountryUseCase(countryStorage domain.CountryStorage) domain.CountryUseCase {
	return countryUsecase{countryStorage: countryStorage}
}

type countryUsecase struct {
	countryStorage domain.CountryStorage
}

func (u countryUsecase) GetCountriesList() (int, []byte) {
	logger := logs.GetLogger()
	var code int
	countries, err := u.countryStorage.GetCountriesList()
	if err != nil {
		logger.Error("error while GetCountriesList: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}
	bytes, err := json.Marshal(countries)
	if err != nil {
		logger.Error("error while marshalling JSON: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}
	code = fasthttp.StatusOK
	return code, bytes
}

func (u countryUsecase) GetById(id int) (int, []byte) {
	logger := logs.GetLogger()
	var code int
	code = fasthttp.StatusOK
	country, err := u.countryStorage.GetById(id)
	if err != nil {
		logger.Error("error while GetById: ", zap.Error(err))
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	if (country == domain.Country{}) {
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	bytes, err := json.Marshal(country)
	if err != nil {
		logger.Error("error while marshalling JSON: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}

	return code, bytes
}

func (u countryUsecase) GetByName(name string) (int, []byte) {
	logger := logs.GetLogger()
	var code int
	code = fasthttp.StatusOK
	country, err := u.countryStorage.GetByName(name)
	if err != nil {
		logger.Error("error while GetByName: ", zap.Error(err))
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	if (country == domain.Country{}) {
		code = fasthttp.StatusNotFound
		return code, []byte("{}")
	}
	bytes, err := json.Marshal(country)
	if err != nil {
		logger.Error("error while marshalling JSON: ", zap.Error(err))
		code = fasthttp.StatusInternalServerError
	}
	return code, bytes
}
