package errors

import "errors"

var (
	UserDoesNotExist    = errors.New("user does not exist")
	SessionDoesNotExist = errors.New("session does not exist")
	WrongUserPassword   = errors.New("wrong user password")
)

var DeniedAccess = errors.New("user doesn't have permission for this action")
var TripNotFound = errors.New("trip not found")

var (
	GatewayErrorMsgSessionNotFound = "сессия отсутствует"
)