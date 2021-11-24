package usecase

import (
	"context"

	"snakealive/m/internal/gateway/country/repository"
	"snakealive/m/internal/models"
)

type CountryUsecase interface {
	GetCountriesList(ctx context.Context) (models.Countries, error)
	GetById(ctx context.Context, id int) (models.Country, error)
	GetByName(ctx context.Context, name string) (models.Country, error)
}

type countryUsecase struct {
	countryRepo repository.CountryStorage
}

func (c *countryUsecase) GetCountriesList(ctx context.Context) (models.Countries, error) {
	return c.countryRepo.GetCountriesList(ctx)
}

func (c *countryUsecase) GetById(ctx context.Context, id int) (models.Country, error) {
	return c.countryRepo.GetById(ctx, id)
}

func (c *countryUsecase) GetByName(ctx context.Context, name string) (models.Country, error) {
	return c.countryRepo.GetByName(ctx, name)
}

func NewCountryUsecase(countryRepo repository.CountryStorage) CountryUsecase {
	return &countryUsecase{countryRepo: countryRepo}
}
