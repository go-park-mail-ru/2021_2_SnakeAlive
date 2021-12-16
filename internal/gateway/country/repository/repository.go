package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"snakealive/m/internal/models"
	"snakealive/m/pkg/errors"
)

type CountryStorage interface {
	GetCountriesList(ctx context.Context) (models.Countries, error)
	GetById(ctx context.Context, id int) (models.Country, error)
	GetByName(ctx context.Context, name string) (models.Country, error)
}

type countryStorage struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (u *countryStorage) GetCountriesList(ctx context.Context) (models.Countries, error) {
	countries := make(models.Countries, 0)

	request := u.queryFactory.CreateGetCountriesList()
	rows, err := u.conn.Query(ctx, request.Request, request.Params...)
	if err != nil {
		return countries, err
	}

	for rows.Next() {
		var country models.Country
		_ = rows.Scan(&country.Id, &country.Name, &country.Description, &country.Photo, &country.Translated)
		countries = append(countries, country)
	}

	return countries, nil
}

func (u *countryStorage) GetById(ctx context.Context, id int) (models.Country, error) {
	request := u.queryFactory.CreateGetCountryByID(id)

	var country models.Country
	if err := u.conn.QueryRow(ctx, request.Request, request.Params...).Scan(
		&country.Id, &country.Name, &country.Description, &country.Photo, &country.Translated,
	); err != nil {
		if err == pgx.ErrNoRows {
			return country, errors.CountryDoesNotExist
		}

		return country, err
	}

	return country, nil
}

func (u *countryStorage) GetByName(ctx context.Context, name string) (models.Country, error) {
	request := u.queryFactory.CreateGetCountryByName(name)

	var country models.Country
	if err := u.conn.QueryRow(ctx, request.Request, request.Params...).Scan(
		&country.Id, &country.Name, &country.Description, &country.Photo, &country.Translated,
	); err != nil {
		if err == pgx.ErrNoRows {
			return country, errors.CountryDoesNotExist
		}

		return country, err
	}

	return country, nil
}

func NewCountryStorage(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) CountryStorage {
	return &countryStorage{
		queryFactory: queryFactory,
		conn:         conn,
	}
}
