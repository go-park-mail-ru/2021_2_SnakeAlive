package repository

const (
	GetCountriesList = `SELECT id, name, description, photo FROM Countries`
	GetCountryById   = `SELECT id, name, description, photo FROM Countries WHERE id = $1`
	GetCountryByName = `SELECT id, name, description, photo FROM Countries WHERE name = $1`
)
