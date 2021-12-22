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