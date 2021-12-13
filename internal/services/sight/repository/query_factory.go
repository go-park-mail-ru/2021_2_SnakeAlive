package repository

import (
	"strconv"
	"strings"

	"snakealive/m/internal/services/sight/models"
	"snakealive/m/pkg/constants"
	"snakealive/m/pkg/query"
)

type QueryFactory interface {
	CreateGetSightByID(id int) *query.Query
	CreateGetSightsByCountry(country string) *query.Query
	CreateGetSightsByIDs(ids []int64) *query.Query
	CreateSearchSights(req *models.SightsSearch) *query.Query
	CreateGetSightsByTag(tag int64) *query.Query
	CreateGetTags() *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateSearchSights(req *models.SightsSearch) *query.Query {
	if req.Limit == 0 {
		req.Limit = constants.DefaultLimit
	}

	pos := 3
	statements := []string{}
	params := []interface{}{
		req.Skip, req.Limit,
	}

	if len(req.Countries) != 0 {
		statements = append(statements, Country(pos))
		params = append(params, req.Countries)
		pos++
	}
	if len(req.Tags) != 0 {
		statements = append(statements, Tags+strconv.Itoa(pos))
		params = append(params, req.Tags)
		pos++
	}
	if req.Search != "" {
		statements = append(statements, SearchStatement(pos))
		params = append(params, strings.ToLower(req.Search), strings.ToLower(req.Search) + "%")
	}

	return &query.Query{
		Request: SearchSights + strings.Join(statements, " AND ") + " " + Offset + strconv.Itoa(1) + " " + Limit + strconv.Itoa(2),
		Params:  params,
	}
}

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

func (q *queryFactory) CreateGetSightsByTag(tag int64) *query.Query {
	return &query.Query{
		Request: GetSightsByTag,
		Params:  []interface{}{tag},
	}
}

func (q *queryFactory) CreateGetTags() *query.Query {
	return &query.Query{
		Request: GetTags,
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
