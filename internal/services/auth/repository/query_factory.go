package repository

import (
	"strings"

	"snakealive/m/internal/services/auth/models"
	"snakealive/m/pkg/query"
)

type QueryFactory interface {
	CreateGetUserByEmail(email string) *query.Query
	CreateGetUserByID(ID int64) *query.Query
	CreateCreateUser(user *models.User) *query.Query
	CreateUpdateUser(user *models.User) *query.Query
	CreateCreateUserSession(userID int64, hash string) *query.Query
	CreateValidateUserSession(hash string) *query.Query
	CreateRemoveUserSession(hash string) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateCreateUser(user *models.User) *query.Query {
	return &query.Query{
		Request: createUserRequest,
		Params: []interface{}{
			user.Name, user.Surname, user.Email, user.Password,
		},
	}
}

func (q *queryFactory) CreateGetUserByEmail(email string) *query.Query {
	return &query.Query{
		Request: getUserByEmailRequest,
		Params:  []interface{}{email},
	}
}

func (q *queryFactory) CreateGetUserByID(ID int64) *query.Query {
	return &query.Query{
		Request: getUserByIDRequest,
		Params:  []interface{}{ID},
	}
}

func (q *queryFactory) CreateUpdateUser(user *models.User) *query.Query {
	params := []interface{}{
		user.ID,
	}
	optionals := make([]string, 0)

	if user.Name != "" {
		params = append(params, user.Name)
		optionals = append(optionals, updateUserName)
	}
	if user.Surname != "" {
		params = append(params, user.Surname)
		optionals = append(optionals, updateUserSurname)
	}
	if user.Password != "" {
		params = append(params, user.Password)
		optionals = append(optionals, updateUserPass)
	}
	if user.Email != "" {
		params = append(params, user.Email)
		optionals = append(optionals, updateUserEmail)
	}
	if user.Description != "" {
		params = append(params, user.Description)
		optionals = append(optionals, updateUserDescription)
	}

	return &query.Query{
		Request: updateUserRequest + strings.Join(optionals, ",") + " " + updateUserReturning,
		Params:  params,
	}
}

func (q *queryFactory) CreateCreateUserSession(userID int64, hash string) *query.Query {
	return &query.Query{
		Request: createUserSession,
		Params:  []interface{}{userID, hash},
	}
}

func (q *queryFactory) CreateValidateUserSession(hash string) *query.Query {
	return &query.Query{
		Request: validateUserSession,
		Params:  []interface{}{hash},
	}
}

func (q *queryFactory) CreateRemoveUserSession(hash string) *query.Query {
	return &query.Query{
		Request: removeUserSession,
		Params:  []interface{}{hash},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
