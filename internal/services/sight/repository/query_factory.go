package repository

import "snakealive/m/pkg/query"

type QueryFactory interface {
	CreateGetSightByID(id int) *query.Query
	CreateGetSightsByCountry(country string) *query.Query
	CreateGetSightsByIDs(ids []int64) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetSightByID(id int) *query.Query {
	return &query.Query{
		Request: GetSightByIdQuery,
		Params:  []interface{}{id},
	}
}

func (q *queryFactory) CreateGetSightsByCountry(country string) *query.Query {
	return &query.Query{
		Request: GetSightsByCountryQuery,
		Params:  []interface{}{country},
	}
}

func (q *queryFactory) CreateGetSightsByIDs(ids []int64) *query.Query {
	return &query.Query{
		Request: GetSightsByIDs,
		Params:  []interface{}{ids},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
