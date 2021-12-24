package errors

import "errors"

var (
	UserDoesNotExist    = errors.New("user does not exist")
	SessionDoesNotExist = errors.New("session does not exist")
	WrongUserPassword   = errors.New("wrong user password")
	SightDoesNotExist   = errors.New("sight does not exist")
	CountryDoesNotExist = errors.New("country does not exist")
)

var DeniedAccess = errors.New("user doesn't have permission for this action")
var TripNotFound = errors.New("trip not found")
var UserIsAlreadyAuthor = errors.New("user is already author")

var (
	GatewayErrorMsgSessionNotFound = "сессия отсутствует"
)
