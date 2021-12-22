package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGetCountriesList(t *testing.T) {
	f := NewQueryFactory()
	q := f.CreateGetCountriesList()
	assert.Equal(t, q.Request, GetCountriesList)
	assert.Equal(t, q.Params, []interface{}{})
}

func TestCreateGetCountryByID(t *testing.T) {
	f := NewQueryFactory()
	q := f.CreateGetCountryByID(1)
	assert.Equal(t, q.Request, GetCountryById)
	assert.Equal(t, q.Params, []interface{}{1})
}

func TestCreateGetCountryByName(t *testing.T) {
	f := NewQueryFactory()
	q := f.CreateGetCountryByName("name")
	assert.Equal(t, q.Request, GetCountryByName)
	assert.Equal(t, q.Params, []interface{}{"name"})
}
