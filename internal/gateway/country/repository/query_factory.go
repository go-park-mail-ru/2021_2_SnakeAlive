package repository

import "snakealive/m/pkg/query"

type QueryFactory interface {
	CreateGetCountriesList() *query.Query
	CreateGetCountryByID(id int) *query.Query
	CreateGetCountryByName(name string) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetCountriesList() *query.Query {
	return &query.Query{
		Request: GetCountriesList,
		Params:  []interface{}{},
	}
}

func (q *queryFactory) CreateGetCountryByID(id int) *query.Query {
	return &query.Query{
		Request: GetCountryById,
		Params:  []interface{}{id},
	}
}

func (q *queryFactory) CreateGetCountryByName(name string) *query.Query {
	return &query.Query{
		Request: GetCountryByName,
		Params:  []interface{}{name},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
