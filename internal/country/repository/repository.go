package countryRepository

import (
	"context"
	"snakealive/m/internal/domain"
	logs "snakealive/m/pkg/logger"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type countryStorage struct {
	dataHolder *pgxpool.Pool
}

func NewCountryStorage(DB *pgxpool.Pool) domain.CountryStorage {
	return &countryStorage{dataHolder: DB}
}

func (u *countryStorage) GetCountriesList() (domain.Countries, error) {
	logger := logs.GetLogger()
	countries := make(domain.Countries, 0)
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return countries, err
	}
	defer conn.Release()

	const GetCountriesListQuery = `SELECT id, name, description, photo FROM Countries`
	rows, err := conn.Query(context.Background(), GetCountriesListQuery)
	if err != nil {
		logger.Error("error while getting places from database")
		return countries, err
	}
	var country domain.Country

	for rows.Next() {
		err = rows.Scan(&country.Id, &country.Name, &country.Description, &country.Photo)
		countries = append(countries, country)
	}
	if rows.Err() != nil {
		logger.Error("error while scanning places from database")
		return countries, err
	}
	if len(countries) == 0 {
		logger.Error("no countries found")
		return countries, err
	}

	return countries, err
}

func (u *countryStorage) GetById(id int) (domain.Country, error) {
	logger := logs.GetLogger()
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return domain.Country{}, err
	}
	defer conn.Release()

	const GetCountryByIdQuery = `SELECT id, name, description, photo FROM Countries WHERE id = $1`
	var country domain.Country

	err = conn.Conn().QueryRow(context.Background(),
		GetCountryByIdQuery, id).Scan(&country.Id, &country.Name, &country.Description, &country.Photo)

	if err != nil {
		logger.Error("error while scanning country: ", zap.Error(err))
		return country, err
	}
	return country, err
}

func (u *countryStorage) GetByName(name string) (domain.Country, error) {
	logger := logs.GetLogger()
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return domain.Country{}, err
	}
	defer conn.Release()

	const GetCountryByNameQuery = `SELECT id, name, description, photo FROM Countries WHERE name = $1`
	var country domain.Country

	err = conn.Conn().QueryRow(context.Background(),
		GetCountryByNameQuery, name).Scan(&country.Id, &country.Name, &country.Description, &country.Photo)

	if err != nil {
		logger.Error("error while scanning country: ", zap.Error(err))
		return country, err
	}
	return country, err
}
